package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"google.golang.org/grpc"
	"log"
	"mall/msg/internal/config"
	"mall/msg/internal/logic"
	"mall/msg/internal/task"
	"mall/msg/proto"
	"net"
)

var configFile = flag.String("f", "msg/etc/msg.yaml", "the config file")

func main() {

	flag.Parse()
	conf.MustLoad(*configFile, &config.Conf)

	go task.InitTask()

	listen, _ := net.Listen("tcp", "127.0.0.1:9090")
	grpcServer := grpc.NewServer()
	proto.RegisterMsgServer(grpcServer, &logic.Server{})

	err := grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动失败")
	}

	//flag.Parse()
	//
	//var c config.Config
	//conf.MustLoad(*configFile, &c)
	//
	//server := rest.MustNewServer(c.RestConf)
	//defer server.Stop()
	//
	//ctx := svc.NewServiceContext(c)
	//handler.RegisterHandlers(server, ctx)
	//
	//fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//server.Start()

}
