package video

import (
	"context"
	"fmt"
	"time"

	"MyDouSheng/app/user/cmd/rpc/userservice"
	"MyDouSheng/app/video/cmd/api/internal/svc"
	"MyDouSheng/app/video/cmd/api/internal/types"

	"MyDouSheng/app/video/cmd/rpc/pb"
	"MyDouSheng/app/video/cmd/rpc/videoservice"
	"MyDouSheng/common/ctxdata"
	"MyDouSheng/common/utils"
	"MyDouSheng/common/xerr"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishvideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// publishvideo
func NewPublishvideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishvideoLogic {
	return &PublishvideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishvideoLogic) Publishvideo(req *types.PublishVideoReq) (resp *types.PublishVideoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	title := req.Title
	videoUUID, _ := uuid.NewV4()
	videoDir := time.Now().Format("2002-12-16") + "/" + videoUUID.String() + ".mp4"
	videoUrl := "https://" + "mydousheng" + ".oss-cn-hangzhou.aliyuncs.com/" + videoDir
	pictureUUID, _ := uuid.NewV4()
	pictureDir := time.Now().Format("2002-12-16") + "/" + pictureUUID.String() + ".jpg"
	pictureUrl := "https://" + "mydousheng" + ".oss-cn-hangzhou.aliyuncs.com/" + pictureDir
	videoId := time.Now().Format("2002-12-16") + videoUUID.String()
	err = utils.Upload(videoDir, []byte(req.Data))
	if err != nil {
		return nil, err
	}
	pictureBytes, _ := utils.ReadFrameAsJpeg(videoUrl)
	go func() {
		err = utils.Upload(pictureDir, pictureBytes)
		if err != nil {
			utils.Delete(videoDir)
		}
	}()
	video := &videoservice.Video{
		Id:       videoId,
		PlayUrl:  videoUrl,
		CoverUrl: pictureUrl,
		Title:    title,
		Author: &videoservice.User1{
			UserId: userId,
		},
	}
	videoRpcBusiServer, err := l.svcCtx.Config.VideoserviceRpcConf.BuildTarget()
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
		Add(videoRpcBusiServer+"/pb.videoservice/publishVideo", videoRpcBusiServer+"/pb.videoservice/publishVideoRollback", &pb.PublishVideoReq{
			UserId: userId,
			Data:   video,
		}).
		Add(userRpcBusiServer+"/pb.userservice/updateWorkcount", userRpcBusiServer+"/pb.userservice/updateWorkcountRollback", &userservice.UpdateWorkCountReq{
			UserId: userId,
		})
	err = saga.Submit()
	if err != nil {
		fmt.Printf("err is %v\n", err)
	}

	return &types.PublishVideoResp{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	}, nil
}
