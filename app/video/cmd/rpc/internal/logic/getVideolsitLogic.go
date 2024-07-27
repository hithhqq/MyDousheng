package logic

import (
	"context"
	"strconv"
	"time"

	"MyDouSheng/app/video/cmd/rpc/internal/svc"
	"MyDouSheng/app/video/cmd/rpc/pb"
	"MyDouSheng/app/video/model"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultPageSize = 10
	defaultLimit    = 10
	expireTime      = 3600
	ZSetKey         = "keyZset"
)

type GetVideolsitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideolsitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideolsitLogic {
	return &GetVideolsitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideolsitLogic) GetVideolsit(in *pb.GetVideolistReq) (*pb.GetVideolistResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "category not found")
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}
	if in.Ps == 0 {
		in.Ps = defaultPageSize
	}
	pids, _ := l.cacheProductList(l.ctx)
	return &pb.GetVideolistResp{}, nil
}

func (l *GetVideolsitLogic) cacheProductList(ctx context.Context, cid int32, cursor, ps int64) ([]int64, error) {
	pairs, err := l.svcCtx.Redis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, ZSetKey, cursor, 0, 0, int(ps))
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, pair := range pairs {
		id, _ := strconv.ParseInt(pair.Key, 10, 64)
		ids = append(ids, id)
	}
	return ids, nil
}
func (l *GetVideolsitLogic) addCacheProductList(ctx context.Context, vidos []*model.Video) error {
	if len(vidos) == 0 {
		return nil
	}
	for _, p := range vidos {
		score := p.FavoriteCount
		if score < 0 {
			score = 0
		}
		_, err := l.svcCtx.Redis.ZaddCtx(ctx, ZSetKey, score, p.VideoId)
		if err != nil {
			return err
		}
	}
}
