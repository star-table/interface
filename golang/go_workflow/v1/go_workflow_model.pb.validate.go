// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: go_workflow_model.proto

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

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Node) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Node with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NodeMultiError, or nil if none found.
func (m *Node) ValidateAll() error {
	return m.validate(true)
}

func (m *Node) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := NodeValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Node_Type_InLookup[m.GetType()]; !ok {
		err := NodeValidationError{
			field:  "Type",
			reason: "value must be in list [start route condition approver notifier]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetNodeId()) < 1 {
		err := NodeValidationError{
			field:  "NodeId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PrevId

	if all {
		switch v := interface{}(m.GetChildNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "ChildNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "ChildNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetChildNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "ChildNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetConditionNodes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  fmt.Sprintf("ConditionNodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NodeValidationError{
						field:  fmt.Sprintf("ConditionNodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("ConditionNodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodeValidationError{
					field:  "Properties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "Properties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NodeMultiError(errors)
	}

	return nil
}

// NodeMultiError is an error wrapping multiple validation errors returned by
// Node.ValidateAll() if the designated constraints aren't met.
type NodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMultiError) AllErrors() []error { return m }

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

var _Node_Type_InLookup = map[string]struct{}{
	"start":     {},
	"route":     {},
	"condition": {},
	"approver":  {},
	"notifier":  {},
}

// Validate checks the field values on NodeProperties with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeProperties) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeProperties with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NodePropertiesMultiError,
// or nil if none found.
func (m *NodeProperties) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeProperties) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ActivateType

	// no validation rules for AgreeAll

	if all {
		switch v := interface{}(m.GetCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodePropertiesValidationError{
					field:  "Condition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodePropertiesValidationError{
					field:  "Condition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodePropertiesValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetActionerRule()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NodePropertiesValidationError{
					field:  "ActionerRule",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NodePropertiesValidationError{
					field:  "ActionerRule",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActionerRule()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodePropertiesValidationError{
				field:  "ActionerRule",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NodePropertiesMultiError(errors)
	}

	return nil
}

// NodePropertiesMultiError is an error wrapping multiple validation errors
// returned by NodeProperties.ValidateAll() if the designated constraints
// aren't met.
type NodePropertiesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodePropertiesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodePropertiesMultiError) AllErrors() []error { return m }

// NodePropertiesValidationError is the validation error returned by
// NodeProperties.Validate if the designated constraints aren't met.
type NodePropertiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodePropertiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodePropertiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodePropertiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodePropertiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodePropertiesValidationError) ErrorName() string { return "NodePropertiesValidationError" }

// Error satisfies the builtin error interface
func (e NodePropertiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeProperties.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodePropertiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodePropertiesValidationError{}

// Validate checks the field values on NodeConditions with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeConditions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeConditions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NodeConditionsMultiError,
// or nil if none found.
func (m *NodeConditions) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeConditions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, NodeConditionsValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NodeConditionsValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeConditionsValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return NodeConditionsMultiError(errors)
	}

	return nil
}

// NodeConditionsMultiError is an error wrapping multiple validation errors
// returned by NodeConditions.ValidateAll() if the designated constraints
// aren't met.
type NodeConditionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeConditionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeConditionsMultiError) AllErrors() []error { return m }

// NodeConditionsValidationError is the validation error returned by
// NodeConditions.Validate if the designated constraints aren't met.
type NodeConditionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeConditionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeConditionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeConditionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeConditionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeConditionsValidationError) ErrorName() string { return "NodeConditionsValidationError" }

// Error satisfies the builtin error interface
func (e NodeConditionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeConditions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeConditionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeConditionsValidationError{}

// Validate checks the field values on NodeCondition with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NodeConditionMultiError, or
// nil if none found.
func (m *NodeCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for ParamKey

	// no validation rules for ParamLabel

	// no validation rules for LowerBound

	// no validation rules for LowerBoundEqual

	// no validation rules for UpperBoundEqual

	// no validation rules for UpperBound

	// no validation rules for BoundEqual

	// no validation rules for Unit

	if len(errors) > 0 {
		return NodeConditionMultiError(errors)
	}

	return nil
}

// NodeConditionMultiError is an error wrapping multiple validation errors
// returned by NodeCondition.ValidateAll() if the designated constraints
// aren't met.
type NodeConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeConditionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeConditionMultiError) AllErrors() []error { return m }

// NodeConditionValidationError is the validation error returned by
// NodeCondition.Validate if the designated constraints aren't met.
type NodeConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeConditionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeConditionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeConditionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeConditionValidationError) ErrorName() string { return "NodeConditionValidationError" }

// Error satisfies the builtin error interface
func (e NodeConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeConditionValidationError{}

// Validate checks the field values on ActionerRule with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ActionerRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActionerRule with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActionerRuleMultiError, or
// nil if none found.
func (m *ActionerRule) ValidateAll() error {
	return m.validate(true)
}

func (m *ActionerRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for MemberCount

	// no validation rules for ActionType

	for idx, item := range m.GetParticipants() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActionerRuleValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActionerRuleValidationError{
						field:  fmt.Sprintf("Participants[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActionerRuleValidationError{
					field:  fmt.Sprintf("Participants[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ActionerRuleMultiError(errors)
	}

	return nil
}

// ActionerRuleMultiError is an error wrapping multiple validation errors
// returned by ActionerRule.ValidateAll() if the designated constraints aren't met.
type ActionerRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActionerRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActionerRuleMultiError) AllErrors() []error { return m }

// ActionerRuleValidationError is the validation error returned by
// ActionerRule.Validate if the designated constraints aren't met.
type ActionerRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionerRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionerRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionerRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionerRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionerRuleValidationError) ErrorName() string { return "ActionerRuleValidationError" }

// Error satisfies the builtin error interface
func (e ActionerRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActionerRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionerRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionerRuleValidationError{}

// Validate checks the field values on Participant with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Participant) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Participant with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ParticipantMultiError, or
// nil if none found.
func (m *Participant) ValidateAll() error {
	return m.validate(true)
}

func (m *Participant) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return ParticipantMultiError(errors)
	}

	return nil
}

// ParticipantMultiError is an error wrapping multiple validation errors
// returned by Participant.ValidateAll() if the designated constraints aren't met.
type ParticipantMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParticipantMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParticipantMultiError) AllErrors() []error { return m }

// ParticipantValidationError is the validation error returned by
// Participant.Validate if the designated constraints aren't met.
type ParticipantValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParticipantValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParticipantValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParticipantValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParticipantValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParticipantValidationError) ErrorName() string { return "ParticipantValidationError" }

// Error satisfies the builtin error interface
func (e ParticipantValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParticipant.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParticipantValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParticipantValidationError{}

// Validate checks the field values on ProcessDefine with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProcessDefine) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessDefine with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProcessDefineMultiError, or
// nil if none found.
func (m *ProcessDefine) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessDefine) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Version

	// no validation rules for Resource

	// no validation rules for UserId

	// no validation rules for UserName

	// no validation rules for CreateAt

	// no validation rules for UpdateAt

	if len(errors) > 0 {
		return ProcessDefineMultiError(errors)
	}

	return nil
}

// ProcessDefineMultiError is an error wrapping multiple validation errors
// returned by ProcessDefine.ValidateAll() if the designated constraints
// aren't met.
type ProcessDefineMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessDefineMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessDefineMultiError) AllErrors() []error { return m }

// ProcessDefineValidationError is the validation error returned by
// ProcessDefine.Validate if the designated constraints aren't met.
type ProcessDefineValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessDefineValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessDefineValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessDefineValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessDefineValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessDefineValidationError) ErrorName() string { return "ProcessDefineValidationError" }

// Error satisfies the builtin error interface
func (e ProcessDefineValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessDefine.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessDefineValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessDefineValidationError{}

// Validate checks the field values on ProcessInstance with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProcessInstance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessInstance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProcessInstanceMultiError, or nil if none found.
func (m *ProcessInstance) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessInstance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ProcessDefineId

	// no validation rules for OrgId

	// no validation rules for NodeId

	// no validation rules for Candidate

	// no validation rules for TaskId

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for StartUserId

	// no validation rules for StartUserName

	// no validation rules for IsFinished

	if len(errors) > 0 {
		return ProcessInstanceMultiError(errors)
	}

	return nil
}

// ProcessInstanceMultiError is an error wrapping multiple validation errors
// returned by ProcessInstance.ValidateAll() if the designated constraints
// aren't met.
type ProcessInstanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessInstanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessInstanceMultiError) AllErrors() []error { return m }

// ProcessInstanceValidationError is the validation error returned by
// ProcessInstance.Validate if the designated constraints aren't met.
type ProcessInstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessInstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessInstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessInstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessInstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessInstanceValidationError) ErrorName() string { return "ProcessInstanceValidationError" }

// Error satisfies the builtin error interface
func (e ProcessInstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessInstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessInstanceValidationError{}