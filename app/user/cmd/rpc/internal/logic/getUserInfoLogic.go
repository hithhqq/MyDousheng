package logic

import (
	"context"

	"MyDouSheng/app/user/cmd/rpc/internal/svc"
	"MyDouSheng/app/user/cmd/rpc/pb"
	"MyDouSheng/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return &pb.GetUserInfoResp{
			StatusCode: 0,
			StatusMsg:  "查询用户信息失败",
		}, err
	}
	var userInfo pb.User
	copier.Copy(&userInfo, user)
	return &pb.GetUserInfoResp{
		StatusCode: 0,
		StatusMsg:  "查询用户信息成功",
		UserInfo:   &userInfo,
	}, nil
}
