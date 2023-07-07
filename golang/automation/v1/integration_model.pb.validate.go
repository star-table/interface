// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integration_model.proto

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

// Validate checks the field values on IntegrationConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IntegrationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationConfigMultiError, or nil if none found.
func (m *IntegrationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Config.(type) {

	case *IntegrationConfig_Smtp:

		if all {
			switch v := interface{}(m.GetSmtp()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Smtp",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Smtp",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSmtp()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IntegrationConfigValidationError{
					field:  "Smtp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *IntegrationConfig_Mysql:

		if all {
			switch v := interface{}(m.GetMysql()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Mysql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Mysql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMysql()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IntegrationConfigValidationError{
					field:  "Mysql",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *IntegrationConfig_Postgres:

		if all {
			switch v := interface{}(m.GetPostgres()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Postgres",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IntegrationConfigValidationError{
						field:  "Postgres",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPostgres()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IntegrationConfigValidationError{
					field:  "Postgres",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return IntegrationConfigMultiError(errors)
	}

	return nil
}

// IntegrationConfigMultiError is an error wrapping multiple validation errors
// returned by IntegrationConfig.ValidateAll() if the designated constraints
// aren't met.
type IntegrationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationConfigMultiError) AllErrors() []error { return m }

// IntegrationConfigValidationError is the validation error returned by
// IntegrationConfig.Validate if the designated constraints aren't met.
type IntegrationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationConfigValidationError) ErrorName() string {
	return "IntegrationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationConfigValidationError{}

// Validate checks the field values on IntegrationInput with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IntegrationInput) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationInput with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationInputMultiError, or nil if none found.
func (m *IntegrationInput) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationInput) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IntegrationInputValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IntegrationInputValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IntegrationInputValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 0 || l > 128 {
		err := IntegrationInputValidationError{
			field:  "Name",
			reason: "value length must be between 0 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDesc()); l < 0 || l > 1024 {
		err := IntegrationInputValidationError{
			field:  "Desc",
			reason: "value length must be between 0 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IntegrationInputMultiError(errors)
	}

	return nil
}

// IntegrationInputMultiError is an error wrapping multiple validation errors
// returned by IntegrationInput.ValidateAll() if the designated constraints
// aren't met.
type IntegrationInputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationInputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationInputMultiError) AllErrors() []error { return m }

// IntegrationInputValidationError is the validation error returned by
// IntegrationInput.Validate if the designated constraints aren't met.
type IntegrationInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationInputValidationError) ErrorName() string { return "IntegrationInputValidationError" }

// Error satisfies the builtin error interface
func (e IntegrationInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationInputValidationError{}

// Validate checks the field values on Integration with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Integration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Integration with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IntegrationMultiError, or
// nil if none found.
func (m *Integration) ValidateAll() error {
	return m.validate(true)
}

func (m *Integration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.GetOrgId() <= 0 {
		err := IntegrationValidationError{
			field:  "OrgId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IntegrationValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IntegrationValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IntegrationValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 0 || l > 128 {
		err := IntegrationValidationError{
			field:  "Name",
			reason: "value length must be between 0 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDesc()); l < 0 || l > 1024 {
		err := IntegrationValidationError{
			field:  "Desc",
			reason: "value length must be between 0 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCreator() < 0 {
		err := IntegrationValidationError{
			field:  "Creator",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdater() < 0 {
		err := IntegrationValidationError{
			field:  "Updater",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCreatedAt() < 0 {
		err := IntegrationValidationError{
			field:  "CreatedAt",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() < 0 {
		err := IntegrationValidationError{
			field:  "UpdatedAt",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CreatorName

	// no validation rules for UpdaterName

	if len(errors) > 0 {
		return IntegrationMultiError(errors)
	}

	return nil
}

// IntegrationMultiError is an error wrapping multiple validation errors
// returned by Integration.ValidateAll() if the designated constraints aren't met.
type IntegrationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationMultiError) AllErrors() []error { return m }

// IntegrationValidationError is the validation error returned by
// Integration.Validate if the designated constraints aren't met.
type IntegrationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationValidationError) ErrorName() string { return "IntegrationValidationError" }

// Error satisfies the builtin error interface
func (e IntegrationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationValidationError{}

// Validate checks the field values on IntegrationMeta with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IntegrationMeta) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationMeta with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationMetaMultiError, or nil if none found.
func (m *IntegrationMeta) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationMeta) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	if l := utf8.RuneCountInString(m.GetName()); l < 0 || l > 128 {
		err := IntegrationMetaValidationError{
			field:  "Name",
			reason: "value length must be between 0 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDesc()); l < 0 || l > 1024 {
		err := IntegrationMetaValidationError{
			field:  "Desc",
			reason: "value length must be between 0 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IntegrationMetaMultiError(errors)
	}

	return nil
}

// IntegrationMetaMultiError is an error wrapping multiple validation errors
// returned by IntegrationMeta.ValidateAll() if the designated constraints
// aren't met.
type IntegrationMetaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationMetaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationMetaMultiError) AllErrors() []error { return m }

// IntegrationMetaValidationError is the validation error returned by
// IntegrationMeta.Validate if the designated constraints aren't met.
type IntegrationMetaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationMetaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationMetaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationMetaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationMetaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationMetaValidationError) ErrorName() string { return "IntegrationMetaValidationError" }

// Error satisfies the builtin error interface
func (e IntegrationMetaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationMeta.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationMetaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationMetaValidationError{}

// Validate checks the field values on IntegrationConfigSmtp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationConfigSmtp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationConfigSmtp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationConfigSmtpMultiError, or nil if none found.
func (m *IntegrationConfigSmtp) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationConfigSmtp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUser()); l < 1 || l > 128 {
		err := IntegrationConfigSmtpValidationError{
			field:  "User",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 1 || l > 128 {
		err := IntegrationConfigSmtpValidationError{
			field:  "Password",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetHost()); l < 1 || l > 128 {
		err := IntegrationConfigSmtpValidationError{
			field:  "Host",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() <= 0 {
		err := IntegrationConfigSmtpValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Ssl

	if len(errors) > 0 {
		return IntegrationConfigSmtpMultiError(errors)
	}

	return nil
}

// IntegrationConfigSmtpMultiError is an error wrapping multiple validation
// errors returned by IntegrationConfigSmtp.ValidateAll() if the designated
// constraints aren't met.
type IntegrationConfigSmtpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationConfigSmtpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationConfigSmtpMultiError) AllErrors() []error { return m }

// IntegrationConfigSmtpValidationError is the validation error returned by
// IntegrationConfigSmtp.Validate if the designated constraints aren't met.
type IntegrationConfigSmtpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationConfigSmtpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationConfigSmtpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationConfigSmtpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationConfigSmtpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationConfigSmtpValidationError) ErrorName() string {
	return "IntegrationConfigSmtpValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationConfigSmtpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationConfigSmtp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationConfigSmtpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationConfigSmtpValidationError{}

// Validate checks the field values on IntegrationConfigMysql with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationConfigMysql) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationConfigMysql with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationConfigMysqlMultiError, or nil if none found.
func (m *IntegrationConfigMysql) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationConfigMysql) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetHost()); l < 1 || l > 128 {
		err := IntegrationConfigMysqlValidationError{
			field:  "Host",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDatabase()); l < 1 || l > 128 {
		err := IntegrationConfigMysqlValidationError{
			field:  "Database",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUser()); l < 1 || l > 128 {
		err := IntegrationConfigMysqlValidationError{
			field:  "User",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 1 || l > 128 {
		err := IntegrationConfigMysqlValidationError{
			field:  "Password",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() <= 0 {
		err := IntegrationConfigMysqlValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IntegrationConfigMysqlMultiError(errors)
	}

	return nil
}

// IntegrationConfigMysqlMultiError is an error wrapping multiple validation
// errors returned by IntegrationConfigMysql.ValidateAll() if the designated
// constraints aren't met.
type IntegrationConfigMysqlMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationConfigMysqlMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationConfigMysqlMultiError) AllErrors() []error { return m }

// IntegrationConfigMysqlValidationError is the validation error returned by
// IntegrationConfigMysql.Validate if the designated constraints aren't met.
type IntegrationConfigMysqlValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationConfigMysqlValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationConfigMysqlValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationConfigMysqlValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationConfigMysqlValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationConfigMysqlValidationError) ErrorName() string {
	return "IntegrationConfigMysqlValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationConfigMysqlValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationConfigMysql.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationConfigMysqlValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationConfigMysqlValidationError{}

// Validate checks the field values on IntegrationConfigPostgres with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IntegrationConfigPostgres) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegrationConfigPostgres with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IntegrationConfigPostgresMultiError, or nil if none found.
func (m *IntegrationConfigPostgres) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegrationConfigPostgres) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetHost()); l < 1 || l > 128 {
		err := IntegrationConfigPostgresValidationError{
			field:  "Host",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDatabase()); l < 1 || l > 128 {
		err := IntegrationConfigPostgresValidationError{
			field:  "Database",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUser()); l < 1 || l > 128 {
		err := IntegrationConfigPostgresValidationError{
			field:  "User",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 1 || l > 128 {
		err := IntegrationConfigPostgresValidationError{
			field:  "Password",
			reason: "value length must be between 1 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() <= 0 {
		err := IntegrationConfigPostgresValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IntegrationConfigPostgresMultiError(errors)
	}

	return nil
}

// IntegrationConfigPostgresMultiError is an error wrapping multiple validation
// errors returned by IntegrationConfigPostgres.ValidateAll() if the
// designated constraints aren't met.
type IntegrationConfigPostgresMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegrationConfigPostgresMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegrationConfigPostgresMultiError) AllErrors() []error { return m }

// IntegrationConfigPostgresValidationError is the validation error returned by
// IntegrationConfigPostgres.Validate if the designated constraints aren't met.
type IntegrationConfigPostgresValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegrationConfigPostgresValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegrationConfigPostgresValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegrationConfigPostgresValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegrationConfigPostgresValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegrationConfigPostgresValidationError) ErrorName() string {
	return "IntegrationConfigPostgresValidationError"
}

// Error satisfies the builtin error interface
func (e IntegrationConfigPostgresValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegrationConfigPostgres.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegrationConfigPostgresValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegrationConfigPostgresValidationError{}
