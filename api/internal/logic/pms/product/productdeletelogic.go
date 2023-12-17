package product

import (
	"context"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.DeleteProductReq) (resp *types.DeleteProductResp, err error) {
	_, err = l.svcCtx.Pms.ProductDelete(l.ctx, &pmsclient.ProductDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteProductResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteProductResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
