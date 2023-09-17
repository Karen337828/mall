package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mall/user/database/sqlx/usermodel"
	"mall/user/internal/models"
	"mall/user/internal/svc"
	"mall/user/internal/types"
	"mall/user/internal/utils"
	"mall/user/proto"
	"strconv"
	"strings"
	"time"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func GetRedisKey(key string) string {
	return fmt.Sprintf("%s_lock", key)
}

func (l *UserLogic) Register(req *types.UserRegisterRequest) (resp *types.Response, err error) {

	if strings.TrimSpace(req.Phone) == "" && strings.TrimSpace(req.Email) == "" {
		return nil, errors.New("请填写手机或邮箱")
	}
	user := new(usermodel.User)
	//拷贝对象属性[类似于java中的BeanUtils.copyProperties]
	deepcopier.Copy(req).To(user)
	if len(user.Email) == 0 {
		format := utils.VerifyPhoneFormat(user.Phone)
		if !format {
			return nil, errors.New("手机号码格式不正确！")
		}
		user.UserName = user.Phone
	} else {
		format := utils.VerifyEmailFormat(user.Email)
		if !format {
			return nil, errors.New("邮箱格式不正确！")
		}
		user.UserName = user.Email
	}

	lock := redis.NewRedisLock(l.svcCtx.Redis, GetRedisKey(user.UserName))
	lock.SetExpire(10)
	acquire, err := lock.Acquire()
	switch {
	case err != nil:
		logx.Errorf("Register Lock Error:", err)

	case !acquire:
		return nil, errors.New("操作频繁，请稍后再试！")

	case acquire:
		defer lock.Release()

		existUser, err := l.svcCtx.UserModel.FindOneByPhoneEmailName(l.ctx, user.UserName)
		if err != nil {
			logx.Errorf("用户注册时FindOneByPhoneEmailName报错:%s", err)
			return types.NewErrorEntityDefault(), nil
		}
		if existUser != nil {
			return nil, errors.New("用户已存在！")
		}

		user.Password = utils.SM4Encrypt(user.Password)

		resp = new(types.Response)
		user.RegisterTime = time.Now()
		user.LastLoginTime = user.RegisterTime
		_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
		if err != nil {
			resp.Code = "1001"
			resp.Message = err.Error()
		}

		thisUser, _ := l.svcCtx.UserModel.FindOneByUserName(l.ctx, user.UserName)
		u := models.User{}
		deepcopier.Copy(thisUser).To(&u)
		u.Token = uuid.New().String()
		resp = types.NewSuccessEntity(u)

		userInfoByte, _ := json.Marshal(thisUser)

		//设置到redis中
		l.svcCtx.Redis.Setex(strconv.FormatInt(thisUser.Id, 10)+":token", u.Token, 60*60*2)
		l.svcCtx.Redis.Setex(strconv.FormatInt(thisUser.Id, 10)+":user", string(userInfoByte), 60*60*2)

		//发送注册邮件
		go sendMsg("SE0001", u)
	}
	return resp, err
}

func sendMsg(sceneNo string, user models.User) {
	//client := proto.NewMsgClient(l.svcCtx.RpcConn)

	client := proto.NewMsgClient(GetGrpcConn())
	msg, err := client.SendMsg(context.Background(), &proto.SendRequest{SceneNo: sceneNo, Phone: user.Phone, Email: user.Email, UserName: user.UserName})
	if err != nil {
		fmt.Println("err====", err)
	}
	fmt.Printf("发送信息 场景码[%s],用户名[%s],rpc调用结果[%s]", sceneNo, user.UserName, msg)
}

func (l *UserLogic) Login(req *types.UserLoginRequest) (resp *types.Response, err error) {

	userName := req.UserName
	password := req.Password

	user, _ := l.svcCtx.UserModel.FindOneByPhoneEmailName(l.ctx, userName)
	if user == nil {
		return nil, errors.New("用户不存在！")
	}

	if utils.SM4Encrypt(password) != user.Password {
		return nil, errors.New("登录异常！")
	}

	u := models.User{}
	deepcopier.Copy(user).To(&u)
	u.Token = uuid.New().String()
	resp = types.NewSuccessEntity(u)

	userInfoByte, _ := json.Marshal(user)

	//设置到redis中
	l.svcCtx.Redis.Setex(strconv.FormatInt(user.Id, 10)+":token", u.Token, 60*60*2)
	l.svcCtx.Redis.Setex(strconv.FormatInt(user.Id, 10)+":user", string(userInfoByte), 60*60*2)

	//发送登录邮件
	go sendMsg("SE0002", u)

	data, _ := json.Marshal(u)
	utils.SendMQMsg(l.svcCtx.Config, "topic-msg", string(data))
	//l.svcCtx.KqPusherClient.Push(string(data))

	return resp, nil
}

func GetGrpcConn() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logx.Errorf("ddddddddddddddddddd")
	}
	return conn
}
