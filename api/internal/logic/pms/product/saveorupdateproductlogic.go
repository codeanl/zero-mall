package product

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"simple_mall_new/rpc/pms/pms"
	"simple_mall_new/rpc/pms/pmsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveOrUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateProductLogic {
	return &SaveOrUpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveOrUpdateProductLogic) SaveOrUpdateProduct(req *types.SaveOrUpdateProductReq) (resp *types.SaveOrUpdateProductResp, err error) {
	var size []*pms.Size
	jsonData, _ := json.Marshal(req.Size)
	err = json.Unmarshal(jsonData, &size)
	//
	data := &pmsclient.SaveOrUpdateProductReq{
		Id:              req.ID,
		CategoryId:      req.CategoryID,
		MerchantId:      req.MerchantID,
		Name:            req.Name,
		Pic:             req.Pic,
		ProductSn:       req.ProductSn,
		Desc:            req.Desc,
		OriginalPrices:  req.OriginalPrice,
		Price:           req.Price,
		IsPublished:     req.IsPublished,
		IsHot:           req.IsHot,
		Size:            size,
		ImgUrl:          req.ImgUrl,
		IntroduceImgUrl: req.IntroduceImgUrl,
	}
	msg := "添加成功"
	if req.ID > 0 {
		msg = "更新成功"
	}
	//
	_, err = l.svcCtx.Pms.SaveOrUpdateProduct(l.ctx, data)
	if err != nil {
		// 解析 gRPC 错误信息
		grpcErr, ok := status.FromError(err)
		if ok {
			return &types.SaveOrUpdateProductResp{
				Code:    400,
				Message: grpcErr.Message(),
			}, nil
		} else {
		}
	}
	//
	return &types.SaveOrUpdateProductResp{
		Code:    200,
		Message: msg,
	}, nil
}
