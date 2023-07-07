package service

import (
	"context"

	pb "github.com/star-table/interface/golang/automation/v1"
)

type AutomationService struct {
	pb.UnimplementedAutomationServer
}

func NewAutomationService() *AutomationService {
	return &AutomationService{}
}

func (s *AutomationService) GetWorkflowsByApp(ctx context.Context, req *pb.GetWorkflowsByAppReq) (*pb.GetWorkflowsByAppReply, error) {
	return &pb.GetWorkflowsByAppReply{}, nil
}
func (s *AutomationService) GetWorkflow(ctx context.Context, req *pb.GetWorkflowReq) (*pb.GetWorkflowReply, error) {
	return &pb.GetWorkflowReply{}, nil
}
func (s *AutomationService) SaveWorkflowMeta(ctx context.Context, req *pb.SaveWorkflowMetaReq) (*pb.SaveWorkflowMetaReply, error) {
	return &pb.SaveWorkflowMetaReply{}, nil
}
func (s *AutomationService) SaveWorkflow(ctx context.Context, req *pb.SaveWorkflowReq) (*pb.SaveWorkflowReply, error) {
	return &pb.SaveWorkflowReply{}, nil
}
func (s *AutomationService) DeleteWorkflow(ctx context.Context, req *pb.DeleteWorkflowReq) (*pb.DeleteWorkflowReply, error) {
	return &pb.DeleteWorkflowReply{}, nil
}
func (s *AutomationService) ActivateWorkflow(ctx context.Context, req *pb.ActivateWorkflowReq) (*pb.ActivateWorkflowReply, error) {
	return &pb.ActivateWorkflowReply{}, nil
}
func (s *AutomationService) DeactivateWorkflow(ctx context.Context, req *pb.DeactivateWorkflowReq) (*pb.DeactivateWorkflowReply, error) {
	return &pb.DeactivateWorkflowReply{}, nil
}
func (s *AutomationService) GetExecutions(ctx context.Context, req *pb.GetExecutionsReq) (*pb.GetExecutionsReply, error) {
	return &pb.GetExecutionsReply{}, nil
}
func (s *AutomationService) GetExecution(ctx context.Context, req *pb.GetExecutionReq) (*pb.GetExecutionReply, error) {
	return &pb.GetExecutionReply{}, nil
}
func (s *AutomationService) CallWebhook(ctx context.Context, req *pb.CallWebhookReq) (*pb.CallWebhookReply, error) {
	return &pb.CallWebhookReply{}, nil
}
func (s *AutomationService) SaveIntegration(ctx context.Context, req *pb.SaveIntegrationReq) (*pb.SaveIntegrationReply, error) {
	return &pb.SaveIntegrationReply{}, nil
}
func (s *AutomationService) GetIntegrations(ctx context.Context, req *pb.GetIntegrationsReq) (*pb.GetIntegrationsReply, error) {
	return &pb.GetIntegrationsReply{}, nil
}
func (s *AutomationService) GetIntegrationsByType(ctx context.Context, req *pb.GetIntegrationsByTypeReq) (*pb.GetIntegrationsByTypeReply, error) {
	return &pb.GetIntegrationsByTypeReply{}, nil
}
func (s *AutomationService) GetIntegration(ctx context.Context, req *pb.GetIntegrationReq) (*pb.GetIntegrationReply, error) {
	return &pb.GetIntegrationReply{}, nil
}
func (s *AutomationService) DeleteIntegration(ctx context.Context, req *pb.DeleteIntegrationReq) (*pb.DeleteIntegrationReply, error) {
	return &pb.DeleteIntegrationReply{}, nil
}
func (s *AutomationService) GetTodosByExecution(ctx context.Context, req *pb.GetTodosByExecutionReq) (*pb.GetTodosByExecutionReply, error) {
	return &pb.GetTodosByExecutionReply{}, nil
}
func (s *AutomationService) TodoStat(ctx context.Context, req *pb.TodoStatReq) (*pb.TodoStatReply, error) {
	return &pb.TodoStatReply{}, nil
}
func (s *AutomationService) TodoFilter(ctx context.Context, req *pb.TodoFilterReq) (*pb.TodoFilterReply, error) {
	return &pb.TodoFilterReply{}, nil
}
func (s *AutomationService) GetTodo(ctx context.Context, req *pb.GetTodoReq) (*pb.GetTodoReply, error) {
	return &pb.GetTodoReply{}, nil
}
