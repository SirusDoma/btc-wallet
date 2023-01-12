package repository

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/SirusDoma/btc-wallet/topup"
)

type historyRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewGetHistoryRepositorySQL(db *gorm.DB, logger *zap.Logger) topup.GetHistoryRepository {
	return &historyRepo{db: db, logger: logger.With(
		zap.String("type", "repository"),
		zap.String("repository", "topup.GetHistoryRepository"),
	)}
}

func (repo *historyRepo) GetHistory(ctx context.Context, req topup.GetHistoryRequest) (topup.HistoryResultSet, error) {
	t := time.Now().UTC()
	query := repo.db.WithContext(ctx).Model(&topup.Topup{}).
		Distinct("DATE_TRUNC('hour', created_at)").
		Where("created_at >= ? AND created_at <= ?",
			req.Start,
			req.End,
		)

	var total int64
	result := query.Count(&total)
	l := repo.logger.With(zap.Duration("duration", time.Since(t)))
	if result.Error != nil {
		l.Error("Failed to retrieve history count", zap.Error(result.Error))
		return topup.HistoryResultSet{}, result.Error
	}

	query = repo.db.WithContext(ctx).Model(&topup.Topup{}).
		Select(
			"DATE_TRUNC('hour', created_at) AS datetime",
			"SUM(amount) AS amount",
		).
		Where("created_at >= ? AND created_at <= ?",
			req.Start,
			req.End,
		).
		Group("datetime").
		Order("datetime")

	if req.Limit > 0 {
		query = query.Limit(int(req.Limit))
	}
	if req.Offset > 0 {
		query = query.Offset(int(req.Offset))
	}

	var histories []topup.History
	result = query.Find(&histories)

	l = repo.logger.With(zap.Duration("duration", time.Since(t)))
	if result.Error != nil {
		l.Error("Failed to retrieve history data", zap.Error(result.Error))
		return topup.HistoryResultSet{}, result.Error
	}

	l.Info("OK")
	return topup.HistoryResultSet{
		ResultSet: topup.ResultSet{
			Offset: req.Offset,
			Limit:  req.Limit,
			Total:  uint(total),
		},
		Histories: histories,
	}, nil
}
