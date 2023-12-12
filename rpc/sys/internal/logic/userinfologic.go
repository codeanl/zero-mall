package logic

import (
	"context"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户信息
func (l *UserInfoLogic) UserInfo(in *sys.UserInfoReq) (*sys.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &sys.UserInfoResp{}, nil
}
