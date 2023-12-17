package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *ProductDeleteLogic) ProductDelete(in *pms.ProductDeleteReq) (*pms.ProductDeleteResp, error) {
	err := l.svcCtx.ProductModel.DeleteProductByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除商品失败")
	}
	for _, i := range in.Ids {
		err = l.svcCtx.SizeModel.DeleteSizeByProID(i)
		if err != nil {
			return nil, errors.New("删除商品规格失败")
		}
		err = l.svcCtx.SkuModel.DeleteSkuByProID(i)
		if err != nil {
			return nil, errors.New("删除商品sku失败")
		}
		//todo 删除关联图片 使用gorm关联
	}

	return &pms.ProductDeleteResp{}, nil
}
