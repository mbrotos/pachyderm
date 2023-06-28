// Code generated by pachgen (etc/proto/pachgen). DO NOT EDIT.

package client

import (
	"context"

	admin_v2 "github.com/pachyderm/pachyderm/v2/src/admin"
	auth_v2 "github.com/pachyderm/pachyderm/v2/src/auth"
	debug_v2 "github.com/pachyderm/pachyderm/v2/src/debug"
	enterprise_v2 "github.com/pachyderm/pachyderm/v2/src/enterprise"
	identity_v2 "github.com/pachyderm/pachyderm/v2/src/identity"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	license_v2 "github.com/pachyderm/pachyderm/v2/src/license"
	pfs_v2 "github.com/pachyderm/pachyderm/v2/src/pfs"
	pps_v2 "github.com/pachyderm/pachyderm/v2/src/pps"
	proxy "github.com/pachyderm/pachyderm/v2/src/proxy"
	taskapi "github.com/pachyderm/pachyderm/v2/src/task"
	transaction_v2 "github.com/pachyderm/pachyderm/v2/src/transaction"
	versionpb_v2 "github.com/pachyderm/pachyderm/v2/src/version/versionpb"

	types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

func unsupportedError(name string) error {
	return errors.Errorf("the '%s' API call is not supported in transactions", name)
}

type unsupportedAdminBuilderClient struct{}

func (c *unsupportedAdminBuilderClient) InspectCluster(_ context.Context, _ *admin_v2.InspectClusterRequest, opts ...grpc.CallOption) (*admin_v2.ClusterInfo, error) {
	return nil, unsupportedError("InspectCluster")
}

type unsupportedAuthBuilderClient struct{}

func (c *unsupportedAuthBuilderClient) Activate(_ context.Context, _ *auth_v2.ActivateRequest, opts ...grpc.CallOption) (*auth_v2.ActivateResponse, error) {
	return nil, unsupportedError("Activate")
}

func (c *unsupportedAuthBuilderClient) Authenticate(_ context.Context, _ *auth_v2.AuthenticateRequest, opts ...grpc.CallOption) (*auth_v2.AuthenticateResponse, error) {
	return nil, unsupportedError("Authenticate")
}

func (c *unsupportedAuthBuilderClient) Authorize(_ context.Context, _ *auth_v2.AuthorizeRequest, opts ...grpc.CallOption) (*auth_v2.AuthorizeResponse, error) {
	return nil, unsupportedError("Authorize")
}

func (c *unsupportedAuthBuilderClient) Deactivate(_ context.Context, _ *auth_v2.DeactivateRequest, opts ...grpc.CallOption) (*auth_v2.DeactivateResponse, error) {
	return nil, unsupportedError("Deactivate")
}

func (c *unsupportedAuthBuilderClient) DeleteExpiredAuthTokens(_ context.Context, _ *auth_v2.DeleteExpiredAuthTokensRequest, opts ...grpc.CallOption) (*auth_v2.DeleteExpiredAuthTokensResponse, error) {
	return nil, unsupportedError("DeleteExpiredAuthTokens")
}

func (c *unsupportedAuthBuilderClient) ExtractAuthTokens(_ context.Context, _ *auth_v2.ExtractAuthTokensRequest, opts ...grpc.CallOption) (*auth_v2.ExtractAuthTokensResponse, error) {
	return nil, unsupportedError("ExtractAuthTokens")
}

func (c *unsupportedAuthBuilderClient) GetConfiguration(_ context.Context, _ *auth_v2.GetConfigurationRequest, opts ...grpc.CallOption) (*auth_v2.GetConfigurationResponse, error) {
	return nil, unsupportedError("GetConfiguration")
}

func (c *unsupportedAuthBuilderClient) GetGroups(_ context.Context, _ *auth_v2.GetGroupsRequest, opts ...grpc.CallOption) (*auth_v2.GetGroupsResponse, error) {
	return nil, unsupportedError("GetGroups")
}

func (c *unsupportedAuthBuilderClient) GetGroupsForPrincipal(_ context.Context, _ *auth_v2.GetGroupsForPrincipalRequest, opts ...grpc.CallOption) (*auth_v2.GetGroupsResponse, error) {
	return nil, unsupportedError("GetGroupsForPrincipal")
}

func (c *unsupportedAuthBuilderClient) GetOIDCLogin(_ context.Context, _ *auth_v2.GetOIDCLoginRequest, opts ...grpc.CallOption) (*auth_v2.GetOIDCLoginResponse, error) {
	return nil, unsupportedError("GetOIDCLogin")
}

func (c *unsupportedAuthBuilderClient) GetPermissions(_ context.Context, _ *auth_v2.GetPermissionsRequest, opts ...grpc.CallOption) (*auth_v2.GetPermissionsResponse, error) {
	return nil, unsupportedError("GetPermissions")
}

func (c *unsupportedAuthBuilderClient) GetPermissionsForPrincipal(_ context.Context, _ *auth_v2.GetPermissionsForPrincipalRequest, opts ...grpc.CallOption) (*auth_v2.GetPermissionsResponse, error) {
	return nil, unsupportedError("GetPermissionsForPrincipal")
}

func (c *unsupportedAuthBuilderClient) GetRobotToken(_ context.Context, _ *auth_v2.GetRobotTokenRequest, opts ...grpc.CallOption) (*auth_v2.GetRobotTokenResponse, error) {
	return nil, unsupportedError("GetRobotToken")
}

func (c *unsupportedAuthBuilderClient) GetRoleBinding(_ context.Context, _ *auth_v2.GetRoleBindingRequest, opts ...grpc.CallOption) (*auth_v2.GetRoleBindingResponse, error) {
	return nil, unsupportedError("GetRoleBinding")
}

func (c *unsupportedAuthBuilderClient) GetRolesForPermission(_ context.Context, _ *auth_v2.GetRolesForPermissionRequest, opts ...grpc.CallOption) (*auth_v2.GetRolesForPermissionResponse, error) {
	return nil, unsupportedError("GetRolesForPermission")
}

func (c *unsupportedAuthBuilderClient) GetUsers(_ context.Context, _ *auth_v2.GetUsersRequest, opts ...grpc.CallOption) (*auth_v2.GetUsersResponse, error) {
	return nil, unsupportedError("GetUsers")
}

func (c *unsupportedAuthBuilderClient) ModifyMembers(_ context.Context, _ *auth_v2.ModifyMembersRequest, opts ...grpc.CallOption) (*auth_v2.ModifyMembersResponse, error) {
	return nil, unsupportedError("ModifyMembers")
}

func (c *unsupportedAuthBuilderClient) ModifyRoleBinding(_ context.Context, _ *auth_v2.ModifyRoleBindingRequest, opts ...grpc.CallOption) (*auth_v2.ModifyRoleBindingResponse, error) {
	return nil, unsupportedError("ModifyRoleBinding")
}

func (c *unsupportedAuthBuilderClient) RestoreAuthToken(_ context.Context, _ *auth_v2.RestoreAuthTokenRequest, opts ...grpc.CallOption) (*auth_v2.RestoreAuthTokenResponse, error) {
	return nil, unsupportedError("RestoreAuthToken")
}

func (c *unsupportedAuthBuilderClient) RevokeAuthToken(_ context.Context, _ *auth_v2.RevokeAuthTokenRequest, opts ...grpc.CallOption) (*auth_v2.RevokeAuthTokenResponse, error) {
	return nil, unsupportedError("RevokeAuthToken")
}

func (c *unsupportedAuthBuilderClient) RevokeAuthTokensForUser(_ context.Context, _ *auth_v2.RevokeAuthTokensForUserRequest, opts ...grpc.CallOption) (*auth_v2.RevokeAuthTokensForUserResponse, error) {
	return nil, unsupportedError("RevokeAuthTokensForUser")
}

func (c *unsupportedAuthBuilderClient) RotateRootToken(_ context.Context, _ *auth_v2.RotateRootTokenRequest, opts ...grpc.CallOption) (*auth_v2.RotateRootTokenResponse, error) {
	return nil, unsupportedError("RotateRootToken")
}

func (c *unsupportedAuthBuilderClient) SetConfiguration(_ context.Context, _ *auth_v2.SetConfigurationRequest, opts ...grpc.CallOption) (*auth_v2.SetConfigurationResponse, error) {
	return nil, unsupportedError("SetConfiguration")
}

func (c *unsupportedAuthBuilderClient) SetGroupsForUser(_ context.Context, _ *auth_v2.SetGroupsForUserRequest, opts ...grpc.CallOption) (*auth_v2.SetGroupsForUserResponse, error) {
	return nil, unsupportedError("SetGroupsForUser")
}

func (c *unsupportedAuthBuilderClient) WhoAmI(_ context.Context, _ *auth_v2.WhoAmIRequest, opts ...grpc.CallOption) (*auth_v2.WhoAmIResponse, error) {
	return nil, unsupportedError("WhoAmI")
}

type unsupportedDebugBuilderClient struct{}

func (c *unsupportedDebugBuilderClient) Binary(_ context.Context, _ *debug_v2.BinaryRequest, opts ...grpc.CallOption) (debug_v2.Debug_BinaryClient, error) {
	return nil, unsupportedError("Binary")
}

func (c *unsupportedDebugBuilderClient) Dump(_ context.Context, _ *debug_v2.DumpRequest, opts ...grpc.CallOption) (debug_v2.Debug_DumpClient, error) {
	return nil, unsupportedError("Dump")
}

func (c *unsupportedDebugBuilderClient) DumpV2(_ context.Context, _ *debug_v2.DumpV2Request, opts ...grpc.CallOption) (debug_v2.Debug_DumpV2Client, error) {
	return nil, unsupportedError("DumpV2")
}

func (c *unsupportedDebugBuilderClient) GetDumpV2Template(_ context.Context, _ *debug_v2.GetDumpV2TemplateRequest, opts ...grpc.CallOption) (*debug_v2.GetDumpV2TemplateResponse, error) {
	return nil, unsupportedError("GetDumpV2Template")
}

func (c *unsupportedDebugBuilderClient) Profile(_ context.Context, _ *debug_v2.ProfileRequest, opts ...grpc.CallOption) (debug_v2.Debug_ProfileClient, error) {
	return nil, unsupportedError("Profile")
}

func (c *unsupportedDebugBuilderClient) SetLogLevel(_ context.Context, _ *debug_v2.SetLogLevelRequest, opts ...grpc.CallOption) (*debug_v2.SetLogLevelResponse, error) {
	return nil, unsupportedError("SetLogLevel")
}

type unsupportedEnterpriseBuilderClient struct{}

func (c *unsupportedEnterpriseBuilderClient) Activate(_ context.Context, _ *enterprise_v2.ActivateRequest, opts ...grpc.CallOption) (*enterprise_v2.ActivateResponse, error) {
	return nil, unsupportedError("Activate")
}

func (c *unsupportedEnterpriseBuilderClient) Deactivate(_ context.Context, _ *enterprise_v2.DeactivateRequest, opts ...grpc.CallOption) (*enterprise_v2.DeactivateResponse, error) {
	return nil, unsupportedError("Deactivate")
}

func (c *unsupportedEnterpriseBuilderClient) GetActivationCode(_ context.Context, _ *enterprise_v2.GetActivationCodeRequest, opts ...grpc.CallOption) (*enterprise_v2.GetActivationCodeResponse, error) {
	return nil, unsupportedError("GetActivationCode")
}

func (c *unsupportedEnterpriseBuilderClient) GetState(_ context.Context, _ *enterprise_v2.GetStateRequest, opts ...grpc.CallOption) (*enterprise_v2.GetStateResponse, error) {
	return nil, unsupportedError("GetState")
}

func (c *unsupportedEnterpriseBuilderClient) Heartbeat(_ context.Context, _ *enterprise_v2.HeartbeatRequest, opts ...grpc.CallOption) (*enterprise_v2.HeartbeatResponse, error) {
	return nil, unsupportedError("Heartbeat")
}

func (c *unsupportedEnterpriseBuilderClient) Pause(_ context.Context, _ *enterprise_v2.PauseRequest, opts ...grpc.CallOption) (*enterprise_v2.PauseResponse, error) {
	return nil, unsupportedError("Pause")
}

func (c *unsupportedEnterpriseBuilderClient) PauseStatus(_ context.Context, _ *enterprise_v2.PauseStatusRequest, opts ...grpc.CallOption) (*enterprise_v2.PauseStatusResponse, error) {
	return nil, unsupportedError("PauseStatus")
}

func (c *unsupportedEnterpriseBuilderClient) Unpause(_ context.Context, _ *enterprise_v2.UnpauseRequest, opts ...grpc.CallOption) (*enterprise_v2.UnpauseResponse, error) {
	return nil, unsupportedError("Unpause")
}

type unsupportedIdentityBuilderClient struct{}

func (c *unsupportedIdentityBuilderClient) CreateIDPConnector(_ context.Context, _ *identity_v2.CreateIDPConnectorRequest, opts ...grpc.CallOption) (*identity_v2.CreateIDPConnectorResponse, error) {
	return nil, unsupportedError("CreateIDPConnector")
}

func (c *unsupportedIdentityBuilderClient) CreateOIDCClient(_ context.Context, _ *identity_v2.CreateOIDCClientRequest, opts ...grpc.CallOption) (*identity_v2.CreateOIDCClientResponse, error) {
	return nil, unsupportedError("CreateOIDCClient")
}

func (c *unsupportedIdentityBuilderClient) DeleteAll(_ context.Context, _ *identity_v2.DeleteAllRequest, opts ...grpc.CallOption) (*identity_v2.DeleteAllResponse, error) {
	return nil, unsupportedError("DeleteAll")
}

func (c *unsupportedIdentityBuilderClient) DeleteIDPConnector(_ context.Context, _ *identity_v2.DeleteIDPConnectorRequest, opts ...grpc.CallOption) (*identity_v2.DeleteIDPConnectorResponse, error) {
	return nil, unsupportedError("DeleteIDPConnector")
}

func (c *unsupportedIdentityBuilderClient) DeleteOIDCClient(_ context.Context, _ *identity_v2.DeleteOIDCClientRequest, opts ...grpc.CallOption) (*identity_v2.DeleteOIDCClientResponse, error) {
	return nil, unsupportedError("DeleteOIDCClient")
}

func (c *unsupportedIdentityBuilderClient) GetIDPConnector(_ context.Context, _ *identity_v2.GetIDPConnectorRequest, opts ...grpc.CallOption) (*identity_v2.GetIDPConnectorResponse, error) {
	return nil, unsupportedError("GetIDPConnector")
}

func (c *unsupportedIdentityBuilderClient) GetIdentityServerConfig(_ context.Context, _ *identity_v2.GetIdentityServerConfigRequest, opts ...grpc.CallOption) (*identity_v2.GetIdentityServerConfigResponse, error) {
	return nil, unsupportedError("GetIdentityServerConfig")
}

func (c *unsupportedIdentityBuilderClient) GetOIDCClient(_ context.Context, _ *identity_v2.GetOIDCClientRequest, opts ...grpc.CallOption) (*identity_v2.GetOIDCClientResponse, error) {
	return nil, unsupportedError("GetOIDCClient")
}

func (c *unsupportedIdentityBuilderClient) ListIDPConnectors(_ context.Context, _ *identity_v2.ListIDPConnectorsRequest, opts ...grpc.CallOption) (*identity_v2.ListIDPConnectorsResponse, error) {
	return nil, unsupportedError("ListIDPConnectors")
}

func (c *unsupportedIdentityBuilderClient) ListOIDCClients(_ context.Context, _ *identity_v2.ListOIDCClientsRequest, opts ...grpc.CallOption) (*identity_v2.ListOIDCClientsResponse, error) {
	return nil, unsupportedError("ListOIDCClients")
}

func (c *unsupportedIdentityBuilderClient) SetIdentityServerConfig(_ context.Context, _ *identity_v2.SetIdentityServerConfigRequest, opts ...grpc.CallOption) (*identity_v2.SetIdentityServerConfigResponse, error) {
	return nil, unsupportedError("SetIdentityServerConfig")
}

func (c *unsupportedIdentityBuilderClient) UpdateIDPConnector(_ context.Context, _ *identity_v2.UpdateIDPConnectorRequest, opts ...grpc.CallOption) (*identity_v2.UpdateIDPConnectorResponse, error) {
	return nil, unsupportedError("UpdateIDPConnector")
}

func (c *unsupportedIdentityBuilderClient) UpdateOIDCClient(_ context.Context, _ *identity_v2.UpdateOIDCClientRequest, opts ...grpc.CallOption) (*identity_v2.UpdateOIDCClientResponse, error) {
	return nil, unsupportedError("UpdateOIDCClient")
}

type unsupportedLicenseBuilderClient struct{}

func (c *unsupportedLicenseBuilderClient) Activate(_ context.Context, _ *license_v2.ActivateRequest, opts ...grpc.CallOption) (*license_v2.ActivateResponse, error) {
	return nil, unsupportedError("Activate")
}

func (c *unsupportedLicenseBuilderClient) AddCluster(_ context.Context, _ *license_v2.AddClusterRequest, opts ...grpc.CallOption) (*license_v2.AddClusterResponse, error) {
	return nil, unsupportedError("AddCluster")
}

func (c *unsupportedLicenseBuilderClient) DeleteAll(_ context.Context, _ *license_v2.DeleteAllRequest, opts ...grpc.CallOption) (*license_v2.DeleteAllResponse, error) {
	return nil, unsupportedError("DeleteAll")
}

func (c *unsupportedLicenseBuilderClient) DeleteCluster(_ context.Context, _ *license_v2.DeleteClusterRequest, opts ...grpc.CallOption) (*license_v2.DeleteClusterResponse, error) {
	return nil, unsupportedError("DeleteCluster")
}

func (c *unsupportedLicenseBuilderClient) GetActivationCode(_ context.Context, _ *license_v2.GetActivationCodeRequest, opts ...grpc.CallOption) (*license_v2.GetActivationCodeResponse, error) {
	return nil, unsupportedError("GetActivationCode")
}

func (c *unsupportedLicenseBuilderClient) Heartbeat(_ context.Context, _ *license_v2.HeartbeatRequest, opts ...grpc.CallOption) (*license_v2.HeartbeatResponse, error) {
	return nil, unsupportedError("Heartbeat")
}

func (c *unsupportedLicenseBuilderClient) ListClusters(_ context.Context, _ *license_v2.ListClustersRequest, opts ...grpc.CallOption) (*license_v2.ListClustersResponse, error) {
	return nil, unsupportedError("ListClusters")
}

func (c *unsupportedLicenseBuilderClient) ListUserClusters(_ context.Context, _ *license_v2.ListUserClustersRequest, opts ...grpc.CallOption) (*license_v2.ListUserClustersResponse, error) {
	return nil, unsupportedError("ListUserClusters")
}

func (c *unsupportedLicenseBuilderClient) UpdateCluster(_ context.Context, _ *license_v2.UpdateClusterRequest, opts ...grpc.CallOption) (*license_v2.UpdateClusterResponse, error) {
	return nil, unsupportedError("UpdateCluster")
}

type unsupportedPfsBuilderClient struct{}

func (c *unsupportedPfsBuilderClient) ActivateAuth(_ context.Context, _ *pfs_v2.ActivateAuthRequest, opts ...grpc.CallOption) (*pfs_v2.ActivateAuthResponse, error) {
	return nil, unsupportedError("ActivateAuth")
}

func (c *unsupportedPfsBuilderClient) AddFileSet(_ context.Context, _ *pfs_v2.AddFileSetRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("AddFileSet")
}

func (c *unsupportedPfsBuilderClient) CheckStorage(_ context.Context, _ *pfs_v2.CheckStorageRequest, opts ...grpc.CallOption) (*pfs_v2.CheckStorageResponse, error) {
	return nil, unsupportedError("CheckStorage")
}

func (c *unsupportedPfsBuilderClient) ClearCache(_ context.Context, _ *pfs_v2.ClearCacheRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("ClearCache")
}

func (c *unsupportedPfsBuilderClient) ClearCommit(_ context.Context, _ *pfs_v2.ClearCommitRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("ClearCommit")
}

func (c *unsupportedPfsBuilderClient) ComposeFileSet(_ context.Context, _ *pfs_v2.ComposeFileSetRequest, opts ...grpc.CallOption) (*pfs_v2.CreateFileSetResponse, error) {
	return nil, unsupportedError("ComposeFileSet")
}

func (c *unsupportedPfsBuilderClient) CreateBranch(_ context.Context, _ *pfs_v2.CreateBranchRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("CreateBranch")
}

func (c *unsupportedPfsBuilderClient) CreateFileSet(_ context.Context, opts ...grpc.CallOption) (pfs_v2.API_CreateFileSetClient, error) {
	return nil, unsupportedError("CreateFileSet")
}

func (c *unsupportedPfsBuilderClient) CreateProject(_ context.Context, _ *pfs_v2.CreateProjectRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("CreateProject")
}

func (c *unsupportedPfsBuilderClient) CreateRepo(_ context.Context, _ *pfs_v2.CreateRepoRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("CreateRepo")
}

func (c *unsupportedPfsBuilderClient) DeleteAll(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteAll")
}

func (c *unsupportedPfsBuilderClient) DeleteBranch(_ context.Context, _ *pfs_v2.DeleteBranchRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteBranch")
}

func (c *unsupportedPfsBuilderClient) DeleteProject(_ context.Context, _ *pfs_v2.DeleteProjectRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteProject")
}

func (c *unsupportedPfsBuilderClient) DeleteRepo(_ context.Context, _ *pfs_v2.DeleteRepoRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteRepo")
}

func (c *unsupportedPfsBuilderClient) DeleteRepos(_ context.Context, _ *pfs_v2.DeleteReposRequest, opts ...grpc.CallOption) (*pfs_v2.DeleteReposResponse, error) {
	return nil, unsupportedError("DeleteRepos")
}

func (c *unsupportedPfsBuilderClient) DiffFile(_ context.Context, _ *pfs_v2.DiffFileRequest, opts ...grpc.CallOption) (pfs_v2.API_DiffFileClient, error) {
	return nil, unsupportedError("DiffFile")
}

func (c *unsupportedPfsBuilderClient) DropCommitSet(_ context.Context, _ *pfs_v2.DropCommitSetRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DropCommitSet")
}

func (c *unsupportedPfsBuilderClient) Egress(_ context.Context, _ *pfs_v2.EgressRequest, opts ...grpc.CallOption) (*pfs_v2.EgressResponse, error) {
	return nil, unsupportedError("Egress")
}

func (c *unsupportedPfsBuilderClient) FindCommits(_ context.Context, _ *pfs_v2.FindCommitsRequest, opts ...grpc.CallOption) (pfs_v2.API_FindCommitsClient, error) {
	return nil, unsupportedError("FindCommits")
}

func (c *unsupportedPfsBuilderClient) FinishCommit(_ context.Context, _ *pfs_v2.FinishCommitRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("FinishCommit")
}

func (c *unsupportedPfsBuilderClient) Fsck(_ context.Context, _ *pfs_v2.FsckRequest, opts ...grpc.CallOption) (pfs_v2.API_FsckClient, error) {
	return nil, unsupportedError("Fsck")
}

func (c *unsupportedPfsBuilderClient) GetCache(_ context.Context, _ *pfs_v2.GetCacheRequest, opts ...grpc.CallOption) (*pfs_v2.GetCacheResponse, error) {
	return nil, unsupportedError("GetCache")
}

func (c *unsupportedPfsBuilderClient) GetFile(_ context.Context, _ *pfs_v2.GetFileRequest, opts ...grpc.CallOption) (pfs_v2.API_GetFileClient, error) {
	return nil, unsupportedError("GetFile")
}

func (c *unsupportedPfsBuilderClient) GetFileSet(_ context.Context, _ *pfs_v2.GetFileSetRequest, opts ...grpc.CallOption) (*pfs_v2.CreateFileSetResponse, error) {
	return nil, unsupportedError("GetFileSet")
}

func (c *unsupportedPfsBuilderClient) GetFileTAR(_ context.Context, _ *pfs_v2.GetFileRequest, opts ...grpc.CallOption) (pfs_v2.API_GetFileTARClient, error) {
	return nil, unsupportedError("GetFileTAR")
}

func (c *unsupportedPfsBuilderClient) GlobFile(_ context.Context, _ *pfs_v2.GlobFileRequest, opts ...grpc.CallOption) (pfs_v2.API_GlobFileClient, error) {
	return nil, unsupportedError("GlobFile")
}

func (c *unsupportedPfsBuilderClient) InspectBranch(_ context.Context, _ *pfs_v2.InspectBranchRequest, opts ...grpc.CallOption) (*pfs_v2.BranchInfo, error) {
	return nil, unsupportedError("InspectBranch")
}

func (c *unsupportedPfsBuilderClient) InspectCommit(_ context.Context, _ *pfs_v2.InspectCommitRequest, opts ...grpc.CallOption) (*pfs_v2.CommitInfo, error) {
	return nil, unsupportedError("InspectCommit")
}

func (c *unsupportedPfsBuilderClient) InspectCommitSet(_ context.Context, _ *pfs_v2.InspectCommitSetRequest, opts ...grpc.CallOption) (pfs_v2.API_InspectCommitSetClient, error) {
	return nil, unsupportedError("InspectCommitSet")
}

func (c *unsupportedPfsBuilderClient) InspectFile(_ context.Context, _ *pfs_v2.InspectFileRequest, opts ...grpc.CallOption) (*pfs_v2.FileInfo, error) {
	return nil, unsupportedError("InspectFile")
}

func (c *unsupportedPfsBuilderClient) InspectProject(_ context.Context, _ *pfs_v2.InspectProjectRequest, opts ...grpc.CallOption) (*pfs_v2.ProjectInfo, error) {
	return nil, unsupportedError("InspectProject")
}

func (c *unsupportedPfsBuilderClient) InspectRepo(_ context.Context, _ *pfs_v2.InspectRepoRequest, opts ...grpc.CallOption) (*pfs_v2.RepoInfo, error) {
	return nil, unsupportedError("InspectRepo")
}

func (c *unsupportedPfsBuilderClient) ListBranch(_ context.Context, _ *pfs_v2.ListBranchRequest, opts ...grpc.CallOption) (pfs_v2.API_ListBranchClient, error) {
	return nil, unsupportedError("ListBranch")
}

func (c *unsupportedPfsBuilderClient) ListCommit(_ context.Context, _ *pfs_v2.ListCommitRequest, opts ...grpc.CallOption) (pfs_v2.API_ListCommitClient, error) {
	return nil, unsupportedError("ListCommit")
}

func (c *unsupportedPfsBuilderClient) ListCommitSet(_ context.Context, _ *pfs_v2.ListCommitSetRequest, opts ...grpc.CallOption) (pfs_v2.API_ListCommitSetClient, error) {
	return nil, unsupportedError("ListCommitSet")
}

func (c *unsupportedPfsBuilderClient) ListFile(_ context.Context, _ *pfs_v2.ListFileRequest, opts ...grpc.CallOption) (pfs_v2.API_ListFileClient, error) {
	return nil, unsupportedError("ListFile")
}

func (c *unsupportedPfsBuilderClient) ListProject(_ context.Context, _ *pfs_v2.ListProjectRequest, opts ...grpc.CallOption) (pfs_v2.API_ListProjectClient, error) {
	return nil, unsupportedError("ListProject")
}

func (c *unsupportedPfsBuilderClient) ListRepo(_ context.Context, _ *pfs_v2.ListRepoRequest, opts ...grpc.CallOption) (pfs_v2.API_ListRepoClient, error) {
	return nil, unsupportedError("ListRepo")
}

func (c *unsupportedPfsBuilderClient) ListTask(_ context.Context, _ *taskapi.ListTaskRequest, opts ...grpc.CallOption) (pfs_v2.API_ListTaskClient, error) {
	return nil, unsupportedError("ListTask")
}

func (c *unsupportedPfsBuilderClient) ModifyFile(_ context.Context, opts ...grpc.CallOption) (pfs_v2.API_ModifyFileClient, error) {
	return nil, unsupportedError("ModifyFile")
}

func (c *unsupportedPfsBuilderClient) PutCache(_ context.Context, _ *pfs_v2.PutCacheRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("PutCache")
}

func (c *unsupportedPfsBuilderClient) RenewFileSet(_ context.Context, _ *pfs_v2.RenewFileSetRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("RenewFileSet")
}

func (c *unsupportedPfsBuilderClient) RunLoadTest(_ context.Context, _ *pfs_v2.RunLoadTestRequest, opts ...grpc.CallOption) (*pfs_v2.RunLoadTestResponse, error) {
	return nil, unsupportedError("RunLoadTest")
}

func (c *unsupportedPfsBuilderClient) RunLoadTestDefault(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*pfs_v2.RunLoadTestResponse, error) {
	return nil, unsupportedError("RunLoadTestDefault")
}

func (c *unsupportedPfsBuilderClient) ShardFileSet(_ context.Context, _ *pfs_v2.ShardFileSetRequest, opts ...grpc.CallOption) (*pfs_v2.ShardFileSetResponse, error) {
	return nil, unsupportedError("ShardFileSet")
}

func (c *unsupportedPfsBuilderClient) SquashCommitSet(_ context.Context, _ *pfs_v2.SquashCommitSetRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("SquashCommitSet")
}

func (c *unsupportedPfsBuilderClient) StartCommit(_ context.Context, _ *pfs_v2.StartCommitRequest, opts ...grpc.CallOption) (*pfs_v2.Commit, error) {
	return nil, unsupportedError("StartCommit")
}

func (c *unsupportedPfsBuilderClient) SubscribeCommit(_ context.Context, _ *pfs_v2.SubscribeCommitRequest, opts ...grpc.CallOption) (pfs_v2.API_SubscribeCommitClient, error) {
	return nil, unsupportedError("SubscribeCommit")
}

func (c *unsupportedPfsBuilderClient) WalkFile(_ context.Context, _ *pfs_v2.WalkFileRequest, opts ...grpc.CallOption) (pfs_v2.API_WalkFileClient, error) {
	return nil, unsupportedError("WalkFile")
}

type unsupportedPpsBuilderClient struct{}

func (c *unsupportedPpsBuilderClient) ActivateAuth(_ context.Context, _ *pps_v2.ActivateAuthRequest, opts ...grpc.CallOption) (*pps_v2.ActivateAuthResponse, error) {
	return nil, unsupportedError("ActivateAuth")
}

func (c *unsupportedPpsBuilderClient) CreatePipeline(_ context.Context, _ *pps_v2.CreatePipelineRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("CreatePipeline")
}

func (c *unsupportedPpsBuilderClient) CreateSecret(_ context.Context, _ *pps_v2.CreateSecretRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("CreateSecret")
}

func (c *unsupportedPpsBuilderClient) DeleteAll(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteAll")
}

func (c *unsupportedPpsBuilderClient) DeleteJob(_ context.Context, _ *pps_v2.DeleteJobRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteJob")
}

func (c *unsupportedPpsBuilderClient) DeletePipeline(_ context.Context, _ *pps_v2.DeletePipelineRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeletePipeline")
}

func (c *unsupportedPpsBuilderClient) DeletePipelines(_ context.Context, _ *pps_v2.DeletePipelinesRequest, opts ...grpc.CallOption) (*pps_v2.DeletePipelinesResponse, error) {
	return nil, unsupportedError("DeletePipelines")
}

func (c *unsupportedPpsBuilderClient) DeleteSecret(_ context.Context, _ *pps_v2.DeleteSecretRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteSecret")
}

func (c *unsupportedPpsBuilderClient) GetClusterDefaults(_ context.Context, _ *pps_v2.GetClusterDefaultsRequest, opts ...grpc.CallOption) (*pps_v2.GetClusterDefaultsResponse, error) {
	return nil, unsupportedError("GetClusterDefaults")
}

func (c *unsupportedPpsBuilderClient) GetKubeEvents(_ context.Context, _ *pps_v2.LokiRequest, opts ...grpc.CallOption) (pps_v2.API_GetKubeEventsClient, error) {
	return nil, unsupportedError("GetKubeEvents")
}

func (c *unsupportedPpsBuilderClient) GetLogs(_ context.Context, _ *pps_v2.GetLogsRequest, opts ...grpc.CallOption) (pps_v2.API_GetLogsClient, error) {
	return nil, unsupportedError("GetLogs")
}

func (c *unsupportedPpsBuilderClient) InspectDatum(_ context.Context, _ *pps_v2.InspectDatumRequest, opts ...grpc.CallOption) (*pps_v2.DatumInfo, error) {
	return nil, unsupportedError("InspectDatum")
}

func (c *unsupportedPpsBuilderClient) InspectJob(_ context.Context, _ *pps_v2.InspectJobRequest, opts ...grpc.CallOption) (*pps_v2.JobInfo, error) {
	return nil, unsupportedError("InspectJob")
}

func (c *unsupportedPpsBuilderClient) InspectJobSet(_ context.Context, _ *pps_v2.InspectJobSetRequest, opts ...grpc.CallOption) (pps_v2.API_InspectJobSetClient, error) {
	return nil, unsupportedError("InspectJobSet")
}

func (c *unsupportedPpsBuilderClient) InspectPipeline(_ context.Context, _ *pps_v2.InspectPipelineRequest, opts ...grpc.CallOption) (*pps_v2.PipelineInfo, error) {
	return nil, unsupportedError("InspectPipeline")
}

func (c *unsupportedPpsBuilderClient) InspectSecret(_ context.Context, _ *pps_v2.InspectSecretRequest, opts ...grpc.CallOption) (*pps_v2.SecretInfo, error) {
	return nil, unsupportedError("InspectSecret")
}

func (c *unsupportedPpsBuilderClient) ListDatum(_ context.Context, _ *pps_v2.ListDatumRequest, opts ...grpc.CallOption) (pps_v2.API_ListDatumClient, error) {
	return nil, unsupportedError("ListDatum")
}

func (c *unsupportedPpsBuilderClient) ListJob(_ context.Context, _ *pps_v2.ListJobRequest, opts ...grpc.CallOption) (pps_v2.API_ListJobClient, error) {
	return nil, unsupportedError("ListJob")
}

func (c *unsupportedPpsBuilderClient) ListJobSet(_ context.Context, _ *pps_v2.ListJobSetRequest, opts ...grpc.CallOption) (pps_v2.API_ListJobSetClient, error) {
	return nil, unsupportedError("ListJobSet")
}

func (c *unsupportedPpsBuilderClient) ListPipeline(_ context.Context, _ *pps_v2.ListPipelineRequest, opts ...grpc.CallOption) (pps_v2.API_ListPipelineClient, error) {
	return nil, unsupportedError("ListPipeline")
}

func (c *unsupportedPpsBuilderClient) ListSecret(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*pps_v2.SecretInfos, error) {
	return nil, unsupportedError("ListSecret")
}

func (c *unsupportedPpsBuilderClient) ListTask(_ context.Context, _ *taskapi.ListTaskRequest, opts ...grpc.CallOption) (pps_v2.API_ListTaskClient, error) {
	return nil, unsupportedError("ListTask")
}

func (c *unsupportedPpsBuilderClient) QueryLoki(_ context.Context, _ *pps_v2.LokiRequest, opts ...grpc.CallOption) (pps_v2.API_QueryLokiClient, error) {
	return nil, unsupportedError("QueryLoki")
}

func (c *unsupportedPpsBuilderClient) RenderTemplate(_ context.Context, _ *pps_v2.RenderTemplateRequest, opts ...grpc.CallOption) (*pps_v2.RenderTemplateResponse, error) {
	return nil, unsupportedError("RenderTemplate")
}

func (c *unsupportedPpsBuilderClient) RestartDatum(_ context.Context, _ *pps_v2.RestartDatumRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("RestartDatum")
}

func (c *unsupportedPpsBuilderClient) RunCron(_ context.Context, _ *pps_v2.RunCronRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("RunCron")
}

func (c *unsupportedPpsBuilderClient) RunLoadTest(_ context.Context, _ *pps_v2.RunLoadTestRequest, opts ...grpc.CallOption) (*pps_v2.RunLoadTestResponse, error) {
	return nil, unsupportedError("RunLoadTest")
}

func (c *unsupportedPpsBuilderClient) RunLoadTestDefault(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*pps_v2.RunLoadTestResponse, error) {
	return nil, unsupportedError("RunLoadTestDefault")
}

func (c *unsupportedPpsBuilderClient) RunPipeline(_ context.Context, _ *pps_v2.RunPipelineRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("RunPipeline")
}

func (c *unsupportedPpsBuilderClient) SetClusterDefaults(_ context.Context, _ *pps_v2.SetClusterDefaultsRequest, opts ...grpc.CallOption) (*pps_v2.SetClusterDefaultsResponse, error) {
	return nil, unsupportedError("SetClusterDefaults")
}

func (c *unsupportedPpsBuilderClient) StartPipeline(_ context.Context, _ *pps_v2.StartPipelineRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("StartPipeline")
}

func (c *unsupportedPpsBuilderClient) StopJob(_ context.Context, _ *pps_v2.StopJobRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("StopJob")
}

func (c *unsupportedPpsBuilderClient) StopPipeline(_ context.Context, _ *pps_v2.StopPipelineRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("StopPipeline")
}

func (c *unsupportedPpsBuilderClient) SubscribeJob(_ context.Context, _ *pps_v2.SubscribeJobRequest, opts ...grpc.CallOption) (pps_v2.API_SubscribeJobClient, error) {
	return nil, unsupportedError("SubscribeJob")
}

func (c *unsupportedPpsBuilderClient) UpdateJobState(_ context.Context, _ *pps_v2.UpdateJobStateRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("UpdateJobState")
}

type unsupportedProxyBuilderClient struct{}

func (c *unsupportedProxyBuilderClient) Listen(_ context.Context, _ *proxy.ListenRequest, opts ...grpc.CallOption) (proxy.API_ListenClient, error) {
	return nil, unsupportedError("Listen")
}

type unsupportedTransactionBuilderClient struct{}

func (c *unsupportedTransactionBuilderClient) BatchTransaction(_ context.Context, _ *transaction_v2.BatchTransactionRequest, opts ...grpc.CallOption) (*transaction_v2.TransactionInfo, error) {
	return nil, unsupportedError("BatchTransaction")
}

func (c *unsupportedTransactionBuilderClient) DeleteAll(_ context.Context, _ *transaction_v2.DeleteAllRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteAll")
}

func (c *unsupportedTransactionBuilderClient) DeleteTransaction(_ context.Context, _ *transaction_v2.DeleteTransactionRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	return nil, unsupportedError("DeleteTransaction")
}

func (c *unsupportedTransactionBuilderClient) FinishTransaction(_ context.Context, _ *transaction_v2.FinishTransactionRequest, opts ...grpc.CallOption) (*transaction_v2.TransactionInfo, error) {
	return nil, unsupportedError("FinishTransaction")
}

func (c *unsupportedTransactionBuilderClient) InspectTransaction(_ context.Context, _ *transaction_v2.InspectTransactionRequest, opts ...grpc.CallOption) (*transaction_v2.TransactionInfo, error) {
	return nil, unsupportedError("InspectTransaction")
}

func (c *unsupportedTransactionBuilderClient) ListTransaction(_ context.Context, _ *transaction_v2.ListTransactionRequest, opts ...grpc.CallOption) (*transaction_v2.TransactionInfos, error) {
	return nil, unsupportedError("ListTransaction")
}

func (c *unsupportedTransactionBuilderClient) StartTransaction(_ context.Context, _ *transaction_v2.StartTransactionRequest, opts ...grpc.CallOption) (*transaction_v2.Transaction, error) {
	return nil, unsupportedError("StartTransaction")
}

type unsupportedVersionpbBuilderClient struct{}

func (c *unsupportedVersionpbBuilderClient) GetVersion(_ context.Context, _ *types.Empty, opts ...grpc.CallOption) (*versionpb_v2.Version, error) {
	return nil, unsupportedError("GetVersion")
}
