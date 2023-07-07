// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: officeapi.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on WOPIRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WOPIRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WOPIRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WOPIRequestMultiError, or
// nil if none found.
func (m *WOPIRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WOPIRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileName

	// no validation rules for AccessToken

	// no validation rules for Permission

	if len(errors) > 0 {
		return WOPIRequestMultiError(errors)
	}

	return nil
}

// WOPIRequestMultiError is an error wrapping multiple validation errors
// returned by WOPIRequest.ValidateAll() if the designated constraints aren't met.
type WOPIRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WOPIRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WOPIRequestMultiError) AllErrors() []error { return m }

// WOPIRequestValidationError is the validation error returned by
// WOPIRequest.Validate if the designated constraints aren't met.
type WOPIRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WOPIRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WOPIRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WOPIRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WOPIRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WOPIRequestValidationError) ErrorName() string { return "WOPIRequestValidationError" }

// Error satisfies the builtin error interface
func (e WOPIRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWOPIRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WOPIRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WOPIRequestValidationError{}

// Validate checks the field values on WOPIRequestFileContent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *WOPIRequestFileContent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WOPIRequestFileContent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WOPIRequestFileContentMultiError, or nil if none found.
func (m *WOPIRequestFileContent) ValidateAll() error {
	return m.validate(true)
}

func (m *WOPIRequestFileContent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileName

	// no validation rules for AccessToken

	// no validation rules for Permission

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WOPIRequestFileContentValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WOPIRequestFileContentValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WOPIRequestFileContentValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WOPIRequestFileContentMultiError(errors)
	}

	return nil
}

// WOPIRequestFileContentMultiError is an error wrapping multiple validation
// errors returned by WOPIRequestFileContent.ValidateAll() if the designated
// constraints aren't met.
type WOPIRequestFileContentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WOPIRequestFileContentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WOPIRequestFileContentMultiError) AllErrors() []error { return m }

// WOPIRequestFileContentValidationError is the validation error returned by
// WOPIRequestFileContent.Validate if the designated constraints aren't met.
type WOPIRequestFileContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WOPIRequestFileContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WOPIRequestFileContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WOPIRequestFileContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WOPIRequestFileContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WOPIRequestFileContentValidationError) ErrorName() string {
	return "WOPIRequestFileContentValidationError"
}

// Error satisfies the builtin error interface
func (e WOPIRequestFileContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWOPIRequestFileContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WOPIRequestFileContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WOPIRequestFileContentValidationError{}

// Validate checks the field values on FormConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FormConfigResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FormConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FormConfigResponseMultiError, or nil if none found.
func (m *FormConfigResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FormConfigResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FormConfigResponseMultiError(errors)
	}

	return nil
}

// FormConfigResponseMultiError is an error wrapping multiple validation errors
// returned by FormConfigResponse.ValidateAll() if the designated constraints
// aren't met.
type FormConfigResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FormConfigResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FormConfigResponseMultiError) AllErrors() []error { return m }

// FormConfigResponseValidationError is the validation error returned by
// FormConfigResponse.Validate if the designated constraints aren't met.
type FormConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormConfigResponseValidationError) ErrorName() string {
	return "FormConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FormConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormConfigResponseValidationError{}

// Validate checks the field values on WOPIEditableResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WOPIEditableResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WOPIEditableResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WOPIEditableRespMultiError, or nil if none found.
func (m *WOPIEditableResp) ValidateAll() error {
	return m.validate(true)
}

func (m *WOPIEditableResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BaseFileName

	// no validation rules for OwnerId

	// no validation rules for UserId

	// no validation rules for Size

	// no validation rules for UserFriendlyName

	// no validation rules for UserCanWrite

	// no validation rules for UserCanNotWriteRelative

	// no validation rules for Status

	// no validation rules for LastModifiedTime

	if len(errors) > 0 {
		return WOPIEditableRespMultiError(errors)
	}

	return nil
}

// WOPIEditableRespMultiError is an error wrapping multiple validation errors
// returned by WOPIEditableResp.ValidateAll() if the designated constraints
// aren't met.
type WOPIEditableRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WOPIEditableRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WOPIEditableRespMultiError) AllErrors() []error { return m }

// WOPIEditableRespValidationError is the validation error returned by
// WOPIEditableResp.Validate if the designated constraints aren't met.
type WOPIEditableRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WOPIEditableRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WOPIEditableRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WOPIEditableRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WOPIEditableRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WOPIEditableRespValidationError) ErrorName() string { return "WOPIEditableRespValidationError" }

// Error satisfies the builtin error interface
func (e WOPIEditableRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWOPIEditableResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WOPIEditableRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WOPIEditableRespValidationError{}

// Validate checks the field values on GetConfigRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetConfigRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConfigRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConfigRequestMultiError, or nil if none found.
func (m *GetConfigRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConfigRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetConfigRequestMultiError(errors)
	}

	return nil
}

// GetConfigRequestMultiError is an error wrapping multiple validation errors
// returned by GetConfigRequest.ValidateAll() if the designated constraints
// aren't met.
type GetConfigRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConfigRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConfigRequestMultiError) AllErrors() []error { return m }

// GetConfigRequestValidationError is the validation error returned by
// GetConfigRequest.Validate if the designated constraints aren't met.
type GetConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConfigRequestValidationError) ErrorName() string { return "GetConfigRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConfigRequestValidationError{}

// Validate checks the field values on GetConfigResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetConfigResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConfigResponseMultiError, or nil if none found.
func (m *GetConfigResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConfigResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for UrlPrefix

	// no validation rules for FileExtensions

	if len(errors) > 0 {
		return GetConfigResponseMultiError(errors)
	}

	return nil
}

// GetConfigResponseMultiError is an error wrapping multiple validation errors
// returned by GetConfigResponse.ValidateAll() if the designated constraints
// aren't met.
type GetConfigResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConfigResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConfigResponseMultiError) AllErrors() []error { return m }

// GetConfigResponseValidationError is the validation error returned by
// GetConfigResponse.Validate if the designated constraints aren't met.
type GetConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConfigResponseValidationError) ErrorName() string {
	return "GetConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConfigResponseValidationError{}
