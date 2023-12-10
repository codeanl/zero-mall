package menu

import (
	"context"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuDeleteLogic) MenuDelete(req *types.DeleteMenuReq) (resp *types.DeleteMenuResp, err error) {
	_, err = l.svcCtx.Sys.MenuDelete(l.ctx, &sysclient.MenuDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteMenuResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteMenuResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
