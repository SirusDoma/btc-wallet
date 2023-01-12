package topup

import (
	"context"
)

type CreateTopupService interface {
	Create(ctx context.Context, req CreateTopupRequest) (Topup, error)
}

type GetHistoryService interface {
	GetHistory(ctx context.Context, req GetHistoryRequest) (HistoryResultSet, error)
}
