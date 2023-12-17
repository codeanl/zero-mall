package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品列表
func (l *ProductListLogic) ProductList(in *pms.ProductListReq) (*pms.ProductListResp, error) {
	all, total, err := l.svcCtx.ProductModel.GetProductList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.ProductListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	return &pms.ProductListResp{
		Total: total,
		List:  list,
	}, nil
}
