package logic

import (
	"MyDouSheng/app/identity/cmd/rpc/identity"
	"MyDouSheng/app/user/cmd/rpc/internal/svc"
	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/app/user/model"
	"MyDouSheng/common/tool"
	"MyDouSheng/common/xerr"
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *userservice.RegisterReq) (*userservice.RegisterResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s, err:%v", in.Username, err)
	}
	if user != nil {
		return &userservice.RegisterResp{
			StatusCode: -1,
			StatusMsg:  "用户名已存在",
			Userid:     -1,
		}, nil
	}
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Password:   tool.Md5ByString(in.Password),
		Createtime: time.Now(),
	})
	if err != nil {
		return &userservice.RegisterResp{
			StatusCode: -1,
			StatusMsg:  "注册失败",
			Userid:     -1,
		}, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
	}
	lastId, _ := insertResult.LastInsertId()
	userId := lastId
	tokenResp, err := l.svcCtx.IdentifyRpc.GenerateToken(l.ctx, &identity.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &userservice.RegisterResp{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		Userid:     userId,
		Token:      tokenResp.AccessToken,
	}, nil
}
