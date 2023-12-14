package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductDeleteLogic {
	return &SubjectProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题商品
func (l *SubjectProductDeleteLogic) SubjectProductDelete(in *sms.SubjectProductDeleteAddReq) (*sms.SubjectProductDeleteResp, error) {
	err := l.svcCtx.SubjectProductModel.DeleteSubjectProductByIds(in.IDs)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sms.SubjectProductDeleteResp{}, nil
}
