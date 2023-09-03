package logic

import (
	"context"

	"mall/msg/internal/svc"
	"mall/msg/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *proto.SendRequest) (*proto.SendResponse, error) {

	logx.Infov("=====sendmsglogic====" + in.Email)
	return &proto.SendResponse{}, nil
}
