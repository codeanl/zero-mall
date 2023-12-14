package logic

import (
	"context"

	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateSubjectProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateSubjectProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateSubjectProductLogic {
	return &SaveOrUpdateSubjectProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOrUpdateSubjectProduct 添加｜｜更新专题商品
func (l *SaveOrUpdateSubjectProductLogic) SaveOrUpdateSubjectProduct(in *sms.SaveOrUpdateSubjectProductReq) (*sms.SaveOrUpdateSubjectProductResp, error) {
	// todo: add your logic here and delete this line

	return &sms.SaveOrUpdateSubjectProductResp{}, nil
}
