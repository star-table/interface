package service

import (
	"context"

	pb "github.com/star-table/interface/golang/push/v1"
)

type PushService struct {
	pb.UnimplementedPushServer
}

func NewPushService() *PushService {
	return &PushService{}
}

func (s *PushService) PushMqtt(ctx context.Context, req *pb.PushMqttReq) (*pb.PushMqttReply, error) {
	return &pb.PushMqttReply{}, nil
}
func (s *PushService) GenerateMqttKey(ctx context.Context, req *pb.GenerateMqttKeyReq) (*pb.GenerateMqttKeyReply, error) {
	return &pb.GenerateMqttKeyReply{}, nil
}
func (s *PushService) PushMail(ctx context.Context, req *pb.PushMailReq) (*pb.PushMailReply, error) {
	return &pb.PushMailReply{}, nil
}
func (s *PushService) PushSms(ctx context.Context, req *pb.PushSmsReq) (*pb.PushSmsReply, error) {
	return &pb.PushSmsReply{}, nil
}
func (s *PushService) PushCard(ctx context.Context, req *pb.PushCardReq) (*pb.PushCardReply, error) {
	return &pb.PushCardReply{}, nil
}
func (s *PushService) PushCardSimple(ctx context.Context, req *pb.PushCardSimpleReq) (*pb.PushCardSimpleReply, error) {
	return &pb.PushCardSimpleReply{}, nil
}
func (s *PushService) GenerateCard(ctx context.Context, req *pb.GenerateCardReq) (*pb.GenerateCardReply, error) {
	return &pb.GenerateCardReply{}, nil
}
