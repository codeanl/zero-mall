package logic

import (
	"context"
	"errors"
	"simple_mall_new/rpc/sys/model"

	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveOrUpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveOrUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveOrUpdateRoleLogic {
	return &SaveOrUpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveOrUpdateRole 添加｜｜更新角色
func (l *SaveOrUpdateRoleLogic) SaveOrUpdateRole(in *sys.SaveOrUpdateRoleReq) (*sys.SaveOrUpdateRoleResp, error) {
	println(in.Id)
	role, exist, _ := l.svcCtx.RoleModel.GetRoleByName(in.Name)
	if exist && int64(role.ID) != in.Id {
		return nil, errors.New("角色名已存在")
	}
	err := l.svcCtx.RoleModel.SaveOrUpdateRole(in.Id, &model.Role{
		Name:     in.Name,
		Remark:   in.Remark,
		CreateBy: in.CreateBy,
		UpdateBy: in.UpdateBy,
	})
	if err != nil {
		msgErr := "添加失败"
		if in.Id > 0 {
			msgErr = "更新失败"
		}
		return nil, errors.New(msgErr)
	}
	return &sys.SaveOrUpdateRoleResp{}, nil
}
