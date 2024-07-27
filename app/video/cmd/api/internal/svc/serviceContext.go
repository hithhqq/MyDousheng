package svc

import (
	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/app/video/cmd/api/internal/config"
	"MyDouSheng/app/video/cmd/rpc/videoservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	UserserviceRpc  userservice.Userservice
	VideoserviceRpc videoservice.Videoservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserserviceRpc:  userservice.NewUserservice(zrpc.MustNewClient(c.UserServiceRpcConf)),
		VideoserviceRpc: videoservice.NewVideoservice(zrpc.MustNewClient(c.VideoserviceRpcConf)),
	}
}
