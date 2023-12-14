package logic

import (
	"context"
	"errors"

	"simple_mall_new/rpc/sms/internal/svc"
	"simple_mall_new/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseDeleteLogic {
	return &HomeAdvertiseDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除广告
func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(in *sms.HomeAdvertiseDeleteReq) (*sms.HomeAdvertiseDeleteResp, error) {
	err := l.svcCtx.HomeAdvertiseModel.DeleteHomeAdvertiseByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &sms.HomeAdvertiseDeleteResp{}, nil
}
