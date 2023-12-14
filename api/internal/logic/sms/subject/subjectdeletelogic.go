package subject

import (
	"context"
	"simple_mall_new/rpc/sms/smsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectDeleteLogic {
	return &SubjectDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubjectDeleteLogic) SubjectDelete(req *types.DeleteSubjectReq) (resp *types.DeleteSubjectResp, err error) {
	_, err = l.svcCtx.Sms.SubjectDelete(l.ctx, &smsclient.SubjectDeleteAddReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteSubjectResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteSubjectResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
