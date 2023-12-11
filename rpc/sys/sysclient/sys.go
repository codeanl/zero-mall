// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package sysclient

import (
	"context"

	"simple_mall_new/rpc/sys/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginLogDeleteReq      = sys.LoginLogDeleteReq
	LoginLogDeleteResp     = sys.LoginLogDeleteResp
	LoginLogListData       = sys.LoginLogListData
	LoginLogListReq        = sys.LoginLogListReq
	LoginLogListResp       = sys.LoginLogListResp
	LoginReq               = sys.LoginReq
	LoginResp              = sys.LoginResp
	MenuDeleteReq          = sys.MenuDeleteReq
	MenuDeleteResp         = sys.MenuDeleteResp
	MenuList               = sys.MenuList
	MenuListReq            = sys.MenuListReq
	MenuListResp           = sys.MenuListResp
	OperationLogAddReq     = sys.OperationLogAddReq
	OperationLogAddResp    = sys.OperationLogAddResp
	OperationLogDeleteReq  = sys.OperationLogDeleteReq
	OperationLogDeleteResp = sys.OperationLogDeleteResp
	OperationLogListReq    = sys.OperationLogListReq
	OperationLogListResp   = sys.OperationLogListResp
	RoleDeleteReq          = sys.RoleDeleteReq
	RoleDeleteResp         = sys.RoleDeleteResp
	RoleListData           = sys.RoleListData
	RoleListReq            = sys.RoleListReq
	RoleListResp           = sys.RoleListResp
	SaveOrUpdateMenuReq    = sys.SaveOrUpdateMenuReq
	SaveOrUpdateMenuResp   = sys.SaveOrUpdateMenuResp
	SaveOrUpdateRoleReq    = sys.SaveOrUpdateRoleReq
	SaveOrUpdateRoleResp   = sys.SaveOrUpdateRoleResp
	SysLogListData         = sys.SysLogListData
	User                   = sys.User

	Sys interface {
		// UserLogin 用户登录
		UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// SaveOrUpdateRole 添加｜｜更新角色
		SaveOrUpdateRole(ctx context.Context, in *SaveOrUpdateRoleReq, opts ...grpc.CallOption) (*SaveOrUpdateRoleResp, error)
		// RoleDelete 删除角色
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		// RoleList 角色列表
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		// SaveOrUpdateMenu 添加｜｜更新菜单
		SaveOrUpdateMenu(ctx context.Context, in *SaveOrUpdateMenuReq, opts ...grpc.CallOption) (*SaveOrUpdateMenuResp, error)
		// MenuList 菜单列表
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		// MenuDelete 删除菜单
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
		// LoginLogList 登录日志列表
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		// LoginLogDelete 删除登录日志
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		// 添加日志
		OperationLogAdd(ctx context.Context, in *OperationLogAddReq, opts ...grpc.CallOption) (*OperationLogAddResp, error)
		// 日志列表
		OperationLogList(ctx context.Context, in *OperationLogListReq, opts ...grpc.CallOption) (*OperationLogListResp, error)
		// 删除日志
		OperationLogDelete(ctx context.Context, in *OperationLogDeleteReq, opts ...grpc.CallOption) (*OperationLogDeleteResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

// UserLogin 用户登录
func (m *defaultSys) UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// SaveOrUpdateRole 添加｜｜更新角色
func (m *defaultSys) SaveOrUpdateRole(ctx context.Context, in *SaveOrUpdateRoleReq, opts ...grpc.CallOption) (*SaveOrUpdateRoleResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SaveOrUpdateRole(ctx, in, opts...)
}

// RoleDelete 删除角色
func (m *defaultSys) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

// RoleList 角色列表
func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

// SaveOrUpdateMenu 添加｜｜更新菜单
func (m *defaultSys) SaveOrUpdateMenu(ctx context.Context, in *SaveOrUpdateMenuReq, opts ...grpc.CallOption) (*SaveOrUpdateMenuResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.SaveOrUpdateMenu(ctx, in, opts...)
}

// MenuList 菜单列表
func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

// MenuDelete 删除菜单
func (m *defaultSys) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

// LoginLogList 登录日志列表
func (m *defaultSys) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

// LoginLogDelete 删除登录日志
func (m *defaultSys) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

// 添加日志
func (m *defaultSys) OperationLogAdd(ctx context.Context, in *OperationLogAddReq, opts ...grpc.CallOption) (*OperationLogAddResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.OperationLogAdd(ctx, in, opts...)
}

// 日志列表
func (m *defaultSys) OperationLogList(ctx context.Context, in *OperationLogListReq, opts ...grpc.CallOption) (*OperationLogListResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.OperationLogList(ctx, in, opts...)
}

// 删除日志
func (m *defaultSys) OperationLogDelete(ctx context.Context, in *OperationLogDeleteReq, opts ...grpc.CallOption) (*OperationLogDeleteResp, error) {
	client := sys.NewSysClient(m.cli.Conn())
	return client.OperationLogDelete(ctx, in, opts...)
}
