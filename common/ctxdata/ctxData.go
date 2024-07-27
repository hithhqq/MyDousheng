package ctxdata

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if uidVal, ok := ctx.Value(CtxKeyJwtUserId).(int64); ok {
		uid = uidVal
	} else {
		logx.WithContext(ctx).Errorf("failed to get uid from context")
	}
	return uid
}
