package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分类列表
func (l *CategoryListLogic) CategoryList(in *pms.CategoryListReq) (*pms.CategoryListResp, error) {
	all, total, err := l.svcCtx.CategoryModel.GetCategoryList()
	if err != nil {
		return nil, err
	}
	var list []*pms.CategoryListData
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	return &pms.CategoryListResp{
		Total: total,
		List:  list,
	}, nil
}
