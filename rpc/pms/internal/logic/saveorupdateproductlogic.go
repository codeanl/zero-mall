package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/model"
	"simple_mall_new/rpc/pms/pms"
	"strings"
)

type SaveOrUpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateProductLogic {
	return &SaveOrUpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加||更新商品
func (l *SaveOrUpdateProductLogic) SaveOrUpdateProduct(in *pms.SaveOrUpdateProductReq) (*pms.SaveOrUpdateProductResp, error) {
	data, err := l.svcCtx.ProductModel.SaveOrUpdateProduct(in.Id, &model.Product{
		CategoryID:    in.CategoryId,
		MerchantID:    in.MerchantId,
		Name:          in.Name,
		Pic:           in.Pic,
		ProductSn:     in.ProductSn,
		Desc:          in.Desc,
		OriginalPrice: in.OriginalPrices,
		Price:         in.Price,
		IsHot:         in.IsHot,
		IsPublished:   in.IsPublished,
	})
	if err != nil {
		return nil, errors.New("失败")
	}
	//
	if in.Id > 0 {
		l.svcCtx.ProductImgModel.DeleteProductImgBySpuID(in.Id)
		l.svcCtx.ProductIntroduceImgModel.DeleteProductIntroduceImgBySpuID(in.Id)
	}

	for _, i := range in.ImgUrl {
		l.svcCtx.ProductImgModel.AddProductImg(&model.ProductImg{
			ProductID: int64(data.ID),
			Url:       i,
		})
	}
	for _, i := range in.IntroduceImgUrl {
		l.svcCtx.ProductIntroduceImgModel.AddProductIntroduceImg(&model.ProductIntroduceImg{
			ProductID: int64(data.ID),
			Url:       i,
		})
	}
	//
	l.svcCtx.SizeModel.DeleteSizeByProID(int64(data.ID))
	for _, i := range in.Size {
		l.svcCtx.SizeModel.AddSize(&model.Size{
			ProductID: int64(data.ID),
			Name:      i.Name,
			Value:     i.Value,
		})
	}
	//
	result := generateCombinations(in.Size)
	l.svcCtx.SkuModel.DeleteSkuByProID(int64(data.ID))
	for index, i := range result {
		jsonData, _ := json.Marshal(i)
		l.svcCtx.SkuModel.AddSku(&model.Sku{
			ProductID: int64(data.ID),
			Name:      data.Name,
			Pic:       data.Pic,
			SkuSn:     data.ProductSn + string(index),
			Desc:      data.Desc,
			Price:     data.Price,
			Stock:     999,
			Tag:       string(jsonData),
		})
	}

	return &pms.SaveOrUpdateProductResp{}, nil
}

type Combination map[string]string

func generateCombinationsHelper(keys []string, values [][]string, current Combination, result *[]Combination) {
	if len(keys) == 0 {
		*result = append(*result, current)
		return
	}
	key := keys[0]
	remainingKeys := keys[1:]
	keyValues := values[0]

	for _, value := range keyValues {
		newCombination := make(Combination)
		if current != nil {
			for k, v := range current {
				newCombination[k] = v
			}
		}
		newCombination[key] = value

		generateCombinationsHelper(remainingKeys, values[1:], newCombination, result)
	}
}
func generateCombinations(data []*pms.Size) []Combination {
	if len(data) == 0 {
		return nil
	}
	var result []Combination
	var keys []string
	var values [][]string
	// 获取所有键和对应的值列表
	for _, entry := range data {
		keys = append(keys, entry.Name)
		values = append(values, strings.Split(entry.Value, ","))
	}
	generateCombinationsHelper(keys, values, nil, &result)
	return result
}
