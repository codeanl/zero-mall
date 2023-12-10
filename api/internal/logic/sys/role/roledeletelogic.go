package role

import (
	"context"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.DeleteRoleReq) (resp *types.DeleteRoleResp, err error) {
	_, err = l.svcCtx.Sys.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Ids: req.IDs,
	})
	if err != nil {
		return &types.DeleteRoleResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteRoleResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
