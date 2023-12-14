package subject

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateSubjectLogic {
	return &SaveOrUpdateSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateSubjectLogic) SaveOrUpdateSubject(req *types.SaveOrUpdateSubjectReq) (resp *types.SaveOrUpdateSubjectResp, err error) {
	_, err = l.svcCtx.Sms.SaveOrUpdateSubject(l.ctx, &smsclient.SaveOrUpdateSubjectReq{
		Id:     req.ID,
		Name:   req.Name,
		Pic:    req.Pic,
		Status: req.Status,
		Sort:   req.Sort,
	})
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateSubjectResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.SaveOrUpdateSubjectResp{
		Code:    200,
		Message: msg,
	}, nil
}
