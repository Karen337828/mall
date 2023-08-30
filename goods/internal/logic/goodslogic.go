package logic

import (
	"context"

	"mall/goods/internal/svc"
	"mall/goods/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GoodsLogic) Goods(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
