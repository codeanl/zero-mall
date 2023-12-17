package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/model"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuUpdateLogic {
	return &SkuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新Sku
func (l *SkuUpdateLogic) SkuUpdate(in *pms.SkuUpdateReq) (*pms.SkuUpdateResp, error) {
	err := l.svcCtx.SkuModel.UpdateSku(in.Id, &model.Sku{
		Name:  in.Name,
		Pic:   in.Pic,
		Desc:  in.Desc,
		Stock: in.Stock,
		Price: in.Price,
		SkuSn: in.SkuSn,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}
	return &pms.SkuUpdateResp{}, nil
}
