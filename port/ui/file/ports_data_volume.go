package file

import (
	"io"
)

type PortsDataVolume interface {
	Open(string) (io.ReadCloser, error)
}
