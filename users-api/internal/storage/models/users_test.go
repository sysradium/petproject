package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToProto(t *testing.T) {

	u := &User{
		Id:    uuid.New(),
		Login: "peter",
		Email: "peter@mail.com",
	}

	p := u.ToProto()
	require.NotNil(t, p)

	assert.Equal(t, u.Login, p.Username)
	assert.Equal(t, u.Email, p.Email)
	assert.Equal(t, u.Id.String(), p.Id)
}
