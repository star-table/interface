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

type BatchHTTPServer interface {
	AddDashBord(context.Context, *HelloRequest) (*HelloReply, error)
	AddHasOverdueView(context.Context, *HelloRequest) (*HelloReply, error)
	AddPGIndex(context.Context, *HelloRequest) (*HelloReply, error)
	AddPgStatisticsTable(context.Context, *HelloRequest) (*HelloReply, error)
	AutomationRebuildActive(context.Context, *HelloRequest) (*HelloReply, error)
	ChangeDomain(context.Context, *HelloRequest) (*HelloReply, error)
	ClearSameUser(context.Context, *HelloRequest) (*HelloReply, error)
	DbMigrations(context.Context, *HelloRequest) (*HelloReply, error)
	ExchangeRelatingDataId(context.Context, *HelloRequest) (*HelloReply, error)
	ExportPGData(context.Context, *HelloRequest) (*HelloReply, error)
	Fix20221104TemplatePath(context.Context, *HelloRequest) (*HelloReply, error)
	Fix20221107AppCreatorAddMember(context.Context, *HelloRequest) (*HelloReply, error)
	FixIssueForDelPro(context.Context, *HelloRequest) (*HelloReply, error)
	FixIssueStatusLostWhenMove(context.Context, *HelloRequest) (*HelloReply, error)
	FixIssueStatusName(context.Context, *HelloRequest) (*HelloReply, error)
	FixPGOrgId(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateAuditStatus(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateCollaborator(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateDropUnusedTables(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateGlobalUser(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateGroupChatSetting(context.Context, *HelloRequest) (*HelloReply, error)
	MigratePath(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateRelating(context.Context, *HelloRequest) (*HelloReply, error)
	MigrateWorkHour(context.Context, *HelloRequest) (*HelloReply, error)
	PgDataTransfer(context.Context, *HelloRequest) (*HelloReply, error)
	SplitTableHeader(context.Context, *HelloRequest) (*HelloReply, error)
	Temp(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterBatchHTTPServer(s *http.Server, srv BatchHTTPServer) {
	r := s.Route("/")
	r.GET("/zo4wrrNYBCx997AS/{orgId}", _Batch_Temp0_HTTP_Handler(srv))
	r.GET("/T33Roh7eqyTgkH9E/{orgId}", _Batch_ExportPGData0_HTTP_Handler(srv))
	r.GET("/1r0fyA2tabUFkeXb/{orgId}", _Batch_SplitTableHeader0_HTTP_Handler(srv))
	r.GET("/WYHN6nAEunvaS3iQ/{orgId}", _Batch_MigrateRelating0_HTTP_Handler(srv))
	r.GET("/O8IuISTBeAfMlWvX", _Batch_MigrateWorkHour0_HTTP_Handler(srv))
	r.GET("/mzg0EAe9oAOX3AUd/{orgId}", _Batch_MigrateCollaborator0_HTTP_Handler(srv))
	r.GET("/ylRhNgOy9XDQkNhE/{orgId}", _Batch_MigrateDropUnusedTables0_HTTP_Handler(srv))
	r.GET("/fP6gFZV3h6aYqHwZ/{orgId}", _Batch_FixIssueStatusLostWhenMove0_HTTP_Handler(srv))
	r.GET("/tgCE4dge24h5xeZO/{orgId}", _Batch_AddPGIndex0_HTTP_Handler(srv))
	r.GET("/5sSPOJy1iipLJheP", _Batch_MigrateGlobalUser0_HTTP_Handler(srv))
	r.GET("/9DDarUt4W4gHAlcm/{orgId}", _Batch_FixPGOrgId0_HTTP_Handler(srv))
	r.GET("/LGRhCYWbFgIiGFlh/{orgId}", _Batch_MigratePath0_HTTP_Handler(srv))
	r.GET("/yo12qhJyT9H3wlja/{orgId}", _Batch_MigrateGroupChatSetting0_HTTP_Handler(srv))
	r.GET("/aE4B6nt3Ul29v5LW/{orgId}", _Batch_FixIssueStatusName0_HTTP_Handler(srv))
	r.GET("/mbo5CUj9eC5J5gKF/{orgId}", _Batch_FixIssueForDelPro0_HTTP_Handler(srv))
	r.GET("/wkxNrIGBkoGfrmYT/{orgId}", _Batch_AddHasOverdueView0_HTTP_Handler(srv))
	r.GET("/BiWSsSuExXGfWFsX/{orgId}", _Batch_AddPgStatisticsTable0_HTTP_Handler(srv))
	r.GET("/jWY49FWFhOg6gY60/{orgId}", _Batch_AddDashBord0_HTTP_Handler(srv))
	r.GET("/ulUntNQbqNTNNjHU/{orgId}", _Batch_ChangeDomain0_HTTP_Handler(srv))
	r.GET("/PggZn5V2D8UngAEw/{orgId}", _Batch_PgDataTransfer0_HTTP_Handler(srv))
	r.GET("/exdIFVTYf3D7Cba3/{orgId}", _Batch_ClearSameUser0_HTTP_Handler(srv))
	r.GET("/ltivMRSdtcgNfMeH/{orgId}", _Batch_MigrateAuditStatus0_HTTP_Handler(srv))
	r.GET("/31LMM8SI61KRmKau/{orgId}", _Batch_DbMigrations0_HTTP_Handler(srv))
	r.GET("/2cl5bZ5gWaA5IE0N/{orgId}", _Batch_Fix20221104TemplatePath0_HTTP_Handler(srv))
	r.GET("/65SjhvFSjtG7fchD/{orgId}", _Batch_Fix20221107AppCreatorAddMember0_HTTP_Handler(srv))
	r.GET("/3y0gVQVAu8hBzh1J/{orgId}", _Batch_AutomationRebuildActive0_HTTP_Handler(srv))
	r.GET("/EfIwZ5oc0dBAGBP7/{orgId}", _Batch_ExchangeRelatingDataId0_HTTP_Handler(srv))
}

func _Batch_Temp0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/Temp")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Temp(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_ExportPGData0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/ExportPGData")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExportPGData(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_SplitTableHeader0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/SplitTableHeader")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SplitTableHeader(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateRelating0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateRelating")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateRelating(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateWorkHour0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateWorkHour")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateWorkHour(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateCollaborator0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateCollaborator")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateCollaborator(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateDropUnusedTables0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateDropUnusedTables")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateDropUnusedTables(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_FixIssueStatusLostWhenMove0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/FixIssueStatusLostWhenMove")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FixIssueStatusLostWhenMove(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_AddPGIndex0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/AddPGIndex")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddPGIndex(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateGlobalUser0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateGlobalUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateGlobalUser(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_FixPGOrgId0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/FixPGOrgId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FixPGOrgId(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigratePath0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigratePath")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigratePath(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateGroupChatSetting0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateGroupChatSetting")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateGroupChatSetting(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_FixIssueStatusName0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/FixIssueStatusName")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FixIssueStatusName(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_FixIssueForDelPro0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/FixIssueForDelPro")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FixIssueForDelPro(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_AddHasOverdueView0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/AddHasOverdueView")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddHasOverdueView(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_AddPgStatisticsTable0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/AddPgStatisticsTable")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddPgStatisticsTable(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_AddDashBord0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/AddDashBord")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddDashBord(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_ChangeDomain0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/ChangeDomain")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeDomain(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_PgDataTransfer0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/PgDataTransfer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PgDataTransfer(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_ClearSameUser0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/ClearSameUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ClearSameUser(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_MigrateAuditStatus0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/MigrateAuditStatus")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MigrateAuditStatus(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_DbMigrations0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/DbMigrations")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DbMigrations(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_Fix20221104TemplatePath0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/Fix20221104TemplatePath")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Fix20221104TemplatePath(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_Fix20221107AppCreatorAddMember0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/Fix20221107AppCreatorAddMember")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Fix20221107AppCreatorAddMember(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_AutomationRebuildActive0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/AutomationRebuildActive")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AutomationRebuildActive(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

func _Batch_ExchangeRelatingDataId0_HTTP_Handler(srv BatchHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/batch.v1.Batch/ExchangeRelatingDataId")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExchangeRelatingDataId(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type BatchHTTPClient interface {
	AddDashBord(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	AddHasOverdueView(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	AddPGIndex(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	AddPgStatisticsTable(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	AutomationRebuildActive(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	ChangeDomain(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	ClearSameUser(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	DbMigrations(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	ExchangeRelatingDataId(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	ExportPGData(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	Fix20221104TemplatePath(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	Fix20221107AppCreatorAddMember(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	FixIssueForDelPro(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	FixIssueStatusLostWhenMove(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	FixIssueStatusName(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	FixPGOrgId(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateAuditStatus(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateCollaborator(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateDropUnusedTables(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateGlobalUser(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateGroupChatSetting(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigratePath(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateRelating(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	MigrateWorkHour(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	PgDataTransfer(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	SplitTableHeader(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
	Temp(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type BatchHTTPClientImpl struct {
	cc *http.Client
}

func NewBatchHTTPClient(client *http.Client) BatchHTTPClient {
	return &BatchHTTPClientImpl{client}
}

func (c *BatchHTTPClientImpl) AddDashBord(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/jWY49FWFhOg6gY60/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/AddDashBord"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) AddHasOverdueView(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/wkxNrIGBkoGfrmYT/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/AddHasOverdueView"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) AddPGIndex(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/tgCE4dge24h5xeZO/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/AddPGIndex"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) AddPgStatisticsTable(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/BiWSsSuExXGfWFsX/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/AddPgStatisticsTable"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) AutomationRebuildActive(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/3y0gVQVAu8hBzh1J/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/AutomationRebuildActive"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) ChangeDomain(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/ulUntNQbqNTNNjHU/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/ChangeDomain"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) ClearSameUser(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/exdIFVTYf3D7Cba3/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/ClearSameUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) DbMigrations(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/31LMM8SI61KRmKau/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/DbMigrations"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) ExchangeRelatingDataId(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/EfIwZ5oc0dBAGBP7/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/ExchangeRelatingDataId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) ExportPGData(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/T33Roh7eqyTgkH9E/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/ExportPGData"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) Fix20221104TemplatePath(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/2cl5bZ5gWaA5IE0N/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/Fix20221104TemplatePath"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) Fix20221107AppCreatorAddMember(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/65SjhvFSjtG7fchD/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/Fix20221107AppCreatorAddMember"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) FixIssueForDelPro(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/mbo5CUj9eC5J5gKF/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/FixIssueForDelPro"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) FixIssueStatusLostWhenMove(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/fP6gFZV3h6aYqHwZ/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/FixIssueStatusLostWhenMove"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) FixIssueStatusName(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/aE4B6nt3Ul29v5LW/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/FixIssueStatusName"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) FixPGOrgId(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/9DDarUt4W4gHAlcm/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/FixPGOrgId"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateAuditStatus(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/ltivMRSdtcgNfMeH/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateAuditStatus"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateCollaborator(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/mzg0EAe9oAOX3AUd/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateCollaborator"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateDropUnusedTables(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/ylRhNgOy9XDQkNhE/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateDropUnusedTables"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateGlobalUser(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/5sSPOJy1iipLJheP"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateGlobalUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateGroupChatSetting(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/yo12qhJyT9H3wlja/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateGroupChatSetting"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigratePath(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/LGRhCYWbFgIiGFlh/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigratePath"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateRelating(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/WYHN6nAEunvaS3iQ/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateRelating"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) MigrateWorkHour(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/O8IuISTBeAfMlWvX"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/MigrateWorkHour"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) PgDataTransfer(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/PggZn5V2D8UngAEw/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/PgDataTransfer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) SplitTableHeader(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/1r0fyA2tabUFkeXb/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/SplitTableHeader"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BatchHTTPClientImpl) Temp(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/zo4wrrNYBCx997AS/{orgId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/batch.v1.Batch/Temp"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}