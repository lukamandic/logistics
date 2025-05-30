package errors

import (
	"encoding/json"
	"net/http"

	"github.com/lukamandic/logistics/backend/internal/api/validation"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func WriteError(w http.ResponseWriter, err error) {
	var response ErrorResponse

	switch e := err.(type) {
	case validation.ValidationErrors:
		response = ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  e,
		}
	case validation.ValidationError:
		response = ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  []validation.ValidationError{e},
		}
	default:
		switch err.Error() {
		case "no parcel sizes found in the table":
			response = ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "No parcel sizes configured.",
			}
		case "multiple items found in the table, expected only one":
			response = ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Database inconsistency detected. Please contact support.",
			}
		default:
			response = ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "An unexpected error occurred",
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
} 