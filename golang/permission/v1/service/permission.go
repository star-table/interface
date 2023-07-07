package service

import (
	"context"

	pb "github.com/star-table/interface/golang/permission/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PermissionService struct {
	pb.UnimplementedPermissionServer
}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

func (s *PermissionService) getCollaboratorUserIdsAndDeptIds(ctx context.Context, req *pb.GetCollaboratorUserIdsAndDeptIdsRequest) (*pb.GetCollaboratorUserIdsAndDeptIdsResponse, error) {
	return &pb.GetCollaboratorUserIdsAndDeptIdsResponse{}, nil
}
func (s *PermissionService) getCollaborators(ctx context.Context, req *pb.GetCollaboratorsRequest) (*pb.GetCollaboratorsResponse, error) {
	return &pb.GetCollaboratorsResponse{}, nil
}
func (s *PermissionService) getUserCollaboratorRoleIds(ctx context.Context, req *pb.GetUserCollaboratorRoleIdsRequest) (*pb.GetUserCollaboratorRoleIdsResponse, error) {
	return &pb.GetUserCollaboratorRoleIdsResponse{}, nil
}
func (s *PermissionService) changeCollaboratorUser(ctx context.Context, req *pb.CollaboratorUserRequest) (*pb.CollaboratorUserResponse, error) {
	return &pb.CollaboratorUserResponse{}, nil
}
func (s *PermissionService) addCollaboratorUser(ctx context.Context, req *pb.CollaboratorUserRequest) (*pb.CollaboratorUserResponse, error) {
	return &pb.CollaboratorUserResponse{}, nil
}
func (s *PermissionService) removeCollaboratorUser(ctx context.Context, req *pb.CollaboratorUserRequest) (*pb.CollaboratorUsersResponse, error) {
	return &pb.CollaboratorUsersResponse{}, nil
}
func (s *PermissionService) syncCollaboratorUsers(ctx context.Context, req *pb.CollaboratorUsersRequest) (*pb.CollaboratorUsersResponse, error) {
	return &pb.CollaboratorUsersResponse{}, nil
}
func (s *PermissionService) batchRemoveCollaboratorUsers(ctx context.Context, req *pb.CollaboratorUsersRequest) (*pb.CollaboratorUsersResponse, error) {
	return &pb.CollaboratorUsersResponse{}, nil
}
func (s *PermissionService) batchRemoveMemberFieldCollaborators(ctx context.Context, req *pb.MemberFieldRequest) (*pb.MemberFieldResponse, error) {
	return &pb.MemberFieldResponse{}, nil
}
func (s *PermissionService) batchMigrateCollaboratorUsersToPermission(ctx context.Context, req *pb.BatchMigrateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
func (s *PermissionService) copyCollaboratorsFromColumnToOther(ctx context.Context, req *pb.CopyColumnRequest) (*pb.CopyColumnResponse, error) {
	return &pb.CopyColumnResponse{}, nil
}
func (s *PermissionService) GetDataCollaborators(ctx context.Context, req *pb.GetDataCollaboratorsRequest) (*pb.GetDataCollaboratorsResponse, error) {
	return &pb.GetDataCollaboratorsResponse{}, nil
}
func (s *PermissionService) GetAppIdsByCollaborators(ctx context.Context, req *pb.GetAppIdsByCollaboratorReq) (*pb.GetAppIdsByCollaboratorResp, error) {
	return &pb.GetAppIdsByCollaboratorResp{}, nil
}
