package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/orders-api/internal/storage/models"
)

type Storage interface {
	Getter
	Lister
	Creator
}

type Getter interface {
	Get(context.Context, uuid.UUID) (models.Order, error)
}

type Lister interface {
	List(context.Context) ([]models.Order, error)
}

type Creator interface {
	Create(context.Context, models.Order) (models.Order, error)
}
