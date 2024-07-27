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

type GetFanlistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFanlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFanlistLogic {
	return &GetFanlistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFanlistLogic) GetFanlist(in *pb.GetFanlistReq) (*pb.GetFanlistResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.RelationModel.FindFanslist(l.ctx, in.UserId)
	if err != nil {
		return nil, fmt.Errorf("rpc: GetFanlist err:%v", err)
	}
	var user_list []*pb.Follow
	err = copier.Copy(&user_list, resp)
	if err != nil {
		return nil, fmt.Errorf("rpc: GetFanlist copy err:%v", err)
	}
	var friends []*pb.Friend
	if len(resp) > 0 {
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, user := range resp {
				source <- user
			}
		}, func(item interface{}, writer mr.Writer[*model.User], cancel func(error)) {
			user := item.(model.User)
			res, err := l.svcCtx.RelationModel.FindOneByFollowerIdFollowingId(l.ctx, user.Userid, in.UserId)
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
	for _, friend := range friends {
		for _, fan := range user_list {
			if friend.UserId == fan.UserId {
				fan.IsFollow = true
			}
		}
	}
	return &pb.GetFanlistResp{
		StatusMsg: "获取粉丝列表成功",
		UserList:  user_list,
	}, nil
}
