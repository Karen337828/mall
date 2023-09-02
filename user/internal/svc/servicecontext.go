package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/user/database/sqlx/usermodel"
	"mall/user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.UserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	redis := *redis.MustNewRedis(c.Redis.RedisConf)

	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewUserModel(sqlConn),
		Redis:     &redis,
	}
}
