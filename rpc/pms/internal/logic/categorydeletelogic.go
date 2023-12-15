package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteLogic {
	return &CategoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除分类
func (l *CategoryDeleteLogic) CategoryDelete(in *pms.CategoryDeleteReq) (*pms.CategoryDeleteResp, error) {
	err := l.svcCtx.CategoryModel.DeleteCategoryByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &pms.CategoryDeleteResp{}, nil
}
