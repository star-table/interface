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

type N8NLesscodeHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterN8NLesscodeHTTPServer(s *http.Server, srv N8NLesscodeHTTPServer) {
	r := s.Route("/")
	r.GET("/hello", _N8NLesscode_SayHello0_HTTP_Handler(srv))
}

func _N8NLesscode_SayHello0_HTTP_Handler(srv N8NLesscodeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/n8n_lesscode.v1.N8nLesscode/SayHello")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type N8NLesscodeHTTPClient interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type N8NLesscodeHTTPClientImpl struct {
	cc *http.Client
}

func NewN8NLesscodeHTTPClient(client *http.Client) N8NLesscodeHTTPClient {
	return &N8NLesscodeHTTPClientImpl{client}
}

func (c *N8NLesscodeHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/hello"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/n8n_lesscode.v1.N8nLesscode/SayHello"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}