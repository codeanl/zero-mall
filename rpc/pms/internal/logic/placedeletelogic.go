package logic

import (
	"context"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceDeleteLogic {
	return &PlaceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除自提点
func (l *PlaceDeleteLogic) PlaceDelete(in *pms.PlaceDeleteReq) (*pms.PlaceDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &pms.PlaceDeleteResp{}, nil
}
