package relation

import (
	"MyDouSheng/app/relation/cmd/api/internal/svc"
	"MyDouSheng/app/relation/cmd/api/internal/types"
	"MyDouSheng/app/relation/cmd/rpc/pb"
	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/common/ctxdata"
	"MyDouSheng/common/xerr"
	"context"
	"fmt"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAttentionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// getattention
func NewGetAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAttentionLogic {
	return &GetAttentionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAttentionLogic) GetAttention(req *types.GetAttentionReq) (*types.GetAttentionResp, error) {
	// todo: add your logic here and delete this line

	userid := ctxdata.GetUidFromCtx(l.ctx)
	releationRpcBusiServer, err := l.svcCtx.Config.RelationserviceRpcConf.BuildTarget()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "关系创建失败")
	}
	userRpcBusiServer, err := l.svcCtx.Config.UserServiceRpcConf.BuildTarget()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "关系创建失败")
	}
	var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(releationRpcBusiServer+"/pb.relationservice/getAttention", releationRpcBusiServer+"/pb.relationservice/getAttentionRollback", &pb.GetAttentionReq{
			Token:      req.Token,
			TouserId:   req.ToUserId,
			FromuserId: userid,
			ActionType: req.ActionType,
		}).
		Add(userRpcBusiServer+"/pb.userservice/updateAttention", userRpcBusiServer+"/pb.userservice/updateAttentionRollback", &userservice.UpdateUserReq{
			FollowingId: userid,
			FollowerId:  req.ToUserId,
			Type:        req.ActionType,
		})
	err = saga.Submit()
	fmt.Printf("error:%v\n", err)
	if req.ActionType == 1 {
		if err != nil {
			return &types.GetAttentionResp{
				StatusMsg: "关注失败",
			}, err
		}
		return &types.GetAttentionResp{
			StatusMsg: "关注成功",
		}, nil
	} else {
		if err != nil {
			return &types.GetAttentionResp{
				StatusMsg: "取关失败",
			}, err
		}
		return &types.GetAttentionResp{
			StatusMsg: "取关成功",
		}, nil
	}
}
