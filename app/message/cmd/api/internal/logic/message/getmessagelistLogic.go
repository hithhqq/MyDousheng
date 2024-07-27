package message

import (
	"context"
	"fmt"

	"MyDouSheng/app/message/cmd/api/internal/svc"
	"MyDouSheng/app/message/cmd/api/internal/types"
	"MyDouSheng/app/message/cmd/rpc/messageservice"
	"MyDouSheng/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetmessagelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getmessagelist
func NewGetmessagelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetmessagelistLogic {
	return &GetmessagelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetmessagelistLogic) Getmessagelist(req *types.GetMessagesReq) (*types.GetMessagesResp, error) {
	// todo: add your logic here and delete this line
	userid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.MessageServiceRpcConf.GetMessageList(l.ctx, &messageservice.GetMessagesReq{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		FromUserId: userid,
	})
	if err != nil {
		return nil, fmt.Errorf("getmessagelist error, err is %v", err)
	}
	var messgaes []types.Message
	err = copier.Copy(&messgaes, resp.Messages)
	if err != nil {
		fmt.Printf("copy err is %v\n", err)
	}
	return &types.GetMessagesResp{
		StatusMsg: "获取列表成功",
		Messages:  messgaes,
	}, nil
}
