package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.Info("Info with customized default logger")

	// Using the SetDefault() method also alters the default log.Logger employed by the log package.
	// This behavior allows existing applications that utilize the older log package to transition to slog

	log.Println("Message with log.Println")
}
