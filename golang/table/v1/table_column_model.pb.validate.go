// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: table_column_model.proto

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

// Validate checks the field values on ColumnPropRelating with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ColumnPropRelating) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ColumnPropRelating with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ColumnPropRelatingMultiError, or nil if none found.
func (m *ColumnPropRelating) ValidateAll() error {
	return m.validate(true)
}

func (m *ColumnPropRelating) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRelateTables() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ColumnPropRelatingValidationError{
						field:  fmt.Sprintf("RelateTables[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ColumnPropRelatingValidationError{
						field:  fmt.Sprintf("RelateTables[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ColumnPropRelatingValidationError{
					field:  fmt.Sprintf("RelateTables[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ColumnPropRelatingMultiError(errors)
	}

	return nil
}

// ColumnPropRelatingMultiError is an error wrapping multiple validation errors
// returned by ColumnPropRelating.ValidateAll() if the designated constraints
// aren't met.
type ColumnPropRelatingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ColumnPropRelatingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ColumnPropRelatingMultiError) AllErrors() []error { return m }

// ColumnPropRelatingValidationError is the validation error returned by
// ColumnPropRelating.Validate if the designated constraints aren't met.
type ColumnPropRelatingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ColumnPropRelatingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ColumnPropRelatingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ColumnPropRelatingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ColumnPropRelatingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ColumnPropRelatingValidationError) ErrorName() string {
	return "ColumnPropRelatingValidationError"
}

// Error satisfies the builtin error interface
func (e ColumnPropRelatingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sColumnPropRelating.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ColumnPropRelatingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ColumnPropRelatingValidationError{}
