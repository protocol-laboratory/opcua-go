package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/protocol-laboratory/opcua-go/opcua"
	"github.com/protocol-laboratory/opcua-go/opcua/ua"
)

func main() {
	logger := slog.Default()

	config := &opcua.ClientConfig{
		Address: opcua.Address{
			Host: "localhost",
			Port: 4840,
		},
		Logger: logger,
	}

	client, err := opcua.NewClient(config)
	if err != nil {
		logger.Error("Failed to start client", slog.String("error", err.Error()))
		os.Exit(1)
	}

	messageAcknowledge, err := client.Hello(&ua.MessageHello{
		Version:           0,
		ReceiveBufferSize: 65535,
		SendBufferSize:    65535,
		MaxMessageSize:    2097152,
		MaxChunkCount:     0,
		EndpointUrl:       "opc.tcp://localhost:4840/opcua",
	})

	if err != nil {
		logger.Error("Failed to hello", slog.String("error", err.Error()))
		os.Exit(1)
	}

	logger.Info("Hello response", slog.Any("response", messageAcknowledge))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Info("Shutting down server...")

	client.Close()
	logger.Info("Client closed gracefully")
}
