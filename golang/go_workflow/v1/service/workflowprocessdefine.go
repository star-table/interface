package service

import (
	"context"

	pb "github.com/star-table/interface/golang/go_workflow/v1"
)

type WorkFlowProcessDefineService struct {
	pb.UnimplementedWorkFlowProcessDefineServer
}

func NewWorkFlowProcessDefineService() *WorkFlowProcessDefineService {
	return &WorkFlowProcessDefineService{}
}

func (s *WorkFlowProcessDefineService) SaveProcessDefine(ctx context.Context, req *pb.SaveProcessDefineRequest) (*pb.SaveProcessDefineReply, error) {
	return &pb.SaveProcessDefineReply{}, nil
}
func (s *WorkFlowProcessDefineService) QueryProcessDefines(ctx context.Context, req *pb.QueryProcessDefinesRequest) (*pb.QueryProcessDefinesReply, error) {
	return &pb.QueryProcessDefinesReply{}, nil
}
func (s *WorkFlowProcessDefineService) DeleteProcessDefine(ctx context.Context, req *pb.DeleteProcessDefineRequest) (*pb.DeleteProcessDefineReply, error) {
	return &pb.DeleteProcessDefineReply{}, nil
}
