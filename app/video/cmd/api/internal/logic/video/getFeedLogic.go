package video

import (
	"context"

	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getfeed
func NewGetFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedLogic {
	return &GetFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFeedLogic) GetFeed(req *types.GetFeedReq) (resp *types.GetFeedResp, err error) {
	// todo: add your logic here and delete this line
	// l.svcCtx.VideoserviceRpc.
	return
}
