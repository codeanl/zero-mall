package menu

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateMenuLogic {
	return &SaveOrUpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateMenuLogic) SaveOrUpdateMenu(req *types.SaveOrUpdateMenuReq) (resp *types.SaveOrUpdateMenuResp, err error) {
	data := &sysclient.SaveOrUpdateMenuReq{
		Id:       req.ID,
		Name:     req.Name,
		ParentId: req.ParentId,
		Url:      req.Url,
		Icon:     req.Icon,
		Type:     req.Type,
		OrderNum: req.OrderNum,
		Tag:      req.TAG,
	}
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	//
	_, err = l.svcCtx.Sys.SaveOrUpdateMenu(l.ctx, data)
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateMenuResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	//
	return &types.SaveOrUpdateMenuResp{
		Code:    200,
		Message: msg,
	}, nil
}
