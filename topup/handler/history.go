package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/SirusDoma/btc-wallet/handler"
	"github.com/SirusDoma/btc-wallet/topup"
)

type historyHandler struct {
	service topup.GetHistoryService
}

func NewGetHistoryHandler(service topup.GetHistoryService) topup.GetHistoryHandler {
	return &historyHandler{service: service}
}

func (h *historyHandler) RegisterRoute(router chi.Router) {
	router.Get("/", h.GetHistory)
}

func (h *historyHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()

	start, err := time.Parse(time.RFC3339, query.Get("startDatetime"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, query.Get("endDatetime"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	offset := 0
	if query.Has("offset") {
		offset, err = strconv.Atoi(query.Get("offset"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	limit := 0
	if query.Has("limit") {
		limit, err = strconv.Atoi(query.Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	req := topup.GetHistoryRequest{
		Start:  start,
		End:    end,
		Offset: uint(offset),
		Limit:  uint(limit),
	}

	ht, err := h.service.GetHistory(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data := make([]topup.HistoryResponse, len(ht.Histories))
	for i, v := range ht.Histories {
		data[i] = topup.SerializeHistory(v)
	}

	response := struct {
		Data []topup.HistoryResponse `json:"data"`
		Meta handler.PaginationMeta  `json:"meta"`
	}{
		Data: data,
		Meta: handler.PaginationMeta{
			Meta:   handler.Meta{Status: http.StatusOK},
			Offset: ht.Offset,
			Limit:  ht.Limit,
			Total:  ht.Total,
		},
	}

	render.Status(r, response.Meta.Status)
	render.JSON(w, r, response)
}
