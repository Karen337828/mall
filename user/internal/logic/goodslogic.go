package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/user/internal/svc"
	"mall/user/internal/types"
	"mall/user/pb/goods"
)

type GoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsLogic {
	return &GoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsLogic) GetGoodsInfo(req *types.GetGoodsInfoRequest) (resp *types.Response, err error) {

	info, err := l.svcCtx.GoodsRpc.GetGoodsInfo(context.Background(), &goods.Request{GoodsId: req.GoodsId})

	if err != nil {
		return nil, err
	}
	resp = types.NewSuccessEntity(info)
	return resp, nil

}

func (l *GoodsLogic) Pusher() error {

	//......业务逻辑....

	data := "zhangSan"
	if err := l.svcCtx.KqPusherClient.Push(data); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}

	return nil
}
