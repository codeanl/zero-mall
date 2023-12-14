package member

import (
	"context"
	"simple_mall_new/rpc/ums/umsclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberDeleteLogic) MemberDelete(req *types.DeleteMemberReq) (resp *types.DeleteMemberResp, err error) {
	_, err = l.svcCtx.Ums.MemberDelete(l.ctx, &umsclient.MemberDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return &types.DeleteMemberResp{
			Code:    400,
			Message: "删除失败",
		}, nil
	}
	return &types.DeleteMemberResp{
		Code:    200,
		Message: "删除成功",
	}, nil
}
