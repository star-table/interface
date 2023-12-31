// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type WorkFlowProcessDefineHTTPServer interface {
	DeleteProcessDefine(context.Context, *DeleteProcessDefineRequest) (*DeleteProcessDefineReply, error)
	QueryProcessDefines(context.Context, *QueryProcessDefinesRequest) (*QueryProcessDefinesReply, error)
	SaveProcessDefine(context.Context, *SaveProcessDefineRequest) (*SaveProcessDefineReply, error)
}

func RegisterWorkFlowProcessDefineHTTPServer(s *http.Server, srv WorkFlowProcessDefineHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/goWorkflow/processDefine/create", _WorkFlowProcessDefine_SaveProcessDefine0_HTTP_Handler(srv))
	r.POST("/v1/goWorkflow/processDefine/query", _WorkFlowProcessDefine_QueryProcessDefines0_HTTP_Handler(srv))
	r.POST("/v1/goWorkflow/processDefine/delete", _WorkFlowProcessDefine_DeleteProcessDefine0_HTTP_Handler(srv))
}

func _WorkFlowProcessDefine_SaveProcessDefine0_HTTP_Handler(srv WorkFlowProcessDefineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveProcessDefineRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/go_workflow.v1.WorkFlowProcessDefine/SaveProcessDefine")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveProcessDefine(ctx, req.(*SaveProcessDefineRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveProcessDefineReply)
		return ctx.Result(200, reply)
	}
}

func _WorkFlowProcessDefine_QueryProcessDefines0_HTTP_Handler(srv WorkFlowProcessDefineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in QueryProcessDefinesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/go_workflow.v1.WorkFlowProcessDefine/QueryProcessDefines")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.QueryProcessDefines(ctx, req.(*QueryProcessDefinesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*QueryProcessDefinesReply)
		return ctx.Result(200, reply)
	}
}

func _WorkFlowProcessDefine_DeleteProcessDefine0_HTTP_Handler(srv WorkFlowProcessDefineHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteProcessDefineRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/go_workflow.v1.WorkFlowProcessDefine/DeleteProcessDefine")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteProcessDefine(ctx, req.(*DeleteProcessDefineRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteProcessDefineReply)
		return ctx.Result(200, reply)
	}
}

type WorkFlowProcessDefineHTTPClient interface {
	DeleteProcessDefine(ctx context.Context, req *DeleteProcessDefineRequest, opts ...http.CallOption) (rsp *DeleteProcessDefineReply, err error)
	QueryProcessDefines(ctx context.Context, req *QueryProcessDefinesRequest, opts ...http.CallOption) (rsp *QueryProcessDefinesReply, err error)
	SaveProcessDefine(ctx context.Context, req *SaveProcessDefineRequest, opts ...http.CallOption) (rsp *SaveProcessDefineReply, err error)
}

type WorkFlowProcessDefineHTTPClientImpl struct {
	cc *http.Client
}

func NewWorkFlowProcessDefineHTTPClient(client *http.Client) WorkFlowProcessDefineHTTPClient {
	return &WorkFlowProcessDefineHTTPClientImpl{client}
}

func (c *WorkFlowProcessDefineHTTPClientImpl) DeleteProcessDefine(ctx context.Context, in *DeleteProcessDefineRequest, opts ...http.CallOption) (*DeleteProcessDefineReply, error) {
	var out DeleteProcessDefineReply
	pattern := "/v1/goWorkflow/processDefine/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/go_workflow.v1.WorkFlowProcessDefine/DeleteProcessDefine"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkFlowProcessDefineHTTPClientImpl) QueryProcessDefines(ctx context.Context, in *QueryProcessDefinesRequest, opts ...http.CallOption) (*QueryProcessDefinesReply, error) {
	var out QueryProcessDefinesReply
	pattern := "/v1/goWorkflow/processDefine/query"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/go_workflow.v1.WorkFlowProcessDefine/QueryProcessDefines"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *WorkFlowProcessDefineHTTPClientImpl) SaveProcessDefine(ctx context.Context, in *SaveProcessDefineRequest, opts ...http.CallOption) (*SaveProcessDefineReply, error) {
	var out SaveProcessDefineReply
	pattern := "/v1/goWorkflow/processDefine/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/go_workflow.v1.WorkFlowProcessDefine/SaveProcessDefine"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
