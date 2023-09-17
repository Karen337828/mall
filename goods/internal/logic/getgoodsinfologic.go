package logic

import (
	"context"

	"mall/goods/internal/svc"
	"mall/goods/pb/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsInfoLogic {
	return &GetGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsInfoLogic) GetGoodsInfo(in *goods.Request) (*goods.Response, error) {
	// todo: add your logic here and delete this line

	return &goods.Response{
		GooodsId:      "1",
		GoodsName:     "测试商品",
		GoodsCategory: "化妆品",
		GoodsNums:     56,
		GoodsPrice:    79,
		GoodsRemark:   "测试rpc",
	}, nil
}
