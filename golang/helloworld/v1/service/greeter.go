package service

import (
	"context"

	pb "github.com/star-table/interface/golang/helloworld/v1"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *GreeterService) SayHelloAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *GreeterService) SayHelloAgainAgain(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
func (s *GreeterService) TransIn(ctx context.Context, req *pb.TransInRequest) (*pb.TransInReply, error) {
	return &pb.TransInReply{}, nil
}
func (s *GreeterService) TransInCompensate(ctx context.Context, req *pb.TransInRequest) (*pb.TransInReply, error) {
	return &pb.TransInReply{}, nil
}
func (s *GreeterService) TransOut(ctx context.Context, req *pb.TransOutRequest) (*pb.TransOutReply, error) {
	return &pb.TransOutReply{}, nil
}
func (s *GreeterService) TransOutCompensate(ctx context.Context, req *pb.TransOutRequest) (*pb.TransOutReply, error) {
	return &pb.TransOutReply{}, nil
}
