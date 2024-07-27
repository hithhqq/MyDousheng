package user

import (
	"context"

	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// userinfo
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (*types.UserInfoRsp, error) {
	// todo: add your logic here and delete this line
	userInfoRsp, err := l.svcCtx.UserserviceRpc.GetUserInfo(l.ctx, &userservice.GetUserInfoReq{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UserInfoRsp
	_ = copier.Copy(&resp, userInfoRsp)
	return &resp, nil
}
