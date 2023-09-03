package main

import (
	"flag"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mall/user/internal/config"
	"mall/user/internal/handler"
	"mall/user/internal/svc"
)

var configFile = flag.String("f", "user/etc/user.yaml", "the config file")

func main() {

	key, _ := sm2.GenerateKey(nil)
	fmt.Printf("-----%s", key)

	//解析命令行参数。可以指定配置文件，否则读取默认的user/etc/user.yaml配置
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//禁用每分钟打印的监控日志[建议打开]
	logx.DisableStat()
	//加载日志配置
	logx.MustSetup(c.Log)
	logx.Infov("我是日志,测试我设置的日志格式是否生效")

	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logx.Errorf("ddddddddddddddddddd")
	}
	defer conn.Close()

	//读取rest http配置
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//注册路由
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//开启服务
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
