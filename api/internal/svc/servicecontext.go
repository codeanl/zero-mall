package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"simple_mall_new/api/internal/config"
	"simple_mall_new/api/middleware"
	"simple_mall_new/rpc/sys/sysclient"
)

type ServiceContext struct {
	Config          config.Config
	Sys             sysclient.Sys
	AddOperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newSys := sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))
	return &ServiceContext{
		Config:          c,
		Sys:             newSys,
		AddOperationLog: middleware.NewAddLogMiddleware(newSys).Handle,
	}
}
