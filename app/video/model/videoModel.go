package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		TxInsert(ctx context.Context, tx *sql.Tx, data *Video) (sql.Result, error)
		TxDelete(ctx context.Context, tx *sql.Tx, videoId string) error
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c, opts...),
	}
}

func (m *defaultVideoModel) TxInsert(ctx context.Context, tx *sql.Tx, data *Video) (sql.Result, error) {
	doushengVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDoushengVideoVideoIdPrefix, data.VideoId)
	resp, err := m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return tx.ExecContext(ctx, query, data.VideoId, data.UserId, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount, data.Title)
	}, doushengVideoVideoIdKey)
	return resp, err
}

func (m *defaultVideoModel) TxDelete(ctx context.Context, tx *sql.Tx, videoId string) error {
	doushengVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDoushengVideoVideoIdPrefix, videoId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `video_id` = ?", m.table)
		return tx.ExecContext(ctx, query, videoId)
	}, doushengVideoVideoIdKey)
	return err
}
