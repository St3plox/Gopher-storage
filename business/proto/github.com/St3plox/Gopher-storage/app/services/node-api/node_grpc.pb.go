// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NodeV1Client is the client API for NodeV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type nodeV1Client struct {
	cc grpc.ClientConnInterface
}

func NewNodeV1Client(cc grpc.ClientConnInterface) NodeV1Client {
	return &nodeV1Client{cc}
}

func (c *nodeV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/NodeV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/NodeV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeV1Server is the server API for NodeV1 service.
// All implementations must embed UnimplementedNodeV1Server
// for forward compatibility
type NodeV1Server interface {
	Create(context.Context, *CreateRequest) (*emptypb.Empty, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedNodeV1Server()
}

// UnimplementedNodeV1Server must be embedded to have forward compatible implementations.
type UnimplementedNodeV1Server struct {
}

func (UnimplementedNodeV1Server) Create(context.Context, *CreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNodeV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNodeV1Server) mustEmbedUnimplementedNodeV1Server() {}

// UnsafeNodeV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeV1Server will
// result in compilation errors.
type UnsafeNodeV1Server interface {
	mustEmbedUnimplementedNodeV1Server()
}

func RegisterNodeV1Server(s grpc.ServiceRegistrar, srv NodeV1Server) {
	s.RegisterService(&NodeV1_ServiceDesc, srv)
}

func _NodeV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeV1_ServiceDesc is the grpc.ServiceDesc for NodeV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NodeV1",
	HandlerType: (*NodeV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NodeV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NodeV1_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
