package order

import (
	"context"

	"github.com/google/uuid"
)

type UpdaterFn func(*Order) (*Order, error)

type Repository interface {
	Getter
	Lister
	Creator
	Updater
}

type Getter interface {
	Get(context.Context, uuid.UUID) (*Order, error)
}

type Lister interface {
	List(context.Context) ([]*Order, error)
}

type Creator interface {
	Create(context.Context, Order) (*Order, error)
}

type Updater interface {
	Update(context.Context, uuid.UUID, UpdaterFn) (*Order, error)
}
