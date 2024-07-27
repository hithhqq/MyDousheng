package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)
var UserRows = userRows

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		TxUpdate(ctx context.Context, tx *sql.Tx, newData *User) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *defaultUserModel) TxUpdate(ctx context.Context, tx *sql.Tx, newData *User) error {
	data, err := m.FindOne(ctx, newData.Userid)
	if err != nil {
		return err
	}
	doushengUserUseridKey := fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, data.Userid)
	doushengUserUsernameKey := fmt.Sprintf("%s%v", cacheDoushengUserUsernamePrefix, data.Username)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `userid` = ?", m.table, userRowsWithPlaceHolder)
		return tx.ExecContext(ctx, query, newData.Username, newData.FollowingCount, newData.FollowerCount, newData.Password, newData.Avatar, newData.Signature, newData.TotalFavorited, newData.Workcount, newData.FavoriteCount, newData.Createtime, newData.BackgroundImage, newData.DeleteAt, newData.DelState, newData.Userid)
	}, doushengUserUseridKey, doushengUserUsernameKey)
	return err
}
