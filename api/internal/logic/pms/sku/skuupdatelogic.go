package sku

import (
	"context"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuUpdateLogic {
	return &SkuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuUpdateLogic) SkuUpdate(req *types.UpdateSkuReq) (resp *types.UpdateSkuResp, err error) {
	_, err = l.svcCtx.Pms.SkuUpdate(l.ctx, &pmsclient.SkuUpdateReq{
		Id:    req.ID,
		Name:  req.Name,
		Pic:   req.Pic,
		Price: req.Price,
		SkuSn: req.SkuSn,
		Desc:  req.Desc,
		Stock: req.Stock,
	})
	if err != nil {
		return &types.UpdateSkuResp{
			Code:    400,
			Message: "更新失败",
		}, nil
	}
	return &types.UpdateSkuResp{
		Code:    200,
		Message: "更新成功",
	}, nil
}
