package menu

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListsLogic {
	return &MenuListsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListsLogic) MenuLists(req *types.ListMenuReq) (resp *types.ListMenuResp, err error) {
	data, err := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{})
	if err != nil {
		return &types.ListMenuResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	//
	var list []types.ListMenuData
	jsonData, err := json.Marshal(data.MenuList)
	err = json.Unmarshal(jsonData, &list)

	menuTree := buildMenuTree(list, 0)

	return &types.ListMenuResp{
		Code:    200,
		Data:    menuTree,
		Message: "查询成功",
	}, nil
}

func buildMenuTree(menuItems []types.ListMenuData, parentID int64) []types.ListMenuData {
	tree := make([]types.ListMenuData, 0)
	for _, item := range menuItems {
		if item.ParentId == parentID {
			children := buildMenuTree(menuItems, item.ID)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
