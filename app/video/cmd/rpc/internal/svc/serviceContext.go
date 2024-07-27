package svc

import (
	"MyDouSheng/app/video/cmd/rpc/internal/config"
	"MyDouSheng/app/video/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideoModel
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(sqlConn, c.Cache),
		Redis:      redis.New(c.Redis.Host, redis.WithPass(c.Redis.Pass)),
	}
}
