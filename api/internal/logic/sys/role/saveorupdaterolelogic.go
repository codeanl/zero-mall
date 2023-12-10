package role

import (
	"context"
	"google.golang.org/grpc/status"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateRoleLogic {
	return &SaveOrUpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateRoleLogic) SaveOrUpdateRole(req *types.SaveOrUpdateRoleReq) (resp *types.SaveOrUpdateRoleResp, err error) {
	data := &sysclient.SaveOrUpdateRoleReq{
		Id:     req.ID,
		Name:   req.Name,
		Remark: req.Remark,
	}
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
		data.UpdateBy = l.ctx.Value("username").(string)
	} else {
		data.CreateBy = l.ctx.Value("username").(string)
	}
	//
	_, err = l.svcCtx.Sys.SaveOrUpdateRole(l.ctx, data)
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateRoleResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	//
	return &types.SaveOrUpdateRoleResp{
		Code:    200,
		Message: msg,
	}, nil
}
