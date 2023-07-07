// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.7.1
// source: process_instance.proto

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

// WorkFlowProcessInstanceClient is the client API for WorkFlowProcessInstance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkFlowProcessInstanceClient interface {
	// Start 开启一个流程
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error)
	StartByConditions(ctx context.Context, in *StartByConditionsRequest, opts ...grpc.CallOption) (*StartByConditionsReply, error)
	QueryProcessInstances(ctx context.Context, in *QueryProcessInstancesRequest, opts ...grpc.CallOption) (*QueryProcessInstancesReply, error)
	DeleteProcessInstance(ctx context.Context, in *DeleteProcessInstanceRequest, opts ...grpc.CallOption) (*DeleteProcessInstanceReply, error)
}

type workFlowProcessInstanceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkFlowProcessInstanceClient(cc grpc.ClientConnInterface) WorkFlowProcessInstanceClient {
	return &workFlowProcessInstanceClient{cc}
}

func (c *workFlowProcessInstanceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := c.cc.Invoke(ctx, "/go_workflow.v1.WorkFlowProcessInstance/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workFlowProcessInstanceClient) StartByConditions(ctx context.Context, in *StartByConditionsRequest, opts ...grpc.CallOption) (*StartByConditionsReply, error) {
	out := new(StartByConditionsReply)
	err := c.cc.Invoke(ctx, "/go_workflow.v1.WorkFlowProcessInstance/StartByConditions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workFlowProcessInstanceClient) QueryProcessInstances(ctx context.Context, in *QueryProcessInstancesRequest, opts ...grpc.CallOption) (*QueryProcessInstancesReply, error) {
	out := new(QueryProcessInstancesReply)
	err := c.cc.Invoke(ctx, "/go_workflow.v1.WorkFlowProcessInstance/QueryProcessInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workFlowProcessInstanceClient) DeleteProcessInstance(ctx context.Context, in *DeleteProcessInstanceRequest, opts ...grpc.CallOption) (*DeleteProcessInstanceReply, error) {
	out := new(DeleteProcessInstanceReply)
	err := c.cc.Invoke(ctx, "/go_workflow.v1.WorkFlowProcessInstance/DeleteProcessInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkFlowProcessInstanceServer is the server API for WorkFlowProcessInstance service.
// All implementations must embed UnimplementedWorkFlowProcessInstanceServer
// for forward compatibility
type WorkFlowProcessInstanceServer interface {
	// Start 开启一个流程
	Start(context.Context, *StartRequest) (*StartReply, error)
	StartByConditions(context.Context, *StartByConditionsRequest) (*StartByConditionsReply, error)
	QueryProcessInstances(context.Context, *QueryProcessInstancesRequest) (*QueryProcessInstancesReply, error)
	DeleteProcessInstance(context.Context, *DeleteProcessInstanceRequest) (*DeleteProcessInstanceReply, error)
	mustEmbedUnimplementedWorkFlowProcessInstanceServer()
}

// UnimplementedWorkFlowProcessInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkFlowProcessInstanceServer struct {
}

func (UnimplementedWorkFlowProcessInstanceServer) Start(context.Context, *StartRequest) (*StartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedWorkFlowProcessInstanceServer) StartByConditions(context.Context, *StartByConditionsRequest) (*StartByConditionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartByConditions not implemented")
}
func (UnimplementedWorkFlowProcessInstanceServer) QueryProcessInstances(context.Context, *QueryProcessInstancesRequest) (*QueryProcessInstancesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProcessInstances not implemented")
}
func (UnimplementedWorkFlowProcessInstanceServer) DeleteProcessInstance(context.Context, *DeleteProcessInstanceRequest) (*DeleteProcessInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProcessInstance not implemented")
}
func (UnimplementedWorkFlowProcessInstanceServer) mustEmbedUnimplementedWorkFlowProcessInstanceServer() {
}

// UnsafeWorkFlowProcessInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkFlowProcessInstanceServer will
// result in compilation errors.
type UnsafeWorkFlowProcessInstanceServer interface {
	mustEmbedUnimplementedWorkFlowProcessInstanceServer()
}

func RegisterWorkFlowProcessInstanceServer(s grpc.ServiceRegistrar, srv WorkFlowProcessInstanceServer) {
	s.RegisterService(&WorkFlowProcessInstance_ServiceDesc, srv)
}

func _WorkFlowProcessInstance_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowProcessInstanceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_workflow.v1.WorkFlowProcessInstance/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowProcessInstanceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkFlowProcessInstance_StartByConditions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartByConditionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowProcessInstanceServer).StartByConditions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_workflow.v1.WorkFlowProcessInstance/StartByConditions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowProcessInstanceServer).StartByConditions(ctx, req.(*StartByConditionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkFlowProcessInstance_QueryProcessInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProcessInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowProcessInstanceServer).QueryProcessInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_workflow.v1.WorkFlowProcessInstance/QueryProcessInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowProcessInstanceServer).QueryProcessInstances(ctx, req.(*QueryProcessInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkFlowProcessInstance_DeleteProcessInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProcessInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkFlowProcessInstanceServer).DeleteProcessInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go_workflow.v1.WorkFlowProcessInstance/DeleteProcessInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkFlowProcessInstanceServer).DeleteProcessInstance(ctx, req.(*DeleteProcessInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkFlowProcessInstance_ServiceDesc is the grpc.ServiceDesc for WorkFlowProcessInstance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkFlowProcessInstance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_workflow.v1.WorkFlowProcessInstance",
	HandlerType: (*WorkFlowProcessInstanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _WorkFlowProcessInstance_Start_Handler,
		},
		{
			MethodName: "StartByConditions",
			Handler:    _WorkFlowProcessInstance_StartByConditions_Handler,
		},
		{
			MethodName: "QueryProcessInstances",
			Handler:    _WorkFlowProcessInstance_QueryProcessInstances_Handler,
		},
		{
			MethodName: "DeleteProcessInstance",
			Handler:    _WorkFlowProcessInstance_DeleteProcessInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process_instance.proto",
}