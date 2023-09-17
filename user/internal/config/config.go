package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Logx logx.LogConf

	Redis struct {
		redis.RedisConf
	}

	Mysql struct {
		DataSource string
	}

	GoodsRpc zrpc.RpcClientConf

	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
