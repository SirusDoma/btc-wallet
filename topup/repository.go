package topup

import "context"

type CreateTopupRepository interface {
	Create(ctx context.Context, req CreateTopupRequest) (Topup, error)
}

type GetHistoryRepository interface {
	GetHistory(ctx context.Context, req GetHistoryRequest) (HistoryResultSet, error)
}
