package service

import (
	"context"

	pb "officeapi/api/officeapi/v1"
)

type OfficeApiService struct {
	pb.UnimplementedOfficeApiServer
}

func NewOfficeApiService() *OfficeApiService {
	return &OfficeApiService{}
}

func (s *OfficeApiService) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	return &pb.GetConfigResponse{}, nil
}
func (s *OfficeApiService) GetWopiFileInfoEditable(ctx context.Context, req *pb.WOPIRequest) (*pb.WOPIEditableResp, error) {
	return &pb.WOPIEditableResp{}, nil
}
