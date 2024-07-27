package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		FindMessages(ctx context.Context, fromUserid int64, toUserid int64) ([]Message, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c, opts...),
	}
}

func (m *defaultMessageModel) FindMessages(ctx context.Context, fromUserid int64, toUserid int64) ([]Message, error) {
	var messages []Message
	err := m.QueryRowsNoCacheCtx(ctx, &messages, fmt.Sprintf("SELECT %s FROM %s WHERE (from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)", messageRows, m.table), fromUserid, toUserid, toUserid, fromUserid)
	switch err {
	case nil:
		return messages, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
