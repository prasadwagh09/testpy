// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/listener.proto

package envoy_api_v3alpha_listener

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _listener_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Filter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return FilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *Filter_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Filter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on FilterChainMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FilterChainMatch) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetDestinationPort(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 65535 {
			return FilterChainMatchValidationError{
				field:  "DestinationPort",
				reason: "value must be inside range [1, 65535]",
			}
		}

	}

	for idx, item := range m.GetPrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					field:  fmt.Sprintf("PrefixRanges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AddressSuffix

	if v, ok := interface{}(m.GetSuffixLen()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainMatchValidationError{
				field:  "SuffixLen",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := FilterChainMatch_ConnectionSourceType_name[int32(m.GetSourceType())]; !ok {
		return FilterChainMatchValidationError{
			field:  "SourceType",
			reason: "value must be one of the defined enum values",
		}
	}

	for idx, item := range m.GetSourcePrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					field:  fmt.Sprintf("SourcePrefixRanges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSourcePorts() {
		_, _ = idx, item

		if val := item; val < 1 || val > 65535 {
			return FilterChainMatchValidationError{
				field:  fmt.Sprintf("SourcePorts[%v]", idx),
				reason: "value must be inside range [1, 65535]",
			}
		}

	}

	// no validation rules for TransportProtocol

	return nil
}

// FilterChainMatchValidationError is the validation error returned by
// FilterChainMatch.Validate if the designated constraints aren't met.
type FilterChainMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterChainMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterChainMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterChainMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterChainMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterChainMatchValidationError) ErrorName() string { return "FilterChainMatchValidationError" }

// Error satisfies the builtin error interface
func (e FilterChainMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChainMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterChainMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterChainMatchValidationError{}

// Validate checks the field values on FilterChain with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilterChain) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilterChainMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "FilterChainMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTlsContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "TlsContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseProxyProto()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "UseProxyProto",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTransportSocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "TransportSocket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	return nil
}

// FilterChainValidationError is the validation error returned by
// FilterChain.Validate if the designated constraints aren't met.
type FilterChainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterChainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterChainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterChainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterChainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterChainValidationError) ErrorName() string { return "FilterChainValidationError" }

// Error satisfies the builtin error interface
func (e FilterChainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterChainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterChainValidationError{}

// Validate checks the field values on ListenerFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListenerFilter) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return ListenerFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *ListenerFilter_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListenerFilter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListenerFilterValidationError is the validation error returned by
// ListenerFilter.Validate if the designated constraints aren't met.
type ListenerFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerFilterValidationError) ErrorName() string { return "ListenerFilterValidationError" }

// Error satisfies the builtin error interface
func (e ListenerFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerFilterValidationError{}
