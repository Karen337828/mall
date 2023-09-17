package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mall/user/database/sqlx/usermodel"
	"mall/user/internal/config"
	"mall/user/rpcclient/goodsclient"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      usermodel.UserModel
	Redis          *redis.Redis
	GoodsRpc       goodsclient.Goods
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redis := *redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:         c,
		UserModel:      usermodel.NewUserModel(sqlConn),
		Redis:          &redis,
		GoodsRpc:       goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
