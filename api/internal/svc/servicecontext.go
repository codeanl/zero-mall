package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"simple_mall_new/api/internal/config"
	"simple_mall_new/rpc/sys/sysclient"
)

type ServiceContext struct {
	Config config.Config
	Sys    sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Sys:    sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
