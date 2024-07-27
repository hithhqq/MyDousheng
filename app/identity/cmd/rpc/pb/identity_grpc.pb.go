// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: identity.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Identity_GenerateToken_FullMethodName = "/pb.identity/generateToken"
	Identity_ClearToken_FullMethodName    = "/pb.identity/clearToken"
	Identity_ValidateToken_FullMethodName = "/pb.identity/validateToken"
)

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityClient interface {
	// 生成token，只针对用户服务开放访问
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	// 清除token，只针对用户服务开放访问
	ClearToken(ctx context.Context, in *ClearTokenReq, opts ...grpc.CallOption) (*ClearTokenResp, error)
	// validateToken ，只很对用户服务、授权服务api开放
	ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*ValidateTokenResp, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, Identity_GenerateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ClearToken(ctx context.Context, in *ClearTokenReq, opts ...grpc.CallOption) (*ClearTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearTokenResp)
	err := c.cc.Invoke(ctx, Identity_ClearToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ValidateToken(ctx context.Context, in *ValidateTokenReq, opts ...grpc.CallOption) (*ValidateTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateTokenResp)
	err := c.cc.Invoke(ctx, Identity_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
// All implementations must embed UnimplementedIdentityServer
// for forward compatibility
type IdentityServer interface {
	// 生成token，只针对用户服务开放访问
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	// 清除token，只针对用户服务开放访问
	ClearToken(context.Context, *ClearTokenReq) (*ClearTokenResp, error)
	// validateToken ，只很对用户服务、授权服务api开放
	ValidateToken(context.Context, *ValidateTokenReq) (*ValidateTokenResp, error)
	mustEmbedUnimplementedIdentityServer()
}

// UnimplementedIdentityServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (UnimplementedIdentityServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedIdentityServer) ClearToken(context.Context, *ClearTokenReq) (*ClearTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearToken not implemented")
}
func (UnimplementedIdentityServer) ValidateToken(context.Context, *ValidateTokenReq) (*ValidateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedIdentityServer) mustEmbedUnimplementedIdentityServer() {}

// UnsafeIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServer will
// result in compilation errors.
type UnsafeIdentityServer interface {
	mustEmbedUnimplementedIdentityServer()
}

func RegisterIdentityServer(s grpc.ServiceRegistrar, srv IdentityServer) {
	s.RegisterService(&Identity_ServiceDesc, srv)
}

func _Identity_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identity_GenerateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ClearToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ClearToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identity_ClearToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ClearToken(ctx, req.(*ClearTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Identity_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ValidateToken(ctx, req.(*ValidateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Identity_ServiceDesc is the grpc.ServiceDesc for Identity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "generateToken",
			Handler:    _Identity_GenerateToken_Handler,
		},
		{
			MethodName: "clearToken",
			Handler:    _Identity_ClearToken_Handler,
		},
		{
			MethodName: "validateToken",
			Handler:    _Identity_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}
