package validate

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ValidationErrors represents a collection of validation error messages
type ValidationErrors struct {
	Messages []string `json:"error_messages"`
}

// NewValidationErrors creates a new instance of ValidationErrors with the provided error messages.
func NewValidationErrors(messages ...string) *ValidationErrors {
	return &ValidationErrors{
		Messages: messages,
	}
}

func (v *ValidationErrors) Error() string {
	data, err := json.Marshal(v.Messages)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

// IsValidationError checks if the given error is of type ValidationErrors.
func IsValidationError(err error) bool {
	var ve *ValidationErrors
	return errors.As(err, &ve)
}
