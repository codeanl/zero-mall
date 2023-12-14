package coupon

import (
	"context"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDeleteLogic {
	return &CouponDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponDeleteLogic) CouponDelete(req *types.DeleteCouponReq) (resp *types.DeleteCouponResp, err error) {
	_, err = l.svcCtx.Sms.CouponDelete(l.ctx, &smsclient.CouponDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteCouponResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteCouponResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
