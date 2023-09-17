package utils

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/user/internal/config"
)

func SendMQMsg(c config.Config, topic string, data string) {
	pusher := kq.NewPusher(c.KqPusherConf.Brokers, topic)
	pusher.Push(data)
	logx.Infof("MQ发送消息成功:topic[%s],msg[%s]", topic, data)
}
