package user

import (
	"context"
	"fmt"

	"MyDouSheng/app/user/cmd/api/internal/svc"
	"MyDouSheng/app/user/cmd/api/internal/types"
	"MyDouSheng/app/user/cmd/rpc/userservice"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// redister
func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (*types.UserRegisterRsp, error) {
	fmt.Printf("register\n")
	registerResp, err := l.svcCtx.UserserviceRpc.Register(l.ctx, &userservice.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	if registerResp == nil {
		return nil, errors.New("register response is nil")
	}
	var resp types.UserRegisterRsp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
