package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/SirusDoma/btc-wallet/handler"
	"github.com/SirusDoma/btc-wallet/topup"
)

type createHandler struct {
	service topup.CreateTopupService
}

func NewCreateTopupHandler(service topup.CreateTopupService) topup.CreateTopupHandler {
	return &createHandler{service: service}
}

func (h *createHandler) RegisterRoute(router chi.Router) {
	router.Post("/", h.CreateTopup)
}

func (h *createHandler) CreateTopup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := topup.CreateTopupRequest{}

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := h.service.Create(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	response := struct {
		Data interface{}  `json:"data"`
		Meta handler.Meta `json:"meta"`
	}{
		Data: topup.SerializeTopup(t),
		Meta: handler.Meta{
			Status: http.StatusCreated,
		},
	}

	render.Status(r, response.Meta.Status)
	render.JSON(w, r, response)
}
