package errsx

import (
	"encoding/json"
	"errors"
	"fmt"

	"connectrpc.com/connect"
)

type VerdantErrDetails struct {
	errors ErrorMap
}

type VerdantErrBuilder struct {
	code    connect.Code
	msg     string
	cause   error
	label   string
	details VerdantErrDetails
}

// NewVerdantErrDetails is a constructor for VerdantErrDetails
func NewVerdantErrDetails(errors ErrorMap) VerdantErrDetails {
	return VerdantErrDetails{errors: errors}
}

// NewVerdantErrBuilder is a constructor for VerdantErrBuilder
func NewVerdantErrBuilder() *VerdantErrBuilder {
	return &VerdantErrBuilder{}
}

// MarshalJSON implements the json.Marshaler interface.
func (builder *VerdantErrBuilder) MarshalJSON() ([]byte, error) {
	// use json.Marshal to convert the error message to a JSON byte slice
	byteBuffer, err := json.Marshal(map[string]interface{}{
		"code":    builder.code,
		"message": builder.msg,
		"cause":   builder.cause.Error(),
		"label":   builder.label,
		"details": builder.details,
	})
	if err != nil {
		return nil, err
	}

	return byteBuffer, nil
}

// Error is a method to return an error, this is an implementation of the error interface.
func (builder *VerdantErrBuilder) Error() string {

	// validate the error instance, if it is nil, return nil
	if builder.code == 0 || builder.msg == "" {
		builder.code = connect.CodeInternal
		builder.msg = "Internal Server Error"
	}

	// if the cause is nil, set the cause to the error message
	if builder.cause == nil {
		builder.cause = errors.New(builder.msg)
	}

	// if the label is empty, set the label to the String of the error code
	if builder.label == "" {
		builder.label = builder.code.String()
	}

	// convert the builder instance to a formatted error message and return it
	return fmt.Sprintf("code: %d, label: %s, message: %s, cause: %s, details: %v",
		builder.code, builder.label, builder.msg, builder.cause.Error(), builder.details)
}

// WithCode is a method to set the error code.
func (builder *VerdantErrBuilder) WithCode(code connect.Code) *VerdantErrBuilder {
	builder.code = code
	return builder
}

// WithMsg is a method to set the error message.
func (builder *VerdantErrBuilder) WithMsg(msg string) *VerdantErrBuilder {
	builder.msg = msg
	return builder
}

// WithCause is a method to set the error cause.
func (builder *VerdantErrBuilder) WithCause(cause error) *VerdantErrBuilder {
	builder.cause = cause
	return builder
}

// WithDetails is a method to set the error details.
func (builder *VerdantErrBuilder) WithDetails(details VerdantErrDetails) *VerdantErrBuilder {
	builder.details = details
	return builder
}

// ErrDetails is a method to return the error details as a map.
func (err *VerdantErrDetails) ErrDetails() (ErrorMap, error) {
	if err.errors == nil {
		return nil, errors.New("no error details found")
	}
	return err.errors, nil
}
