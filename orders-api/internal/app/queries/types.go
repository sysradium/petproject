package queries

import "github.com/google/uuid"

type Order struct {
	UserID uuid.UUID
	Name   string
}
