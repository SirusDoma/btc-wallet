package service

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/SirusDoma/btc-wallet/topup"
)

type createSvc struct {
	repo   topup.CreateTopupRepository
	logger *zap.Logger
}

func NewCreateTopupService(repo topup.CreateTopupRepository, logger *zap.Logger) topup.CreateTopupService {
	return &createSvc{
		repo: repo,
		logger: logger.With(
			zap.String("type", "service"),
			zap.String("service", "topup.CreateTopupService"),
		),
	}
}

func (svc *createSvc) Create(ctx context.Context, req topup.CreateTopupRequest) (topup.Topup, error) {
	if req.Amount < 0 {
		err := errors.New("top-up amount must be greater than zero")
		svc.logger.Error("Failed to create top-up",
			zap.String("field", "amount"),
			zap.Error(err),
		)

		return topup.Topup{}, err
	}

	t := time.Now().UTC()
	r, err := svc.repo.Create(ctx, req)

	l := svc.logger.With(zap.Duration("duration", time.Since(t)))
	if err != nil {
		l.Error(err.Error())
		return topup.Topup{}, err
	}

	l.Info("OK")
	return r, err
}
