package sys

import (
	"github.com/bifidokk/awesome-chat/auth/internal/sys/codes"
	"github.com/pkg/errors"
)

type commonError struct {
	msg  string
	code codes.Code
}

// NewCommonError creates a new commonError with the specified message and code.
func NewCommonError(msg string, code codes.Code) *commonError {
	return &commonError{msg, code}
}

// Error returns the error message of commonError.
func (r *commonError) Error() string {
	return r.msg
}

// Code returns the error code of commonError.
func (r *commonError) Code() codes.Code {
	return r.code
}

// IsCommonError checks if the given error is of type commonError.
func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

// GetCommonError retrieves the commonError from the given error if it exists.
func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
