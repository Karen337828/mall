package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/msg/database/sqlx/msgmodel"
	"mall/msg/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel msgmodel.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(config.Conf.MySQL.DataSource)

	return &ServiceContext{
		Config:       c,
		MessageModel: msgmodel.NewMessageModel(sqlConn),
	}
}
