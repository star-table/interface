// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package ping

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

type PingHTTPServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

func RegisterPingHTTPServer(s *http.Server, srv PingHTTPServer) {
	r := s.Route("/")
	r.GET("/ping", _Ping_Ping0_HTTP_Handler(srv))
}

func _Ping_Ping0_HTTP_Handler(srv PingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/ping.Ping/Ping")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*PingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingReply)
		return ctx.Result(200, reply)
	}
}

type PingHTTPClient interface {
	Ping(ctx context.Context, req *PingRequest, opts ...http.CallOption) (rsp *PingReply, err error)
}

type PingHTTPClientImpl struct {
	cc *http.Client
}

func NewPingHTTPClient(client *http.Client) PingHTTPClient {
	return &PingHTTPClientImpl{client}
}

func (c *PingHTTPClientImpl) Ping(ctx context.Context, in *PingRequest, opts ...http.CallOption) (*PingReply, error) {
	var out PingReply
	pattern := "/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/ping.Ping/Ping"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
