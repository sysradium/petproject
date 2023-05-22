package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Name   string
}
