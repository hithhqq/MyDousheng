package svc

import (
	"MyDouSheng/app/relation/cmd/api/internal/config"
	"MyDouSheng/app/relation/cmd/rpc/relationservice"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	RelationserviceRpc relationservice.Relationservice
	UserserviceRpc     userservice.Userservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		RelationserviceRpc: relationservice.NewRelationservice(zrpc.MustNewClient(c.RelationserviceRpcConf)),
		UserserviceRpc:     userservice.NewUserservice(zrpc.MustNewClient(c.UserServiceRpcConf)),
	}
}
