package logic

import (
	"context"
	"fmt"

	"MyDouSheng/app/relation/cmd/rpc/internal/svc"
	"MyDouSheng/app/relation/cmd/rpc/pb"
	"MyDouSheng/app/user/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsLogic {
	return &GetFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsLogic) GetFriends(in *pb.GetFriendsReq) (*pb.GetFriendsResp, error) {
	resp, err := l.svcCtx.RelationModel.FindFollowlist(l.ctx, in.UserId)
	if err != nil {
		return nil, fmt.Errorf("rpc GetFriends FindFollowlist error:%v", err)
	}
	var friends []*pb.Friend
	if len(resp) > 0 {
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, user := range resp {
				source <- user
			}
		}, func(item interface{}, writer mr.Writer[*model.User], cancel func(error)) {
			user := item.(model.User)
			res, err := l.svcCtx.RelationModel.FindOneByFollowerIdFollowingId(l.ctx, in.UserId, user.Userid)
			if err != nil {
				logx.WithContext(l.ctx).Errorf("FindOneByFollowerIdFollowingId err id:%v, err:%v\n", user.Userid, err)
				return
			}
			if res != nil {
				writer.Write(&user)
			}
		}, func(pipe <-chan *model.User, cancel func(error)) {
			for user := range pipe {
				friend := new(pb.Friend)
				err := copier.Copy(friend, user)
				if err != nil {
					logx.WithContext(l.ctx).Errorf("copy err friend:%+v, err:%v\n", user, err)
					return
				}
				friend.IsFollow = true
				friends = append(friends, friend)
				fmt.Printf("friends is %+v\n", friends)
			}
		})
	}
	return &pb.GetFriendsResp{
		StatusMsg: "获取好友列表成功",
		UserList:  friends,
	}, nil
}
