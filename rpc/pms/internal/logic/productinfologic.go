package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"simple_mall_new/rpc/pms/internal/svc"
	"simple_mall_new/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoLogic {
	return &ProductInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询商品详情
func (l *ProductInfoLogic) ProductInfo(in *pms.ProductInfoReq) (*pms.ProductInfoResp, error) {
	var ImgUrl []string
	img, _ := l.svcCtx.ProductImgModel.GetImgtByProducID(in.ID)
	for _, i := range img {
		ImgUrl = append(ImgUrl, i.Url)
	}
	//
	var IntroduceImgUrl []string
	introduceImgUrl, _ := l.svcCtx.ProductIntroduceImgModel.GetImgtByProducID(in.ID)
	for _, i := range introduceImgUrl {
		IntroduceImgUrl = append(IntroduceImgUrl, i.Url)
	}
	//
	productInfo, _ := l.svcCtx.ProductModel.GetProductById(in.ID)
	var productInfoData *pms.ProductListData
	jsonData, _ := json.Marshal(productInfo)
	_ = json.Unmarshal(jsonData, &productInfoData)
	//
	sizeList, _, _ := l.svcCtx.SizeModel.GetSizeList(in.ID)
	var sizeListData []*pms.Size
	jsonData1, _ := json.Marshal(sizeList)
	_ = json.Unmarshal(jsonData1, &sizeListData)
	fmt.Println(sizeListData)
	//
	//
	skuList, _, _ := l.svcCtx.SkuModel.GetSkuList(in.ID)
	var skListData []*pms.SkuListData
	jsonData2, _ := json.Marshal(skuList)
	_ = json.Unmarshal(jsonData2, &skListData)
	//
	//
	return &pms.ProductInfoResp{
		ProductInfo:     productInfoData,
		ImgUrl:          ImgUrl,
		IntroduceImgUrl: IntroduceImgUrl,
		Size:            sizeListData,
		SkuList:         skListData,
	}, nil
}
