package logic

import (
	"context"
	"fmt"

	"MyDouSheng/app/message/cmd/rpc/internal/svc"
	"MyDouSheng/app/message/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *pb.GetMessagesReq) (*pb.GetMessagesResp, error) {
	resp, err := l.svcCtx.MessageModel.FindMessages(l.ctx, in.FromUserId, in.ToUserId)
	if err != nil {
		return nil, err
	}
	var messages []*pb.Message
	err = copier.Copy(&messages, resp)
	if err != nil {
		return nil, fmt.Errorf("rpc GetMessageList copy err:%v", err)
	}
	return &pb.GetMessagesResp{
		StatusMsg: "获取消息列表成功",
		Messages:  messages,
	}, nil
}
