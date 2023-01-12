package service

import (
	"context"
	"errors"
	"github.com/SirusDoma/btc-wallet/topup"
	"go.uber.org/zap"
	"time"
)

type historySvc struct {
	repo   topup.GetHistoryRepository
	logger *zap.Logger
}

func NewGetHistoryService(repo topup.GetHistoryRepository, logger *zap.Logger) topup.GetHistoryService {
	return &historySvc{
		repo: repo,
		logger: logger.With(
			zap.String("type", "service"),
			zap.String("service", "topup.GetHistoryService"),
		),
	}
}

func (svc *historySvc) GetHistory(ctx context.Context, req topup.GetHistoryRequest) (topup.HistoryResultSet, error) {
	if req.Start.IsZero() {
		err := errors.New("start date cannot be empty")
		svc.logger.Error("Failed to get top-up histories",
			zap.String("field", "start"),
			zap.Error(err),
		)

		return topup.HistoryResultSet{}, err
	}
	if req.End.IsZero() {
		err := errors.New("end date cannot be empty")
		svc.logger.Error("Failed to get top-up histories",
			zap.String("field", "end"),
			zap.Error(err),
		)

		return topup.HistoryResultSet{}, err
	}
	if req.End.Before(req.Start) {
		err := errors.New("end date must be greater than start date")
		svc.logger.Error("Failed to get top-up histories",
			zap.String("field", "end"),
			zap.Error(err),
		)

		return topup.HistoryResultSet{}, err
	}

	t := time.Now()
	r, err := svc.repo.GetHistory(ctx, req)
	l := svc.logger.With(
		zap.Time("start", req.Start),
		zap.Time("end", req.End),
		zap.Duration("duration", time.Since(t)),
	)

	if err != nil {
		l.Error(err.Error())
		return topup.HistoryResultSet{}, err
	}

	l.Info("OK")
	return r, err
}
