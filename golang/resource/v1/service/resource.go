package service

import (
	"context"

	pb "github.com/star-table/interface/golang/resource/v1"
)

type ResourceService struct {
	pb.UnimplementedResourceServer
}

func NewResourceService() *ResourceService {
	return &ResourceService{}
}

func (s *ResourceService) CreateAttachment(ctx context.Context, req *pb.CreateAttachmentRequest) (*pb.CreateAttachmentResponse, error) {
	return &pb.CreateAttachmentResponse{}, nil
}
func (s *ResourceService) DeleteAttachment(ctx context.Context, req *pb.DeleteAttachmentRequest) (*pb.DeleteAttachmentResponse, error) {
	return &pb.DeleteAttachmentResponse{}, nil
}
func (s *ResourceService) ListAttachments(ctx context.Context, req *pb.ListAttachmentsRequest) (*pb.ListAttachmentsResponse, error) {
	return &pb.ListAttachmentsResponse{}, nil
}
func (s *ResourceService) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	return &pb.ListFilesResponse{}, nil
}
func (s *ResourceService) GetPolicy(ctx context.Context, req *pb.GetPolicyRequest) (*pb.GetPolicyResponse, error) {
	return &pb.GetPolicyResponse{}, nil
}
