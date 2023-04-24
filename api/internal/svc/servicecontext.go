package svc

import (
	"go-zero-demo/api/internal/config"
	"go-zero-demo/api/internal/middleware"
	"go-zero-demo/rpc/sys/sys"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Sys      sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
		Sys:      sys.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
