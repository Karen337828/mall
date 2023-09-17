package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

var Conf Config

type Config struct {
	rest.RestConf
	MySQL struct {
		DataSource string
	}
	KqConsumerConf kq.KqConf
}
