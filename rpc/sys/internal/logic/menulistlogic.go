package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 菜单列表
func (l *MenuListLogic) MenuList(in *sys.MenuListReq) (*sys.MenuListResp, error) {
	all, _, err := l.svcCtx.MenuModel.GetMenuList()
	if err != nil {
		return nil, err
	}
	//
	var list []*sys.MenuList
	jsonData, err := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	//

	return &sys.MenuListResp{
		MenuList: list,
	}, nil
}
