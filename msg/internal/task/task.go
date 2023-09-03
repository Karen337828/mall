package task

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/msg/database/sqlx/msgmodel"
	"mall/msg/internal/config"
	"net/smtp"
	"time"
)

func InitTask() {
	// 新建一个定时任务对象,根据cron表达式进行时间调度,cron可以精确到秒,大部分表达式格式也是从秒开始
	// 默认从分开始进行时间调度
	// cronTab := cron.New()
	// 精确到秒
	cronTab := cron.New(cron.WithSeconds())
	// 定义定时器调用的任务函数
	task := func() {
		fmt.Println("hello world", time.Now())

		sqlConn := sqlx.NewMysql(config.Conf.MySQL.DataSource)
		list, err := msgmodel.NewMessageModel(sqlConn).FindList(context.Background())
		if err != nil {
			fmt.Println("doTask query error", err)
		}

		if list == nil {
			fmt.Println("doTask query is empty", err)

		}

		for _, message := range *list {
			logx.Infov(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			message.Status = doSendMsg(message)

			msgmodel.NewMessageModel(sqlConn).Update(context.Background(), &message)

			logx.Infov("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		}

	}
	// 定时任务,cron表达式,每10秒一次
	spec := "*/ * * * * ?"
	// 添加定时任务
	cronTab.AddFunc(spec, task)
	// 启动定时器
	cronTab.Start()
	select {}
}

func doSendMsg(message msgmodel.Message) int {

	if len(message.Email) > 0 {

		e := email.NewEmail()
		//设置发送方的邮箱
		e.From = "2459212504@qq.com"
		// 设置接收方的邮箱
		e.To = []string{message.Email}

		sceneNo := message.SceneNo

		if sceneNo == "SE0001" {
			//注册场景
			e.Subject = "注册成功，欢迎成为尊贵会员！"
			e.Text = []byte(fmt.Sprintf("尊敬的%s,注册成功，恭喜恭喜", message.UserName))

		} else if sceneNo == "SE0002" {
			//登录场景
			e.Subject = "登录成功，欢迎体验！"
			e.Text = []byte(fmt.Sprintf("尊敬的%s,登录成功，欢迎光临", message.UserName))
		}

		//设置服务器相关的配置
		err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2459212504@qq.com", "xupzhvrwiunxdijb", "smtp.qq.com"))
		if err != nil {
			logx.Infov("发送邮件失败，请检查邮件服务器配置是否正确！")
			return 2
		}
	}
	return 1
}
