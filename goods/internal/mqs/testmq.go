package mqs

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/goods/internal/svc"
)

type TestMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *TestMq {
	return &TestMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestMq) Consume(key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
