package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gfelixc/maritime/port"
	"github.com/gfelixc/maritime/port/ui/file"
	"github.com/gfelixc/maritime/port/v1"
	"go.uber.org/zap"
)

type PortDomainGRPCServer struct {
	createUpdatePort *port.CreateUpdatePort
	portsDataVolume  file.PortsDataVolume
	logger           *zap.Logger
	portv1.UnimplementedPortDomainServiceServer
}

func NewPortDomainGRPCServer(
	createUpdatePort *port.CreateUpdatePort,
	portsDataVolume file.PortsDataVolume,
	logger *zap.Logger,
) PortDomainGRPCServer {
	return PortDomainGRPCServer{
		createUpdatePort: createUpdatePort,
		portsDataVolume:  portsDataVolume,
		logger:           logger,
	}
}

func (p *PortDomainGRPCServer) CreateOrUpdateFromPortsDataFile(ctx context.Context, request *portv1.CreateOrUpdateFromPortsDataFileRequest) (response *portv1.CreateOrUpdateFromPortsDataFileResponse, err error) {
	source, err := p.portsDataVolume.Open(request.Filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closingErr := source.Close(); closingErr != nil {
			err = errors.Join(err, closingErr)
		}
	}()

	parser := file.NewPortsDataParser(source)

	response = &portv1.CreateOrUpdateFromPortsDataFileResponse{}
	for {
		nextPort, err := parser.NextPort()
		if errors.Is(err, io.EOF) {
			return response, nil
		}
		if err != nil {
			return nil, err
		}

		response.PortsProcessed = response.PortsProcessed + 1

		if len(nextPort.Coordinates) != 2 {
			nextPort.Coordinates = []float64{0, 0}
		}

		command := port.CreateUpdatePortCommand{
			UNLOC:                nextPort.FieldName,
			Name:                 nextPort.Name,
			City:                 nextPort.City,
			Country:              nextPort.Country,
			Alias:                nextPort.Alias,
			Regions:              nextPort.Regions,
			CoordinatesLongitude: nextPort.Coordinates[0],
			CoordinatesLatitude:  nextPort.Coordinates[1],
			Province:             nextPort.Province,
			Timezone:             nextPort.Timezone,
			UNLOCS:               nextPort.Unlocs,
			Code:                 nextPort.Code,
		}

		err = p.createUpdatePort.Create(ctx, command)
		if err != nil {
			response.PortsFailed = response.PortsFailed + 1
			cmd, marshalErr := json.Marshal(command)
			if marshalErr != nil {
				p.logger.Error("unable to marshal port.CreateUpdatePortCommand", zap.Error(marshalErr), zap.Any("command", command))
				cmd = []byte{}
			}

			response.Errors = append(response.Errors, fmt.Sprintf("%s | %s", err.Error(), string(cmd)))

			continue
		}

		response.PortsSucceeded = response.PortsSucceeded + 1
	}
}
