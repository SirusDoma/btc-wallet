package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/SirusDoma/btc-wallet/database"
	topup "github.com/SirusDoma/btc-wallet/topup/server"
)

func main() {
	// Setup logging preference
	logConf := zap.NewProductionConfig()
	logConf.EncoderConfig.CallerKey = ""
	logConf.EncoderConfig.TimeKey = "timestamp"
	logConf.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	logConf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Initialize the logger
	logger, err := logConf.Build()
	if err != nil {
		fmt.Printf("Failed to initialize logger: %s", err.Error())
		return
	}
	defer logger.Sync()

	// Load configuration (.env file)
	_ = godotenv.Load()

	// Initialize the database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		IgnoreRelationshipsWhenMigrating:         true,
	})
	if err != nil {
		logger.Fatal("Failed to connect to the database server.", zap.Error(err))
		return
	}

	// For convenient sake, automatically migrate database table
	database.MigrateUp(db)

	// Initialize the router
	router := chi.NewRouter()
	topup.Register(router, db, logger)

	// Initialize http server
	rto, _ := strconv.Atoi(os.Getenv("HTTP_READ_TIMEOUT"))
	wto, _ := strconv.Atoi(os.Getenv("HTTP_WRITE_TIMEOUT"))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		ReadTimeout:  time.Duration(rto) * time.Second,
		WriteTimeout: time.Duration(wto) * time.Second,
		Handler:      router,
	}

	// Server run context
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for syscall signals for process to quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(ctx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("Graceful shutdown timed out.")
			}
		}()

		// Trigger graceful shutdown
		err = srv.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		cancel()
	}()

	// Run the server
	logger.Info("Server is running at " + srv.Addr)
	if err = srv.ListenAndServe(); err != nil {
		logger.Fatal("Failed to start the http server.", zap.Error(err))
	}

	// Wait for server context to be stopped
	<-ctx.Done()
}
