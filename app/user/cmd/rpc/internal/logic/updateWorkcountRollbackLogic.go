package logic

import (
	"context"
	"database/sql"
	"fmt"

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

type UpdateWorkcountRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWorkcountRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkcountRollbackLogic {
	return &UpdateWorkcountRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWorkcountRollbackLogic) UpdateWorkcountRollback(in *pb.UpdateWorkCountReq) (*pb.UpdateWorkCountResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("UpdateWorkcountRollback start..\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		logx.Error("findone user err:"+err.Error()+", user is %+v", err, user)
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	user.Workcount -= 1
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "User Database Exception user : %+v, err :%v", user, err)
		}
		if err := l.svcCtx.UserModel.TxUpdate(l.ctx, tx, user); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "User Database Exception user : %+v, err :%v", user, err)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Printf("UpdateWorkcountRollback end..\n")
	return &pb.UpdateWorkCountResp{}, nil
}
