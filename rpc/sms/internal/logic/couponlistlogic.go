package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 优惠券列表
func (l *CouponListLogic) CouponList(in *sms.CouponListReq) (*sms.CouponListResp, error) {
	all, total, err := l.svcCtx.CouponModel.GetCouponList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.CouponListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return &sms.CouponListResp{
		Total: total,
		List:  list,
	}, nil
}
