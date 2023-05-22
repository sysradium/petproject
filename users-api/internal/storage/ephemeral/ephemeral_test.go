package ephemeral

import (
	"context"
	"testing"

	"github.com/google/uuid"
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

func TestGet(t *testing.T) {
	e := New()

	u := &models.User{}
	rsp, err := e.Create(context.Background(), u)
	require.NoError(t, err)

	gRsp, err := e.Get(context.Background(), rsp.Id)
	require.NoError(t, err)
	assert.Equal(t, rsp.Id, gRsp.Id)
}

func TestCreate(t *testing.T) {
	e := New()

	u := &models.User{}
	rsp, err := e.Create(context.Background(), u)

	require.NoError(t, err)
	assert.NotEmpty(t, rsp)
}

func TestList(t *testing.T) {
	e := New()

	ids := make(map[uuid.UUID]struct{}, 2)

	for _, email := range []string{"s@s.com", "s1@s.com"} {
		rsp, err := e.Create(
			context.Background(),
			&models.User{Email: email},
		)
		require.NoError(t, err)
		ids[rsp.Id] = struct{}{}
	}

	users, err := e.List(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)

	assert.Contains(t, ids, users[0].Id)

}
