// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"simple_mall_new/rpc/sys/internal/logic"
	"simple_mall_new/rpc/sys/internal/svc"
	"simple_mall_new/rpc/sys/sys"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

// UserLogin 用户登录
func (s *SysServer) UserLogin(ctx context.Context, in *sys.LoginReq) (*sys.LoginResp, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

// SaveOrUpdateRole 添加｜｜更新角色
func (s *SysServer) SaveOrUpdateRole(ctx context.Context, in *sys.SaveOrUpdateRoleReq) (*sys.SaveOrUpdateRoleResp, error) {
	l := logic.NewSaveOrUpdateRoleLogic(ctx, s.svcCtx)
	return l.SaveOrUpdateRole(in)
}

// RoleDelete 删除角色
func (s *SysServer) RoleDelete(ctx context.Context, in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	l := logic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

// RoleList 角色列表
func (s *SysServer) RoleList(ctx context.Context, in *sys.RoleListReq) (*sys.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

// SaveOrUpdateMenu 添加｜｜更新菜单
func (s *SysServer) SaveOrUpdateMenu(ctx context.Context, in *sys.SaveOrUpdateMenuReq) (*sys.SaveOrUpdateMenuResp, error) {
	l := logic.NewSaveOrUpdateMenuLogic(ctx, s.svcCtx)
	return l.SaveOrUpdateMenu(in)
}

// 菜单列表
func (s *SysServer) MenuList(ctx context.Context, in *sys.MenuListReq) (*sys.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

// 删除菜单
func (s *SysServer) MenuDelete(ctx context.Context, in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}
