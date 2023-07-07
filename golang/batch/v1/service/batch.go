package service

import (
	"context"

	pb "github.com/star-table/interface/golang/batch/v1"
)

type BatchService struct {
	pb.UnimplementedBatchServer
}

func NewBatchService() *BatchService {
	return &BatchService{}
}

func (s *BatchService) Temp(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) ExportPGData(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) SplitTableHeader(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateRelating(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateWorkHour(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateCollaborator(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateDropUnusedTables(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) FixIssueStatusLostWhenMove(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) AddPGIndex(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateGlobalUser(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) FixPGOrgId(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigratePath(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateGroupChatSetting(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) FixIssueStatusName(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) FixIssueForDelPro(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) AddHasOverdueView(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) AddPgStatisticsTable(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) AddDashBord(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) ChangeDomain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) PgDataTransfer(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) ClearSameUser(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) MigrateAuditStatus(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) DbMigrations(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) Fix20221104TemplatePath(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) Fix20221107AppCreatorAddMember(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) AutomationRebuildActive(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *BatchService) ExchangeRelatingDataId(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
