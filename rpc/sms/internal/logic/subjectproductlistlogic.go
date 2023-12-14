package logic

import (
	"context"

	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductListLogic {
	return &SubjectProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 专题列表商品 todo
func (l *SubjectProductListLogic) SubjectProductList(in *sms.SubjectProductListReq) (*sms.SubjectProductListResp, error) {

	return &sms.SubjectProductListResp{}, nil
}
