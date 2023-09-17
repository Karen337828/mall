package mqs

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/msg/internal/config"
	"mall/msg/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	logx.Info("我是11111")
	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewMessageSuccess(ctx, svcContext)),
		//.....
	}

}
