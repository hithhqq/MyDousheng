package user

import (
	"context"

	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (*types.UserLoginRsp, error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UserserviceRpc.Login(l.ctx, &userservice.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var resp types.UserLoginRsp
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}
