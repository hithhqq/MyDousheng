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
	relationFieldNames          = builder.RawFieldNames(&Relation{})
	relationRows                = strings.Join(relationFieldNames, ",")
	relationRowsExpectAutoSet   = strings.Join(stringx.Remove(relationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	relationRowsWithPlaceHolder = strings.Join(stringx.Remove(relationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDoushengRelationIdPrefix                    = "cache:dousheng:relation:id:"
	cacheDoushengRelationFollowerIdFollowingIdPrefix = "cache:dousheng:relation:followerId:followingId:"
)

type (
	relationModel interface {
		Insert(ctx context.Context, data *Relation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Relation, error)
		FindOneByFollowerIdFollowingId(ctx context.Context, followerId int64, followingId int64) (*Relation, error)
		Update(ctx context.Context, data *Relation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRelationModel struct {
		sqlc.CachedConn
		table string
	}

	Relation struct {
		Id          int64     `db:"id"`
		FollowerId  int64     `db:"follower_id"`  // 被关注的人
		FollowingId int64     `db:"following_id"` // 发起关注的人
		CreateAt    time.Time `db:"create_at"`
	}
)

func newRelationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRelationModel {
	return &defaultRelationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`relation`",
	}
}

func (m *defaultRelationModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, data.FollowerId, data.FollowingId)
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, doushengRelationFollowerIdFollowingIdKey, doushengRelationIdKey)
	return err
}

func (m *defaultRelationModel) FindOne(ctx context.Context, id int64) (*Relation, error) {
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, id)
	var resp Relation
	err := m.QueryRowCtx(ctx, &resp, doushengRelationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultRelationModel) FindOneByFollowerIdFollowingId(ctx context.Context, followerId int64, followingId int64) (*Relation, error) {
	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, followerId, followingId)
	var resp Relation
	err := m.QueryRowIndexCtx(ctx, &resp, doushengRelationFollowerIdFollowingIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `follower_id` = ? and `following_id` = ? limit 1", relationRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, followerId, followingId); err != nil {
			return nil, err
		}
		return resp.Id, nil
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

func (m *defaultRelationModel) Insert(ctx context.Context, data *Relation) (sql.Result, error) {
	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, data.FollowerId, data.FollowingId)
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, relationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.FollowerId, data.FollowingId)
	}, doushengRelationFollowerIdFollowingIdKey, doushengRelationIdKey)
	return ret, err
}

func (m *defaultRelationModel) Update(ctx context.Context, newData *Relation) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, data.FollowerId, data.FollowingId)
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, relationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.FollowerId, newData.FollowingId, newData.Id)
	}, doushengRelationFollowerIdFollowingIdKey, doushengRelationIdKey)
	return err
}

func (m *defaultRelationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, primary)
}

func (m *defaultRelationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", relationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRelationModel) tableName() string {
	return m.table
}
