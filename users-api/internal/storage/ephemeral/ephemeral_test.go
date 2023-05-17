package ephemeral

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sysradium/petproject/users-api/internal/storage/models"
	"github.com/tj/assert"
)

func TestInitiallyHasNoObjects(t *testing.T) {
	e := New()

	u, err := e.List(context.Background())
	require.NoError(t, err)
	assert.Empty(t, u)
}

func TestCreate(t *testing.T) {
	e := New()

	u := &models.User{}
	rsp, err := e.Create(context.Background(), u)

	require.NoError(t, err)
	assert.Equal(t, "1", rsp)
}

func TestList(t *testing.T) {
	e := New()

	for _, email := range []string{"s@s.com", "s1@s.com"} {
		_, err := e.Create(
			context.Background(),
			&models.User{Email: email},
		)
		require.NoError(t, err)
	}

	users, err := e.List(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)

	assert.Equal(t, "1", users[0].Id)

}
