package logic

import (
	"context"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 自提点列表
func (l *PlaceListLogic) PlaceList(in *pms.PlaceListReq) (*pms.PlaceListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.PlaceListResp{}, nil
}
