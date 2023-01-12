package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/SirusDoma/btc-wallet/topup"
)

type createRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewCreateTopupRepositorySQL(db *gorm.DB, logger *zap.Logger) topup.CreateTopupRepository {
	return &createRepo{db: db, logger: logger.With(
		zap.String("type", "repository"),
		zap.String("repository", "topup.CreateTopupRepository"),
	)}
}

func (repo *createRepo) Create(ctx context.Context, req topup.CreateTopupRequest) (topup.Topup, error) {
	t := time.Now().UTC()
	r := topup.Topup{
		Amount:    req.Amount,
		CreatedAt: t,
		UpdatedAt: t,
	}

	result := repo.db.WithContext(ctx).Create(&r)
	l := repo.logger.With(
		zap.Duration("duration", time.Since(t)),
	)

	if result.Error != nil || result.RowsAffected == 0 {
		l.Error(result.Error.Error())
		return topup.Topup{}, result.Error
	}

	l.Info("OK")
	return r, nil
}
