//go:build k8s

package server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/go-units"
	"github.com/pachyderm/pachyderm/v2/src/debug"
	"github.com/pachyderm/pachyderm/v2/src/internal/minikubetestenv"
	"github.com/pachyderm/pachyderm/v2/src/internal/pachsql"
	"github.com/pachyderm/pachyderm/v2/src/internal/pctx"
	"github.com/pachyderm/pachyderm/v2/src/internal/storage/fileset/index"
	"github.com/pachyderm/pachyderm/v2/src/internal/tarutil"
	"github.com/pachyderm/pachyderm/v2/src/internal/testsnowflake"
	"google.golang.org/protobuf/proto"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/pachyderm/pachyderm/v2/src/internal/client"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/errutil"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"
	tu "github.com/pachyderm/pachyderm/v2/src/internal/testutil"
	"github.com/pachyderm/pachyderm/v2/src/pfs"
	"github.com/pachyderm/pachyderm/v2/src/pps"
)

func newCountBreakFunc(maxCount int) func(func() error) error {
	var count int
	return func(cb func() error) error {
		if err := cb(); err != nil {
			return err
		}
		count++
		if count == maxCount {
			return errutil.ErrBreak
		}
		return nil
	}
}

func basicPipelineReq(name, input string) *pps.CreatePipelineRequest {
	return &pps.CreatePipelineRequest{
		Pipeline: client.NewPipeline(pfs.DefaultProjectName, name),
		Transform: &pps.Transform{
			Cmd: []string{"bash"},
			Stdin: []string{
				fmt.Sprintf("cp /pfs/%s/* /pfs/out/", input),
			},
		},
		ParallelismSpec: &pps.ParallelismSpec{
			Constant: 1,
		},
		Input: client.NewPFSInput(pfs.DefaultProjectName, input, "/*"),
	}
}

func createTestCommits(t *testing.T, repoName, branchName string, numCommits int, c *client.APIClient) {
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repoName), "pods should be available to create a repo against")

	for i := 0; i < numCommits; i++ {
		commit, err := c.StartCommit(pfs.DefaultProjectName, repoName, branchName)
		require.NoError(t, err, "creating a test commit should succeed")

		err = c.PutFile(commit, "/file", bytes.NewBufferString("file contents"))
		require.NoError(t, err, "should be able to add a file to a commit")

		err = c.FinishCommit(pfs.DefaultProjectName, repoName, branchName, commit.Id)
		require.NoError(t, err, "finishing a commit should succeed")
	}
}

func deleteEtcd(t *testing.T, ctx context.Context, namespace string) {
	t.Helper()
	kubeClient := tu.GetKubeClient(t)
	label := "app=etcd"
	etcd, err := kubeClient.CoreV1().Pods(namespace).
		List(ctx, metav1.ListOptions{LabelSelector: label})
	require.NoError(t, err, "Error attempting to find etcd pod.")
	require.Equal(t, 1, len(etcd.Items), "etcd did not have the correct number of pods.")

	watcher, err := kubeClient.CoreV1().Pods(namespace).
		Watch(ctx, metav1.ListOptions{LabelSelector: label})
	defer watcher.Stop()
	require.NoError(t, err, "Starting etcd pod watch failed.")

	require.NoError(t, kubeClient.CoreV1().Pods(namespace).Delete(
		context.Background(),
		etcd.Items[0].ObjectMeta.Name, metav1.DeleteOptions{}), "Deleting etcd pod failed.")

	for event := range watcher.ResultChan() {
		if event.Type == watch.Deleted {
			break
		}
	}
}

// Wait for at least one pod with the given selector to be running and ready.
func waitForOnePodReady(t testing.TB, ctx context.Context, namespace string, label string) {
	t.Helper()
	kubeClient := tu.GetKubeClient(t)
	require.NoErrorWithinTRetryConstant(t, 1*time.Minute, func() error {
		pods, err := kubeClient.CoreV1().Pods(namespace).
			List(ctx, metav1.ListOptions{LabelSelector: label})
		if err != nil {
			return errors.EnsureStack(err)
		}

		if len(pods.Items) < 1 {
			return errors.Errorf("pod with label %s has not yet been restarted.", label)
		}
		for _, item := range pods.Items {
			if item.Status.Phase == v1.PodRunning {
				for _, c := range item.Status.Conditions {
					if c.Type == v1.PodReady && c.Status == v1.ConditionTrue {
						return nil
					}
				}
			}
		}
		return errors.Errorf("one pod with label %s is not yet running and ready.", label)
	}, 5*time.Second)
}

func TestDatumSetCache(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	c = c.WithDefaultTransformUser("1000")
	dataRepo := tu.UniqueString("TestDatumSetCache_data")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	masterCommit := client.NewCommit(pfs.DefaultProjectName, dataRepo, "master", "")
	require.NoError(t, c.WithModifyFileClient(masterCommit, func(mfc client.ModifyFile) error {
		for i := 0; i < 90; i++ {
			require.NoError(t, mfc.PutFile(strconv.Itoa(i), strings.NewReader("")))
		}
		return nil
	}))
	pipeline := tu.UniqueString("TestDatumSetCache")
	_, err := c.PpsAPIClient.CreatePipeline(context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipeline),
			Transform: &pps.Transform{
				Cmd: []string{"bash"},
				Stdin: []string{
					fmt.Sprintf("cp /pfs/%s/* /pfs/out/", dataRepo),
					"sleep 1",
				},
			},
			Input:        client.NewPFSInput(pfs.DefaultProjectName, dataRepo, "/*"),
			DatumSetSpec: &pps.DatumSetSpec{Number: 1},
		})
	require.NoError(t, err)
	ctx, cancel := pctx.WithCancel(context.Background())
	defer cancel()
	go func() {
		ticker := time.NewTimer(50 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				deleteEtcd(t, ctx, ns)
				waitForOnePodReady(t, ctx, ns, "app=etcd")

				select {
				case <-ctx.Done():
					return
				default:
					ticker = time.NewTimer(50 * time.Second)
				}
			}
		}
	}()
	commitInfo, err := c.InspectCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	require.NoErrorWithinTRetry(t, 60*time.Second, func() error {
		_, err = c.WaitCommitSetAll(commitInfo.Commit.Id)
		return err
	})
	for i := 0; i < 5; i++ {
		_, err := c.InspectFile(commitInfo.Commit, strconv.Itoa(i))
		require.NoError(t, err)
	}
}

func TestTemporaryDuplicatedPath(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	repo := tu.UniqueString("repo")
	other := tu.UniqueString("other")
	pipeline := tu.UniqueString("pipeline")

	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repo))
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, other))

	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, other, "master", ""), "empty",
		strings.NewReader("")))

	// add an output file bigger than the sharding threshold so that two in a row
	// will fall on either side of a naive shard division
	bigSize := index.DefaultShardSizeThreshold * 5 / 4
	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, repo, "master", ""), "a",
		strings.NewReader(strconv.Itoa(bigSize/units.MB))))

	// add some more files to push the fileset that deletes the old datums out of the first fan-in
	for i := 0; i < 10; i++ {
		require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, repo, "master", ""), fmt.Sprintf("b-%d", i),
			strings.NewReader("1")))
	}

	req := basicPipelineReq(pipeline, repo)
	req.DatumSetSpec = &pps.DatumSetSpec{
		Number: 1,
	}
	req.Transform.Stdin = []string{
		fmt.Sprintf("for f in /pfs/%s/* ; do", repo),
		"dd if=/dev/zero of=/pfs/out/$(basename $f) bs=1M count=$(cat $f)",
		"done",
	}
	_, err := c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err := c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	require.Equal(t, "", commitInfo.Error)

	// update the pipeline to take new input, but produce identical output
	// the filesets in the output of the resulting job should look roughly like
	// [old output] [new output] [deletion of old output]
	// If a path is included in multiple shards, it would appear twice for both
	// the old and new datums, while only one copy would be deleted,
	// leading to either a validation error or a repeated path/datum pair
	req.Update = true
	req.Input = client.NewCrossInput(
		client.NewPFSInput(pfs.DefaultProjectName, repo, "/*"),
		client.NewPFSInput(pfs.DefaultProjectName, other, "/*"))
	_, err = c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err = c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	require.Equal(t, "", commitInfo.Error)
}

func TestValidationFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	repo := tu.UniqueString(t.Name())
	pipeline := tu.UniqueString("pipeline-" + t.Name())

	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repo))

	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, repo, "master", ""), "foo", strings.NewReader("baz")))
	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, repo, "master", ""), "bar", strings.NewReader("baz")))

	req := basicPipelineReq(pipeline, repo)
	req.Transform.Stdin = []string{
		fmt.Sprintf("cat /pfs/%s/* > /pfs/out/overlap", repo), // write datums to the same path
	}
	_, err := c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err := c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	require.NotEqual(t, "", commitInfo.Error)

	// update the pipeline to fix things
	req.Update = true
	req.Transform.Stdin = []string{
		fmt.Sprintf("cp /pfs/%s/* /pfs/out", repo),
	}

	_, err = c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err = c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	require.Equal(t, "", commitInfo.Error)
}

// TestPPSEgressToSnowflake tests basic Egress functionality via PPS mechanism,
// and how it handles inserting the same primary key twice.
func TestPPSEgressToSnowflake(t *testing.T) {
	// setup
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	// create input repo with CSV
	repo := tu.UniqueString(t.Name())
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repo))

	// create output database, and destination table
	ctx := pctx.TestContext(t)
	db, dbName := testsnowflake.NewEphemeralSnowflakeDB(ctx, t)
	require.NoError(t, pachsql.CreateTestTable(db, "test_table", struct {
		Id int    `column:"ID" dtype:"INT" constraint:"PRIMARY KEY"`
		A  string `column:"A" dtype:"VARCHAR(100)"`
	}{}))

	// create K8s secrets
	b := []byte(fmt.Sprintf(`
	{
		"apiVersion": "v1",
		"kind": "Secret",
		"stringData": {
			"PACHYDERM_SQL_PASSWORD": "%s"
		},
		"metadata": {
			"name": "egress-secret",
			"creationTimestamp": null
		}
	}`, os.Getenv("SNOWSQL_PWD")))
	require.NoError(t, c.CreateSecret(b))

	// create a pipeline with egress
	pipeline := tu.UniqueString("egress")
	dsn, err := testsnowflake.DSN()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.PpsAPIClient.CreatePipeline(
		c.Ctx(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipeline),
			Transform: &pps.Transform{
				Image: tu.DefaultTransformImage,
				Cmd:   []string{"bash"},
				Stdin: []string{"cp -r /pfs/in/* /pfs/out/"},
			},
			Input: &pps.Input{Pfs: &pps.PFSInput{
				Repo: repo,
				Glob: "/",
				Name: "in",
			}},
			Egress: &pps.Egress{
				Target: &pps.Egress_SqlDatabase{SqlDatabase: &pfs.SQLDatabaseEgress{
					Url: fmt.Sprintf("%s/%s", dsn, dbName),
					FileFormat: &pfs.SQLDatabaseEgress_FileFormat{
						Type: pfs.SQLDatabaseEgress_FileFormat_CSV,
					},
					Secret: &pfs.SQLDatabaseEgress_Secret{
						Name: "egress-secret",
						Key:  "PACHYDERM_SQL_PASSWORD",
					},
				},
				},
			},
		},
	); err != nil {
		t.Fatal(err)
	}

	// Initial load
	master := client.NewCommit(pfs.DefaultProjectName, repo, "master", "")
	require.NoError(t, c.PutFile(master, "/test_table/0000", strings.NewReader("1,Foo\n2,Bar")))
	commitInfo, err := c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	_, err = c.InspectJob(pfs.DefaultProjectName, pipeline, commitInfo.Commit.Id, false)
	require.NoError(t, err)
	// query db for results
	var count, expected int
	expected = 2
	require.NoError(t, db.QueryRow("select count(*) from test_table").Scan(&count))
	require.Equal(t, expected, count)

	// Add a new row, and test whether primary key conflicts
	require.NoError(t, c.PutFile(master, "/test_table/0000", strings.NewReader("1,Foo\n2,Bar\n3,ABC")))
	commitInfo, err = c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	_, err = c.InspectJob(pfs.DefaultProjectName, pipeline, commitInfo.Commit.Id, false)
	require.NoError(t, err)
	// query db for results
	expected = 3
	require.NoError(t, db.QueryRow("select count(*) from test_table").Scan(&count))
	require.Equal(t, expected, count)
}

func TestMissingSecretFailure(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	kc := tu.GetKubeClient(t).CoreV1().Secrets(ns)

	secret := tu.UniqueString(strings.ToLower(t.Name() + "-secret"))
	repo := tu.UniqueString(t.Name() + "-data")
	pipeline := tu.UniqueString(t.Name())
	_, err := kc.Create(c.Ctx(), &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: secret},
		StringData: map[string]string{"foo": "bar"}}, metav1.CreateOptions{})
	require.NoError(t, err)

	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repo))
	req := basicPipelineReq(pipeline, repo)
	req.Transform.Secrets = append(req.Transform.Secrets, &pps.SecretMount{
		Name:   secret,
		Key:    "foo",
		EnvVar: "MY_SECRET_ENV_VAR",
	})
	req.Autoscaling = true
	_, err = c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	// wait for system to go into standby
	require.NoErrorWithinTRetry(t, 60*time.Second, func() error {
		_, err = c.WaitCommit(pfs.DefaultProjectName, repo, "master", "")
		if err != nil {
			return err
		}
		info, err := c.InspectPipeline(pfs.DefaultProjectName, pipeline, false)
		if err != nil {
			return err
		}
		if info.State != pps.PipelineState_PIPELINE_STANDBY {
			return errors.Errorf("pipeline in %s instead of standby", info.State)
		}
		return nil
	})

	require.NoError(t, kc.Delete(c.Ctx(), secret, metav1.DeleteOptions{}))
	// force pipeline to come back up, and check for failure
	require.NoError(t, c.PutFile(
		client.NewCommit(pfs.DefaultProjectName, repo, "master", ""), "foo", strings.NewReader("bar")))
	require.NoErrorWithinTRetry(t, 30*time.Second, func() error {
		info, err := c.InspectPipeline(pfs.DefaultProjectName, pipeline, false)
		if err != nil {
			return err
		}
		if info.State != pps.PipelineState_PIPELINE_CRASHING {
			return errors.Errorf("pipeline in %s instead of crashing", info.State)
		}
		return nil
	})
}
func TestDatabaseStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	type rowCountResults struct {
		NLiveTup int    `json:"n_live_tup"`
		RelName  string `json:"relname"`
	}

	type commitResults struct {
		Key string `json:"key"`
	}

	repoName := tu.UniqueString("TestDatabaseStats-repo")
	branchName := "master"
	numCommits := 100
	buf := &bytes.Buffer{}
	filter := &debug.Filter{
		Filter: &debug.Filter_Database{Database: true},
	}

	c, _ := minikubetestenv.AcquireCluster(t)

	createTestCommits(t, repoName, branchName, numCommits, c)
	time.Sleep(5 * time.Second) // give some time for the stats collector to run.
	require.NoError(t, c.Dump(filter, 100, buf), "dumping database files should succeed")
	gr, err := gzip.NewReader(buf)
	require.NoError(t, err)

	var rows []commitResults
	foundCommitsJSON, foundRowCounts := false, false
	require.NoError(t, tarutil.Iterate(gr, func(f tarutil.File) error {
		fileContents := &bytes.Buffer{}
		if err := f.Content(fileContents); err != nil {
			return errors.EnsureStack(err)
		}

		hdr, err := f.Header()
		require.NoError(t, err, "getting database tar file header should succeed")
		require.NotMatch(t, "^[a-zA-Z0-9_\\-\\/]+\\.error$", hdr.Name)
		switch hdr.Name {
		case "database/row-counts.json":
			var rows []rowCountResults
			require.NoError(t, json.Unmarshal(fileContents.Bytes(), &rows),
				"unmarshalling row-counts.json should succeed")

			for _, row := range rows {
				if row.RelName == "commits" {
					require.NotEqual(t, 0, row.NLiveTup,
						"some commits from createTestCommits should be accounted for")
					foundRowCounts = true
				}
			}
		case "database/tables/collections/commits.json":
			require.NoError(t, json.Unmarshal(fileContents.Bytes(), &rows),
				"unmarshalling commits.json should succeed")
			require.Equal(t, numCommits, len(rows), "number of commits should match number of rows")
			foundCommitsJSON = true
		}

		return nil
	}))

	require.Equal(t, true, foundRowCounts,
		"we should have an entry in row-counts.json for commits")
	require.Equal(t, true, foundCommitsJSON,
		"checks for commits.json should succeed")
}

func TestSimplePipelineNonRoot(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	dataRepo := tu.UniqueString("TestSimplePipeline_data")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))

	commit1, err := c.StartCommit(pfs.DefaultProjectName, dataRepo, "master")
	require.NoError(t, err)
	require.NoError(t, c.PutFile(commit1, "file", strings.NewReader("foo"), client.WithAppendPutFile()))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, dataRepo, commit1.Branch.Name, commit1.Id))
	pipeline := tu.UniqueString("TestSimplePipeline")
	req := basicPipelineReq(pipeline, dataRepo)
	req.Transform.User = "65534"
	_, err = c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err := c.InspectCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	commitInfos, err := c.WaitCommitSetAll(commitInfo.Commit.Id)
	require.NoError(t, err)
	// The commitset should have a commit in: data, spec, pipeline, meta
	// the last two are dependent upon the first two, so should come later
	// in topological ordering
	require.Equal(t, 4, len(commitInfos))
	var commitRepos []*pfs.Repo
	for _, info := range commitInfos {
		commitRepos = append(commitRepos, info.Commit.Branch.Repo)
	}
	require.EqualOneOf(t, commitRepos[:2], client.NewRepo(pfs.DefaultProjectName, dataRepo))
	require.EqualOneOf(t, commitRepos[:2], client.NewSystemRepo(pfs.DefaultProjectName, pipeline, pfs.SpecRepoType))
	require.EqualOneOf(t, commitRepos[2:], client.NewRepo(pfs.DefaultProjectName, pipeline))
	require.EqualOneOf(t, commitRepos[2:], client.NewSystemRepo(pfs.DefaultProjectName, pipeline, pfs.MetaRepoType))

	var buf bytes.Buffer
	for _, info := range commitInfos {
		if proto.Equal(info.Commit.Branch.Repo, client.NewRepo(pfs.DefaultProjectName, pipeline)) {
			require.NoError(t, c.GetFile(info.Commit, "file", &buf))
			require.Equal(t, "foo", buf.String())
		}
	}
}

func TestSimplePipelinePodPatchNonRoot(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	dataRepo := tu.UniqueString("TestSimplePipeline_data")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))

	commit1, err := c.StartCommit(pfs.DefaultProjectName, dataRepo, "master")
	require.NoError(t, err)
	require.NoError(t, c.PutFile(commit1, "file", strings.NewReader("foo"), client.WithAppendPutFile()))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, dataRepo, commit1.Branch.Name, commit1.Id))
	pipeline := tu.UniqueString("TestSimplePipeline")
	req := basicPipelineReq(pipeline, dataRepo)

	req.PodPatch = "[{\"op\":\"add\",\"path\":\"/securityContext\",\"value\":{}},{\"op\":\"add\",\"path\":\"/securityContext\",\"value\":{\"runAsGroup\":1000,\"runAsUser\":1000,\"fsGroup\":1000,\"runAsNonRoot\":true,\"seccompProfile\":{\"type\":\"RuntimeDefault\"}}},{\"op\":\"add\",\"path\":\"/containers/0/securityContext\",\"value\":{}},{\"op\":\"add\",\"path\":\"/containers/0/securityContext\",\"value\":{\"runAsGroup\":1000,\"runAsUser\":1000,\"allowPrivilegeEscalation\":false,\"capabilities\":{\"drop\":[\"all\"]},\"readOnlyRootFilesystem\":true}}]"
	_, err = c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)

	commitInfo, err := c.InspectCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	commitInfos, err := c.WaitCommitSetAll(commitInfo.Commit.Id)
	require.NoError(t, err)
	// The commitset should have a commit in: data, spec, pipeline, meta
	// the last two are dependent upon the first two, so should come later
	// in topological ordering
	require.Equal(t, 4, len(commitInfos))
	var commitRepos []*pfs.Repo
	for _, info := range commitInfos {
		commitRepos = append(commitRepos, info.Commit.Branch.Repo)
	}
	require.EqualOneOf(t, commitRepos[:2], client.NewRepo(pfs.DefaultProjectName, dataRepo))
	require.EqualOneOf(t, commitRepos[:2], client.NewSystemRepo(pfs.DefaultProjectName, pipeline, pfs.SpecRepoType))
	require.EqualOneOf(t, commitRepos[2:], client.NewRepo(pfs.DefaultProjectName, pipeline))
	require.EqualOneOf(t, commitRepos[2:], client.NewSystemRepo(pfs.DefaultProjectName, pipeline, pfs.MetaRepoType))

	var buf bytes.Buffer
	for _, info := range commitInfos {
		if proto.Equal(info.Commit.Branch.Repo, client.NewRepo(pfs.DefaultProjectName, pipeline)) {
			require.NoError(t, c.GetFile(info.Commit, "file", &buf))
			require.Equal(t, "foo", buf.String())
		}
	}
}

func TestZombieCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	repo := tu.UniqueString(t.Name() + "-data")
	pipeline := tu.UniqueString(t.Name())
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, repo))

	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, repo, "master", ""),
		"foo", strings.NewReader("baz")))

	req := basicPipelineReq(pipeline, repo)
	_, err := c.PpsAPIClient.CreatePipeline(c.Ctx(), req)
	require.NoError(t, err)
	require.NoErrorWithinT(t, 30*time.Second, func() error {
		_, err := c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
		return err
	})
	// fsck should succeed (including zombie check, on by default)
	require.NoError(t, c.Fsck(false, func(response *pfs.FsckResponse) error {
		return errors.Errorf("got fsck error: %s", response.Error)
	}, client.WithZombieCheckAll()))

	// stop pipeline so we can modify
	require.NoError(t, c.StopPipeline(pfs.DefaultProjectName, pipeline))
	// create new commits on output and meta
	_, err = c.ExecuteInTransaction(func(c *client.APIClient) error {
		if _, err := c.StartCommit(pfs.DefaultProjectName, pipeline, "master"); err != nil {
			return errors.EnsureStack(err)
		}
		metaBranch := client.NewSystemRepo(pfs.DefaultProjectName, pipeline, pfs.MetaRepoType).NewBranch("master")
		if _, err := c.PfsAPIClient.StartCommit(c.Ctx(), &pfs.StartCommitRequest{Branch: metaBranch}); err != nil {
			return errors.EnsureStack(err)
		}
		_, err = c.PfsAPIClient.FinishCommit(c.Ctx(), &pfs.FinishCommitRequest{Commit: metaBranch.NewCommit("")})
		return errors.EnsureStack(err)
	})
	require.NoError(t, err)
	// add a file to the output with a datum that doesn't exist
	require.NoError(t, c.PutFile(client.NewCommit(pfs.DefaultProjectName, pipeline, "master", ""),
		"zombie", strings.NewReader("zombie"), client.WithDatumPutFile("zombie")))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, pipeline, "master", ""))
	_, err = c.WaitCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	var messages []string
	// fsck should notice the zombie file
	require.NoError(t, c.Fsck(false, func(response *pfs.FsckResponse) error {
		messages = append(messages, response.Error)
		return nil
	}, client.WithZombieCheckAll()))
	require.Equal(t, 1, len(messages))
	require.Matches(t, "zombie", messages[0])
}
