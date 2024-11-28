package logger

import (
	"io"
	"log/slog"
	"os"
)

const (
	logFormatJSON = "json"
	logFormatText = "text"
)

func SetDefaultLogger(logFile *os.File, logLevel, logFormat string) {
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	var slogLevel slog.Level
	switch logLevel {
	case slog.LevelDebug.String():
		slogLevel = slog.LevelDebug
	case slog.LevelInfo.String():
		slogLevel = slog.LevelInfo
	case slog.LevelWarn.String():
		slogLevel = slog.LevelWarn
	case slog.LevelError.String():
		slogLevel = slog.LevelError
	default:
		panic("unknown log level:" + logLevel)
	}

	var logger *slog.Logger

	switch logFormat {
	case logFormatText:
		logger = slog.New(
			slog.NewTextHandler(multiWriter, &slog.HandlerOptions{Level: slogLevel}))
	case logFormatJSON:
		logger = slog.New(
			slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{Level: slogLevel}))
	default:
		panic("unknown log format:" + logFormat)
	}

	slog.SetDefault(logger)
}
