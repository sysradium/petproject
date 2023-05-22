package ephemeral

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/sysradium/petproject/orders-api/internal/storage"
	"github.com/tj/assert"
)

func TestContainsNothing(t *testing.T) {
	s := New()

	_, err := s.Get(context.Background(), uuid.New())
	require.ErrorIs(t, err, storage.ErrNotFound)
}

func TestListReturnsNothing(t *testing.T) {
	s := New()

	o, err := s.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, o)
	assert.Empty(t, o)
}

func TestAdd(t *testing.T) {

}
