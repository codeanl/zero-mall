package product

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoLogic {
	return &ProductInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductInfoLogic) ProductInfo(req *types.ProductInfoReq) (resp *types.ProductInfoResp, err error) {
	data, err := l.svcCtx.Pms.ProductInfo(l.ctx, &pmsclient.ProductInfoReq{
		ID: req.ID,
	})
	var info types.InfoData
	jsonData, err := json.Marshal(data)
	err = json.Unmarshal(jsonData, &info)
	return &types.ProductInfoResp{
		Code:    200,
		Message: "查询成功",
		Data:    info,
	}, nil
}
