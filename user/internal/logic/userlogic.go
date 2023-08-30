package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/user/internal/models"
	"mall/user/internal/svc"
	"mall/user/internal/types"
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

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	if req.Name == "wanchangxuan" {
		return types.NewErrorEntity("1001", "用户不存在"), nil
	}

	user := models.User{
		UserName:      "万昌轩",
		Age:           28,
		NickName:      "万贞煜",
		LastLoginTime: time.Now(),
	}

	return types.NewSuccessEntity(user), nil
}

func (l *UserLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = req.Name
	return
}

func (l *UserLogic) Login(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
