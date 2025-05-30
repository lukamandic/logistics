package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lukamandic/logistics/backend/internal/api/errors"
	"github.com/lukamandic/logistics/backend/internal/api/validation"
	"github.com/lukamandic/logistics/backend/internal/service"
)

type ParcelHandler struct {
	parcelService *service.ParcelService
}

func NewParcelHandler(parcelService *service.ParcelService) *ParcelHandler {
	return &ParcelHandler{
		parcelService: parcelService,
	}
}

func (h *ParcelHandler) HandleParcelRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetParcelSizes(w, r)
	case http.MethodPost, http.MethodPut:
		h.UpdateParcelSizes(w, r)
	default:
		errors.WriteError(w, fmt.Errorf("method %s not allowed", r.Method))
	}
}

func (h *ParcelHandler) GetParcelSizes(w http.ResponseWriter, r *http.Request) {
	items, err := h.parcelService.GetAllItems(r.Context())
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	if len(items) == 0 {
		errors.WriteError(w, fmt.Errorf("no parcel sizes found in the table"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		errors.WriteError(w, fmt.Errorf("error encoding response: %w", err))
		return
	}
}

type UpdateParcelSizesRequest struct {
	ID          string        `json:"id"`
	ParcelSizes service.ParcelSize `json:"parcel_sizes"`
}

func (h *ParcelHandler) UpdateParcelSizes(w http.ResponseWriter, r *http.Request) {
	var req UpdateParcelSizesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteError(w, fmt.Errorf("invalid request body: %w", err))
		return
	}

	if err := validation.ValidateID(req.ID); err != nil {
		errors.WriteError(w, err)
		return
	}

	if err := validation.ValidateParcelSizes(req.ParcelSizes); err != nil {
		errors.WriteError(w, err)
		return
	}

	if err := h.parcelService.UpdateParcelSizes(r.Context(), req.ID, req.ParcelSizes); err != nil {
		errors.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Parcel sizes updated successfully",
		"id":      req.ID,
	})
}

func (h *ParcelHandler) HandleCalculateDistribution(w http.ResponseWriter, r *http.Request) {
	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		errors.WriteError(w, validation.ValidationError{
			Field:   "amount",
			Message: "is required",
		})
		return
	}

	// Parse amount to float64
	var amount float64
	if _, err := fmt.Sscanf(amountStr, "%f", &amount); err != nil {
		errors.WriteError(w, validation.ValidationError{
			Field:   "amount",
			Message: "must be a valid number",
		})
		return
	}

	if err := validation.ValidateAmount(amount); err != nil {
		errors.WriteError(w, err)
		return
	}

	distribution, err := h.parcelService.CalculateDistribution(r.Context(), amount)
	if err != nil {
		errors.WriteError(w, err)
		return
	}

	stringDistribution := make(map[string]int)
	for size, count := range distribution {
		stringDistribution[fmt.Sprintf("%g", size)] = count
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stringDistribution); err != nil {
		fmt.Println("Error calculating distribution:", err)
		errors.WriteError(w, fmt.Errorf("error encoding response: %w", err))
		return
	}
}