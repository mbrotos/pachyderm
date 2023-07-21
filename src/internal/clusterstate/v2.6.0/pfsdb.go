package v2_6_0

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	v2_5_0 "github.com/pachyderm/pachyderm/v2/src/internal/clusterstate/v2.5.0"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/log"
	"github.com/pachyderm/pachyderm/v2/src/internal/pachsql"
	"github.com/pachyderm/pachyderm/v2/src/internal/storage/fileset"
	"github.com/pachyderm/pachyderm/v2/src/pfs"
)

const (
	maxStmts = 100
)

func validateExistingDAGs(cis []*v2_5_0.CommitInfo) error {
	// group duplicate commits by branchless key
	duplicates := make(map[string]map[string]*v2_5_0.CommitInfo) // branchless commit key -> { old commit key ->  commit info }
	for _, ci := range cis {
		if _, ok := duplicates[commitBranchlessKey(ci.Commit)]; !ok {
			duplicates[commitBranchlessKey(ci.Commit)] = make(map[string]*v2_5_0.CommitInfo)
		}
		duplicates[commitBranchlessKey(ci.Commit)][oldCommitKey(ci.Commit)] = ci
	}
	// the only duplicate commits we allow and handle in the migration are those connected through direct ancestry.
	// the allowed duplicate commits are expected to arise from branch triggers / deferred processing.
	var badCommitSets []string
	for _, dups := range duplicates {
		if len(dups) <= 1 {
			continue
		}
		seen := make(map[string]struct{})
		var commitSet string
		for _, d := range dups {
			commitSet = d.Commit.Id
			if d.ParentCommit != nil {
				if _, ok := dups[oldCommitKey(d.ParentCommit)]; ok {
					seen[oldCommitKey(d.Commit)] = struct{}{}
				}
			}
		}
		if len(dups)-len(seen) > 1 {
			badCommitSets = append(badCommitSets, commitSet)
		}
	}
	if badCommitSets != nil {
		return errors.Errorf("invalid commit sets detected: %v", badCommitSets)
	}
	return nil
}

func removeAliasCommits(ctx context.Context, tx *pachsql.Tx) error {
	cis, err := listCollectionProtos(ctx, tx, "commits", &v2_5_0.CommitInfo{})
	if err != nil {
		return err
	}
	type commit struct {
		id   int
		info *v2_5_0.CommitInfo
	}
	commits := make(map[string]*commit)
	id := 1
	for _, ci := range cis {
		c := &commit{info: ci}
		if ci.Origin.Kind != 4 {
			c.id = id
			id++
		}
		commits[oldCommitKey(ci.Commit)] = c
	}
	getCommitId := func(c *pfs.Commit) int {
		return commits[oldCommitKey(c)].id
	}
	getCommitInfo := func(c *pfs.Commit) *v2_5_0.CommitInfo {
		return commits[oldCommitKey(c)].info
	}
	var deleteCommits []*v2_5_0.CommitInfo
	// Insert the commits into pfs.commits.
	if err := func() (retErr error) {
		ctx, end := log.SpanContext(ctx, "insertCommits")
		defer end(log.Errorp(&retErr))
		batcher := NewPostgresBatcher(ctx, tx, maxStmts)
		for _, ci := range cis {
			// Handle alias commit.
			if ci.Origin.Kind == 4 {
				deleteCommits = append(deleteCommits, ci)
				continue
			}
			stmt := fmt.Sprintf(`INSERT INTO pfs.commits (commit_id, commit_set_id) VALUES ('%v', '%v')`, oldCommitKey(ci.Commit), ci.Commit.Id)
			if err := batcher.Add(stmt); err != nil {
				return err
			}
		}
		return batcher.Close()
	}(); err != nil {
		return err
	}
	realAncestorCommits := make(map[string]*pfs.Commit)
	// Update the ancestry and provenance of the commits.
	if err := func() (retErr error) {
		ctx, end := log.SpanContext(ctx, "updateCommits")
		defer end(log.Errorp(&retErr))
		batcher := NewPostgresBatcher(ctx, tx, maxStmts)
		for _, ci := range cis {
			// Skip alias commit.
			if ci.Origin.Kind == 4 {
				continue
			}
			ci := proto.Clone(ci).(*v2_5_0.CommitInfo)
			// Update the parent and child commits.
			if ci.ParentCommit != nil {
				ci.ParentCommit = getRealAncestorCommit(getCommitInfo, realAncestorCommits, ci.ParentCommit)
			}
			var childCommits []*pfs.Commit
			for _, child := range ci.ChildCommits {
				childCommits = append(childCommits, getRealDescendantCommits(getCommitInfo, child)...)
			}
			ci.ChildCommits = childCommits
			// Update the direct provenance.
			var directProvenance []*pfs.Commit
			for _, b := range ci.DirectProvenance {
				toC := b.NewCommit(ci.Commit.Id)
				// If the provenant commit's repo was deleted, its
				// reference may still exist in the old provenance, so
				// we skip mapping it over to the new model.
				if _, ok := commits[oldCommitKey(toC)]; !ok {
					continue
				}
				toC = getRealAncestorCommit(getCommitInfo, realAncestorCommits, toC)
				directProvenance = append(directProvenance, toC)
				from := strconv.Itoa(getCommitId(ci.Commit))
				to := strconv.Itoa(getCommitId(toC))
				stmt := fmt.Sprintf(`INSERT INTO pfs.commit_provenance(from_id, to_id) VALUES (%v, %v)`, from, to)
				if err := batcher.Add(stmt); err != nil {
					return err
				}
			}
			newCI := convertCommitInfoToV2_6_0(ci)
			newCI.DirectProvenance = directProvenance
			data, err := proto.Marshal(newCI)
			if err != nil {
				return errors.EnsureStack(err)
			}
			stmt := fmt.Sprintf("UPDATE collections.commits SET proto=decode('%v', 'hex') WHERE key='%v'", hex.EncodeToString(data), oldCommitKey(newCI.Commit))
			if err := batcher.Add(stmt); err != nil {
				return err
			}
		}
		return batcher.Close()
	}(); err != nil {
		return err
	}
	// Validate the alias commits that will be deleted.
	var count int
	for _, ci := range deleteCommits {
		log.Info(ctx, "validating deleted alias commit",
			zap.String("commit", oldCommitKey(ci.Commit)),
			zap.String("progress", fmt.Sprintf("%v/%v", count, len(deleteCommits))),
		)
		same, err := sameFileSets(ctx, tx, ci.Commit, ci.ParentCommit)
		if err != nil {
			return err
		}
		if !same {
			return errors.Errorf("commit %q is listed as ALIAS but has a different ID than its first real ancestor",
				oldCommitKey(ci.Commit))
		}
		count++
	}
	// Delete the alias commits.
	if err := func() (retErr error) {
		ctx, end := log.SpanContext(ctx, "deleteAliasCommits")
		defer end(log.Errorp(&retErr))
		batcher := NewPostgresBatcher(ctx, tx, maxStmts)
		for _, ci := range deleteCommits {
			key := oldCommitKey(ci.Commit)
			stmt := fmt.Sprintf(`DELETE FROM collections.commits WHERE key='%v'`, key)
			if err := batcher.Add(stmt); err != nil {
				return err
			}
			stmt = fmt.Sprintf(`DELETE FROM pfs.commit_diffs WHERE commit_id='%v'`, key)
			if err := batcher.Add(stmt); err != nil {
				return err
			}
			stmt = fmt.Sprintf(`DELETE FROM pfs.commit_totals WHERE commit_id='%v'`, key)
			if err := batcher.Add(stmt); err != nil {
				return err
			}
			// TODO: Why does LIKE perform so poorly?
			stmt = fmt.Sprintf(`DELETE FROM storage.tracker_objects WHERE str_id > 'commit/%v' AND str_id < 'commit/%v/z'`, key, key)
			if err := batcher.Add(stmt); err != nil {
				return err
			}
		}
		return batcher.Close()
	}(); err != nil {
		return err
	}
	bis, err := listCollectionProtos(ctx, tx, "branches", &pfs.BranchInfo{})
	if err != nil {
		return err
	}
	for _, bi := range bis {
		if err := updateBranch(ctx, tx, bi.Branch, func(bi *pfs.BranchInfo) {
			bi.Head = getRealAncestorCommit(getCommitInfo, realAncestorCommits, bi.Head)
		}); err != nil {
			return errors.Wrap(err, "update headless branches")
		}
	}
	return nil
}

func getRealAncestorCommit(getCommitInfo func(*pfs.Commit) *v2_5_0.CommitInfo, realAncestorCommits map[string]*pfs.Commit, c *pfs.Commit) *pfs.Commit {
	realAncestorCommit, ok := realAncestorCommits[oldCommitKey(c)]
	if ok {
		return realAncestorCommit
	}
	ci := getCommitInfo(c)
	if ci.Origin.Kind != 4 {
		return c
	}
	realAncestorCommit = getRealAncestorCommit(getCommitInfo, realAncestorCommits, ci.ParentCommit)
	realAncestorCommits[oldCommitKey(c)] = realAncestorCommit
	return realAncestorCommit
}

func getRealDescendantCommits(getCommitInfo func(*pfs.Commit) *v2_5_0.CommitInfo, c *pfs.Commit) []*pfs.Commit {
	ci := getCommitInfo(c)
	if ci.Origin.Kind != 4 {
		return []*pfs.Commit{c}
	}
	var childCommits []*pfs.Commit
	for _, child := range ci.ChildCommits {
		childCommits = append(childCommits, getRealDescendantCommits(getCommitInfo, child)...)
	}
	return childCommits
}

func convertCommitInfoToV2_6_0(ci *v2_5_0.CommitInfo) *pfs.CommitInfo {
	return &pfs.CommitInfo{
		Commit:              ci.Commit,
		Origin:              ci.Origin,
		Description:         ci.Description,
		ParentCommit:        ci.ParentCommit,
		ChildCommits:        ci.ChildCommits,
		Started:             ci.Started,
		Finishing:           ci.Finishing,
		Finished:            ci.Finished,
		Error:               ci.Error,
		SizeBytesUpperBound: ci.SizeBytesUpperBound,
		Details:             ci.Details,
	}
}

func sameFileSets(ctx context.Context, tx *pachsql.Tx, c1 *pfs.Commit, c2 *pfs.Commit) (bool, error) {
	md1, err := getFileSetMd(ctx, tx, c1)
	if err != nil {
		return false, err
	}
	md2, err := getFileSetMd(ctx, tx, c2)
	if err != nil {
		return false, err
	}
	if md1 == nil || md2 == nil {
		// the semantics here are a little odd - if either commit doesn't have a fileset yet,
		// we assume that they aren't different and are therefore the same.
		return true, nil
	}
	if md1.GetPrimitive() != nil && md2.GetPrimitive() != nil {
		ser1, err := proto.Marshal(md1.GetPrimitive())
		if err != nil {
			return false, errors.Wrap(err, "marshal first primitive fileset")
		}
		ser2, err := proto.Marshal(md2.GetPrimitive())
		if err != nil {
			return false, errors.Wrap(err, "marshal second primitive fileset")
		}
		return bytes.Equal(ser1, ser2), nil
	} else if md1.GetComposite() != nil && md2.GetComposite() != nil {
		comp1Layers := md1.GetComposite().Layers
		comp2Layers := md2.GetComposite().Layers
		if len(comp1Layers) != len(comp2Layers) {
			return false, nil
		}
		for i, l := range comp1Layers {
			if l != comp2Layers[i] {
				return false, nil
			}
		}
		return true, nil
	} else {
		return false, nil
	}
}

func getFileSetMd(ctx context.Context, tx *pachsql.Tx, c *pfs.Commit) (*fileset.Metadata, error) {
	var mdData []byte
	if err := tx.GetContext(ctx, &mdData, `SELECT metadata_pb FROM storage.filesets JOIN pfs.commit_totals ON id = fileset_id WHERE commit_id = $1`, oldCommitKey(c)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.EnsureStack(err)
	}
	md := &fileset.Metadata{}
	if err := proto.Unmarshal(mdData, md); err != nil {
		return nil, errors.EnsureStack(err)
	}
	return md, nil
}

func updateBranch(ctx context.Context, tx *pachsql.Tx, b *pfs.Branch, f func(bi *pfs.BranchInfo)) error {
	k := branchKey(b)
	bi := &pfs.BranchInfo{}
	if err := getCollectionProto(ctx, tx, "branches", k, bi); err != nil {
		return errors.Wrapf(err, "get branch info for branch: %q", b)
	}
	f(bi)
	return updateCollectionProto(ctx, tx, "branches", k, k, bi)
}

func deleteDanglingCommitRefs(ctx context.Context, tx *pachsql.Tx) (retErr error) {
	log.Info(ctx, "checking for dangling commits")
	parseRepo := func(key string) *pfs.Repo {
		slashSplit := strings.Split(key, "/")
		dotSplit := strings.Split(slashSplit[1], ".")
		return &pfs.Repo{
			Project: &pfs.Project{Name: slashSplit[0]},
			Name:    dotSplit[0],
			Type:    dotSplit[1],
		}
	}
	parseBranch := func(key string) *pfs.Branch {
		split := strings.Split(key, "@")
		return &pfs.Branch{
			Repo: parseRepo(split[0]),
			Name: split[1],
		}
	}
	listRepoKeys := func(tx *pachsql.Tx) (map[string]struct{}, error) {
		var keys []string
		if err := tx.Select(&keys, `SELECT key FROM collections.repos`); err != nil {
			return nil, errors.Wrap(err, "select keys from collections.repos")
		}
		rs := make(map[string]struct{})
		for _, k := range keys {
			rs[k] = struct{}{}
		}
		return rs, nil
	}
	parseCommit_2_5 := func(key string) (*pfs.Commit, error) {
		split := strings.Split(key, "=")
		if len(split) != 2 {
			return nil, errors.Errorf("parsing commit key with 2.6.x+ structure %q", key)
		}
		b := parseBranch(split[0])
		return &pfs.Commit{
			Repo:   b.Repo,
			Branch: b,
			Id:     split[1],
		}, nil
	}
	listReferencedCommits := func(tx *pachsql.Tx) (map[string]*pfs.Commit, error) {
		cs := make(map[string]*pfs.Commit)
		var err error
		var ids []string
		if err := tx.Select(&ids, `SELECT commit_id FROM pfs.commit_totals`); err != nil {
			return nil, errors.Wrap(err, "select commit ids from pfs.commit_totals")
		}
		for _, id := range ids {
			cs[id], err = parseCommit_2_5(id)
			if err != nil {
				return nil, err
			}
		}
		ids = make([]string, 0)
		if err := tx.Select(&ids, `SELECT commit_id FROM pfs.commit_diffs`); err != nil {
			return nil, errors.Wrap(err, "select commit ids from pfs.commit_diffs")
		}
		for _, id := range ids {
			cs[id], err = parseCommit_2_5(id)
			if err != nil {
				return nil, err
			}
		}
		return cs, nil
	}
	cs, err := listReferencedCommits(tx)
	if err != nil {
		return errors.Wrap(err, "list referenced commits")
	}
	rs, err := listRepoKeys(tx)
	if err != nil {
		return errors.Wrap(err, "list repos")
	}
	var dangCommitKeys []string
	for _, c := range cs {
		if _, ok := rs[repoKey(c.Repo)]; !ok {
			dangCommitKeys = append(dangCommitKeys, oldCommitKey(c))
		}
	}
	if len(dangCommitKeys) > 0 {
		log.Info(ctx, "detected dangling commit references", zap.Any("references", dangCommitKeys))
	}
	ctx, end := log.SpanContext(ctx, "deleteDanglingCommits")
	defer end(log.Errorp(&retErr))
	batcher := NewPostgresBatcher(ctx, tx, maxStmts)
	for _, id := range dangCommitKeys {
		stmt := fmt.Sprintf(`DELETE FROM pfs.commit_totals WHERE commit_id = '%v'`, id)
		if err := batcher.Add(stmt); err != nil {
			return err
		}
		stmt = fmt.Sprintf(`DELETE FROM pfs.commit_diffs WHERE commit_id = '%v'`, id)
		if err := batcher.Add(stmt); err != nil {
			return err
		}
	}
	return batcher.Close()
}

// the goal of this migration is to migrate commits to be unqiuely indentified by a (repo, UUID), instead of by (repo, branch, UUID)
//
// 1) commits with equal (repo, UUID) pairs must be de-duplicated. This is primarily expected to happen in deferred processing cases.
// In such cases, the deferred branch's commit is expected to be a child commit of the leading branch's commit. So to make the substitution,
// we can simply delete the deferred branch's commit and re-point it's direct subvenant commits to point to the master commit as their
// direct provenance commits.
//
// 2) assert that we didn't miss any duplicate commits, and error the migration if we did.
//
// 3) commit keys can now be substituted from <project>/<repo>@<branch>=<id> -> <project>/<repo>@<id>
func branchlessCommitsPFS(ctx context.Context, tx *pachsql.Tx) error {
	// Update commits table.
	cis, err := listCollectionProtos(ctx, tx, "commits", &pfs.CommitInfo{})
	if err != nil {
		return err
	}
	if err := func() (retErr error) {
		ctx, end := log.SpanContext(ctx, "branchlessUpdateCommits")
		defer end(log.Errorp(&retErr))
		batcher := NewPostgresBatcher(ctx, tx, maxStmts)
		for _, ci := range cis {
			stmt := fmt.Sprintf(`UPDATE pfs.commits SET commit_id='%v' WHERE commit_id='%v'`, commitBranchlessKey(ci.Commit), oldCommitKey(ci.Commit))
			if err := batcher.Add(stmt); err != nil {
				return err
			}
			ci.Commit.Repo = ci.Commit.Branch.Repo
			if ci.ParentCommit != nil {
				ci.ParentCommit.Repo = ci.ParentCommit.Branch.Repo
			}
			for _, child := range ci.ChildCommits {
				child.Repo = child.Branch.Repo
			}
			for _, prov := range ci.DirectProvenance {
				prov.Repo = prov.Branch.Repo
			}
			data, err := proto.Marshal(ci)
			if err != nil {
				return errors.EnsureStack(err)
			}
			stmt = fmt.Sprintf(`UPDATE collections.commits SET key='%v', proto=decode('%v', 'hex') WHERE key='%v'`, commitBranchlessKey(ci.Commit), hex.EncodeToString(data), oldCommitKey(ci.Commit))
			if err := batcher.Add(stmt); err != nil {
				return err
			}
		}
		return batcher.Close()
	}(); err != nil {
		return err
	}
	// Update branches table.
	bis, err := listCollectionProtos(ctx, tx, "branches", &pfs.BranchInfo{})
	if err != nil {
		return err
	}
	for _, bi := range bis {
		if err := updateBranch(ctx, tx, bi.Branch, func(bi *pfs.BranchInfo) {
			bi.Head.Repo = bi.Head.Branch.Repo
		}); err != nil {
			return errors.Wrap(err, "update headless branches")
		}
	}
	// Update everything else.
	return branchlessCommitKeysPFS(ctx, tx)
}

// map <project>/<repo>@<branch>=<id> -> <project>/<repo>@<id>; basically replace '@[-a-zA-Z0-9_]+)=' -> '@'
func branchlessCommitKeysPFS(ctx context.Context, tx *pachsql.Tx) (retErr error) {
	ctx, end := log.SpanContext(ctx, "branchlessCommitKeysPFS")
	defer end(log.Errorp(&retErr))
	log.Info(ctx, "removing branch from pfs.commit_totals")
	if _, err := tx.ExecContext(ctx, `UPDATE pfs.commit_totals SET commit_id = regexp_replace(commit_id, '@(.+)=', '@');`); err != nil {
		return errors.Wrap(err, "update pfs.commit_totals to branchless commit ids")
	}
	log.Info(ctx, "removing branch from pfs.commit_diffs")
	if _, err := tx.ExecContext(ctx, `UPDATE pfs.commit_diffs SET commit_id = regexp_replace(commit_id, '@(.+)=', '@');`); err != nil {
		return errors.Wrap(err, "update pfs.commit_diffs to branchless commits ids")
	}
	log.Info(ctx, "removing branch from storage.tracker_objects")
	if _, err := tx.ExecContext(ctx, `UPDATE storage.tracker_objects SET str_id = regexp_replace(str_id, '@(.+)=', '@') WHERE str_id LIKE 'commit/%';`); err != nil {
		return errors.Wrap(err, "update storage.tracker_objects to branchless commit ids")
	}
	return nil
}
