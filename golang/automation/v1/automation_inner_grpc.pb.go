// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: automation_inner.proto

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

// AutomationInnerClient is the client API for AutomationInner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutomationInnerClient interface {
	// 应用模板（isCreate为true时，为保存模板）
	ApplyTemplate(ctx context.Context, in *ApplyTemplateReq, opts ...grpc.CallOption) (*ApplyTemplateReply, error)
	// 执行回调
	HandleWorkflowExecute(ctx context.Context, in *HandleWorkflowExecuteReq, opts ...grpc.CallOption) (*HandleWorkflowExecuteReply, error)
	// 执行失败回调
	HandleWorkflowError(ctx context.Context, in *HandleWorkflowErrorReq, opts ...grpc.CallOption) (*HandleWorkflowErrorReply, error)
	// 重建所有n8n workflow
	BatchRebuildActive(ctx context.Context, in *BatchRebuildActiveReq, opts ...grpc.CallOption) (*BatchRebuildActiveReply, error)
	// 关闭整个组织工作流开关
	GlobalSwitchOff(ctx context.Context, in *GlobalSwitchOffReq, opts ...grpc.CallOption) (*GlobalSwitchOffReply, error)
	// 创建待办
	CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoReply, error)
	// 更新待办
	UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoReply, error)
}

type automationInnerClient struct {
	cc grpc.ClientConnInterface
}

func NewAutomationInnerClient(cc grpc.ClientConnInterface) AutomationInnerClient {
	return &automationInnerClient{cc}
}

func (c *automationInnerClient) ApplyTemplate(ctx context.Context, in *ApplyTemplateReq, opts ...grpc.CallOption) (*ApplyTemplateReply, error) {
	out := new(ApplyTemplateReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/ApplyTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) HandleWorkflowExecute(ctx context.Context, in *HandleWorkflowExecuteReq, opts ...grpc.CallOption) (*HandleWorkflowExecuteReply, error) {
	out := new(HandleWorkflowExecuteReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/HandleWorkflowExecute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) HandleWorkflowError(ctx context.Context, in *HandleWorkflowErrorReq, opts ...grpc.CallOption) (*HandleWorkflowErrorReply, error) {
	out := new(HandleWorkflowErrorReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/HandleWorkflowError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) BatchRebuildActive(ctx context.Context, in *BatchRebuildActiveReq, opts ...grpc.CallOption) (*BatchRebuildActiveReply, error) {
	out := new(BatchRebuildActiveReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/BatchRebuildActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) GlobalSwitchOff(ctx context.Context, in *GlobalSwitchOffReq, opts ...grpc.CallOption) (*GlobalSwitchOffReply, error) {
	out := new(GlobalSwitchOffReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/GlobalSwitchOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoReply, error) {
	out := new(CreateTodoReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationInnerClient) UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoReply, error) {
	out := new(UpdateTodoReply)
	err := c.cc.Invoke(ctx, "/automation.v1.AutomationInner/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutomationInnerServer is the server API for AutomationInner service.
// All implementations must embed UnimplementedAutomationInnerServer
// for forward compatibility
type AutomationInnerServer interface {
	// 应用模板（isCreate为true时，为保存模板）
	ApplyTemplate(context.Context, *ApplyTemplateReq) (*ApplyTemplateReply, error)
	// 执行回调
	HandleWorkflowExecute(context.Context, *HandleWorkflowExecuteReq) (*HandleWorkflowExecuteReply, error)
	// 执行失败回调
	HandleWorkflowError(context.Context, *HandleWorkflowErrorReq) (*HandleWorkflowErrorReply, error)
	// 重建所有n8n workflow
	BatchRebuildActive(context.Context, *BatchRebuildActiveReq) (*BatchRebuildActiveReply, error)
	// 关闭整个组织工作流开关
	GlobalSwitchOff(context.Context, *GlobalSwitchOffReq) (*GlobalSwitchOffReply, error)
	// 创建待办
	CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoReply, error)
	// 更新待办
	UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoReply, error)
	mustEmbedUnimplementedAutomationInnerServer()
}

// UnimplementedAutomationInnerServer must be embedded to have forward compatible implementations.
type UnimplementedAutomationInnerServer struct {
}

func (UnimplementedAutomationInnerServer) ApplyTemplate(context.Context, *ApplyTemplateReq) (*ApplyTemplateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyTemplate not implemented")
}
func (UnimplementedAutomationInnerServer) HandleWorkflowExecute(context.Context, *HandleWorkflowExecuteReq) (*HandleWorkflowExecuteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleWorkflowExecute not implemented")
}
func (UnimplementedAutomationInnerServer) HandleWorkflowError(context.Context, *HandleWorkflowErrorReq) (*HandleWorkflowErrorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleWorkflowError not implemented")
}
func (UnimplementedAutomationInnerServer) BatchRebuildActive(context.Context, *BatchRebuildActiveReq) (*BatchRebuildActiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchRebuildActive not implemented")
}
func (UnimplementedAutomationInnerServer) GlobalSwitchOff(context.Context, *GlobalSwitchOffReq) (*GlobalSwitchOffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GlobalSwitchOff not implemented")
}
func (UnimplementedAutomationInnerServer) CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedAutomationInnerServer) UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedAutomationInnerServer) mustEmbedUnimplementedAutomationInnerServer() {}

// UnsafeAutomationInnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutomationInnerServer will
// result in compilation errors.
type UnsafeAutomationInnerServer interface {
	mustEmbedUnimplementedAutomationInnerServer()
}

func RegisterAutomationInnerServer(s grpc.ServiceRegistrar, srv AutomationInnerServer) {
	s.RegisterService(&AutomationInner_ServiceDesc, srv)
}

func _AutomationInner_ApplyTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).ApplyTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/ApplyTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).ApplyTemplate(ctx, req.(*ApplyTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_HandleWorkflowExecute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleWorkflowExecuteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).HandleWorkflowExecute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/HandleWorkflowExecute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).HandleWorkflowExecute(ctx, req.(*HandleWorkflowExecuteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_HandleWorkflowError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleWorkflowErrorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).HandleWorkflowError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/HandleWorkflowError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).HandleWorkflowError(ctx, req.(*HandleWorkflowErrorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_BatchRebuildActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRebuildActiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).BatchRebuildActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/BatchRebuildActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).BatchRebuildActive(ctx, req.(*BatchRebuildActiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_GlobalSwitchOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalSwitchOffReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).GlobalSwitchOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/GlobalSwitchOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).GlobalSwitchOff(ctx, req.(*GlobalSwitchOffReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).CreateTodo(ctx, req.(*CreateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationInner_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationInnerServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automation.v1.AutomationInner/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationInnerServer).UpdateTodo(ctx, req.(*UpdateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AutomationInner_ServiceDesc is the grpc.ServiceDesc for AutomationInner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutomationInner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "automation.v1.AutomationInner",
	HandlerType: (*AutomationInnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyTemplate",
			Handler:    _AutomationInner_ApplyTemplate_Handler,
		},
		{
			MethodName: "HandleWorkflowExecute",
			Handler:    _AutomationInner_HandleWorkflowExecute_Handler,
		},
		{
			MethodName: "HandleWorkflowError",
			Handler:    _AutomationInner_HandleWorkflowError_Handler,
		},
		{
			MethodName: "BatchRebuildActive",
			Handler:    _AutomationInner_BatchRebuildActive_Handler,
		},
		{
			MethodName: "GlobalSwitchOff",
			Handler:    _AutomationInner_GlobalSwitchOff_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _AutomationInner_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _AutomationInner_UpdateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automation_inner.proto",
}
