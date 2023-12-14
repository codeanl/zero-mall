package home_advertise

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateHomeAdvertiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateHomeAdvertiseLogic {
	return &SaveOrUpdateHomeAdvertiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateHomeAdvertiseLogic) SaveOrUpdateHomeAdvertise(req *types.SaveOrUpdateHomeAdvertiseReq) (resp *types.SaveOrUpdateHomeAdvertiseResp, err error) {
	_, err = l.svcCtx.Sms.SaveOrUpdateHomeAdvertise(l.ctx, &smsclient.SaveOrUpdateHomeAdvertiseReq{
		Id:     req.ID,
		Name:   req.Name,
		Pic:    req.Pic,
		Status: req.Status,
		Sort:   req.Sort,
		Url:    req.Url,
		Note:   req.Note,
	})
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateHomeAdvertiseResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.SaveOrUpdateHomeAdvertiseResp{
		Code:    200,
		Message: msg,
	}, nil
}
