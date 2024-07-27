package logic

import (
	"context"
	"database/sql"
	"fmt"

	"MyDouSheng/app/video/cmd/rpc/internal/svc"
	"MyDouSheng/app/video/cmd/rpc/pb"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PublishVideoRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoRollbackLogic {
	return &PublishVideoRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoRollbackLogic) PublishVideoRollback(in *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("PublishVideoRollback start..\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		err := l.svcCtx.VideoModel.TxDelete(l.ctx, tx, in.Data.Id)
		if err != nil {
			fmt.Printf("err is %v\n", err)
			return fmt.Errorf("delete error :%v, video:%+v", err, in.Data.Id)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Printf("PublishVideoRollback end..\n")
	return &pb.PublishVideoResp{}, nil
}
