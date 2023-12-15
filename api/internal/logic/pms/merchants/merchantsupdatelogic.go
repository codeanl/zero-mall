package merchants

import (
	"context"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsUpdateLogic {
	return &MerchantsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantsUpdateLogic) MerchantsUpdate(req *types.UpdateMerchantsReq) (resp *types.UpdateMerchantsResp, err error) {
	_, err = l.svcCtx.Pms.MerchantsUpdate(l.ctx, &pmsclient.MerchantsUpdateReq{
		Id:             req.ID,
		Name:           req.Name,
		PrincipalPhone: req.PrincipalPhone,
		PrincipalName:  req.PrincipalName,
		Address:        req.Address,
		Pic:            req.Pic,
	})
	if err != nil {
		return &types.UpdateMerchantsResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	//
	return &types.UpdateMerchantsResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
