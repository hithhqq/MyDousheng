package svc

import (
	"MyDouSheng/app/user/cmd/api/internal/config"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserserviceRpc userservice.Userservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserserviceRpc: userservice.NewUserservice(zrpc.MustNewClient(c.UserserviceRpcConf)),
	}
}
