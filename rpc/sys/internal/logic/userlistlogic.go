package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"
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
	jsonData, _ := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)
	for index, i := range all {
		list[index].CreatedAt = i.CreatedAt.Format("2006-01-02 15:04:05")
	}
	fmt.Println(list[0].CreatedAt)
	return &sys.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
