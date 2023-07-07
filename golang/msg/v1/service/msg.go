package service

import (
	pb "github.com/star-table/interface/golang/msg/v1"
)

type MsgService struct {
	pb.UnimplementedMsgServer
}

func NewMsgService() *MsgService {
	return &MsgService{}
}
