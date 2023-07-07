package service

import (
	"context"

	pb "github.com/star-table/interface/golang/health/v1"
)

type HealthService struct {
	pb.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{}, nil
}
func (s *HealthService) Watch(req *pb.HealthCheckRequest, conn pb.Health_WatchServer) error {
	for {
		err := conn.Send(&pb.HealthCheckResponse{})
		if err != nil {
			return err
		}
	}
}
