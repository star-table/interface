// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type PermissionHTTPServer interface {
	AddCollaboratorUser(context.Context, *CollaboratorUserRequest) (*CollaboratorUserResponse, error)
	BatchMigrateCollaboratorUsersToPermission(context.Context, *BatchMigrateRequest) (*emptypb.Empty, error)
	BatchRemoveCollaboratorUsers(context.Context, *CollaboratorUsersRequest) (*CollaboratorUsersResponse, error)
	BatchRemoveMemberFieldCollaborators(context.Context, *MemberFieldRequest) (*MemberFieldResponse, error)
	ChangeCollaboratorUser(context.Context, *CollaboratorUserRequest) (*CollaboratorUserResponse, error)
	CopyCollaboratorsFromColumnToOther(context.Context, *CopyColumnRequest) (*CopyColumnResponse, error)
	GetAppIdsByCollaborators(context.Context, *GetAppIdsByCollaboratorReq) (*GetAppIdsByCollaboratorResp, error)
	GetCollaboratorUserIdsAndDeptIds(context.Context, *GetCollaboratorUserIdsAndDeptIdsRequest) (*GetCollaboratorUserIdsAndDeptIdsResponse, error)
	GetCollaborators(context.Context, *GetCollaboratorsRequest) (*GetCollaboratorsResponse, error)
	GetDataCollaborators(context.Context, *GetDataCollaboratorsRequest) (*GetDataCollaboratorsResponse, error)
	GetUserCollaboratorRoleIds(context.Context, *GetUserCollaboratorRoleIdsRequest) (*GetUserCollaboratorRoleIdsResponse, error)
	RemoveCollaboratorUser(context.Context, *CollaboratorUserRequest) (*CollaboratorUsersResponse, error)
	SyncCollaboratorUsers(context.Context, *CollaboratorUsersRequest) (*CollaboratorUsersResponse, error)
}

func RegisterPermissionHTTPServer(s *http.Server, srv PermissionHTTPServer) {
	r := s.Route("/")
	r.POST("/permission/api/v1/collaboratorPermission/collaboratorUserIdsAndDeptIds", _Permission_GetCollaboratorUserIdsAndDeptIds0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/collaborators", _Permission_GetCollaborators0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/getUserCollaboratorRoleIds", _Permission_GetUserCollaboratorRoleIds0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/changeCollaboratorUser", _Permission_ChangeCollaboratorUser0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/addCollaboratorUser", _Permission_AddCollaboratorUser0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/removeCollaboratorUser", _Permission_RemoveCollaboratorUser0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/syncCollaboratorUsers", _Permission_SyncCollaboratorUsers0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/batchRemoveCollaboratorUsers", _Permission_BatchRemoveCollaboratorUsers0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/batchRemoveMemberFieldCollaborators", _Permission_BatchRemoveMemberFieldCollaborators0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/batchMigrateCollaboratorUsersToPermission", _Permission_BatchMigrateCollaboratorUsersToPermission0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/copyColumn", _Permission_CopyCollaboratorsFromColumnToOther0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/getDataCollaborators", _Permission_GetDataCollaborators0_HTTP_Handler(srv))
	r.POST("/permission/api/v1/collaboratorPermission/getAppIdsByCollaborators", _Permission_GetAppIdsByCollaborators0_HTTP_Handler(srv))
}

func _Permission_GetCollaboratorUserIdsAndDeptIds0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCollaboratorUserIdsAndDeptIdsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/GetCollaboratorUserIdsAndDeptIds")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCollaboratorUserIdsAndDeptIds(ctx, req.(*GetCollaboratorUserIdsAndDeptIdsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCollaboratorUserIdsAndDeptIdsResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_GetCollaborators0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCollaboratorsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/GetCollaborators")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCollaborators(ctx, req.(*GetCollaboratorsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCollaboratorsResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_GetUserCollaboratorRoleIds0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserCollaboratorRoleIdsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/GetUserCollaboratorRoleIds")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserCollaboratorRoleIds(ctx, req.(*GetUserCollaboratorRoleIdsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserCollaboratorRoleIdsResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_ChangeCollaboratorUser0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollaboratorUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/ChangeCollaboratorUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeCollaboratorUser(ctx, req.(*CollaboratorUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollaboratorUserResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_AddCollaboratorUser0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollaboratorUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/AddCollaboratorUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCollaboratorUser(ctx, req.(*CollaboratorUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollaboratorUserResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_RemoveCollaboratorUser0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollaboratorUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/RemoveCollaboratorUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveCollaboratorUser(ctx, req.(*CollaboratorUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollaboratorUsersResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_SyncCollaboratorUsers0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollaboratorUsersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/SyncCollaboratorUsers")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SyncCollaboratorUsers(ctx, req.(*CollaboratorUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollaboratorUsersResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_BatchRemoveCollaboratorUsers0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CollaboratorUsersRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/BatchRemoveCollaboratorUsers")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchRemoveCollaboratorUsers(ctx, req.(*CollaboratorUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CollaboratorUsersResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_BatchRemoveMemberFieldCollaborators0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MemberFieldRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/BatchRemoveMemberFieldCollaborators")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchRemoveMemberFieldCollaborators(ctx, req.(*MemberFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MemberFieldResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_BatchMigrateCollaboratorUsersToPermission0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BatchMigrateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/BatchMigrateCollaboratorUsersToPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BatchMigrateCollaboratorUsersToPermission(ctx, req.(*BatchMigrateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Permission_CopyCollaboratorsFromColumnToOther0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CopyColumnRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/CopyCollaboratorsFromColumnToOther")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CopyCollaboratorsFromColumnToOther(ctx, req.(*CopyColumnRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CopyColumnResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_GetDataCollaborators0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDataCollaboratorsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/GetDataCollaborators")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDataCollaborators(ctx, req.(*GetDataCollaboratorsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDataCollaboratorsResponse)
		return ctx.Result(200, reply)
	}
}

func _Permission_GetAppIdsByCollaborators0_HTTP_Handler(srv PermissionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAppIdsByCollaboratorReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/permission.v1.Permission/GetAppIdsByCollaborators")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAppIdsByCollaborators(ctx, req.(*GetAppIdsByCollaboratorReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAppIdsByCollaboratorResp)
		return ctx.Result(200, reply)
	}
}

type PermissionHTTPClient interface {
	AddCollaboratorUser(ctx context.Context, req *CollaboratorUserRequest, opts ...http.CallOption) (rsp *CollaboratorUserResponse, err error)
	BatchMigrateCollaboratorUsersToPermission(ctx context.Context, req *BatchMigrateRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	BatchRemoveCollaboratorUsers(ctx context.Context, req *CollaboratorUsersRequest, opts ...http.CallOption) (rsp *CollaboratorUsersResponse, err error)
	BatchRemoveMemberFieldCollaborators(ctx context.Context, req *MemberFieldRequest, opts ...http.CallOption) (rsp *MemberFieldResponse, err error)
	ChangeCollaboratorUser(ctx context.Context, req *CollaboratorUserRequest, opts ...http.CallOption) (rsp *CollaboratorUserResponse, err error)
	CopyCollaboratorsFromColumnToOther(ctx context.Context, req *CopyColumnRequest, opts ...http.CallOption) (rsp *CopyColumnResponse, err error)
	GetAppIdsByCollaborators(ctx context.Context, req *GetAppIdsByCollaboratorReq, opts ...http.CallOption) (rsp *GetAppIdsByCollaboratorResp, err error)
	GetCollaboratorUserIdsAndDeptIds(ctx context.Context, req *GetCollaboratorUserIdsAndDeptIdsRequest, opts ...http.CallOption) (rsp *GetCollaboratorUserIdsAndDeptIdsResponse, err error)
	GetCollaborators(ctx context.Context, req *GetCollaboratorsRequest, opts ...http.CallOption) (rsp *GetCollaboratorsResponse, err error)
	GetDataCollaborators(ctx context.Context, req *GetDataCollaboratorsRequest, opts ...http.CallOption) (rsp *GetDataCollaboratorsResponse, err error)
	GetUserCollaboratorRoleIds(ctx context.Context, req *GetUserCollaboratorRoleIdsRequest, opts ...http.CallOption) (rsp *GetUserCollaboratorRoleIdsResponse, err error)
	RemoveCollaboratorUser(ctx context.Context, req *CollaboratorUserRequest, opts ...http.CallOption) (rsp *CollaboratorUsersResponse, err error)
	SyncCollaboratorUsers(ctx context.Context, req *CollaboratorUsersRequest, opts ...http.CallOption) (rsp *CollaboratorUsersResponse, err error)
}

type PermissionHTTPClientImpl struct {
	cc *http.Client
}

func NewPermissionHTTPClient(client *http.Client) PermissionHTTPClient {
	return &PermissionHTTPClientImpl{client}
}

func (c *PermissionHTTPClientImpl) AddCollaboratorUser(ctx context.Context, in *CollaboratorUserRequest, opts ...http.CallOption) (*CollaboratorUserResponse, error) {
	var out CollaboratorUserResponse
	pattern := "/permission/api/v1/collaboratorPermission/addCollaboratorUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/AddCollaboratorUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) BatchMigrateCollaboratorUsersToPermission(ctx context.Context, in *BatchMigrateRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/permission/api/v1/collaboratorPermission/batchMigrateCollaboratorUsersToPermission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/BatchMigrateCollaboratorUsersToPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) BatchRemoveCollaboratorUsers(ctx context.Context, in *CollaboratorUsersRequest, opts ...http.CallOption) (*CollaboratorUsersResponse, error) {
	var out CollaboratorUsersResponse
	pattern := "/permission/api/v1/collaboratorPermission/batchRemoveCollaboratorUsers"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/BatchRemoveCollaboratorUsers"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) BatchRemoveMemberFieldCollaborators(ctx context.Context, in *MemberFieldRequest, opts ...http.CallOption) (*MemberFieldResponse, error) {
	var out MemberFieldResponse
	pattern := "/permission/api/v1/collaboratorPermission/batchRemoveMemberFieldCollaborators"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/BatchRemoveMemberFieldCollaborators"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) ChangeCollaboratorUser(ctx context.Context, in *CollaboratorUserRequest, opts ...http.CallOption) (*CollaboratorUserResponse, error) {
	var out CollaboratorUserResponse
	pattern := "/permission/api/v1/collaboratorPermission/changeCollaboratorUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/ChangeCollaboratorUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) CopyCollaboratorsFromColumnToOther(ctx context.Context, in *CopyColumnRequest, opts ...http.CallOption) (*CopyColumnResponse, error) {
	var out CopyColumnResponse
	pattern := "/permission/api/v1/collaboratorPermission/copyColumn"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/CopyCollaboratorsFromColumnToOther"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) GetAppIdsByCollaborators(ctx context.Context, in *GetAppIdsByCollaboratorReq, opts ...http.CallOption) (*GetAppIdsByCollaboratorResp, error) {
	var out GetAppIdsByCollaboratorResp
	pattern := "/permission/api/v1/collaboratorPermission/getAppIdsByCollaborators"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/GetAppIdsByCollaborators"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) GetCollaboratorUserIdsAndDeptIds(ctx context.Context, in *GetCollaboratorUserIdsAndDeptIdsRequest, opts ...http.CallOption) (*GetCollaboratorUserIdsAndDeptIdsResponse, error) {
	var out GetCollaboratorUserIdsAndDeptIdsResponse
	pattern := "/permission/api/v1/collaboratorPermission/collaboratorUserIdsAndDeptIds"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/GetCollaboratorUserIdsAndDeptIds"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) GetCollaborators(ctx context.Context, in *GetCollaboratorsRequest, opts ...http.CallOption) (*GetCollaboratorsResponse, error) {
	var out GetCollaboratorsResponse
	pattern := "/permission/api/v1/collaboratorPermission/collaborators"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/GetCollaborators"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) GetDataCollaborators(ctx context.Context, in *GetDataCollaboratorsRequest, opts ...http.CallOption) (*GetDataCollaboratorsResponse, error) {
	var out GetDataCollaboratorsResponse
	pattern := "/permission/api/v1/collaboratorPermission/getDataCollaborators"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/GetDataCollaborators"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) GetUserCollaboratorRoleIds(ctx context.Context, in *GetUserCollaboratorRoleIdsRequest, opts ...http.CallOption) (*GetUserCollaboratorRoleIdsResponse, error) {
	var out GetUserCollaboratorRoleIdsResponse
	pattern := "/permission/api/v1/collaboratorPermission/getUserCollaboratorRoleIds"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/GetUserCollaboratorRoleIds"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) RemoveCollaboratorUser(ctx context.Context, in *CollaboratorUserRequest, opts ...http.CallOption) (*CollaboratorUsersResponse, error) {
	var out CollaboratorUsersResponse
	pattern := "/permission/api/v1/collaboratorPermission/removeCollaboratorUser"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/RemoveCollaboratorUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionHTTPClientImpl) SyncCollaboratorUsers(ctx context.Context, in *CollaboratorUsersRequest, opts ...http.CallOption) (*CollaboratorUsersResponse, error) {
	var out CollaboratorUsersResponse
	pattern := "/permission/api/v1/collaboratorPermission/syncCollaboratorUsers"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/permission.v1.Permission/SyncCollaboratorUsers"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
