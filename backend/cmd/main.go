package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lananolana/webpulse/backend/internal/config"
	"github.com/lananolana/webpulse/backend/internal/httpclient"
	"github.com/lananolana/webpulse/backend/internal/httpserver"
	"github.com/lananolana/webpulse/backend/internal/httpserver/handlers"
	"github.com/lananolana/webpulse/backend/pkg/closer"
	"github.com/lananolana/webpulse/backend/pkg/logger"
	"github.com/lananolana/webpulse/backend/pkg/logger/sl"
)

const (
	ConfigPath = "./configs/app.yaml"
	LogPath    = "./logs/app.log"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	// Load config
	cfg := config.MustLoad(ConfigPath)

	// Set default logger and logfile
	logFile, err := os.OpenFile(LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	// Закрытие файла логов при выходе из программы
	closer.Add(func(ctx context.Context) error {
		return logFile.Close()
	})

	logger.SetDefaultLogger(logFile, cfg.App.LogLevel, cfg.App.LogFormat)

	// Handle os signals
	syscallCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	siteClient := httpclient.New()

	// Create new http router
	r := chi.NewRouter()

	// Create http server
	httpSrv := httpserver.New(r, cfg.App.HTTP.ListenAddr)

	// Register handlers
	handlers.NewSiteHandler(r, siteClient, cfg.App.Mock)

	// Run http server
	httpSrv.MustRun()

	slog.Info("app started")

	// Wait for syscall
	<-syscallCtx.Done()

	// ctx for graceful shutdown with timeout
	shutdownDeadlineCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Graceful shutdown init
	slog.Debug("starting app graceful shutdown...")
	if err = closer.Close(shutdownDeadlineCtx); err != nil {
		slog.Error("failed to close resources", sl.Err(err))
		return
	}

	slog.Info("app gracefully stopped")
}
