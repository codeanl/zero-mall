package user

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.UpdatePasswordResp, err error) {
	var id int64
	msg := "更新成功"
	if req.Type == "0" { //修改密码
		myId, _ := l.ctx.Value("id").(json.Number).Int64()
		id = myId
	} else if req.Type == "1" { //重置密码
		id = req.ID
		msg = "重置成功"
	}
	_, err = l.svcCtx.Sys.UpdatePassword(l.ctx, &sysclient.UpdatePasswordReq{
		Id:          id,
		Type:        req.Type,
		NewPassword: req.NewPassword,
		OldPassword: req.OldPassword,
	})
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.UpdatePasswordResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	return &types.UpdatePasswordResp{
		Code:    200,
		Message: msg,
	}, nil
}
