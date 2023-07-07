package service

import (
	"context"

	pb "github.com/star-table/interface/golang/table/v1"
)

type RowsService struct {
	pb.UnimplementedRowsServer
}

func NewRowsService() *RowsService {
	return &RowsService{}
}

func (s *RowsService) RecycleAttachment(ctx context.Context, req *pb.RecycleAttachmentRequest) (*pb.RecycleAttachmentReply, error) {
	return &pb.RecycleAttachmentReply{}, nil
}
func (s *RowsService) RecoverAttachment(ctx context.Context, req *pb.RecoverAttachmentRequest) (*pb.RecoverAttachmentReply, error) {
	return &pb.RecoverAttachmentReply{}, nil
}
func (s *RowsService) DeleteValues(ctx context.Context, req *pb.DeleteValuesRequest) (*pb.DeleteValuesReply, error) {
	return &pb.DeleteValuesReply{}, nil
}
func (s *RowsService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	return &pb.ListReply{}, nil
}
func (s *RowsService) ListRaw(ctx context.Context, req *pb.ListRawRequest) (*pb.ListRawReply, error) {
	return &pb.ListRawReply{}, nil
}
func (s *RowsService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *RowsService) CheckIsAppCollaborator(ctx context.Context, req *pb.CheckIsAppCollaboratorRequest) (*pb.CheckIsAppCollaboratorReply, error) {
	return &pb.CheckIsAppCollaboratorReply{}, nil
}
func (s *RowsService) GetUserAppCollaboratorRoles(ctx context.Context, req *pb.GetUserAppCollaboratorRolesRequest) (*pb.GetUserAppCollaboratorRolesReply, error) {
	return &pb.GetUserAppCollaboratorRolesReply{}, nil
}
func (s *RowsService) GetAppCollaboratorRoles(ctx context.Context, req *pb.GetAppCollaboratorRolesRequest) (*pb.GetAppCollaboratorRolesReply, error) {
	return &pb.GetAppCollaboratorRolesReply{}, nil
}
func (s *RowsService) GetDataCollaborators(ctx context.Context, req *pb.GetDataCollaboratorsRequest) (*pb.GetDataCollaboratorsReply, error) {
	return &pb.GetDataCollaboratorsReply{}, nil
}
func (s *RowsService) ExchangeSummaryCondition(ctx context.Context, req *pb.ExchangeSummaryConditionRequest) (*pb.ExchangeSummaryConditionReply, error) {
	return &pb.ExchangeSummaryConditionReply{}, nil
}
