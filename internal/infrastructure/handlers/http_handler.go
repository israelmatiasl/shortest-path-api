package handlers

import (
	"encoding/json"
	"net/http"

	"shortest-path-api/internal/application"
	"shortest-path-api/internal/domain"
)

type HTTPHandler struct {
	service *application.ShortestPathService
}

func NewHTTPHandler(service *application.ShortestPathService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) ShortestPath(w http.ResponseWriter, r *http.Request) {
	var input domain.Input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.FindClosestDepot(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
