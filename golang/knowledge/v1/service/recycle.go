package service

import (
	"context"

	pb "github.com/star-table/interface/golang/knowledge/v1"
)

type RecycleService struct {
	pb.UnimplementedRecycleServer
}

func NewRecycleService() *RecycleService {
	return &RecycleService{}
}

func (s *RecycleService) RecycleDocument(ctx context.Context, req *pb.RecycleDocumentRequest) (*pb.RecycleDocumentReply, error) {
	return &pb.RecycleDocumentReply{}, nil
}
func (s *RecycleService) Recover(ctx context.Context, req *pb.RecoverRequest) (*pb.RecoverReply, error) {
	return &pb.RecoverReply{}, nil
}
func (s *RecycleService) ListRecycle(ctx context.Context, req *pb.ListRecycleRequest) (*pb.ListRecycleReply, error) {
	return &pb.ListRecycleReply{}, nil
}
