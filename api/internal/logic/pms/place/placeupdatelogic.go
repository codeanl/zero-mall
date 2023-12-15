package place

import (
	"context"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceUpdateLogic) PlaceUpdate(req *types.UpdatePlaceReq) (resp *types.UpdatePlaceResp, err error) {
	_, err = l.svcCtx.Pms.PlaceUpdate(l.ctx, &pmsclient.PlaceUpdateReq{
		Id:             req.ID,
		Name:           req.Name,
		PrincipalPhone: req.PrincipalPhone,
		PrincipalName:  req.PrincipalName,
		Address:        req.Address,
		Pic:            req.Pic,
	})
	if err != nil {
		return &types.UpdatePlaceResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	//
	return &types.UpdatePlaceResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
