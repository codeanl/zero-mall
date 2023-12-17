package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuListLogic {
	return &SkuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Sku列表
func (l *SkuListLogic) SkuList(in *pms.SkuListReq) (*pms.SkuListResp, error) {
	all, total, err := l.svcCtx.SkuModel.GetSkuList(in.ProductIds)
	fmt.Println(all[0].ID)
	if err != nil {
		return nil, err
	}
	var list []*pms.SkuListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	return &pms.SkuListResp{
		Total: total,
		List:  list,
	}, nil
}
