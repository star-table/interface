package service

import (
	"context"

	pb "github.com/star-table/interface/golang/n8n_lesscode/v1"
)

type N8nLesscodeService struct {
	pb.UnimplementedN8nLesscodeServer
}

func NewN8nLesscodeService() *N8nLesscodeService {
	return &N8nLesscodeService{}
}

func (s *N8nLesscodeService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
