package message

import (
	"context"

	"MyDouSheng/app/message/cmd/api/internal/svc"
	"MyDouSheng/app/message/cmd/api/internal/types"
	"MyDouSheng/app/message/cmd/rpc/pb"
	"MyDouSheng/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendmessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// sendmessage
func NewSendmessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendmessageLogic {
	return &SendmessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendmessageLogic) Sendmessage(req *types.SendMessageReq) (*types.SendMessageResp, error) {
	userid := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.MessageServiceRpcConf.SendMessage(l.ctx, &pb.SendMessageReq{
		Token:      req.Token,
		Content:    req.Content,
		ActionType: req.ActionType,
		ToUserId:   req.ToUserId,
		FromUserId: userid,
	})
	if err != nil {
		return nil, err
	}
	return &types.SendMessageResp{
		StatusMsg: "发送成功",
	}, err
}
