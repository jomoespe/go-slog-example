package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Message", "key-1", "value 1", "key-2", "value 2")

	//slog.Info("Message", "key-1", "value 1", "key-2")
	// Since the time_taken_ms key does not have a corresponding value, it will be treated as a value with key !BADKEY.
	// you can run vet command to check this

	// also strongly typed contextual attributes can be used
	slog.Info(
		"Request",
		slog.String("method", "GET"),
		slog.String("path", "/path/to/resource"),
		slog.Group("Response",
			slog.Int("status", 200),
			slog.Duration("duration", 175*time.Millisecond),
		),
	)

}
