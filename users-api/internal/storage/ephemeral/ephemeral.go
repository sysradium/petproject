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
	s map[string]*models.User
}

func (e *Ephemeral) Get(_ context.Context, id uuid.UUID) (models.User, error) {
	if u, ok := e.s[id.String()]; ok {
		return *u, nil
	}

	return models.User{}, storage.ErrNotFound
}

func (e *Ephemeral) Create(_ context.Context, u *models.User) (models.User, error) {
	e.m.Lock()
	defer e.m.Unlock()

	u.Id = uuid.New()

	e.s[u.Id.String()] = u

	return *u, nil
}

func (e *Ephemeral) List(_ context.Context) ([]*models.User, error) {
	e.m.RLock()
	defer e.m.RUnlock()

	var rsp []*models.User

	for _, u := range e.s {
		rsp = append(rsp, u)
	}

	return rsp, nil
}

func New() *Ephemeral {
	return &Ephemeral{
		s: make(map[string]*models.User),
	}
}
