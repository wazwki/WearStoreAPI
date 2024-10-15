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

// @Summary      Get wear by ID
// @Description  Returns a specific wear item by ID
// @Tags         wear
// @Param        id   path   string  true  "Wear ID"
// @Produce      json
// @Success      200  {object}  models.Item
// @Failure      400  {string}  string  "Bad Request"
// @Router       /wear/{id} [get]
// @Failure      400
func (h *ProductHandler) GetWearHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	item, err := h.service.GetWearData(id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(item)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// @Summary      Get all wears
// @Description  Returns all available wear items
// @Tags         wear
// @Produce      json
// @Success      200  {array}   models.Item
// @Failure      400  {string}  string  "Bad Request"
// @Router       /wear [get]
// @Failure      400
func (h *ProductHandler) GetAllWearHandler(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAllWearData()
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(items)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

// @Summary      Create a new wear item
// @Description  Creates a new wear item
// @Tags         wear
// @Accept       json
// @Produce      json
// @Param        item   body   models.Item  true  "New Wear Item"
// @Success      201    {object}  models.Item
// @Failure      400    {string}  string  "Bad Request"
// @Failure      500    {string}  string  "Internal Server Error"
// @Router       /wear [post]
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

// @Summary      Update a wear item
// @Description  Updates an existing wear item by ID
// @Tags         wear
// @Accept       json
// @Produce      json
// @Param        id     path   string      true  "Wear ID"
// @Param        item   body   models.Item  true  "Updated Wear Item"
// @Success      200    {object}  models.Item
// @Failure      400    {string}  string  "Bad Request"
// @Failure      500    {string}  string  "Internal Server Error"
// @Router       /wear/{id} [put]
func (h *ProductHandler) UpdateWearHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var item *models.Item

	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.service.UpdateWear(item, id)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary      Delete a wear item
// @Description  Deletes a wear item by ID
// @Tags         wear
// @Param        id   path   string  true  "Wear ID"
// @Success      204  "No Content"
// @Failure      500  {string}  string  "Internal Server Error"
// @Router       /wear/{id} [delete]
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
