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
	LoginReq             = sys.LoginReq
	LoginResp            = sys.LoginResp
	RoleDeleteReq        = sys.RoleDeleteReq
	RoleDeleteResp       = sys.RoleDeleteResp
	RoleListData         = sys.RoleListData
	RoleListReq          = sys.RoleListReq
	RoleListResp         = sys.RoleListResp
	SaveOrUpdateRoleReq  = sys.SaveOrUpdateRoleReq
	SaveOrUpdateRoleResp = sys.SaveOrUpdateRoleResp

	Sys interface {
		// UserLogin 用户登录
		UserLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// SaveOrUpdateRole 添加｜｜更新角色
		SaveOrUpdateRole(ctx context.Context, in *SaveOrUpdateRoleReq, opts ...grpc.CallOption) (*SaveOrUpdateRoleResp, error)
		// RoleDelete 删除角色
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		// RoleList 角色列表
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
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