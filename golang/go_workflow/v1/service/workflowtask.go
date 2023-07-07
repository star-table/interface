package service

import (
	"context"

	pb "github.com/star-table/interface/golang/go_workflow/v1"
)

type WorkFlowTaskService struct {
	pb.UnimplementedWorkFlowTaskServer
}

func NewWorkFlowTaskService() *WorkFlowTaskService {
	return &WorkFlowTaskService{}
}

func (s *WorkFlowTaskService) Complete(ctx context.Context, req *pb.CompleteRequest) (*pb.CompleteReply, error) {
	return &pb.CompleteReply{}, nil
}
func (s *WorkFlowTaskService) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawReply, error) {
	return &pb.WithdrawReply{}, nil
}
