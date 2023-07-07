package service

import (
	"context"

	pb "github.com/star-table/interface/golang/automation/v1"
)

type AutomationInnerService struct {
	pb.UnimplementedAutomationInnerServer
}

func NewAutomationInnerService() *AutomationInnerService {
	return &AutomationInnerService{}
}

func (s *AutomationInnerService) ApplyTemplate(ctx context.Context, req *pb.ApplyTemplateReq) (*pb.ApplyTemplateReply, error) {
	return &pb.ApplyTemplateReply{}, nil
}
func (s *AutomationInnerService) HandleWorkflowExecute(ctx context.Context, req *pb.HandleWorkflowExecuteReq) (*pb.HandleWorkflowExecuteReply, error) {
	return &pb.HandleWorkflowExecuteReply{}, nil
}
func (s *AutomationInnerService) HandleWorkflowError(ctx context.Context, req *pb.HandleWorkflowErrorReq) (*pb.HandleWorkflowErrorReply, error) {
	return &pb.HandleWorkflowErrorReply{}, nil
}
func (s *AutomationInnerService) BatchRebuildActive(ctx context.Context, req *pb.BatchRebuildActiveReq) (*pb.BatchRebuildActiveReply, error) {
	return &pb.BatchRebuildActiveReply{}, nil
}
func (s *AutomationInnerService) GlobalSwitchOff(ctx context.Context, req *pb.GlobalSwitchOffReq) (*pb.GlobalSwitchOffReply, error) {
	return &pb.GlobalSwitchOffReply{}, nil
}
func (s *AutomationInnerService) CreateTodo(ctx context.Context, req *pb.CreateTodoReq) (*pb.CreateTodoReply, error) {
	return &pb.CreateTodoReply{}, nil
}
func (s *AutomationInnerService) UpdateTodo(ctx context.Context, req *pb.UpdateTodoReq) (*pb.UpdateTodoReply, error) {
	return &pb.UpdateTodoReply{}, nil
}
