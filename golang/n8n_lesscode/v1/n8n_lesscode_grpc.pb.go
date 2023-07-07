// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: n8n_lesscode.proto

package v1

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

// N8NLesscodeClient is the client API for N8NLesscode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type N8NLesscodeClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type n8NLesscodeClient struct {
	cc grpc.ClientConnInterface
}

func NewN8NLesscodeClient(cc grpc.ClientConnInterface) N8NLesscodeClient {
	return &n8NLesscodeClient{cc}
}

func (c *n8NLesscodeClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/n8n_lesscode.v1.N8nLesscode/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// N8NLesscodeServer is the server API for N8NLesscode service.
// All implementations must embed UnimplementedN8NLesscodeServer
// for forward compatibility
type N8NLesscodeServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedN8NLesscodeServer()
}

// UnimplementedN8NLesscodeServer must be embedded to have forward compatible implementations.
type UnimplementedN8NLesscodeServer struct {
}

func (UnimplementedN8NLesscodeServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedN8NLesscodeServer) mustEmbedUnimplementedN8NLesscodeServer() {}

// UnsafeN8NLesscodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to N8NLesscodeServer will
// result in compilation errors.
type UnsafeN8NLesscodeServer interface {
	mustEmbedUnimplementedN8NLesscodeServer()
}

func RegisterN8NLesscodeServer(s grpc.ServiceRegistrar, srv N8NLesscodeServer) {
	s.RegisterService(&N8NLesscode_ServiceDesc, srv)
}

func _N8NLesscode_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(N8NLesscodeServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/n8n_lesscode.v1.N8nLesscode/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(N8NLesscodeServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// N8NLesscode_ServiceDesc is the grpc.ServiceDesc for N8NLesscode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var N8NLesscode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "n8n_lesscode.v1.N8nLesscode",
	HandlerType: (*N8NLesscodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _N8NLesscode_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "n8n_lesscode.proto",
}