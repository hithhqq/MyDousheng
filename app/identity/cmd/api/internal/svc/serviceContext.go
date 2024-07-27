package svc

import (
	"MyDouSheng/app/identity/cmd/api/internal/config"
	"MyDouSheng/app/identity/cmd/rpc/identity"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	IdentityRpc identity.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentifyRpcConf)),
	}
}
