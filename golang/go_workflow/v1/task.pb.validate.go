// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: task.proto

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

// Validate checks the field values on CompleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CompleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CompleteRequestMultiError, or nil if none found.
func (m *CompleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := CompleteRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := CompleteRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUserName()) < 1 {
		err := CompleteRequestValidationError{
			field:  "UserName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Pass

	if m.GetOrgId() <= 0 {
		err := CompleteRequestValidationError{
			field:  "OrgId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ProcessInstanceId

	// no validation rules for Comment

	// no validation rules for Candidate

	if len(errors) > 0 {
		return CompleteRequestMultiError(errors)
	}

	return nil
}

// CompleteRequestMultiError is an error wrapping multiple validation errors
// returned by CompleteRequest.ValidateAll() if the designated constraints
// aren't met.
type CompleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteRequestMultiError) AllErrors() []error { return m }

// CompleteRequestValidationError is the validation error returned by
// CompleteRequest.Validate if the designated constraints aren't met.
type CompleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteRequestValidationError) ErrorName() string { return "CompleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e CompleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteRequestValidationError{}

// Validate checks the field values on CompleteReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CompleteReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CompleteReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CompleteReplyMultiError, or
// nil if none found.
func (m *CompleteReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CompleteReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CompleteReplyMultiError(errors)
	}

	return nil
}

// CompleteReplyMultiError is an error wrapping multiple validation errors
// returned by CompleteReply.ValidateAll() if the designated constraints
// aren't met.
type CompleteReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteReplyMultiError) AllErrors() []error { return m }

// CompleteReplyValidationError is the validation error returned by
// CompleteReply.Validate if the designated constraints aren't met.
type CompleteReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteReplyValidationError) ErrorName() string { return "CompleteReplyValidationError" }

// Error satisfies the builtin error interface
func (e CompleteReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompleteReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteReplyValidationError{}

// Validate checks the field values on WithdrawRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WithdrawRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WithdrawRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WithdrawRequestMultiError, or nil if none found.
func (m *WithdrawRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WithdrawRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TaskId

	if m.GetUserId() <= 0 {
		err := WithdrawRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for UserName

	if m.GetOrgId() <= 0 {
		err := WithdrawRequestValidationError{
			field:  "OrgId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ProcessInstanceId

	// no validation rules for Comment

	// no validation rules for Candidate

	if len(errors) > 0 {
		return WithdrawRequestMultiError(errors)
	}

	return nil
}

// WithdrawRequestMultiError is an error wrapping multiple validation errors
// returned by WithdrawRequest.ValidateAll() if the designated constraints
// aren't met.
type WithdrawRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WithdrawRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WithdrawRequestMultiError) AllErrors() []error { return m }

// WithdrawRequestValidationError is the validation error returned by
// WithdrawRequest.Validate if the designated constraints aren't met.
type WithdrawRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WithdrawRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WithdrawRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WithdrawRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WithdrawRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WithdrawRequestValidationError) ErrorName() string { return "WithdrawRequestValidationError" }

// Error satisfies the builtin error interface
func (e WithdrawRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWithdrawRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WithdrawRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WithdrawRequestValidationError{}

// Validate checks the field values on WithdrawReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WithdrawReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WithdrawReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WithdrawReplyMultiError, or
// nil if none found.
func (m *WithdrawReply) ValidateAll() error {
	return m.validate(true)
}

func (m *WithdrawReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WithdrawReplyMultiError(errors)
	}

	return nil
}

// WithdrawReplyMultiError is an error wrapping multiple validation errors
// returned by WithdrawReply.ValidateAll() if the designated constraints
// aren't met.
type WithdrawReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WithdrawReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WithdrawReplyMultiError) AllErrors() []error { return m }

// WithdrawReplyValidationError is the validation error returned by
// WithdrawReply.Validate if the designated constraints aren't met.
type WithdrawReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WithdrawReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WithdrawReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WithdrawReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WithdrawReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WithdrawReplyValidationError) ErrorName() string { return "WithdrawReplyValidationError" }

// Error satisfies the builtin error interface
func (e WithdrawReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWithdrawReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WithdrawReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WithdrawReplyValidationError{}
