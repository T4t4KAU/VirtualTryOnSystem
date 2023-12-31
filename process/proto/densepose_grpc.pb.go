// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: densepose.proto

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
	Densepose_Generate_FullMethodName = "/Densepose/Generate"
)

// DenseposeClient is the client API for Densepose service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DenseposeClient interface {
	Generate(ctx context.Context, in *DensePoseRequest, opts ...grpc.CallOption) (*DensePoseReply, error)
}

type denseposeClient struct {
	cc grpc.ClientConnInterface
}

func NewDenseposeClient(cc grpc.ClientConnInterface) DenseposeClient {
	return &denseposeClient{cc}
}

func (c *denseposeClient) Generate(ctx context.Context, in *DensePoseRequest, opts ...grpc.CallOption) (*DensePoseReply, error) {
	out := new(DensePoseReply)
	err := c.cc.Invoke(ctx, Densepose_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DenseposeServer is the server API for Densepose service.
// All implementations must embed UnimplementedDenseposeServer
// for forward compatibility
type DenseposeServer interface {
	Generate(context.Context, *DensePoseRequest) (*DensePoseReply, error)
	mustEmbedUnimplementedDenseposeServer()
}

// UnimplementedDenseposeServer must be embedded to have forward compatible implementations.
type UnimplementedDenseposeServer struct {
}

func (UnimplementedDenseposeServer) Generate(context.Context, *DensePoseRequest) (*DensePoseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedDenseposeServer) mustEmbedUnimplementedDenseposeServer() {}

// UnsafeDenseposeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DenseposeServer will
// result in compilation errors.
type UnsafeDenseposeServer interface {
	mustEmbedUnimplementedDenseposeServer()
}

func RegisterDenseposeServer(s grpc.ServiceRegistrar, srv DenseposeServer) {
	s.RegisterService(&Densepose_ServiceDesc, srv)
}

func _Densepose_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DensePoseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DenseposeServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Densepose_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DenseposeServer).Generate(ctx, req.(*DensePoseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Densepose_ServiceDesc is the grpc.ServiceDesc for Densepose service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Densepose_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Densepose",
	HandlerType: (*DenseposeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Densepose_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "densepose.proto",
}
