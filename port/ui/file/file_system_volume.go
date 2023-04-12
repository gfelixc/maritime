package file

import (
	"io"
	"os"
)

type FileSystemVolume struct{}

func (i FileSystemVolume) Open(s string) (io.ReadCloser, error) {
	return os.Open(s)
}
