package logic

import (
	"context"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsListLogic {
	return &MerchantsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商家列表
func (l *MerchantsListLogic) MerchantsList(in *pms.MerchantsListReq) (*pms.MerchantsListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.MerchantsListResp{}, nil
}
