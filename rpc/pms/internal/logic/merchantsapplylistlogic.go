package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantsApplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantsApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantsApplyListLogic {
	return &MerchantsApplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 入驻申请列表
func (l *MerchantsApplyListLogic) MerchantsApplyList(in *pms.MerchantsApplyListReq) (*pms.MerchantsApplyListResp, error) {
	all, total, err := l.svcCtx.MerchantsApplyModel.GetMerchantsApplyList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.MerchantApplysListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return &pms.MerchantsApplyListResp{
		Total: total,
		List:  list,
	}, nil
}
