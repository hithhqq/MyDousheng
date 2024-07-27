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

type GetfriendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getfriends
func NewGetfriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetfriendsLogic {
	return &GetfriendsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetfriendsLogic) Getfriends(req *types.GetFriendsReq) (*types.GetFriendsResp, error) {
	var in pb.GetFriendsReq
	in.Token = req.Token
	in.UserId = req.UserId
	userList, err := l.svcCtx.RelationserviceRpc.GetFriends(l.ctx, &in)
	if err != nil {
		return nil, fmt.Errorf("api Getfriends error:%v", err)
	}
	resp := new(types.GetFriendsResp)
	err = copier.Copy(&resp, userList)
	if err != nil {
		return nil, fmt.Errorf("api Getfriends copy serror:%v", err)
	}
	return resp, nil
}
