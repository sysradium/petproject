package ephemeral

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/users-api/internal/storage"
	"github.com/sysradium/petproject/users-api/internal/storage/models"
)

var _ storage.Storage = (*Ephemeral)(nil)

type Ephemeral struct {
	m sync.RWMutex
	s []*models.User
}

func (e *Ephemeral) Create(_ context.Context, u *models.User) (string, error) {
	e.m.Lock()
	defer e.m.Unlock()

	e.s = append(e.s, u)

	u.RowId = len(e.s)
	u.Id = uuid.New()

	return u.Id.String(), nil
}

func (e *Ephemeral) List(_ context.Context) ([]*models.User, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	return e.s, nil
}

func New() *Ephemeral {
	return &Ephemeral{}
}
