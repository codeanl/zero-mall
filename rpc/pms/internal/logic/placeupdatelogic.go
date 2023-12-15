package logic

import (
	"context"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新自提点
func (l *PlaceUpdateLogic) PlaceUpdate(in *pms.PlaceUpdateReq) (*pms.PlaceUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.PlaceUpdateResp{}, nil
}
