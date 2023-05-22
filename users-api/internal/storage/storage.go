package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/users-api/internal/storage/models"
)

var (
	ErrNotFound = errors.New("user not found")
)

type Storage interface {
	Creator
	Lister
	Getter
}

type Creator interface {
	Create(context.Context, *models.User) (models.User, error)
}

type Getter interface {
	Get(context.Context, uuid.UUID) (models.User, error)
}

type Lister interface {
	List(context.Context) ([]*models.User, error)
}
