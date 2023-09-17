package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"mall/goods/internal/config"
	"mall/goods/internal/mqs"
	"mall/goods/internal/server"
	"mall/goods/internal/svc"
	"mall/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "goods/etc/goods.yaml", "the config file")

func initGRPC(ctx *svc.ServiceContext, c config.Config) {

}

func initMQ(ctx *svc.ServiceContext, c config.Config) {
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
		serviceGroup.Add(mq)
	}
	fmt.Printf("Starting mq consumer at %s...\n", c.ListenOn)
	serviceGroup.Start()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	go initMQ(ctx, c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, server.NewGoodsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()

}
