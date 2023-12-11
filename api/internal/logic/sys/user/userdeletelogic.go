package user

import (
	"context"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	_, err = l.svcCtx.Sys.UserDelete(l.ctx, &sysclient.UserDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteUserResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteUserResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
