package main

import (
	"log/slog"
	"opcua-go/opcua"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.Default()

	config := &opcua.ServerConfig{
		Host:               "localhost",
		Port:               4840,
		ReceiverBufferSize: 1024,
		ReadTimeout:        5 * time.Second,
		Logger:             logger,
	}

	server, err := opcua.NewServer(config)
	if err != nil {
		logger.Error("Failed to create server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	port, err := server.Run()
	if err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Server running", slog.Int("port", port))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Info("Shutting down server...")

	if err := server.Close(); err != nil {
		logger.Error("Error closing server", slog.String("error", err.Error()))
	} else {
		logger.Info("Server stopped gracefully")
	}
}
