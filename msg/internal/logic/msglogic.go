package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/msg/database/sqlx/msgmodel"
	"mall/msg/internal/config"
	"mall/msg/internal/svc"
	"mall/msg/internal/types"
	"mall/msg/proto"
	"net/smtp"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgLogic {
	return &MsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgLogic) Msg(req *types.Request) (resp *types.Response, err error) {
	return
}

type Server struct {
	proto.UnimplementedMsgServer
}

func (s *Server) SendMsg(ctx context.Context, req *proto.SendRequest) (*proto.SendResponse, error) {
	logx.Infof("收到发送信息请求。场景码[%s],用户名[%s],手机[%s],邮箱[%s]", req.SceneNo, req.UserName, req.Phone, req.Email)

	var title string
	var content string
	var msgType = 0
	if req.SceneNo == "SE0001" {
		//注册场景
		title = "注册成功，欢迎成为尊贵会员！"
		content = fmt.Sprintf("尊敬的%s,注册成功，恭喜恭喜", req.UserName)

	} else if req.SceneNo == "SE0002" {
		//登录场景
		title = "登录成功，欢迎体验！"
		content = fmt.Sprintf("尊敬的%s,登录成功，欢迎光临", req.UserName)
	}

	if len(req.Email) > 0 {
		msgType = 1
	} else {
		msgType = 2
	}

	sqlConn := sqlx.NewMysql(config.Conf.MySQL.DataSource)

	_, err := msgmodel.NewMessageModel(sqlConn).Insert(ctx,
		&msgmodel.Message{
			MsgType:  msgType,
			SceneNo:  req.SceneNo,
			UserName: req.UserName,
			Phone:    req.Phone,
			Email:    req.Email,
			Title:    title,
			Content:  content,
			Status:   0,
		})
	if err != nil {
		logx.Error("插入信息表message错误", err)
	}

	return nil, nil
	//return doSendMsg(req)
}

func doSendMsg(req *proto.SendRequest) (*proto.SendResponse, error) {

	if len(req.Email) > 0 {

		e := email.NewEmail()
		//设置发送方的邮箱
		e.From = "2459212504@qq.com"
		// 设置接收方的邮箱
		e.To = []string{req.Email}

		sceneNo := req.SceneNo

		if sceneNo == "SE0001" {
			//注册场景
			e.Subject = "注册成功，欢迎成为尊贵会员！"
			e.Text = []byte(fmt.Sprintf("尊敬的%s,注册成功，恭喜恭喜", req.UserName))

		} else if sceneNo == "SE0002" {
			//登录场景
			e.Subject = "登录成功，欢迎体验！"
			e.Text = []byte(fmt.Sprintf("尊敬的%s,登录成功，欢迎光临", req.UserName))
		}

		//设置服务器相关的配置
		err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2459212504@qq.com", "xupzhvrwiunxdijb", "smtp.qq.com"))
		if err != nil {
			logx.Infov("发送邮件失败，请检查邮件服务器配置是否正确！")
		}
		return &proto.SendResponse{Result: err == nil}, err
	}

	return &proto.SendResponse{Result: false}, errors.New("邮箱不正确，暂时不支持短信发送，后续将会更新")
}
