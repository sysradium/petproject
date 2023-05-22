package ephemeral

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/orders-api/internal/storage"
	"github.com/sysradium/petproject/orders-api/internal/storage/models"
)

var _ storage.Storage = (*Ephemeral)(nil)

type Ephemeral struct {
	m *sync.RWMutex
	s map[string]models.Order
}

func (e *Ephemeral) Get(_ context.Context, id uuid.UUID) (models.Order, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	if o, ok := e.s[id.String()]; ok {
		return o, nil
	}
	return models.Order{}, storage.ErrNotFound
}

func (e *Ephemeral) List(_ context.Context) ([]models.Order, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	orders := make([]models.Order, len(e.s))

	for _, o := range e.s {
		orders = append(orders, o)
	}

	return orders, nil
}

func New() *Ephemeral {
	return &Ephemeral{
		m: &sync.RWMutex{},
		s: make(map[string]models.Order),
	}
}
