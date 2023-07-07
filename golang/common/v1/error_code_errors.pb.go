// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsDuplicateOperation(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_OPERATION.String() && e.Code == 503
}

func ErrorDuplicateOperation(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ErrorReason_DUPLICATE_OPERATION.String(), fmt.Sprintf(format, args...))
}

func IsResourceNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RESOURCE_NOT_EXIST.String() && e.Code == 404
}

func ErrorResourceNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_RESOURCE_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsServerInternal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SERVER_INTERNAL.String() && e.Code == 500
}

func ErrorServerInternal(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SERVER_INTERNAL.String(), fmt.Sprintf(format, args...))
}

func IsDtxNeedRollback(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DTX_NEED_ROLLBACK.String() && e.Code == 409
}

func ErrorDtxNeedRollback(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReason_DTX_NEED_ROLLBACK.String(), fmt.Sprintf(format, args...))
}

func IsParamsNotCorrect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PARAMS_NOT_CORRECT.String() && e.Code == 400
}

func ErrorParamsNotCorrect(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_PARAMS_NOT_CORRECT.String(), fmt.Sprintf(format, args...))
}

func IsForbidenOperation(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_FORBIDEN_OPERATION.String() && e.Code == 403
}

func ErrorForbidenOperation(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_FORBIDEN_OPERATION.String(), fmt.Sprintf(format, args...))
}

func IsPermissionUserIdNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PERMISSION_USER_ID_NOT_FOUND.String() && e.Code == 500
}

func ErrorPermissionUserIdNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_PERMISSION_USER_ID_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsPermissionUserIdParseFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PERMISSION_USER_ID_PARSE_FAILED.String() && e.Code == 500
}

func ErrorPermissionUserIdParseFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_PERMISSION_USER_ID_PARSE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsPermissionOrgIdNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PERMISSION_ORG_ID_NOT_FOUND.String() && e.Code == 500
}

func ErrorPermissionOrgIdNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_PERMISSION_ORG_ID_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsPermissionOrgIdParseFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PERMISSION_ORG_ID_PARSE_FAILED.String() && e.Code == 500
}

func ErrorPermissionOrgIdParseFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_PERMISSION_ORG_ID_PARSE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsPermissionContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PERMISSION_CONTENT_MISSING.String() && e.Code == 500
}

func ErrorPermissionContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_PERMISSION_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsRowDataValidateFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ROW_DATA_VALIDATE_FAILED.String() && e.Code == 400
}

func ErrorRowDataValidateFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_ROW_DATA_VALIDATE_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsCanNotDeleteLastTable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CAN_NOT_DELETE_LAST_TABLE.String() && e.Code == 403
}

func ErrorCanNotDeleteLastTable(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_CAN_NOT_DELETE_LAST_TABLE.String(), fmt.Sprintf(format, args...))
}

func IsCanNotDeleteByUse(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CAN_NOT_DELETE_BY_USE.String() && e.Code == 403
}

func ErrorCanNotDeleteByUse(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_CAN_NOT_DELETE_BY_USE.String(), fmt.Sprintf(format, args...))
}

func IsAutomationInvalidWorkflow(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AUTOMATION_INVALID_WORKFLOW.String() && e.Code == 403
}

func ErrorAutomationInvalidWorkflow(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_AUTOMATION_INVALID_WORKFLOW.String(), fmt.Sprintf(format, args...))
}

func IsAutomationExecutionCanceled(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AUTOMATION_EXECUTION_CANCELED.String() && e.Code == 403
}

func ErrorAutomationExecutionCanceled(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_AUTOMATION_EXECUTION_CANCELED.String(), fmt.Sprintf(format, args...))
}
