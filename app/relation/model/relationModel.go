package model

import (
	"MyDouSheng/app/user/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationModel = (*customRelationModel)(nil)

type (
	// RelationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationModel.
	RelationModel interface {
		relationModel
		FindId(ctx context.Context, follower_id int64, following_id int64) (int64, error)
		DeleteAttention(ctx context.Context, tx *sql.Tx, follower_id, following_id int64) error
		TxInsert(ctx context.Context, tx *sql.Tx, data *Relation) (sql.Result, error)
		FindFollowlist(ctx context.Context, userid int64) ([]model.User, error)
		FindFanslist(ctx context.Context, userid int64) ([]model.User, error)
	}

	customRelationModel struct {
		*defaultRelationModel
	}
)

// NewRelationModel returns a model for the database table.
func NewRelationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RelationModel {
	return &customRelationModel{
		defaultRelationModel: newRelationModel(conn, c, opts...),
	}
}

func (m *defaultRelationModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Relation) (sql.Result, error) {
	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, data.FollowerId, data.FollowingId)
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, data.Id)
	resp, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, relationRowsExpectAutoSet)
		return tx.ExecContext(ctx, query, data.FollowerId, data.FollowingId)
	}, doushengRelationIdKey, doushengRelationFollowerIdFollowingIdKey)
	return resp, err
}

func (m *defaultRelationModel) DeleteAttention(ctx context.Context, tx *sql.Tx, follower_id, following_id int64) error {
	id, err := m.FindId(ctx, follower_id, following_id)
	if err != nil {
		return err
	}
	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, follower_id, following_id)
	doushengRelationIdKey := fmt.Sprintf("%s%v", cacheDoushengRelationIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return tx.ExecContext(ctx, query, id)
	}, doushengRelationIdKey, doushengRelationFollowerIdFollowingIdKey)
	return err
}

func (m *defaultRelationModel) FindId(ctx context.Context, follower_id int64, following_id int64) (int64, error) {
	doushengRelationFollowerIdFollowingIdKey := fmt.Sprintf("%s%v:%v", cacheDoushengRelationFollowerIdFollowingIdPrefix, follower_id, following_id)
	var resp Relation
	err := m.QueryRowCtx(ctx, &resp, doushengRelationFollowerIdFollowingIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `follower_id` = ? and `following_id` = ?", relationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, follower_id, following_id)
	})
	switch err {
	case nil:
		return resp.Id, nil
	case sqlc.ErrNotFound:
		return -1, ErrNotFound
	default:
		return -1, err
	}
}

func (m *defaultRelationModel) FindFollowlist(ctx context.Context, userid int64) ([]model.User, error) {
	var resp []model.User
	err := m.QueryRowsNoCacheCtx(ctx, &resp, fmt.Sprintf("SELECT %s FROM %s u JOIN %s r ON u.userid = r.follower_id WHERE r.following_id = ?", model.UserRows, "`user`", m.table), userid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRelationModel) FindFanslist(ctx context.Context, userid int64) ([]model.User, error) {
	var resp []model.User
	err := m.QueryRowsNoCacheCtx(ctx, &resp, fmt.Sprintf("SELECT %s FROM %s u JOIN %s r ON u.userid = r.following_id WHERE r.follower_id = ?", model.UserRows, "`user`", m.table), userid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
