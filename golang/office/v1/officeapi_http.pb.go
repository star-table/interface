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

type OfficeApiHTTPServer interface {
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	GetWopiFileInfoEditable(context.Context, *WOPIRequest) (*WOPIEditableResp, error)
}

func RegisterOfficeApiHTTPServer(s *http.Server, srv OfficeApiHTTPServer) {
	r := s.Route("/")
	r.GET("/api/config", _OfficeApi_GetConfig0_HTTP_Handler(srv))
	r.GET("/api/wopi/files/{file_name}", _OfficeApi_GetWopiFileInfoEditable0_HTTP_Handler(srv))
}

func _OfficeApi_GetConfig0_HTTP_Handler(srv OfficeApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetConfigRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/officeapi.v1.OfficeApi/GetConfig")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetConfig(ctx, req.(*GetConfigRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetConfigResponse)
		return ctx.Result(200, reply)
	}
}

func _OfficeApi_GetWopiFileInfoEditable0_HTTP_Handler(srv OfficeApiHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WOPIRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/officeapi.v1.OfficeApi/GetWopiFileInfoEditable")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetWopiFileInfoEditable(ctx, req.(*WOPIRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WOPIEditableResp)
		return ctx.Result(200, reply)
	}
}

type OfficeApiHTTPClient interface {
	GetConfig(ctx context.Context, req *GetConfigRequest, opts ...http.CallOption) (rsp *GetConfigResponse, err error)
	GetWopiFileInfoEditable(ctx context.Context, req *WOPIRequest, opts ...http.CallOption) (rsp *WOPIEditableResp, err error)
}

type OfficeApiHTTPClientImpl struct {
	cc *http.Client
}

func NewOfficeApiHTTPClient(client *http.Client) OfficeApiHTTPClient {
	return &OfficeApiHTTPClientImpl{client}
}

func (c *OfficeApiHTTPClientImpl) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...http.CallOption) (*GetConfigResponse, error) {
	var out GetConfigResponse
	pattern := "/api/config"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/officeapi.v1.OfficeApi/GetConfig"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *OfficeApiHTTPClientImpl) GetWopiFileInfoEditable(ctx context.Context, in *WOPIRequest, opts ...http.CallOption) (*WOPIEditableResp, error) {
	var out WOPIEditableResp
	pattern := "/api/wopi/files/{file_name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/officeapi.v1.OfficeApi/GetWopiFileInfoEditable"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
