package logic

import (
	"context"
	"encoding/json"
	"fmt"
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
	//查询用户信息
	user, _ := l.svcCtx.UserModel.GetUserByID(in.Id)

	var UserInfo *sys.User
	jsonData, err := json.Marshal(user)
	err = json.Unmarshal(jsonData, &UserInfo)

	role, err := l.svcCtx.RoleModel.GetRoleByUserID(in.Id)
	if err != nil {
		return nil, err
	}
	var Roles []*sys.RoleListData
	jsonData1, err := json.Marshal(role)
	err = json.Unmarshal(jsonData1, &Roles)
	fmt.Println(Roles)
	return &sys.UserInfoResp{
		User:  UserInfo,
		Roles: Roles,
	}, nil
}
