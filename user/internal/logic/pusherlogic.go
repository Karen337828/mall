package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/user/internal/svc"
)

type PusherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *PusherLogic) Pusher() error {

	//......业务逻辑....

	data := "zhangSan"
	if err := l.svcCtx.KqPusherClient.Push(data); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}

	return nil
}
