package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/model"
	"simple_mall_new/rpc/pms/pms"
)

type SaveOrUpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateCategoryLogic {
	return &SaveOrUpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加||更新分类
func (l *SaveOrUpdateCategoryLogic) SaveOrUpdateCategory(in *pms.SaveOrUpdateCategoryReq) (*pms.SaveOrUpdateCategoryResp, error) {
	//查询用户名是否存在
	userInfo, exist, _ := l.svcCtx.CategoryModel.GetCategoryByName(in.Name)
	if exist && int64(userInfo.ID) != in.Id {
		return nil, errors.New("菜单已存在")
	}
	//
	cate := &model.Category{
		Name:     in.Name,
		ParentId: in.ParentId,
		Status:   in.Status,
		Sort:     in.Sort,
		Icon:     in.Icon,
		Desc:     in.Desc,
	}
	err := l.svcCtx.CategoryModel.SaveOrUpdateCategory(in.Id, cate)
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	return &pms.SaveOrUpdateCategoryResp{}, nil
}
