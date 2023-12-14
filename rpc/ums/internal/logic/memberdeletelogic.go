package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/ums/internal/svc"
	"simple_mall_new/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除会员
func (l *MemberDeleteLogic) MemberDelete(in *ums.MemberDeleteReq) (*ums.MemberDeleteResp, error) {
	err := l.svcCtx.MemberModel.DeleteByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &ums.MemberDeleteResp{}, nil
}
