package service

import (
	"context"

	pb "github.com/star-table/interface/golang/knowledge/v1"
)

type UserActionService struct {
	pb.UnimplementedUserActionServer
}

func NewUserActionService() *UserActionService {
	return &UserActionService{}
}

func (s *UserActionService) LikeDocument(ctx context.Context, req *pb.LikeDocumentRequest) (*pb.LikeDocumentReply, error) {
	return &pb.LikeDocumentReply{}, nil
}
func (s *UserActionService) ListDocumentLikeUser(ctx context.Context, req *pb.ListDocumentLikeUserRequest) (*pb.ListDocumentLikeUserReply, error) {
	return &pb.ListDocumentLikeUserReply{}, nil
}
func (s *UserActionService) TopDocument(ctx context.Context, req *pb.TopDocumentRequest) (*pb.TopDocumentReply, error) {
	return &pb.TopDocumentReply{}, nil
}
func (s *UserActionService) CommentDocument(ctx context.Context, req *pb.CommentDocumentRequest) (*pb.CommentDocumentReply, error) {
	return &pb.CommentDocumentReply{}, nil
}
func (s *UserActionService) UpdateCommentStatus(ctx context.Context, req *pb.UpdateCommentStatusRequest) (*pb.UpdateCommentStatusReply, error) {
	return &pb.UpdateCommentStatusReply{}, nil
}
func (s *UserActionService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, nil
}
func (s *UserActionService) ListDocumentComment(ctx context.Context, req *pb.ListDocumentCommentRequest) (*pb.ListDocumentCommentReply, error) {
	return &pb.ListDocumentCommentReply{}, nil
}
func (s *UserActionService) StarSpace(ctx context.Context, req *pb.StarSpaceRequest) (*pb.StarSpaceReply, error) {
	return &pb.StarSpaceReply{}, nil
}
