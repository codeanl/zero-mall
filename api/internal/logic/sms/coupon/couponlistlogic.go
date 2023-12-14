package coupon

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponListLogic) CouponList(req *types.ListCouponReq) (resp *types.ListCouponResp, err error) {
	data, err := l.svcCtx.Sms.CouponList(l.ctx, &smsclient.CouponListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
		UseType:  req.UseType,
		Type:     req.Type,
	})
	if err != nil {
		return &types.ListCouponResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListCouponData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)

	return &types.ListCouponResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    list,
	}, nil
}
