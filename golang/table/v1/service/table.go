package service

import (
	"context"

	pb "github.com/star-table/interface/golang/table/v1"
)

type TableService struct {
	pb.UnimplementedTableServer
}

func NewTableService() *TableService {
	return &TableService{}
}

func (s *TableService) CreateSummeryTable(ctx context.Context, req *pb.CreateSummeryTableRequest) (*pb.CreateSummeryTableReply, error) {
	return &pb.CreateSummeryTableReply{}, nil
}
func (s *TableService) CreateMultiSummeryTable(ctx context.Context, req *pb.CreateMultiSummeryTableRequest) (*pb.CreateMultiSummeryTableReply, error) {
	return &pb.CreateMultiSummeryTableReply{}, nil
}
func (s *TableService) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.CreateTableReply, error) {
	return &pb.CreateTableReply{}, nil
}
func (s *TableService) RenameTable(ctx context.Context, req *pb.RenameTableRequest) (*pb.RenameTableReply, error) {
	return &pb.RenameTableReply{}, nil
}
func (s *TableService) CopyTables(ctx context.Context, req *pb.CopyTablesRequest) (*pb.CopyTablesReply, error) {
	return &pb.CopyTablesReply{}, nil
}
func (s *TableService) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*pb.DeleteTableReply, error) {
	return &pb.DeleteTableReply{}, nil
}
func (s *TableService) SetAutoSchedule(ctx context.Context, req *pb.SetAutoScheduleRequest) (*pb.SetAutoScheduleReply, error) {
	return &pb.SetAutoScheduleReply{}, nil
}
func (s *TableService) ReadTables(ctx context.Context, req *pb.ReadTablesRequest) (*pb.ReadTablesReply, error) {
	return &pb.ReadTablesReply{}, nil
}
func (s *TableService) ReadTable(ctx context.Context, req *pb.ReadTableRequest) (*pb.ReadTableReply, error) {
	return &pb.ReadTableReply{}, nil
}
func (s *TableService) ReadTablesByApps(ctx context.Context, req *pb.ReadTablesByAppsRequest) (*pb.ReadTablesByAppsReply, error) {
	return &pb.ReadTablesByAppsReply{}, nil
}
func (s *TableService) ReadOrgTables(ctx context.Context, req *pb.ReadOrgTablesRequest) (*pb.ReadOrgTablesReply, error) {
	return &pb.ReadOrgTablesReply{}, nil
}
func (s *TableService) ReadTableSchemas(ctx context.Context, req *pb.ReadTableSchemasRequest) (*pb.ReadTableSchemasReply, error) {
	return &pb.ReadTableSchemasReply{}, nil
}
func (s *TableService) ReadTableSchemasByAppId(ctx context.Context, req *pb.ReadTableSchemasByAppIdRequest) (*pb.ReadTableSchemasByAppIdReply, error) {
	return &pb.ReadTableSchemasByAppIdReply{}, nil
}
func (s *TableService) ReadOrgTableSchemas(ctx context.Context, req *pb.ReadOrgTableSchemasRequest) (*pb.ReadOrgTableSchemasReply, error) {
	return &pb.ReadOrgTableSchemasReply{}, nil
}
func (s *TableService) ReadSummeryTableId(ctx context.Context, req *pb.ReadSummeryTableIdRequest) (*pb.ReadSummeryTableIdReply, error) {
	return &pb.ReadSummeryTableIdReply{}, nil
}
func (s *TableService) InitOrgColumns(ctx context.Context, req *pb.InitOrgColumnsRequest) (*pb.InitOrgColumnsReply, error) {
	return &pb.InitOrgColumnsReply{}, nil
}
func (s *TableService) ReadOrgColumns(ctx context.Context, req *pb.ReadOrgColumnsRequest) (*pb.ReadOrgColumnsReply, error) {
	return &pb.ReadOrgColumnsReply{}, nil
}
func (s *TableService) CreateOrgColumn(ctx context.Context, req *pb.CreateOrgColumnRequest) (*pb.CreateOrgColumnReply, error) {
	return &pb.CreateOrgColumnReply{}, nil
}
func (s *TableService) DeleteOrgColumn(ctx context.Context, req *pb.DeleteOrgColumnRequest) (*pb.DeleteOrgColumnReply, error) {
	return &pb.DeleteOrgColumnReply{}, nil
}
func (s *TableService) CreateColumn(ctx context.Context, req *pb.CreateColumnRequest) (*pb.CreateColumnReply, error) {
	return &pb.CreateColumnReply{}, nil
}
func (s *TableService) CopyColumn(ctx context.Context, req *pb.CopyColumnRequest) (*pb.CopyColumnReply, error) {
	return &pb.CopyColumnReply{}, nil
}
func (s *TableService) UpdateColumn(ctx context.Context, req *pb.UpdateColumnRequest) (*pb.UpdateColumnReply, error) {
	return &pb.UpdateColumnReply{}, nil
}
func (s *TableService) UpdateColumnDescription(ctx context.Context, req *pb.UpdateColumnDescriptionRequest) (*pb.UpdateColumnDescriptionReply, error) {
	return &pb.UpdateColumnDescriptionReply{}, nil
}
func (s *TableService) DeleteColumn(ctx context.Context, req *pb.DeleteColumnRequest) (*pb.DeleteColumnReply, error) {
	return &pb.DeleteColumnReply{}, nil
}
func (s *TableService) CreateRows(ctx context.Context, req *pb.CreateRowsRequest) (*pb.CreateRowsReply, error) {
	return &pb.CreateRowsReply{}, nil
}
func (s *TableService) MoveRow(ctx context.Context, req *pb.MoveRowRequest) (*pb.MoveRowReply, error) {
	return &pb.MoveRowReply{}, nil
}
func (s *TableService) CopyRow(ctx context.Context, req *pb.CopyRowRequest) (*pb.CopyRowReply, error) {
	return &pb.CopyRowReply{}, nil
}
func (s *TableService) DeleteRow(ctx context.Context, req *pb.DeleteRowRequest) (*pb.DeleteRowReply, error) {
	return &pb.DeleteRowReply{}, nil
}
func (s *TableService) UpdateRowRelate(ctx context.Context, req *pb.UpdateRowRelateReq) (*pb.UpdateRowRelateReply, error) {
	return &pb.UpdateRowRelateReply{}, nil
}
func (s *TableService) UpdateRowBeforeAfter(ctx context.Context, req *pb.UpdateRowBeforeAfterReq) (*pb.UpdateRowBeforeAfterReply, error) {
	return &pb.UpdateRowBeforeAfterReply{}, nil
}
func (s *TableService) GetRowRelationList(ctx context.Context, req *pb.GetRelationListReq) (*pb.GetRelationListReply, error) {
	return &pb.GetRelationListReply{}, nil
}
func (s *TableService) GetBeforeAfterRowList(ctx context.Context, req *pb.GetBeforeAfterRowListReq) (*pb.GetBeforeAfterRowListReply, error) {
	return &pb.GetBeforeAfterRowListReply{}, nil
}
