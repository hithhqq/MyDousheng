// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"context"

	"MyDouSheng/app/video/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetVideolistReq  = pb.GetVideolistReq
	GetVideolistResp = pb.GetVideolistResp
	PublishVideoReq  = pb.PublishVideoReq
	PublishVideoResp = pb.PublishVideoResp
	User1            = pb.User1
	Video            = pb.Video

	Videoservice interface {
		PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error)
		PublishVideoRollback(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error)
		GetVideolsit(ctx context.Context, in *GetVideolistReq, opts ...grpc.CallOption) (*GetVideolistResp, error)
	}

	defaultVideoservice struct {
		cli zrpc.Client
	}
)

func NewVideoservice(cli zrpc.Client) Videoservice {
	return &defaultVideoservice{
		cli: cli,
	}
}

func (m *defaultVideoservice) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error) {
	client := pb.NewVideoserviceClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideoservice) PublishVideoRollback(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error) {
	client := pb.NewVideoserviceClient(m.cli.Conn())
	return client.PublishVideoRollback(ctx, in, opts...)
}

func (m *defaultVideoservice) GetVideolsit(ctx context.Context, in *GetVideolistReq, opts ...grpc.CallOption) (*GetVideolistResp, error) {
	client := pb.NewVideoserviceClient(m.cli.Conn())
	return client.GetVideolsit(ctx, in, opts...)
}
