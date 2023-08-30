package logic

import (
	"context"

	"mall/msg/internal/svc"
	"mall/msg/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLogic {
	return &MsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgLogic) Msg(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
