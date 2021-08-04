// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/account/proto/profile/profile_service.proto

package profile

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	entities "github.com/ygpark2/njro/service/account/proto/entities"
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

	_ = entities.Profile_GenderType(0)

	_ = entities.Profile_GenderType(0)
)

// define the regex for a UUID once up-front
var _profile_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetLimit(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 100 {
			err := ListRequestValidationError{
				field:  "Limit",
				reason: "value must be inside range [1, 100]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetPage(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := ListRequestValidationError{
				field:  "Page",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetSort()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Sort",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListRequestValidationError{
					field:  "Sort",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListRequestValidationError{
				field:  "Sort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetPreferredTheme(); wrapper != nil {

		if _, ok := _ListRequest_PreferredTheme_InLookup[wrapper.GetValue()]; !ok {
			err := ListRequestValidationError{
				field:  "PreferredTheme",
				reason: "value must be in list [dark light cosmic corporate]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if _, ok := _ListRequest_Gender_InLookup[m.GetGender()]; !ok {
		err := ListRequestValidationError{
			field:  "Gender",
			reason: "value must be in list [0 1 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}
	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

var _ListRequest_PreferredTheme_InLookup = map[string]struct{}{
	"dark":      {},
	"light":     {},
	"cosmic":    {},
	"corporate": {},
}

var _ListRequest_Gender_InLookup = map[entities.Profile_GenderType]struct{}{
	0: {},
	1: {},
	2: {},
}

// Validate is disabled for ListResponse. This method will always return nil.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll is disabled for ListResponse. This method will always return nil.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Id.(type) {

	case *GetRequest_ProfileId:

		if wrapper := m.GetProfileId(); wrapper != nil {

			if err := m._validateUuid(wrapper.GetValue()); err != nil {
				err = GetRequestValidationError{
					field:  "ProfileId",
					reason: "value must be a valid UUID",
					cause:  err,
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}

	case *GetRequest_UserId:

		if wrapper := m.GetUserId(); wrapper != nil {

			if err := m._validateUuid(wrapper.GetValue()); err != nil {
				err = GetRequestValidationError{
					field:  "UserId",
					reason: "value must be a valid UUID",
					cause:  err,
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}

	}

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}
	return nil
}

func (m *GetRequest) _validateUuid(uuid string) error {
	if matched := _profile_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate is disabled for GetResponse. This method will always return nil.
func (m *GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll is disabled for GetResponse. This method will always return nil.
func (m *GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResponse) validate(all bool) error {
	return nil
}

// GetResponseMultiError is an error wrapping multiple validation errors
// returned by GetResponse.ValidateAll() if the designated constraints aren't met.
type GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResponseMultiError) AllErrors() []error { return m }

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateRequestMultiError, or
// nil if none found.
func (m *CreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetUserId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			err = CreateRequestValidationError{
				field:  "UserId",
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetTz()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Tz",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Tz",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTz()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "Tz",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetAvatar(); wrapper != nil {

		if _, err := url.Parse(wrapper.GetValue()); err != nil {
			err = CreateRequestValidationError{
				field:  "Avatar",
				reason: "value must be a valid URI",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if _, ok := _CreateRequest_Gender_InLookup[m.GetGender()]; !ok {
		err := CreateRequestValidationError{
			field:  "Gender",
			reason: "value must be in list [0 1 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if t := m.GetBirthday(); t != nil {
		ts, err := t.AsTime(), t.CheckValid()
		if err != nil {
			err = CreateRequestValidationError{
				field:  "Birthday",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			now := time.Now()

			if ts.Sub(now) >= 0 {
				err := CreateRequestValidationError{
					field:  "Birthday",
					reason: "value must be less than now",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if wrapper := m.GetPreferredTheme(); wrapper != nil {

		if _, ok := _CreateRequest_PreferredTheme_InLookup[wrapper.GetValue()]; !ok {
			err := CreateRequestValidationError{
				field:  "PreferredTheme",
				reason: "value must be in list [dark light cosmic corporate]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CreateRequestMultiError(errors)
	}
	return nil
}

func (m *CreateRequest) _validateUuid(uuid string) error {
	if matched := _profile_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// CreateRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRequestMultiError) AllErrors() []error { return m }

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

var _CreateRequest_Gender_InLookup = map[entities.Profile_GenderType]struct{}{
	0: {},
	1: {},
	2: {},
}

var _CreateRequest_PreferredTheme_InLookup = map[string]struct{}{
	"dark":      {},
	"light":     {},
	"cosmic":    {},
	"corporate": {},
}

// Validate is disabled for CreateResponse. This method will always return nil.
func (m *CreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll is disabled for CreateResponse. This method will always return nil.
func (m *CreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResponse) validate(all bool) error {
	return nil
}

// CreateResponseMultiError is an error wrapping multiple validation errors
// returned by CreateResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResponseMultiError) AllErrors() []error { return m }

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}
