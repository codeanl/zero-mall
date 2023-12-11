package logic

import (
	"context"
	"encoding/json"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户列表
func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {
	all, total, err := l.svcCtx.UserModel.GetUserList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.UserList
	jsonData, err := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	//
	for index, i := range list {
		role, _ := l.svcCtx.RoleModel.GetRoleByUserID(i.Id)
		var Roles []string
		for _, item := range role {
			Roles = append(Roles, item.Name)
		}
		list[index].RoleName = Roles
	}
	return &sys.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
