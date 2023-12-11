package user

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/sysclient"

	"simple_mall_new/api/internal/svc"
	"simple_mall_new/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO 需要修改
func (l *UserListLogic) UserList(req *types.ListUserReq) (resp *types.ListUserResp, err error) {
	data, err := l.svcCtx.Sys.UserList(l.ctx, &sysclient.UserListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Username: req.Username,
		Status:   req.Status,
		Gender:   req.Gander,
		Email:    req.Email,
	})
	if err != nil {
		return &types.ListUserResp{
			Code:    400,
			Message: "查询失败",
		}, nil
	}
	var list []*types.ListUser
	jsonData, err := json.Marshal(data.List)
	err = json.Unmarshal(jsonData, &list)
	//
	return &types.ListUserResp{
		Code:    200,
		Message: "查询成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
