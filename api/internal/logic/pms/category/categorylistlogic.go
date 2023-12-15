package category

import (
	"context"
	"encoding/json"
	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"
	"simple_mall_new/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.ListCategoryReq) (resp *types.ListCategoryResp, err error) {
	data, err := l.svcCtx.Pms.CategoryList(l.ctx, &pmsclient.CategoryListReq{
		//PageNum:  req.PageNum,
		//PageSize: req.PageSize,
		//Name:     req.Name,
		//ParentId: req.ParentId,
	})
	if err != nil {
		return &types.ListCategoryResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []types.ListCategoryData
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	menuTree := buildTree(list, 0)
	return &types.ListCategoryResp{
		Code:    200,
		Message: "查询成功",
		Total:   data.Total,
		Data:    menuTree,
	}, nil
}
func buildTree(Items []types.ListCategoryData, parentID int64) []types.ListCategoryData {
	tree := make([]types.ListCategoryData, 0)
	for _, item := range Items {
		if item.ParentId == parentID {
			children := buildTree(Items, item.ID)
			item.Children = children
			tree = append(tree, item)
		}
	}
	return tree
}
