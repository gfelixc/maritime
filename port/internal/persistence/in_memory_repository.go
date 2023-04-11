package persistence

import (
	"context"
	"sync"

	"github.com/gfelixc/maritime/port/internal"
)

type InMemoryRepository struct {
	ports map[string]*internal.Port
	mutex sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		ports: map[string]*internal.Port{},
	}
}

func (i *InMemoryRepository) Save(_ context.Context, port *internal.Port) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.ports[port.UNLOC().String()] = port

	return nil
}
