package service

import (
	"context"

	pb "github.com/star-table/interface/golang/batch/v1"
)

type TimingWheelService struct {
	pb.UnimplementedTimingWheelServer
}

func NewTimingWheelService() *TimingWheelService {
	return &TimingWheelService{}
}

func (s *TimingWheelService) SaveTimingWheel(ctx context.Context, req *pb.SaveTimingWheelReq) (*pb.SaveTimingWheelReply, error) {
	return &pb.SaveTimingWheelReply{}, nil
}
func (s *TimingWheelService) DeleteTimingWheel(ctx context.Context, req *pb.DeleteTimingWheelReq) (*pb.DeleteTimingWheelReply, error) {
	return &pb.DeleteTimingWheelReply{}, nil
}
func (s *TimingWheelService) DebugTimingWheel(ctx context.Context, req *pb.DebugTimingWheelReq) (*pb.DebugTimingWheelReply, error) {
	return &pb.DebugTimingWheelReply{}, nil
}
