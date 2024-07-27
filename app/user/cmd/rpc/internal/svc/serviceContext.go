package svc

import (
	"MyDouSheng/app/identity/cmd/rpc/identity"
	"MyDouSheng/app/user/cmd/rpc/internal/config"
	"MyDouSheng/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	IdentifyRpc identity.Identity
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		UserModel:   model.NewUserModel(sqlConn, c.Cache),
		IdentifyRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),
	}
}
