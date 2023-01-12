package topup

import "time"

type CreateTopupRequest struct {
	Amount float64 `json:"amount"`
}

type GetHistoryRequest struct {
	Start  time.Time `json:"startDatetime"`
	End    time.Time `json:"endDatetime"`
	Offset uint      `json:"offset"`
	Limit  uint      `json:"limit"`
}
