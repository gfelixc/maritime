package server

import (
	"context"
	"net"
	"os/signal"

	"github.com/gfelixc/maritime/port"
	"github.com/gfelixc/maritime/port/internal/persistence"
	"github.com/gfelixc/maritime/port/ui/file"
	portv1 "github.com/gfelixc/maritime/port/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BootGRPCServer() error {
	rootCtx := context.Background()
	signalCtx, _ := signal.NotifyContext(rootCtx)

	server := grpc.NewServer()
	reflection.Register(server)
	repository := persistence.NewInMemoryRepository()
	createUpdatePort := port.NewCreateUpdatePort(repository)
	portsDataVolume := file.FileSystemVolume{}
	logger, _ := zap.NewProduction()
	grpcServer := NewPortDomainGRPCServer(createUpdatePort, portsDataVolume, logger)
	portv1.RegisterPortDomainServiceServer(server, &grpcServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Error("Error listening port 8080", zap.Error(err))

		return err
	}

	logger.Info("Starting gRPC server", zap.String("port", listener.Addr().String()))

	go func() {
		err = server.Serve(listener)
		if err != nil {
			logger.Error("error starting server", zap.Error(err))
		}
	}()

	<-signalCtx.Done()

	server.GracefulStop()

	return signalCtx.Err()
}
