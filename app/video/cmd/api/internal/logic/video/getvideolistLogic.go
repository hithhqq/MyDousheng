package video

import (
	"context"

	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetvideolistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getvideolist
func NewGetvideolistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetvideolistLogic {
	return &GetvideolistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetvideolistLogic) Getvideolist(req *types.GetVideoListReq) (resp *types.GetVideoListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
