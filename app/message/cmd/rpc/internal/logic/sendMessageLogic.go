package logic

import (
	"context"
	"fmt"
	"time"

	"MyDouSheng/app/message/cmd/rpc/internal/svc"
	"MyDouSheng/app/message/cmd/rpc/pb"
	"MyDouSheng/app/message/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	data := &model.Message{
		FromUserId: in.FromUserId,
		ToUserId:   in.ToUserId,
		Content:    in.Content,
		CreateAt:   time.Now(),
	}
	_, err := l.svcCtx.MessageModel.Insert(l.ctx, data)
	if err != nil {
		return nil, fmt.Errorf("rpc SendMessage error:%v", err)
	}
	return &pb.SendMessageResp{}, nil
}
