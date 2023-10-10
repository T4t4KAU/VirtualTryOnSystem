// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: vitons.proto

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

const (
	Vitons_Generate_FullMethodName = "/Vitons/Generate"
)

// VitonsClient is the client API for Vitons service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VitonsClient interface {
	Generate(ctx context.Context, in *VitonsRequest, opts ...grpc.CallOption) (*VitonsReply, error)
}

type vitonsClient struct {
	cc grpc.ClientConnInterface
}

func NewVitonsClient(cc grpc.ClientConnInterface) VitonsClient {
	return &vitonsClient{cc}
}

func (c *vitonsClient) Generate(ctx context.Context, in *VitonsRequest, opts ...grpc.CallOption) (*VitonsReply, error) {
	out := new(VitonsReply)
	err := c.cc.Invoke(ctx, Vitons_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VitonsServer is the server API for Vitons service.
// All implementations must embed UnimplementedVitonsServer
// for forward compatibility
type VitonsServer interface {
	Generate(context.Context, *VitonsRequest) (*VitonsReply, error)
	mustEmbedUnimplementedVitonsServer()
}

// UnimplementedVitonsServer must be embedded to have forward compatible implementations.
type UnimplementedVitonsServer struct {
}

func (UnimplementedVitonsServer) Generate(context.Context, *VitonsRequest) (*VitonsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedVitonsServer) mustEmbedUnimplementedVitonsServer() {}

// UnsafeVitonsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VitonsServer will
// result in compilation errors.
type UnsafeVitonsServer interface {
	mustEmbedUnimplementedVitonsServer()
}

func RegisterVitonsServer(s grpc.ServiceRegistrar, srv VitonsServer) {
	s.RegisterService(&Vitons_ServiceDesc, srv)
}

func _Vitons_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VitonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VitonsServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Vitons_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VitonsServer).Generate(ctx, req.(*VitonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vitons_ServiceDesc is the grpc.ServiceDesc for Vitons service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vitons_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Vitons",
	HandlerType: (*VitonsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Vitons_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vitons.proto",
}
