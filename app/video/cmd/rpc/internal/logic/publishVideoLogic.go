package logic

import (
	"context"
	"database/sql"
	"fmt"

	"MyDouSheng/app/video/cmd/rpc/internal/svc"
	"MyDouSheng/app/video/cmd/rpc/pb"
	"MyDouSheng/app/video/model"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var CtxKeyVideoId = "CtxVideoId"

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("PublishVideo start..\n")
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		var video model.Video
		video.VideoId = in.Data.Id
		video.UserId = in.Data.Author.UserId
		video.PlayUrl = in.Data.PlayUrl
		video.CoverUrl = in.Data.CoverUrl
		video.CommentCount = 0
		video.FavoriteCount = 0
		video.Title = in.Data.Title
		_, err := l.svcCtx.VideoModel.TxInsert(l.ctx, tx, &video)
		if err != nil {
			return fmt.Errorf("insert error :%v, video:%+v", err, video)
		}
		return nil
	}); err != nil {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	fmt.Printf("publisVideo end..\n")
	return &pb.PublishVideoResp{}, nil
}
