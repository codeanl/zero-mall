package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/pms/model"

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
	err := l.svcCtx.PlaceModel.UpdatePlace(in.Id, &model.Place{
		Name:           in.Name,
		PrincipalName:  in.PrincipalName,
		PrincipalPhone: in.PrincipalPhone,
		Address:        in.Address,
		Pic:            in.Pic,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}

	return &pms.PlaceUpdateResp{}, nil
}
