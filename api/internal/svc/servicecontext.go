package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"simple_mall_new/api/internal/config"
	"simple_mall_new/api/middleware"
	"simple_mall_new/rpc/sys/sysclient"
	"simple_mall_new/rpc/ums/umsclient"
)

type ServiceContext struct {
	Config          config.Config
	Sys             sysclient.Sys
	Ums             umsclient.Ums
	AddOperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newSys := sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))
	newUms := umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc))
	return &ServiceContext{
		Config:          c,
		Sys:             newSys,
		Ums:             newUms,
		AddOperationLog: middleware.NewAddLogMiddleware(newSys).Handle,
	}
}
