package logic

import (
	"context"
	"database/sql"
	"fmt"

	"MyDouSheng/app/relation/cmd/rpc/internal/svc"
	"MyDouSheng/app/relation/cmd/rpc/pb"
	"MyDouSheng/app/relation/model"
	"MyDouSheng/common/xerr"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetAttentionRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttentionRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttentionRollbackLogic {
	return &GetAttentionRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttentionRollbackLogic) GetAttentionRollback(in *pb.GetAttentionReq) (*pb.GetAttentionResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("关注关系回滚start...\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DTM_ERROR), err.Error())
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		var relation model.Relation
		relation.FollowerId = in.TouserId
		relation.FollowingId = in.FromuserId
		if in.ActionType == 1 {
			err := l.svcCtx.RelationModel.DeleteAttention(l.ctx, tx, in.TouserId, in.FromuserId)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Relation Database Exception releation : %+v, err :%v", relation, err)
			}
		} else {
			_, err := l.svcCtx.RelationModel.TxInsert(l.ctx, tx, &relation)
			if err != nil { // 改成xerr
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Relation Database Exception releation : %+v, err: %v", relation, err)
			}
			// l.svcCtx.RelationModel.Delete()
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	return &pb.GetAttentionResp{}, nil
}
