package server

import (
	"context"
	"net"

	"github.com/gfelixc/maritime/port"
	"github.com/gfelixc/maritime/port/internal/persistence"
	"github.com/gfelixc/maritime/port/ui/file"
	portv1 "github.com/gfelixc/maritime/port/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BootGRPCServer(ctx context.Context) error {
	logger, _ := zap.NewProduction()
	logger.Info("Bootstrapping service")

	server := grpc.NewServer()
	reflection.Register(server)
	repository := persistence.NewInMemoryRepository()
	createUpdatePort := port.NewCreateUpdatePort(repository)
	portsDataVolume := file.FileSystemVolume{}
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
			logger.Error("Error starting grpc server", zap.Error(err))
		}
	}()

	<-ctx.Done()

	logger.Info("Shutting down server", zap.Error(err))
	server.GracefulStop()

	return ctx.Err()
}
