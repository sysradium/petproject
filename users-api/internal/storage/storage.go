package storage

import (
	"context"

	"github.com/sysradium/petproject/users-api/internal/storage/models"
)

type Storage interface {
	Creator
	Lister
}

type Creator interface {
	Create(context.Context, *models.User) (string, error)
}

type Lister interface {
	List(context.Context) ([]*models.User, error)
}
