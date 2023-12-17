package product

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ListProductReq) (resp *types.ListProductResp, err error) {
	data, err := l.svcCtx.Pms.ProductList(l.ctx, &pmsclient.ProductListReq{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		Name:       req.Name,
		CategoryId: req.CategoryId,
		MerchantId: req.MerchantID,
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		SearchType: req.SearchType,
	})
	if err != nil {
		return &types.ListProductResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListProductData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListProductResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
