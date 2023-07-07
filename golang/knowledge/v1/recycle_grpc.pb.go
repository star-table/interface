// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.7.1
// source: recycle.proto

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

// RecycleClient is the client API for Recycle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecycleClient interface {
	// 将文档放回回收站
	RecycleDocument(ctx context.Context, in *RecycleDocumentRequest, opts ...grpc.CallOption) (*RecycleDocumentReply, error)
	// 将文档放回回收站
	Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverReply, error)
	// 回收站列表
	ListRecycle(ctx context.Context, in *ListRecycleRequest, opts ...grpc.CallOption) (*ListRecycleReply, error)
}

type recycleClient struct {
	cc grpc.ClientConnInterface
}

func NewRecycleClient(cc grpc.ClientConnInterface) RecycleClient {
	return &recycleClient{cc}
}

func (c *recycleClient) RecycleDocument(ctx context.Context, in *RecycleDocumentRequest, opts ...grpc.CallOption) (*RecycleDocumentReply, error) {
	out := new(RecycleDocumentReply)
	err := c.cc.Invoke(ctx, "/knowledge.v1.Recycle/RecycleDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleClient) Recover(ctx context.Context, in *RecoverRequest, opts ...grpc.CallOption) (*RecoverReply, error) {
	out := new(RecoverReply)
	err := c.cc.Invoke(ctx, "/knowledge.v1.Recycle/Recover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleClient) ListRecycle(ctx context.Context, in *ListRecycleRequest, opts ...grpc.CallOption) (*ListRecycleReply, error) {
	out := new(ListRecycleReply)
	err := c.cc.Invoke(ctx, "/knowledge.v1.Recycle/ListRecycle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecycleServer is the server API for Recycle service.
// All implementations must embed UnimplementedRecycleServer
// for forward compatibility
type RecycleServer interface {
	// 将文档放回回收站
	RecycleDocument(context.Context, *RecycleDocumentRequest) (*RecycleDocumentReply, error)
	// 将文档放回回收站
	Recover(context.Context, *RecoverRequest) (*RecoverReply, error)
	// 回收站列表
	ListRecycle(context.Context, *ListRecycleRequest) (*ListRecycleReply, error)
	mustEmbedUnimplementedRecycleServer()
}

// UnimplementedRecycleServer must be embedded to have forward compatible implementations.
type UnimplementedRecycleServer struct {
}

func (UnimplementedRecycleServer) RecycleDocument(context.Context, *RecycleDocumentRequest) (*RecycleDocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecycleDocument not implemented")
}
func (UnimplementedRecycleServer) Recover(context.Context, *RecoverRequest) (*RecoverReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recover not implemented")
}
func (UnimplementedRecycleServer) ListRecycle(context.Context, *ListRecycleRequest) (*ListRecycleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRecycle not implemented")
}
func (UnimplementedRecycleServer) mustEmbedUnimplementedRecycleServer() {}

// UnsafeRecycleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecycleServer will
// result in compilation errors.
type UnsafeRecycleServer interface {
	mustEmbedUnimplementedRecycleServer()
}

func RegisterRecycleServer(s grpc.ServiceRegistrar, srv RecycleServer) {
	s.RegisterService(&Recycle_ServiceDesc, srv)
}

func _Recycle_RecycleDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecycleDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleServer).RecycleDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knowledge.v1.Recycle/RecycleDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleServer).RecycleDocument(ctx, req.(*RecycleDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recycle_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleServer).Recover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knowledge.v1.Recycle/Recover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleServer).Recover(ctx, req.(*RecoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recycle_ListRecycle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecycleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleServer).ListRecycle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knowledge.v1.Recycle/ListRecycle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleServer).ListRecycle(ctx, req.(*ListRecycleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Recycle_ServiceDesc is the grpc.ServiceDesc for Recycle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Recycle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "knowledge.v1.Recycle",
	HandlerType: (*RecycleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecycleDocument",
			Handler:    _Recycle_RecycleDocument_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _Recycle_Recover_Handler,
		},
		{
			MethodName: "ListRecycle",
			Handler:    _Recycle_ListRecycle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recycle.proto",
}