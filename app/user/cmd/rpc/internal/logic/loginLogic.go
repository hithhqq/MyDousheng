package logic

import (
	"context"

	"MyDouSheng/app/identity/cmd/rpc/identity"
	"MyDouSheng/app/user/cmd/rpc/internal/svc"

	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/common/tool"
	"MyDouSheng/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userservice.LoginReq) (*userservice.LoginResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据用户名查询信息失败，Username:%s, err:%v", in.Username, err)
	}
	if user == nil {
		return &userservice.LoginResp{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
			Userid:     -1,
		}, nil
	}
	if !(tool.Md5ByString(in.Password) == user.Password) {
		return &userservice.LoginResp{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
			Userid:     -1,
		}, nil
	}
	tokenResp, err := l.svcCtx.IdentifyRpc.GenerateToken(l.ctx, &identity.GenerateTokenReq{
		UserId: user.Userid,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId: %d", user.Userid)
	}
	return &userservice.LoginResp{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		Userid:     user.Userid,
		Token:      tokenResp.AccessToken,
	}, nil
}
