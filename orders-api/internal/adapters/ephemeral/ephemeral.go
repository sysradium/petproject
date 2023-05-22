package ephemeral

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
)

var _ order.Repository = (*Ephemeral)(nil)

type Ephemeral struct {
	m *sync.RWMutex
	s map[string]order.Order

	orderFactory order.Factory
}

func (e *Ephemeral) Create(_ context.Context, o order.Order) (*order.Order, error) {
	e.m.Lock()
	defer e.m.Unlock()

	order, err := e.orderFactory.New(o.UserID, o.Name)
	if err != nil {
		return nil, err
	}

	e.s[order.ID.String()] = *order

	return order, nil

}

func (e *Ephemeral) Get(_ context.Context, id uuid.UUID) (*order.Order, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	if o, ok := e.s[id.String()]; ok {
		return &o, nil
	}

	return nil, order.ErrNotFound
}

func (e *Ephemeral) List(_ context.Context) ([]*order.Order, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	orders := make([]*order.Order, 0, len(e.s))

	for _, o := range e.s {
		o := o
		orders = append(orders, &o)
	}

	return orders, nil
}

func (e *Ephemeral) Update(ctx context.Context, id uuid.UUID, updateFn order.UpdaterFn) (*order.Order, error) {
	order, err := e.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	e.m.Lock()
	defer e.m.Unlock()

	updatedOrder, err := updateFn(order)
	if err != nil {
		return nil, err
	}

	e.s[id.String()] = *updatedOrder

	return updatedOrder, nil
}

func New() *Ephemeral {
	return &Ephemeral{
		m: &sync.RWMutex{},
		s: map[string]order.Order{},
	}
}
