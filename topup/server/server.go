package server

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/SirusDoma/btc-wallet/topup/handler"
	"github.com/SirusDoma/btc-wallet/topup/repository"
	"github.com/SirusDoma/btc-wallet/topup/service"
)

func Register(router chi.Router, db *gorm.DB, logger *zap.Logger) {
	// Initialize repositories
	createRepo := repository.NewCreateTopupRepositorySQL(db, logger)
	historyRepo := repository.NewGetHistoryRepositorySQL(db, logger)

	// Initialize services
	createSvc := service.NewCreateTopupService(createRepo, logger)
	historySvc := service.NewGetHistoryService(historyRepo, logger)

	// Initializes handlers
	createHandler := handler.NewCreateTopupHandler(createSvc)
	historyHandler := handler.NewGetHistoryHandler(historySvc)

	// Register routes
	router.Route("/top-ups", func(r chi.Router) {
		createHandler.RegisterRoute(r)
		historyHandler.RegisterRoute(r)
	})
}
