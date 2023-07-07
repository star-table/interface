package service

import (
	"context"

	pb "github.com/star-table/interface/golang/knowledge/v1"
)

type DocumentService struct {
	pb.UnimplementedDocumentServer
}

func NewDocumentService() *DocumentService {
	return &DocumentService{}
}

func (s *DocumentService) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentReply, error) {
	return &pb.GetDocumentReply{}, nil
}
func (s *DocumentService) ListTreeDocument(ctx context.Context, req *pb.ListTreeDocumentRequest) (*pb.ListTreeDocumentReply, error) {
	return &pb.ListTreeDocumentReply{}, nil
}
func (s *DocumentService) ListChildDocument(ctx context.Context, req *pb.ListChildDocumentRequest) (*pb.ListChildDocumentReply, error) {
	return &pb.ListChildDocumentReply{}, nil
}
func (s *DocumentService) SearchDocument(ctx context.Context, req *pb.SearchDocumentRequest) (*pb.SearchDocumentReply, error) {
	return &pb.SearchDocumentReply{}, nil
}
func (s *DocumentService) SaveDocument(ctx context.Context, req *pb.SaveDocumentRequest) (*pb.SaveDocumentReply, error) {
	return &pb.SaveDocumentReply{}, nil
}
func (s *DocumentService) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*pb.DeleteDocumentReply, error) {
	return &pb.DeleteDocumentReply{}, nil
}
func (s *DocumentService) CopyDocument(ctx context.Context, req *pb.CopyDocumentRequest) (*pb.CopyDocumentReply, error) {
	return &pb.CopyDocumentReply{}, nil
}
func (s *DocumentService) MoveDocument(ctx context.Context, req *pb.MoveDocumentRequest) (*pb.MoveDocumentReply, error) {
	return &pb.MoveDocumentReply{}, nil
}
