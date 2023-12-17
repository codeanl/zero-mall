package sku

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuListLogic {
	return &SkuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SkuListLogic) SkuList(req *types.ListSkuReq) (resp *types.ListSkuResp, err error) {
	data, err := l.svcCtx.Pms.SkuList(l.ctx, &pmsclient.SkuListReq{
		ProductIds: req.ProductID,
	})
	var list []*types.ListSkuData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	return &types.ListSkuResp{
		Code:    200,
		Message: "查询成功",
		Data:    list,
	}, nil
}
