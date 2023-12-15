package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"simple_mall_new/api/internal/config"
	"simple_mall_new/api/middleware"
	"simple_mall_new/rpc/pms/pmsclient"
	"simple_mall_new/rpc/sms/smsclient"
	"simple_mall_new/rpc/sys/sysclient"
	"simple_mall_new/rpc/ums/umsclient"
)

type ServiceContext struct {
	Config          config.Config
	Sys             sysclient.Sys
	Ums             umsclient.Ums
	Sms             smsclient.Sms
	Pms             pmsclient.Pms
	AddOperationLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newSys := sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))
	newUms := umsclient.NewUms(zrpc.MustNewClient(c.UmsRpc))
	newSms := smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc))
	newPms := pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc))
	return &ServiceContext{
		Config:          c,
		Sys:             newSys,
		Ums:             newUms,
		Sms:             newSms,
		Pms:             newPms,
		AddOperationLog: middleware.NewAddLogMiddleware(newSys).Handle,
	}
}
