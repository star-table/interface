package service

import (
	"context"

	pb "github.com/star-table/interface/golang/knowledge/v1"
)

type SpaceService struct {
	pb.UnimplementedSpaceServer
}

func NewSpaceService() *SpaceService {
	return &SpaceService{}
}

func (s *SpaceService) CreateSpace(ctx context.Context, req *pb.CreateSpaceRequest) (*pb.CreateSpaceReply, error) {
	return &pb.CreateSpaceReply{}, nil
}
func (s *SpaceService) ListSpace(ctx context.Context, req *pb.ListSpaceRequest) (*pb.ListSpaceReply, error) {
	return &pb.ListSpaceReply{}, nil
}
