package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToProto(t *testing.T) {

	u := &User{
		Id:    "1-1-1-1",
		Login: "peter",
		Email: "peter@mail.com",
	}

	p := u.ToProto()
	require.NotNil(t, p)

	assert.Equal(t, u.Login, p.Username)
	assert.Equal(t, u.Email, p.Email)
	assert.Equal(t, u.Id, p.Id)
}
