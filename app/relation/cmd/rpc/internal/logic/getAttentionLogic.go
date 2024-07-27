package logic

import (
	"context"
	"database/sql"
	"fmt"

	"MyDouSheng/app/relation/cmd/rpc/internal/svc"
	"MyDouSheng/app/relation/cmd/rpc/pb"
	"MyDouSheng/app/relation/model"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetAttentionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttentionLogic {
	return &GetAttentionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAttentionLogic) GetAttention(in *pb.GetAttentionReq) (*pb.GetAttentionResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("新增关注关系start...\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		var relation model.Relation
		relation.FollowerId = in.TouserId
		relation.FollowingId = in.FromuserId
		if in.ActionType == 1 {
			_, err := l.svcCtx.RelationModel.TxInsert(l.ctx, tx, &relation)
			if err != nil { // 改成xerr
				return fmt.Errorf("insert error :%v, relation:%+v", err, relation)
			}
		} else {
			// l.svcCtx.RelationModel.Delete()
			err := l.svcCtx.RelationModel.DeleteAttention(l.ctx, tx, in.TouserId, in.FromuserId)
			if err != nil {
				return fmt.Errorf("delete error :%v, relation:%+v", err, relation)
			}
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Printf("新增关注关系end...")
	return &pb.GetAttentionResp{}, nil
}
