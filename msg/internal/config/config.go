package config

import (
	"github.com/zeromicro/go-zero/rest"
)

var Conf Config

type Config struct {
	rest.RestConf

	MySQL struct {
		DataSource string
	}

	Email struct {
		From     string
		Addr     string
		Identity string
		Username string
		Password string
		Host     string
	}
}
