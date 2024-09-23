package main

import (
	"github.com/shoothzj/gox/netx"
	"log/slog"
	"opcua-go/opcua"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.Default()

	config := &opcua.ClientConfig{
		Address: netx.Address{
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

	messageAcknowledge, err := client.Hello(&opcua.MessageHello{
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
