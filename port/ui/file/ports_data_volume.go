package file

import (
	"context"
	"io"
)

type PortsDataVolume interface {
	Open(context.Context, string) (io.ReadCloser, error)
}
