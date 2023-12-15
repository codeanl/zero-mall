package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/pms/model"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsUpdateLogic {
	return &MerchantsUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商家
func (l *MerchantsUpdateLogic) MerchantsUpdate(in *pms.MerchantsUpdateReq) (*pms.MerchantsUpdateResp, error) {
	err := l.svcCtx.MerchantsModel.UpdateMerchants(in.Id, &model.Merchants{
		Name:           in.Name,
		PrincipalName:  in.PrincipalName,
		PrincipalPhone: in.PrincipalPhone,
		Address:        in.Address,
		Pic:            in.Pic,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.MerchantsUpdateResp{}, nil
}
