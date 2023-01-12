package topup

import (
	"time"
)

type TopupResponse struct {
	ID        uint64    `json:"id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type HistoryResponse struct {
	Amount float64   `json:"amount"`
	Period time.Time `json:"datetime"`
}

func SerializeTopup(topup Topup) TopupResponse {
	return TopupResponse{
		ID:        topup.ID,
		Amount:    topup.Amount,
		CreatedAt: topup.CreatedAt,
		UpdatedAt: topup.UpdatedAt,
	}
}

func SerializeHistory(history History) HistoryResponse {
	return HistoryResponse{
		Amount: history.Amount,
		Period: history.Period,
	}
}
