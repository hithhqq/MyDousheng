// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`userid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`userid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDoushengUserUseridPrefix   = "cache:dousheng:user:userid:"
	cacheDoushengUserUsernamePrefix = "cache:dousheng:user:username:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, userid int64) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, userid int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Userid          int64     `db:"userid"`
		Username        string    `db:"username"`
		FollowingCount  int64     `db:"following_count"`
		FollowerCount   int64     `db:"follower_count"`
		Password        string    `db:"password"`
		Avatar          string    `db:"avatar"`
		Signature       string    `db:"signature"`
		TotalFavorited  int64     `db:"total_favorited"`
		Workcount       int64     `db:"workcount"`
		FavoriteCount   int64     `db:"favorite_count"`
		Createtime      time.Time `db:"createtime"`
		BackgroundImage string    `db:"background_image"`
		DeleteAt        time.Time `db:"delete_at"`
		DelState        int64     `db:"del_state"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, userid int64) error {
	data, err := m.FindOne(ctx, userid)
	if err != nil {
		return err
	}

	doushengUserUseridKey := fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, userid)
	doushengUserUsernameKey := fmt.Sprintf("%s%v", cacheDoushengUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `userid` = ?", m.table)
		return conn.ExecCtx(ctx, query, userid)
	}, doushengUserUseridKey, doushengUserUsernameKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, userid int64) (*User, error) {
	doushengUserUseridKey := fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, userid)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, doushengUserUseridKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	doushengUserUsernameKey := fmt.Sprintf("%s%v", cacheDoushengUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, doushengUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Userid, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	doushengUserUseridKey := fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, data.Userid)
	doushengUserUsernameKey := fmt.Sprintf("%s%v", cacheDoushengUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.FollowingCount, data.FollowerCount, data.Password, data.Avatar, data.Signature, data.TotalFavorited, data.Workcount, data.FavoriteCount, data.Createtime, data.BackgroundImage, data.DeleteAt, data.DelState)
	}, doushengUserUseridKey, doushengUserUsernameKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Userid)
	if err != nil {
		return err
	}

	doushengUserUseridKey := fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, data.Userid)
	doushengUserUsernameKey := fmt.Sprintf("%s%v", cacheDoushengUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `userid` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Username, newData.FollowingCount, newData.FollowerCount, newData.Password, newData.Avatar, newData.Signature, newData.TotalFavorited, newData.Workcount, newData.FavoriteCount, newData.Createtime, newData.BackgroundImage, newData.DeleteAt, newData.DelState, newData.Userid)
	}, doushengUserUseridKey, doushengUserUsernameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDoushengUserUseridPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
