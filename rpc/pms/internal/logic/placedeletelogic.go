package logic

import (
	"context"
	"errors"

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
	err := l.svcCtx.PlaceModel.DeletePlaceByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.PlaceDeleteResp{}, nil
}
