// Code generated by goctl. DO NOT EDIT.
// Source: identity.proto

package server

import (
	"context"

	"MyDouSheng/app/identity/cmd/rpc/internal/logic"
	"MyDouSheng/app/identity/cmd/rpc/internal/svc"
	"MyDouSheng/app/identity/cmd/rpc/pb"
)

type IdentityServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedIdentityServer
}

func NewIdentityServer(svcCtx *svc.ServiceContext) *IdentityServer {
	return &IdentityServer{
		svcCtx: svcCtx,
	}
}

// 生成token，只针对用户服务开放访问
func (s *IdentityServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

// 清除token，只针对用户服务开放访问
func (s *IdentityServer) ClearToken(ctx context.Context, in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {
	l := logic.NewClearTokenLogic(ctx, s.svcCtx)
	return l.ClearToken(in)
}

// validateToken ，只很对用户服务、授权服务api开放
func (s *IdentityServer) ValidateToken(ctx context.Context, in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	l := logic.NewValidateTokenLogic(ctx, s.svcCtx)
	return l.ValidateToken(in)
}
