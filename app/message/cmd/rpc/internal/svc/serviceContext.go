package svc

import (
	"MyDouSheng/app/message/cmd/rpc/internal/config"
	"MyDouSheng/app/message/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(sqlConn, c.Cache),
	}
}
