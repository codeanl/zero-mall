package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsDeleteLogic {
	return &MerchantsDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商家
func (l *MerchantsDeleteLogic) MerchantsDelete(in *pms.MerchantsDeleteReq) (*pms.MerchantsDeleteResp, error) {
	err := l.svcCtx.MerchantsModel.DeleteMerchantsByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &pms.MerchantsDeleteResp{}, nil
}
