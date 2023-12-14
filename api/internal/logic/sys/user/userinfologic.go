package user

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	data, err := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.UserInfoReq{Id: id})
	println(data.Roles[1].Id)

	var UserInfo types.UserInfoData
	jsonData, err := json.Marshal(data)
	err = json.Unmarshal(jsonData, &UserInfo)

	return &types.UserInfoResp{
		Code:    200,
		Data:    UserInfo,
		Message: "查询成功",
	}, nil
}
