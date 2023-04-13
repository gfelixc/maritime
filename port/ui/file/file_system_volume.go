package file

import (
	"context"
	"io"
	"os"
)

type FileSystemVolume struct{}

func (i FileSystemVolume) Open(ctx context.Context, s string) (io.ReadCloser, error) {
	return os.Open(s)
}
