package login_log

import (
	"context"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogDeleteLogic {
	return &LoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(req *types.DeleteLoginLogReq) (resp *types.DeleteLoginLogResp, err error) {
	_, err = l.svcCtx.Sys.LoginLogDelete(l.ctx, &sysclient.LoginLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteLoginLogResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteLoginLogResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
