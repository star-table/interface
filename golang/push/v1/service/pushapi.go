package service

import (
	"context"

	pb "github.com/star-table/interface/golang/push/v1"
)

type PushApiService struct {
	pb.UnimplementedPushApiServer
}

func NewPushApiService() *PushApiService {
	return &PushApiService{}
}

func (s *PushApiService) SendCard(ctx context.Context, req *pb.SendCardReq) (*pb.SendCardReply, error) {
	return &pb.SendCardReply{}, nil
}
