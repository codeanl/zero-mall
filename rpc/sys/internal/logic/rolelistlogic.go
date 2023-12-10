package logic

import (
	"context"
	"encoding/json"
	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleList 角色列表
func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	//获取角色列表
	all, total, err := l.svcCtx.RoleModel.GetRoleList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.RoleListData

	jsonData, err := json.Marshal(all)
	err = json.Unmarshal(jsonData, &list)

	return &sys.RoleListResp{
		Total: total,
		List:  list,
	}, nil
}
