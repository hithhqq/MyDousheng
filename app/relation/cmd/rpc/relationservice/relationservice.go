// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package relationservice

import (
	"context"

	"MyDouSheng/app/relation/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Follow            = pb.Follow
	Friend            = pb.Friend
	GetAttentionReq   = pb.GetAttentionReq
	GetAttentionResp  = pb.GetAttentionResp
	GetFanlistReq     = pb.GetFanlistReq
	GetFanlistResp    = pb.GetFanlistResp
	GetFollowlistReq  = pb.GetFollowlistReq
	GetFollowlistResp = pb.GetFollowlistResp
	GetFriendsReq     = pb.GetFriendsReq
	GetFriendsResp    = pb.GetFriendsResp

	Relationservice interface {
		GetAttention(ctx context.Context, in *GetAttentionReq, opts ...grpc.CallOption) (*GetAttentionResp, error)
		GetAttentionRollback(ctx context.Context, in *GetAttentionReq, opts ...grpc.CallOption) (*GetAttentionResp, error)
		GetFollowlist(ctx context.Context, in *GetFollowlistReq, opts ...grpc.CallOption) (*GetFollowlistResp, error)
		GetFanlist(ctx context.Context, in *GetFanlistReq, opts ...grpc.CallOption) (*GetFanlistResp, error)
		GetFriends(ctx context.Context, in *GetFriendsReq, opts ...grpc.CallOption) (*GetFriendsResp, error)
	}

	defaultRelationservice struct {
		cli zrpc.Client
	}
)

func NewRelationservice(cli zrpc.Client) Relationservice {
	return &defaultRelationservice{
		cli: cli,
	}
}

func (m *defaultRelationservice) GetAttention(ctx context.Context, in *GetAttentionReq, opts ...grpc.CallOption) (*GetAttentionResp, error) {
	client := pb.NewRelationserviceClient(m.cli.Conn())
	return client.GetAttention(ctx, in, opts...)
}

func (m *defaultRelationservice) GetAttentionRollback(ctx context.Context, in *GetAttentionReq, opts ...grpc.CallOption) (*GetAttentionResp, error) {
	client := pb.NewRelationserviceClient(m.cli.Conn())
	return client.GetAttentionRollback(ctx, in, opts...)
}

func (m *defaultRelationservice) GetFollowlist(ctx context.Context, in *GetFollowlistReq, opts ...grpc.CallOption) (*GetFollowlistResp, error) {
	client := pb.NewRelationserviceClient(m.cli.Conn())
	return client.GetFollowlist(ctx, in, opts...)
}

func (m *defaultRelationservice) GetFanlist(ctx context.Context, in *GetFanlistReq, opts ...grpc.CallOption) (*GetFanlistResp, error) {
	client := pb.NewRelationserviceClient(m.cli.Conn())
	return client.GetFanlist(ctx, in, opts...)
}

func (m *defaultRelationservice) GetFriends(ctx context.Context, in *GetFriendsReq, opts ...grpc.CallOption) (*GetFriendsResp, error) {
	client := pb.NewRelationserviceClient(m.cli.Conn())
	return client.GetFriends(ctx, in, opts...)
}
