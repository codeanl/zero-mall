package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *MenuDeleteLogic) MenuDelete(in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	err := l.svcCtx.MenuModel.DeleteMenuByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sys.MenuDeleteResp{}, nil
}
