package handlers

import (
	"WearStoreAPI/internal/models"
	"WearStoreAPI/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetWearHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, err := h.service.GetWearData(id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(item)
	w.Header().Set("Content-Type", "application/json")
}

func (h *ProductHandler) GetAllWearHandler(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAllWearData()
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(items)
	w.Header().Set("Content-Type", "application/json")
}

func (h *ProductHandler) PostWearHandler(w http.ResponseWriter, r *http.Request) {
	var item *models.Item

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.CreateWear(item)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) PatchWearHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var item *models.Item

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.PatchWear(item, id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteWearHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteWear(id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
