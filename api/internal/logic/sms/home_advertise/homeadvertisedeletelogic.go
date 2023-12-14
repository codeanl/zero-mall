package home_advertise

import (
	"context"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseDeleteLogic {
	return &HomeAdvertiseDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(req *types.DeleteHomeAdvertiseReq) (resp *types.DeleteHomeAdvertiseResp, err error) {
	_, err = l.svcCtx.Sms.HomeAdvertiseDelete(l.ctx, &smsclient.HomeAdvertiseDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteHomeAdvertiseResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteHomeAdvertiseResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
