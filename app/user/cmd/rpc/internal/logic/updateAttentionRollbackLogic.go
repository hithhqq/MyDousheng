package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"MyDouSheng/app/user/cmd/rpc/internal/svc"
	"MyDouSheng/app/user/cmd/rpc/pb"
	"MyDouSheng/common/xerr"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateAttentionRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAttentionRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAttentionRollbackLogic {
	return &UpdateAttentionRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAttentionRollbackLogic) UpdateAttentionRollback(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("关注数回滚start...\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DTM_ERROR), err.Error())
	}
	following_user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.FollowingId)
	if err != nil {
		logx.Error("findone user err:"+err.Error()+", user is %+v", err, following_user)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	follower_user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.FollowerId)
	if err != nil {
		logx.Error("findone user err:"+err.Error()+", user is %+v", err, follower_user)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if in.Type == 1 {
		if following_user.FollowingCount > 0 {
			following_user.FollowingCount -= 1
		}
		if follower_user.FollowerCount > 0 {
			follower_user.FollowerCount -= 1
		}
	} else {
		following_user.FollowingCount += 1
		follower_user.FollowerCount += 1
	}
	following_user.DeleteAt = time.Now()
	follower_user.DeleteAt = time.Now()
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "User Database Exception user : %+v, err :%v", following_user, err)
		}
		if err := l.svcCtx.UserModel.TxUpdate(l.ctx, tx, following_user); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "User Database Exception user : %+v, err :%v", following_user, err)
		}
		if err := l.svcCtx.UserModel.TxUpdate(l.ctx, tx, follower_user); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "User Database Exception user : %+v, err :%v", follower_user, err)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Printf("关注数回滚end...\n")
	return &pb.UpdateUserResp{}, nil
}
