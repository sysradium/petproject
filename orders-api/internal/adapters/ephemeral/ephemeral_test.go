package ephemeral

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
	"github.com/tj/assert"
)

func TestContainsNothing(t *testing.T) {
	s := New()

	_, err := s.Get(context.Background(), uuid.New())
	require.ErrorIs(t, err, order.ErrNotFound)
}

func TestListReturnsNothing(t *testing.T) {
	s := New()

	o, err := s.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, o)
	assert.Empty(t, o)
}

func TestCreate(t *testing.T) {

	s := New()

	var uuids []uuid.UUID
	for _, name := range []string{"fruit", "vegetable"} {
		o, err := s.Create(context.Background(), order.Order{
			Name: name,
		})

		require.NoError(t, err)
		require.NotNil(t, o)

		uuids = append(uuids, o.ID)
	}

	for _, u := range uuids {
		rsp, err := s.Get(context.Background(), u)
		require.NoError(t, err)
		assert.Equal(t, u.String(), rsp.ID.String())
	}
}

func TestUpdate(t *testing.T) {
	s := New()

	ctx := context.Background()

	newOrder, err := s.Create(ctx, order.Order{
		Name: "something",
	})

	require.NoError(t, err)
	require.NotNil(t, newOrder)
	assert.Equal(t, "something", newOrder.Name)

	updatedOrder, err := s.Update(
		ctx,
		newOrder.ID,
		func(o *order.Order) (*order.Order, error) {
			o.Name = "something else"
			return o, nil
		})

	require.NoError(t, err)
	assert.Equal(t, "something else", updatedOrder.Name)
}
