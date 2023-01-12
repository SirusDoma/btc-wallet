package topup

import (
	"net/http"

	"github.com/SirusDoma/btc-wallet/handler"
)

type CreateTopupHandler interface {
	handler.Handler
	CreateTopup(w http.ResponseWriter, r *http.Request)
}

type GetHistoryHandler interface {
	handler.Handler
	GetHistory(w http.ResponseWriter, r *http.Request)
}
