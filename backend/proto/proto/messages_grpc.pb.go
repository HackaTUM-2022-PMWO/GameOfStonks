// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/messages.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StonkServerClient is the client API for StonkServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StonkServerClient interface {
	GetStonkByName(ctx context.Context, in *StonkRequest, opts ...grpc.CallOption) (*Stonk, error)
}

type stonkServerClient struct {
	cc grpc.ClientConnInterface
}

func NewStonkServerClient(cc grpc.ClientConnInterface) StonkServerClient {
	return &stonkServerClient{cc}
}

func (c *stonkServerClient) GetStonkByName(ctx context.Context, in *StonkRequest, opts ...grpc.CallOption) (*Stonk, error) {
	out := new(Stonk)
	err := c.cc.Invoke(ctx, "/StonkServer/GetStonkByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StonkServerServer is the server API for StonkServer service.
// All implementations must embed UnimplementedStonkServerServer
// for forward compatibility
type StonkServerServer interface {
	GetStonkByName(context.Context, *StonkRequest) (*Stonk, error)
	mustEmbedUnimplementedStonkServerServer()
}

// UnimplementedStonkServerServer must be embedded to have forward compatible implementations.
type UnimplementedStonkServerServer struct {
}

func (UnimplementedStonkServerServer) GetStonkByName(context.Context, *StonkRequest) (*Stonk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStonkByName not implemented")
}
func (UnimplementedStonkServerServer) mustEmbedUnimplementedStonkServerServer() {}

// UnsafeStonkServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StonkServerServer will
// result in compilation errors.
type UnsafeStonkServerServer interface {
	mustEmbedUnimplementedStonkServerServer()
}

func RegisterStonkServerServer(s grpc.ServiceRegistrar, srv StonkServerServer) {
	s.RegisterService(&StonkServer_ServiceDesc, srv)
}

func _StonkServer_GetStonkByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StonkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StonkServerServer).GetStonkByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StonkServer/GetStonkByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StonkServerServer).GetStonkByName(ctx, req.(*StonkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StonkServer_ServiceDesc is the grpc.ServiceDesc for StonkServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StonkServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StonkServer",
	HandlerType: (*StonkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStonkByName",
			Handler:    _StonkServer_GetStonkByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/messages.proto",
}
