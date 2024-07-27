package logic

import (
	"context"
	"fmt"

	"MyDouSheng/app/identity/cmd/rpc/internal/svc"
	"MyDouSheng/app/identity/cmd/rpc/pb"
	"MyDouSheng/common/globalkey"
	"MyDouSheng/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrClearTokenError = xerr.NewErrMsg("退出token失败")

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清除token，只针对用户服务开放访问
func (l *ClearTokenLogic) ClearToken(in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {
	// todo: add your logic here and delete this line
	userTokenKey := fmt.Sprintf(globalkey.CacheUserTokenKey, in.UserId)
	if _, err := l.svcCtx.RedisClient.DelCtx(l.ctx, userTokenKey); err != nil {
		return nil, errors.Wrapf(ErrClearTokenError, "userId:%d, err:%v", in.UserId, err)
	}
	return &pb.ClearTokenResp{}, nil
}
