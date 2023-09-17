package mqs

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/msg/internal/svc"
)

type MessageMQ struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *MessageMQ {
	return &MessageMQ{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageMQ) Consume(key, val string) error {
	logx.Infof("我是22222 key :%s , val :%s", key, val)
	//todo 存储到数据库
	return nil
}
