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

// SaveOrUpdateUser添加｜｜更新用户
func (s *SysServer) SaveOrUpdateUser(ctx context.Context, in *sys.SaveOrUpdateUserReq) (*sys.SaveOrUpdateUserResp, error) {
	l := logic.NewSaveOrUpdateUserLogic(ctx, s.svcCtx)
	return l.SaveOrUpdateUser(in)
}

// 删除用户
func (s *SysServer) UserDelete(ctx context.Context, in *sys.UserDeleteReq) (*sys.UserDeleteResp, error) {
	l := logic.NewUserDeleteLogic(ctx, s.svcCtx)
	return l.UserDelete(in)
}

// 用户列表
func (s *SysServer) UserList(ctx context.Context, in *sys.UserListReq) (*sys.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
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

// MenuList 菜单列表
func (s *SysServer) MenuList(ctx context.Context, in *sys.MenuListReq) (*sys.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

// MenuDelete 删除菜单
func (s *SysServer) MenuDelete(ctx context.Context, in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	l := logic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}

// LoginLogList 登录日志列表
func (s *SysServer) LoginLogList(ctx context.Context, in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

// LoginLogDelete 删除登录日志
func (s *SysServer) LoginLogDelete(ctx context.Context, in *sys.LoginLogDeleteReq) (*sys.LoginLogDeleteResp, error) {
	l := logic.NewLoginLogDeleteLogic(ctx, s.svcCtx)
	return l.LoginLogDelete(in)
}

// 添加日志
func (s *SysServer) OperationLogAdd(ctx context.Context, in *sys.OperationLogAddReq) (*sys.OperationLogAddResp, error) {
	l := logic.NewOperationLogAddLogic(ctx, s.svcCtx)
	return l.OperationLogAdd(in)
}

// 日志列表
func (s *SysServer) OperationLogList(ctx context.Context, in *sys.OperationLogListReq) (*sys.OperationLogListResp, error) {
	l := logic.NewOperationLogListLogic(ctx, s.svcCtx)
	return l.OperationLogList(in)
}

// 删除日志
func (s *SysServer) OperationLogDelete(ctx context.Context, in *sys.OperationLogDeleteReq) (*sys.OperationLogDeleteResp, error) {
	l := logic.NewOperationLogDeleteLogic(ctx, s.svcCtx)
	return l.OperationLogDelete(in)
}
