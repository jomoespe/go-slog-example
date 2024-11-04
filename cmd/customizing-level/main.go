package main

import (
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)

	logger.Debug("Debug message") // Level is debug, so this message will be logged
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")
}
