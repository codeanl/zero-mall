package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pms.MerchantsUpdateResp{}, nil
}
