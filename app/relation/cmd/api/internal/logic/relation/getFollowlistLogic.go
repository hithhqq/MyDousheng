package relation

import (
	"context"
	"fmt"

	"MyDouSheng/app/relation/cmd/api/internal/svc"
	"MyDouSheng/app/relation/cmd/api/internal/types"
	"MyDouSheng/app/relation/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getfollowlist
func NewGetFollowlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowlistLogic {
	return &GetFollowlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowlistLogic) GetFollowlist(req *types.GetFollowlistReq) (*types.GetFollowlistResp, error) {
	// todo: add your logic here and delete this line
	var in pb.GetFollowlistReq
	in.Token = req.Token
	in.UserId = req.UserId
	userList, err := l.svcCtx.RelationserviceRpc.GetFollowlist(l.ctx, &in)
	if err != nil {
		return &types.GetFollowlistResp{
			StatusMsg: "获取用户关注列表失败",
		}, err
	}
	var resp types.GetFollowlistResp
	// resp.StatusMsg = "获取关注列表成功"
	err = copier.Copy(&resp, userList)
	if err != nil {
		return nil, fmt.Errorf("api:GetFollowlist copy err:%v", err)
	}
	return &resp, nil
}
