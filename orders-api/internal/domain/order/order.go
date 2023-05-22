package order

import (
	"github.com/google/uuid"
)

type Order struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Name   string
}

type Factory struct {
}

func (f Factory) New(userID uuid.UUID, name string) (*Order, error) {
	return &Order{
		ID:     uuid.New(),
		UserID: userID,
		Name:   name,
	}, nil
}
