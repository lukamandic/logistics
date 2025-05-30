package validation

import (
	"fmt"
	"strings"

	"github.com/lukamandic/logistics/backend/internal/service"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var messages []string
	for _, err := range e {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}

func ValidateParcelSizes(sizes service.ParcelSize) error {
	var errors ValidationErrors

	if len(sizes) == 0 {
		errors = append(errors, ValidationError{
			Field:   "parcel_sizes",
			Message: "array cannot be empty",
		})
		return errors
	}

	hasInvalidSize := false
	for _, size := range sizes {
		if size <= 0 {
			hasInvalidSize = true
			break
		}
	}

	if hasInvalidSize {
		errors = append(errors, ValidationError{
			Field:   "parcel_sizes",
			Message: "all values must be greater than zero",
		})
	}

	if len(errors) > 0 {
		return errors
	}
	return nil
}

func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return ValidationError{
			Field:   "amount",
			Message: "value must be greater than zero",
		}
	}
	return nil
}

func ValidateID(id string) error {
	if id == "" {
		return ValidationError{
			Field:   "id",
			Message: "is required",
		}
	}
	return nil
} 