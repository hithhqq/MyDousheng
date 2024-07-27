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

type GetFunlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getfanlist
func NewGetFunlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFunlistLogic {
	return &GetFunlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFunlistLogic) GetFunlist(req *types.GetFanlistReq) (*types.GetFanlistResp, error) {
	// todo: add your logic here and delete this line
	var in pb.GetFanlistReq
	in.Token = req.Token
	in.UserId = req.UserId
	userList, err := l.svcCtx.RelationserviceRpc.GetFanlist(l.ctx, &in)
	if err != nil {
		return &types.GetFanlistResp{
			StatusMsg: "获取粉丝列表失败",
		}, err
	}
	var resp types.GetFanlistResp
	err = copier.Copy(&resp, userList)
	if err != nil {
		return nil, fmt.Errorf("api: GetFunlist copy err:%v", err)
	}
	return &resp, nil
}
