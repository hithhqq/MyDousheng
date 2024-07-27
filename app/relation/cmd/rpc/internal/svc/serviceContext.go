package svc

import (
	"MyDouSheng/app/relation/cmd/rpc/internal/config"
	"MyDouSheng/app/relation/model"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel model.RelationModel
	UserRpc       userservice.Userservice
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		RelationModel: model.NewRelationModel(sqlConn, c.Cache),
	}
}
