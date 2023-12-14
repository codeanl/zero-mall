package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sms/model"

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
	if in.Id > 0 {
		err := l.svcCtx.SubjectProductModel.UpdateSubjectProduct(in.Id, &model.SubjectProduct{
			Sort: in.Sort,
		})
		if err != nil {
			return nil, errors.New("更新失败")
		}
	} else {
		for _, i := range in.ProductId {
			exist, _ := l.svcCtx.SubjectProductModel.GetSubjectProductExist(in.SubjectId, 1)
			if exist {
				return nil, errors.New("该商品已存在")
			}
			err := l.svcCtx.SubjectProductModel.AddSubjectProduct(&model.SubjectProduct{
				SubjectID: in.SubjectId,
				ProductID: i,
				Sort:      in.Sort,
			})
			if err != nil {
				return nil, errors.New("添加失败")
			}
		}
	}
	return &sms.SaveOrUpdateSubjectProductResp{}, nil
}
