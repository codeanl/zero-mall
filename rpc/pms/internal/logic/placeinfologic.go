package logic

import (
	"context"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceInfoLogic {
	return &PlaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点详情
func (l *PlaceInfoLogic) PlaceInfo(in *pms.PlaceInfoReq) (*pms.PlaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pms.PlaceInfoResp{}, nil
}
