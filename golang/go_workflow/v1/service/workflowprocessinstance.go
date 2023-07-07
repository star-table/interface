package service

import (
	"context"

	pb "github.com/star-table/interface/golang/go_workflow/v1"
)

type WorkFlowProcessInstanceService struct {
	pb.UnimplementedWorkFlowProcessInstanceServer
}

func NewWorkFlowProcessInstanceService() *WorkFlowProcessInstanceService {
	return &WorkFlowProcessInstanceService{}
}

func (s *WorkFlowProcessInstanceService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartReply, error) {
	return &pb.StartReply{}, nil
}
func (s *WorkFlowProcessInstanceService) StartByConditions(ctx context.Context, req *pb.StartByConditionsRequest) (*pb.StartByConditionsReply, error) {
	return &pb.StartByConditionsReply{}, nil
}
func (s *WorkFlowProcessInstanceService) QueryProcessInstances(ctx context.Context, req *pb.QueryProcessInstancesRequest) (*pb.QueryProcessInstancesReply, error) {
	return &pb.QueryProcessInstancesReply{}, nil
}
func (s *WorkFlowProcessInstanceService) DeleteProcessInstance(ctx context.Context, req *pb.DeleteProcessInstanceRequest) (*pb.DeleteProcessInstanceReply, error) {
	return &pb.DeleteProcessInstanceReply{}, nil
}
