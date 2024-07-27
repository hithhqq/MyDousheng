package svc

import (
	"MyDouSheng/app/message/cmd/api/internal/config"
	"MyDouSheng/app/message/cmd/api/internal/middleware"
	"MyDouSheng/app/message/cmd/rpc/messageservice"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	MessageServiceRpcConf messageservice.Messageservice
	JwtauthMiddleware     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		JwtauthMiddleware:     middleware.NewJwtauthMiddleware().Handle,
		MessageServiceRpcConf: messageservice.NewMessageservice(zrpc.MustNewClient(c.MessageServiceRpcConf)),
	}
}
