package service

import (
	"context"

	pb "github.com/star-table/interface/golang/ping"
)

type PingService struct {
	pb.UnimplementedPingServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (s *PingService) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingReply, error) {
	return &pb.PingReply{}, nil
}
