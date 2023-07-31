//go:build k8s

package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/pachyderm/pachyderm/v2/src/auth"
	"github.com/pachyderm/pachyderm/v2/src/internal/backoff"
	"github.com/pachyderm/pachyderm/v2/src/internal/minikubetestenv"
	"github.com/pachyderm/pachyderm/v2/src/internal/pctx"
	"github.com/pachyderm/pachyderm/v2/src/internal/tarutil"
	"path"
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

func TestPipelineJobHasAuthToken(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)
	tu.ActivateAuthClient(t, c)
	rc := tu.AuthenticateClient(t, c, auth.RootUser)

	dataRepo := tu.UniqueString("TestPipelineJobAuthToken_data")
	require.NoError(t, rc.CreateRepo(pfs.DefaultProjectName, dataRepo))

	commit, err := rc.StartCommit(pfs.DefaultProjectName, dataRepo, "master")
	require.NoError(t, err)
	require.NoError(t, rc.PutFile(commit, "file", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, rc.FinishCommit(pfs.DefaultProjectName, dataRepo, commit.Branch.Name, commit.Id))

	pipeline := tu.UniqueString("pipeline")
	require.NoError(t, rc.CreatePipeline(pfs.DefaultProjectName,
		pipeline,
		"",
		[]string{"bash"},
		[]string{
			"echo $PACH_TOKEN | pachctl auth use-auth-token",
			"pachctl config update context --pachd-address $(echo grpc://pachd.$PACH_NAMESPACE.svc.cluster.local:30660)",
			"echo $(pachctl auth whoami) >/pfs/out/file2",
			"pachctl list repo",
		},
		&pps.ParallelismSpec{
			Constant: 1,
		},
		client.NewPFSInput(pfs.DefaultProjectName, dataRepo, "/*"),
		"",
		false,
	))
	var jobInfos []*pps.JobInfo
	require.NoError(t, backoff.Retry(func() error {
		jobInfos, err = c.ListJob(pfs.DefaultProjectName, pipeline, nil, -1, true)
		require.NoError(t, err)
		if len(jobInfos) != 1 {
			return errors.Errorf("expected 1 jobs, got %d", len(jobInfos))
		}
		return nil
	}, backoff.NewTestingBackOff()))
	jobInfo, err := rc.WaitJob(pfs.DefaultProjectName, pipeline, jobInfos[0].Job.Id, false)
	require.NoError(t, err)
	require.Equal(t, pps.JobState_JOB_SUCCESS, jobInfo.State)
	buffer := bytes.Buffer{}
	require.NoError(t, c.GetFile(jobInfo.OutputCommit, "file2", &buffer))
	require.Equal(t, fmt.Sprintf("You are \"pipeline:default/%s\"\n", jobInfo.Job.Pipeline.Name), buffer.String())
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

func testGetLogs(t *testing.T, useLoki bool) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	var opts []minikubetestenv.Option
	if !useLoki {
		opts = append(opts, minikubetestenv.SkipLokiOption)
	}
	c, _ := minikubetestenv.AcquireCluster(t, opts...)
	iter := c.GetLogs(pfs.DefaultProjectName, "", "", nil, "", false, false, 0)
	for iter.Next() {
	}
	require.NoError(t, iter.Err())
	// create repos
	dataRepo := tu.UniqueString("data")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	dataCommit := client.NewCommit(pfs.DefaultProjectName, dataRepo, "master", "")
	// create pipeline
	pipelineName := tu.UniqueString("pipeline")
	_, err := c.PpsAPIClient.CreatePipeline(context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
			Transform: &pps.Transform{
				Cmd: []string{"sh"},
				Stdin: []string{
					fmt.Sprintf("cp /pfs/%s/file /pfs/out/file", dataRepo),
					"echo foo",
					"echo %s", // %s tests a formatting bug we had (#2729)
				},
			},
			Input: client.NewPFSInput(pfs.DefaultProjectName, dataRepo, "/*"),
			ParallelismSpec: &pps.ParallelismSpec{
				Constant: 4,
			},
		})
	require.NoError(t, err)

	// Commit data to repo and flush commit
	commit, err := c.StartCommit(pfs.DefaultProjectName, dataRepo, "master")
	require.NoError(t, err)
	require.NoError(t, c.PutFile(commit, "file", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, dataRepo, "master", ""))
	_, err = c.WaitJobSetAll(commit.Id, false)
	require.NoError(t, err)

	// Get logs from pipeline, using a pipeline that doesn't exist. There should
	// be an error
	iter = c.GetLogs(pfs.DefaultProjectName, "__DOES_NOT_EXIST__", "", nil, "", false, false, 0)
	require.False(t, iter.Next())
	require.YesError(t, iter.Err())
	require.Matches(t, "could not get", iter.Err().Error())

	// Get logs from pipeline, using a job that doesn't exist. There should
	// be an error
	iter = c.GetLogs(pfs.DefaultProjectName, pipelineName, "__DOES_NOT_EXIST__", nil, "", false, false, 0)
	require.False(t, iter.Next())
	require.YesError(t, iter.Err())
	require.Matches(t, "could not get", iter.Err().Error())

	// This is put in a backoff because there's the possibility that pod was
	// evicted from k8s and is being re-initialized, in which case `GetLogs`
	// will appropriately fail. With the loki logging backend enabled the
	// eviction worry goes away, but is replaced with there being a window when
	// Loki hasn't scraped the logs yet so they don't show up.
	require.NoError(t, backoff.Retry(func() error {
		// Get logs from pipeline, using pipeline
		iter = c.GetLogs(pfs.DefaultProjectName, pipelineName, "", nil, "", false, false, 0)
		var numLogs, totalLogs int
		for iter.Next() {
			totalLogs++
			if !iter.Message().User {
				continue
			}
			numLogs++
			require.True(t, iter.Message().Message != "")
			require.False(t, strings.Contains(iter.Message().Message, "MISSING"), iter.Message().Message)
		}
		if numLogs < 2 {
			return errors.Errorf("didn't get enough log lines (%d total logs, %d non-user logs)", totalLogs, numLogs)
		}
		if err := iter.Err(); err != nil {
			return err
		}

		// Get logs from pipeline, using job
		// (1) Get job ID, from pipeline that just ran
		jobInfos, err := c.ListJob(pfs.DefaultProjectName, pipelineName, nil, -1, true)
		if err != nil {
			return err
		}
		require.Equal(t, 2, len(jobInfos))
		// (2) Get logs using extracted job ID
		// wait for logs to be collected
		time.Sleep(10 * time.Second)
		iter = c.GetLogs(pfs.DefaultProjectName, pipelineName, jobInfos[0].Job.Id, nil, "", false, false, 0)
		numLogs = 0
		for iter.Next() {
			numLogs++
			require.True(t, iter.Message().Message != "")
		}
		// Make sure that we've seen some logs
		if err = iter.Err(); err != nil {
			return err
		}
		require.True(t, numLogs > 0)

		// Get logs for datums but don't specify pipeline or job. These should error
		iter = c.GetLogs(pfs.DefaultProjectName, "", "", []string{"/foo"}, "", false, false, 0)
		require.False(t, iter.Next())
		require.YesError(t, iter.Err())

		dis, err := c.ListDatumAll(pfs.DefaultProjectName, jobInfos[0].Job.Pipeline.Name, jobInfos[0].Job.Id)
		if err != nil {
			return err
		}
		require.True(t, len(dis) > 0)
		iter = c.GetLogs(pfs.DefaultProjectName, "", "", nil, dis[0].Datum.Id, false, false, 0)
		require.False(t, iter.Next())
		require.YesError(t, iter.Err())

		// Filter logs based on input (using file that exists). Get logs using file
		// path, hex hash, and base64 hash, and make sure you get the same log lines
		fileInfo, err := c.InspectFile(dataCommit, "/file")
		if err != nil {
			return err
		}

		pathLog := c.GetLogs(pfs.DefaultProjectName, pipelineName, jobInfos[0].Job.Id, []string{"/file"}, "", false, false, 0)

		base64Hash := "kstrTGrFE58QWlxEpCRBt3aT8NJPNY0rso6XK7a4+wM="
		require.Equal(t, base64Hash, base64.StdEncoding.EncodeToString(fileInfo.Hash))
		base64Log := c.GetLogs(pfs.DefaultProjectName, pipelineName, jobInfos[0].Job.Id, []string{base64Hash}, "", false, false, 0)

		numLogs = 0
		for {
			havePathLog, haveBase64Log := pathLog.Next(), base64Log.Next()
			if havePathLog != haveBase64Log {
				return errors.Errorf("Unequal log lengths")
			}
			if !havePathLog {
				break
			}
			numLogs++
			if pathLog.Message().Message != base64Log.Message().Message {
				return errors.Errorf(
					"unequal logs, pathLogs: \"%s\" base64Log: \"%s\"",
					pathLog.Message().Message,
					base64Log.Message().Message)
			}
		}
		for _, logsiter := range []*client.LogsIter{pathLog, base64Log} {
			if logsiter.Err() != nil {
				return logsiter.Err()
			}
		}
		if numLogs == 0 {
			return errors.Errorf("no logs found")
		}

		// Filter logs based on input (using file that doesn't exist). There should
		// be no logs
		iter = c.GetLogs(pfs.DefaultProjectName, pipelineName, jobInfos[0].Job.Id, []string{"__DOES_NOT_EXIST__"}, "", false, false, 0)
		require.False(t, iter.Next())
		if err = iter.Err(); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		iter = c.WithCtx(ctx).GetLogs(pfs.DefaultProjectName, pipelineName, "", nil, "", false, false, 0)
		numLogs = 0
		for iter.Next() {
			numLogs++
			if numLogs == 8 {
				// Do another commit so there's logs to receive with follow
				_, err = c.StartCommit(pfs.DefaultProjectName, dataRepo, "master")
				if err != nil {
					return err
				}
				if err := c.PutFile(dataCommit, "file", strings.NewReader("bar\n"), client.WithAppendPutFile()); err != nil {
					return err
				}
				if err = c.FinishCommit(pfs.DefaultProjectName, dataRepo, "master", ""); err != nil {
					return err
				}
			}
			require.True(t, iter.Message().Message != "")
			if numLogs == 16 {
				break
			}
		}
		if err := iter.Err(); err != nil {
			return err
		}

		time.Sleep(time.Second * 30)

		numLogs = 0
		iter = c.WithCtx(ctx).GetLogs(pfs.DefaultProjectName, pipelineName, "", nil, "", false, false, 15*time.Second)
		for iter.Next() {
			numLogs++
		}
		if err := iter.Err(); err != nil {
			return err
		}
		if numLogs != 0 {
			return errors.Errorf("shouldn't return logs due to since time")
		}
		return nil
	}, backoff.NewTestingBackOff()))
}
func TestGetLogsWithoutLoki(t *testing.T) {
	testGetLogs(t, false)
}

func TestGetLogs(t *testing.T) {
	testGetLogs(t, true)
}

func TestAllDatumsAreProcessed(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)

	dataRepo1 := tu.UniqueString("TestAllDatumsAreProcessed_data1")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo1))
	dataRepo2 := tu.UniqueString("TestAllDatumsAreProcessed_data2")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo2))

	commit1, err := c.StartCommit(pfs.DefaultProjectName, dataRepo1, "master")
	require.NoError(t, err)
	require.NoError(t, c.PutFile(commit1, "file1", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, c.PutFile(commit1, "file2", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, dataRepo1, "master", ""))

	commit2, err := c.StartCommit(pfs.DefaultProjectName, dataRepo2, "master")
	require.NoError(t, err)
	require.NoError(t, c.PutFile(commit2, "file1", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, c.PutFile(commit2, "file2", strings.NewReader("foo\n"), client.WithAppendPutFile()))
	require.NoError(t, c.FinishCommit(pfs.DefaultProjectName, dataRepo2, "master", ""))

	pipeline := tu.UniqueString("pipeline")
	require.NoError(t, c.CreatePipeline(pfs.DefaultProjectName,
		pipeline,
		"",
		[]string{"bash"},
		[]string{
			fmt.Sprintf("cat /pfs/%s/* /pfs/%s/* > /pfs/out/file", dataRepo1, dataRepo2),
		},
		nil,
		client.NewCrossInput(
			client.NewPFSInput(pfs.DefaultProjectName, dataRepo1, "/*"),
			client.NewPFSInput(pfs.DefaultProjectName, dataRepo2, "/*"),
		),
		"",
		false,
	))

	commitInfo, err := c.InspectCommit(pfs.DefaultProjectName, pipeline, "master", "")
	require.NoError(t, err)
	commitInfos, err := c.WaitCommitSetAll(commitInfo.Commit.Id)
	require.NoError(t, err)
	require.Equal(t, 5, len(commitInfos))

	var buf bytes.Buffer
	rc, err := c.GetFileTAR(commitInfo.Commit, "file")
	require.NoError(t, err)
	defer rc.Close()
	require.NoError(t, tarutil.ConcatFileContent(&buf, rc))
	// should be 8 because each file gets copied twice due to cross product
	require.Equal(t, strings.Repeat("foo\n", 8), buf.String())
}

// TestSystemResourceRequest doesn't create any jobs or pipelines, it
// just makes sure that when pachyderm is deployed, we give pachd,
// and etcd default resource requests. This prevents them from overloading
// nodes and getting evicted, which can slow down or break a cluster.
func TestSystemResourceRequests(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	t.Parallel()
	_, ns := minikubetestenv.AcquireCluster(t)
	kubeClient := tu.GetKubeClient(t)

	// Expected resource requests for pachyderm system pods:
	defaultLocalMem := map[string]string{
		"pachd": "512M",
		"etcd":  "512M",
	}
	defaultLocalCPU := map[string]string{
		"pachd": "250m",
		"etcd":  "250m",
	}
	defaultCloudMem := map[string]string{
		"pachd": "3G",
		"etcd":  "2G",
	}
	defaultCloudCPU := map[string]string{
		"pachd": "1",
		"etcd":  "1",
	}
	// Get Pod info for 'app' from k8s
	var c v1.Container
	for _, app := range []string{"pachd", "etcd"} {
		err := backoff.Retry(func() error {
			podList, err := kubeClient.CoreV1().Pods(ns).List(
				context.Background(),
				metav1.ListOptions{
					LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
						map[string]string{"app": app, "suite": "pachyderm"},
					)),
				})
			if err != nil {
				return errors.EnsureStack(err)
			}
			if len(podList.Items) < 1 {
				return errors.Errorf("could not find pod for %s", app) // retry
			}
			c = podList.Items[0].Spec.Containers[0]
			return nil
		}, backoff.NewTestingBackOff())
		require.NoError(t, err)

		// Make sure the pod's container has resource requests
		cpu, ok := c.Resources.Requests[v1.ResourceCPU]
		require.True(t, ok, "could not get CPU request for "+app)
		require.True(t, cpu.String() == defaultLocalCPU[app] ||
			cpu.String() == defaultCloudCPU[app])
		mem, ok := c.Resources.Requests[v1.ResourceMemory]
		require.True(t, ok, "could not get memory request for "+app)
		require.True(t, mem.String() == defaultLocalMem[app] ||
			mem.String() == defaultCloudMem[app])
	}
}

// TestPipelineResourceRequest creates a pipeline with a resource request, and
// makes sure that's passed to k8s (by inspecting the pipeline's pods)
func TestPipelineResourceRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	// create repos
	dataRepo := tu.UniqueString("repo")
	pipelineName := tu.UniqueString("pipeline")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	// Resources are not yet in client.CreatePipeline() (we may add them later)
	_, err := c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
			Transform: &pps.Transform{
				Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
			},
			ParallelismSpec: &pps.ParallelismSpec{
				Constant: 1,
			},
			ResourceRequests: &pps.ResourceSpec{
				Memory: "100M",
				Cpu:    0.5,
				Disk:   "10M",
			},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)

	// Get info about the pipeline pods from k8s & check for resources
	pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, pipelineName, false)
	require.NoError(t, err)

	var container v1.Container
	kubeClient := tu.GetKubeClient(t)
	require.NoError(t, backoff.Retry(func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{
						"app":             "pipeline",
						"pipelineName":    pipelineInfo.Pipeline.Name,
						"pipelineVersion": fmt.Sprint(pipelineInfo.Version),
					},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err) // retry
		}
		if len(podList.Items) != 1 || len(podList.Items[0].Spec.Containers) == 0 {
			return errors.Errorf("could not find single container for pipeline %s", pipelineInfo.Pipeline.Name)
		}
		container = podList.Items[0].Spec.Containers[0]
		return nil // no more retries
	}, backoff.NewTestingBackOff()))
	// Make sure a CPU and Memory request are both set
	cpu, ok := container.Resources.Requests[v1.ResourceCPU]
	require.True(t, ok)
	require.Equal(t, "500m", cpu.String())
	mem, ok := container.Resources.Requests[v1.ResourceMemory]
	require.True(t, ok)
	require.Equal(t, "100M", mem.String())
	disk, ok := container.Resources.Requests[v1.ResourceEphemeralStorage]
	require.True(t, ok)
	require.Equal(t, "10M", disk.String())
}

func TestPipelineResourceLimit(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	// create repos
	dataRepo := tu.UniqueString("TestPipelineResourceLimit")
	pipelineName := tu.UniqueString("pipeline")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	// Resources are not yet in client.CreatePipeline() (we may add them later)
	_, err := c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
			Transform: &pps.Transform{
				Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
			},
			ParallelismSpec: &pps.ParallelismSpec{
				Constant: 1,
			},
			ResourceLimits: &pps.ResourceSpec{
				Memory: "100M",
				Cpu:    0.5,
			},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)

	// Get info about the pipeline pods from k8s & check for resources
	pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, pipelineName, false)
	require.NoError(t, err)

	var container v1.Container
	kubeClient := tu.GetKubeClient(t)
	err = backoff.Retry(func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{
						"app":             "pipeline",
						"pipelineName":    pipelineInfo.Pipeline.Name,
						"pipelineVersion": fmt.Sprint(pipelineInfo.Version),
						"suite":           "pachyderm"},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err) // retry
		}
		if len(podList.Items) != 1 || len(podList.Items[0].Spec.Containers) == 0 {
			return errors.Errorf("could not find single container for pipeline %s", pipelineInfo.Pipeline.Name)
		}
		container = podList.Items[0].Spec.Containers[0]
		return nil // no more retries
	}, backoff.NewTestingBackOff())
	require.NoError(t, err)
	// Make sure a CPU and Memory request are both set
	cpu, ok := container.Resources.Limits[v1.ResourceCPU]
	require.True(t, ok)
	require.Equal(t, "500m", cpu.String())
	mem, ok := container.Resources.Limits[v1.ResourceMemory]
	require.True(t, ok)
	require.Equal(t, "100M", mem.String())
}

func TestPipelineResourceLimitDefaults(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	// create repos
	dataRepo := tu.UniqueString("TestPipelineResourceLimit")
	pipelineName := tu.UniqueString("pipeline")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	// Resources are not yet in client.CreatePipeline() (we may add them later)
	_, err := c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
			Transform: &pps.Transform{
				Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
			},
			ParallelismSpec: &pps.ParallelismSpec{
				Constant: 1,
			},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)

	// Get info about the pipeline pods from k8s & check for resources
	pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, pipelineName, false)
	require.NoError(t, err)

	var container v1.Container
	kubeClient := tu.GetKubeClient(t)
	err = backoff.Retry(func() error {
		podList, err := kubeClient.CoreV1().Pods(ns).List(
			context.Background(),
			metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
					map[string]string{
						"app":             "pipeline",
						"pipelineName":    pipelineInfo.Pipeline.Name,
						"pipelineVersion": fmt.Sprint(pipelineInfo.Version),
						"suite":           "pachyderm"},
				)),
			})
		if err != nil {
			return errors.EnsureStack(err) // retry
		}
		if len(podList.Items) != 1 || len(podList.Items[0].Spec.Containers) == 0 {
			return errors.Errorf("could not find single container for pipeline %s", pipelineInfo.Pipeline.Name)
		}
		container = podList.Items[0].Spec.Containers[0]
		return nil // no more retries
	}, backoff.NewTestingBackOff())
	require.NoError(t, err)
	require.Nil(t, container.Resources.Limits)
}

func TestPipelinePartialResourceRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, _ := minikubetestenv.AcquireCluster(t)
	// create repos
	dataRepo := tu.UniqueString("TestPipelinePartialResourceRequest")
	pipelineName := tu.UniqueString("pipeline")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	// Resources are not yet in client.CreatePipeline() (we may add them later)
	_, err := c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, fmt.Sprintf("%s-%d", pipelineName, 0)),
			Transform: &pps.Transform{
				Cmd: []string{"true"},
			},
			ResourceRequests: &pps.ResourceSpec{
				Cpu:    0.25,
				Memory: "100M",
			},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)
	_, err = c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, fmt.Sprintf("%s-%d", pipelineName, 1)),
			Transform: &pps.Transform{
				Cmd: []string{"true"},
			},
			ResourceRequests: &pps.ResourceSpec{
				Memory: "100M",
			},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)
	_, err = c.PpsAPIClient.CreatePipeline(
		context.Background(),
		&pps.CreatePipelineRequest{
			Pipeline: client.NewPipeline(pfs.DefaultProjectName, fmt.Sprintf("%s-%d", pipelineName, 2)),
			Transform: &pps.Transform{
				Cmd: []string{"true"},
			},
			ResourceRequests: &pps.ResourceSpec{},
			Input: &pps.Input{
				Pfs: &pps.PFSInput{
					Repo:   dataRepo,
					Branch: "master",
					Glob:   "/*",
				},
			},
		})
	require.NoError(t, err)
	require.NoError(t, backoff.Retry(func() error {
		for i := 0; i < 3; i++ {
			pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, fmt.Sprintf("%s-%d", pipelineName, i), false)
			require.NoError(t, err)
			if pipelineInfo.State != pps.PipelineState_PIPELINE_RUNNING {
				return errors.Errorf("pipeline not in running state")
			}
		}
		return nil
	}, backoff.NewTestingBackOff()))
}

func TestPodOpts(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	t.Parallel()
	c, ns := minikubetestenv.AcquireCluster(t)
	// create repos
	dataRepo := tu.UniqueString("TestPodSpecOpts_data")
	require.NoError(t, c.CreateRepo(pfs.DefaultProjectName, dataRepo))
	t.Run("Validation", func(t *testing.T) {
		pipelineName := tu.UniqueString("TestPodSpecOpts")
		_, err := c.PpsAPIClient.CreatePipeline(
			context.Background(),
			&pps.CreatePipelineRequest{
				Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
				Transform: &pps.Transform{
					Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
				},
				Input: &pps.Input{
					Pfs: &pps.PFSInput{
						Repo:   dataRepo,
						Branch: "master",
						Glob:   "/*",
					},
				},
				PodSpec: "not-json",
			})
		require.YesError(t, err)
		_, err = c.PpsAPIClient.CreatePipeline(
			context.Background(),
			&pps.CreatePipelineRequest{
				Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
				Transform: &pps.Transform{
					Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
				},
				Input: &pps.Input{
					Pfs: &pps.PFSInput{
						Repo:   dataRepo,
						Branch: "master",
						Glob:   "/*",
					},
				},
				PodPatch: "also-not-json",
			})
		require.YesError(t, err)
	})
	t.Run("Spec", func(t *testing.T) {
		pipelineName := tu.UniqueString("TestPodSpecOpts")
		_, err := c.PpsAPIClient.CreatePipeline(
			context.Background(),
			&pps.CreatePipelineRequest{
				Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
				Transform: &pps.Transform{
					Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
				},
				ParallelismSpec: &pps.ParallelismSpec{
					Constant: 1,
				},
				Input: &pps.Input{
					Pfs: &pps.PFSInput{
						Repo:   dataRepo,
						Branch: "master",
						Glob:   "/*",
					},
				},
				SchedulingSpec: &pps.SchedulingSpec{
					// This NodeSelector will cause the worker pod to fail to
					// schedule, but the test can still pass because we just check
					// for values on the pod, it doesn't need to actually come up.
					NodeSelector: map[string]string{
						"foo": "bar",
					},
				},
				PodSpec: `{
				"hostname": "hostname"
			}`,
			})
		require.NoError(t, err)

		// Get info about the pipeline pods from k8s & check for resources
		pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, pipelineName, false)
		require.NoError(t, err)

		var pod v1.Pod
		kubeClient := tu.GetKubeClient(t)
		err = backoff.Retry(func() error {
			podList, err := kubeClient.CoreV1().Pods(ns).List(
				context.Background(),
				metav1.ListOptions{
					LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
						map[string]string{
							"app":             "pipeline",
							"pipelineName":    pipelineInfo.Pipeline.Name,
							"pipelineVersion": fmt.Sprint(pipelineInfo.Version),
							"suite":           "pachyderm"},
					)),
				})
			if err != nil {
				return errors.EnsureStack(err) // retry
			}
			if len(podList.Items) != 1 || len(podList.Items[0].Spec.Containers) == 0 {
				return errors.Errorf("could not find single container for pipeline %s", pipelineInfo.Pipeline.Name)
			}
			pod = podList.Items[0]
			return nil // no more retries
		}, backoff.NewTestingBackOff())
		require.NoError(t, err)
		// Make sure a CPU and Memory request are both set
		require.Equal(t, "bar", pod.Spec.NodeSelector["foo"])
		require.Equal(t, "hostname", pod.Spec.Hostname)
	})
	t.Run("Patch", func(t *testing.T) {
		pipelineName := tu.UniqueString("TestPodSpecOpts")
		_, err := c.PpsAPIClient.CreatePipeline(
			context.Background(),
			&pps.CreatePipelineRequest{
				Pipeline: client.NewPipeline(pfs.DefaultProjectName, pipelineName),
				Transform: &pps.Transform{
					Cmd: []string{"cp", path.Join("/pfs", dataRepo, "file"), "/pfs/out/file"},
				},
				ParallelismSpec: &pps.ParallelismSpec{
					Constant: 1,
				},
				Input: &pps.Input{
					Pfs: &pps.PFSInput{
						Repo:   dataRepo,
						Branch: "master",
						Glob:   "/*",
					},
				},
				SchedulingSpec: &pps.SchedulingSpec{
					// This NodeSelector will cause the worker pod to fail to
					// schedule, but the test can still pass because we just check
					// for values on the pod, it doesn't need to actually come up.
					NodeSelector: map[string]string{
						"foo": "bar",
					},
				},
				PodPatch: `[
					{ "op": "add", "path": "/hostname", "value": "hostname" }
			]`,
			})
		require.NoError(t, err)

		// Get info about the pipeline pods from k8s & check for resources
		pipelineInfo, err := c.InspectPipeline(pfs.DefaultProjectName, pipelineName, false)
		require.NoError(t, err)

		var pod v1.Pod
		kubeClient := tu.GetKubeClient(t)
		err = backoff.Retry(func() error {
			podList, err := kubeClient.CoreV1().Pods(ns).List(
				context.Background(),
				metav1.ListOptions{
					LabelSelector: metav1.FormatLabelSelector(metav1.SetAsLabelSelector(
						map[string]string{
							"app":             "pipeline",
							"pipelineName":    pipelineInfo.Pipeline.Name,
							"pipelineVersion": fmt.Sprint(pipelineInfo.Version),
							"suite":           "pachyderm"},
					)),
				})
			if err != nil {
				return errors.EnsureStack(err) // retry
			}
			if len(podList.Items) != 1 || len(podList.Items[0].Spec.Containers) == 0 {
				return errors.Errorf("could not find single container for pipeline %s", pipelineInfo.Pipeline.Name)
			}
			pod = podList.Items[0]
			return nil // no more retries
		}, backoff.NewTestingBackOff())
		require.NoError(t, err)
		// Make sure a CPU and Memory request are both set
		require.Equal(t, "bar", pod.Spec.NodeSelector["foo"])
		require.Equal(t, "hostname", pod.Spec.Hostname)
	})
}
