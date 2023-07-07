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

type UserActionHTTPServer interface {
	CommentDocument(context.Context, *CommentDocumentRequest) (*CommentDocumentReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentReply, error)
	LikeDocument(context.Context, *LikeDocumentRequest) (*LikeDocumentReply, error)
	ListDocumentComment(context.Context, *ListDocumentCommentRequest) (*ListDocumentCommentReply, error)
	ListDocumentLikeUser(context.Context, *ListDocumentLikeUserRequest) (*ListDocumentLikeUserReply, error)
	StarSpace(context.Context, *StarSpaceRequest) (*StarSpaceReply, error)
	TopDocument(context.Context, *TopDocumentRequest) (*TopDocumentReply, error)
	UpdateCommentStatus(context.Context, *UpdateCommentStatusRequest) (*UpdateCommentStatusReply, error)
}

func RegisterUserActionHTTPServer(s *http.Server, srv UserActionHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/action/document/like", _UserAction_LikeDocument0_HTTP_Handler(srv))
	r.POST("/v1/action/document/like/user/list", _UserAction_ListDocumentLikeUser0_HTTP_Handler(srv))
	r.POST("/v1/action/document/top", _UserAction_TopDocument0_HTTP_Handler(srv))
	r.POST("/v1/action/document/comment", _UserAction_CommentDocument0_HTTP_Handler(srv))
	r.POST("/v1/action/document/comment/status/update", _UserAction_UpdateCommentStatus0_HTTP_Handler(srv))
	r.POST("/v1/action/comment/delete", _UserAction_DeleteComment0_HTTP_Handler(srv))
	r.POST("/v1/action/document/comment/list", _UserAction_ListDocumentComment0_HTTP_Handler(srv))
	r.POST("/v1/action/space/star", _UserAction_StarSpace0_HTTP_Handler(srv))
}

func _UserAction_LikeDocument0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LikeDocumentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/LikeDocument")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LikeDocument(ctx, req.(*LikeDocumentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LikeDocumentReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_ListDocumentLikeUser0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDocumentLikeUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/ListDocumentLikeUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDocumentLikeUser(ctx, req.(*ListDocumentLikeUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDocumentLikeUserReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_TopDocument0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TopDocumentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/TopDocument")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TopDocument(ctx, req.(*TopDocumentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TopDocumentReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_CommentDocument0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CommentDocumentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/CommentDocument")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CommentDocument(ctx, req.(*CommentDocumentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommentDocumentReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_UpdateCommentStatus0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCommentStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/UpdateCommentStatus")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCommentStatus(ctx, req.(*UpdateCommentStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCommentStatusReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_DeleteComment0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/DeleteComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCommentReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_ListDocumentComment0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDocumentCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/ListDocumentComment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDocumentComment(ctx, req.(*ListDocumentCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDocumentCommentReply)
		return ctx.Result(200, reply)
	}
}

func _UserAction_StarSpace0_HTTP_Handler(srv UserActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StarSpaceRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/knowledge.v1.UserAction/StarSpace")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StarSpace(ctx, req.(*StarSpaceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StarSpaceReply)
		return ctx.Result(200, reply)
	}
}

type UserActionHTTPClient interface {
	CommentDocument(ctx context.Context, req *CommentDocumentRequest, opts ...http.CallOption) (rsp *CommentDocumentReply, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *DeleteCommentReply, err error)
	LikeDocument(ctx context.Context, req *LikeDocumentRequest, opts ...http.CallOption) (rsp *LikeDocumentReply, err error)
	ListDocumentComment(ctx context.Context, req *ListDocumentCommentRequest, opts ...http.CallOption) (rsp *ListDocumentCommentReply, err error)
	ListDocumentLikeUser(ctx context.Context, req *ListDocumentLikeUserRequest, opts ...http.CallOption) (rsp *ListDocumentLikeUserReply, err error)
	StarSpace(ctx context.Context, req *StarSpaceRequest, opts ...http.CallOption) (rsp *StarSpaceReply, err error)
	TopDocument(ctx context.Context, req *TopDocumentRequest, opts ...http.CallOption) (rsp *TopDocumentReply, err error)
	UpdateCommentStatus(ctx context.Context, req *UpdateCommentStatusRequest, opts ...http.CallOption) (rsp *UpdateCommentStatusReply, err error)
}

type UserActionHTTPClientImpl struct {
	cc *http.Client
}

func NewUserActionHTTPClient(client *http.Client) UserActionHTTPClient {
	return &UserActionHTTPClientImpl{client}
}

func (c *UserActionHTTPClientImpl) CommentDocument(ctx context.Context, in *CommentDocumentRequest, opts ...http.CallOption) (*CommentDocumentReply, error) {
	var out CommentDocumentReply
	pattern := "/v1/action/document/comment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/CommentDocument"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*DeleteCommentReply, error) {
	var out DeleteCommentReply
	pattern := "/v1/action/comment/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/DeleteComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) LikeDocument(ctx context.Context, in *LikeDocumentRequest, opts ...http.CallOption) (*LikeDocumentReply, error) {
	var out LikeDocumentReply
	pattern := "/v1/action/document/like"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/LikeDocument"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) ListDocumentComment(ctx context.Context, in *ListDocumentCommentRequest, opts ...http.CallOption) (*ListDocumentCommentReply, error) {
	var out ListDocumentCommentReply
	pattern := "/v1/action/document/comment/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/ListDocumentComment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) ListDocumentLikeUser(ctx context.Context, in *ListDocumentLikeUserRequest, opts ...http.CallOption) (*ListDocumentLikeUserReply, error) {
	var out ListDocumentLikeUserReply
	pattern := "/v1/action/document/like/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/ListDocumentLikeUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) StarSpace(ctx context.Context, in *StarSpaceRequest, opts ...http.CallOption) (*StarSpaceReply, error) {
	var out StarSpaceReply
	pattern := "/v1/action/space/star"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/StarSpace"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) TopDocument(ctx context.Context, in *TopDocumentRequest, opts ...http.CallOption) (*TopDocumentReply, error) {
	var out TopDocumentReply
	pattern := "/v1/action/document/top"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/TopDocument"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserActionHTTPClientImpl) UpdateCommentStatus(ctx context.Context, in *UpdateCommentStatusRequest, opts ...http.CallOption) (*UpdateCommentStatusReply, error) {
	var out UpdateCommentStatusReply
	pattern := "/v1/action/document/comment/status/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/knowledge.v1.UserAction/UpdateCommentStatus"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
