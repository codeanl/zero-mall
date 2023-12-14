package coupon

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateCouponLogic {
	return &SaveOrUpdateCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateCouponLogic) SaveOrUpdateCoupon(req *types.SaveOrUpdateCouponReq) (resp *types.SaveOrUpdateCouponResp, err error) {
	_, err = l.svcCtx.Sms.SaveOrUpdateCoupon(l.ctx, &smsclient.SaveOrUpdateCouponReq{
		Id:         req.ID,
		Type:       req.Type,
		Name:       req.Name,
		Count:      req.Count,
		Amount:     float32(req.Amount),
		PerLimit:   req.PerLimit,
		MinPoint:   float32(req.MinPoint),
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		UseType:    req.UseType,
		Note:       req.Note,
		EnableTime: req.EnableTime,
		Code:       req.Code,
	})
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateCouponResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.SaveOrUpdateCouponResp{
		Code:    200,
		Message: msg,
	}, nil
}
