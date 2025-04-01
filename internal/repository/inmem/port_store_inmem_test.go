package inmem

import (
	"context"
	"testing"

	"github.com/evt/port-service/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPortStore_CreateOrUpdatePort(t *testing.T) {
	t.Parallel()
	store := NewPortStore()

	t.Run("create port", func(t *testing.T) {
		t.Parallel()

		randomPort := newRandomDomainPort(t)

		err := store.CreateOrUpdatePort(context.Background(), randomPort)
		require.NoError(t, err)

		port, err := store.GetPort(context.Background(), randomPort.ID())
		require.NoError(t, err)

		require.Equal(t, port, randomPort)
	})
	t.Run("update port", func(t *testing.T) {
		t.Parallel()

		randomPort := newRandomDomainPort(t)

		err := store.CreateOrUpdatePort(context.Background(), randomPort)
		require.NoError(t, err)

		beforeUpdatePort, err := store.GetPort(context.Background(), randomPort.ID())
		require.NoError(t, err)

		require.Equal(t, beforeUpdatePort, randomPort)

		err = randomPort.SetName("updated name")
		require.NoError(t, err)

		err = store.CreateOrUpdatePort(context.Background(), randomPort)
		require.NoError(t, err)

		updatedPort, err := store.GetPort(context.Background(), randomPort.ID())
		require.NoError(t, err)

		require.NotEqual(t, beforeUpdatePort.Name(), updatedPort.Name())
	})
	t.Run("nil port", func(t *testing.T) {
		t.Parallel()
		err := store.CreateOrUpdatePort(context.Background(), nil)
		require.ErrorIs(t, err, domain.ErrNil)
	})
}

func newRandomDomainPort(t *testing.T) *domain.Port {
	t.Helper()
	randomID := uuid.New().String()
	port, err := domain.NewPort(randomID, randomID, randomID, randomID, randomID,
		[]string{randomID}, []string{randomID}, []float64{1.0, 2.0}, randomID, randomID, nil)
	require.NoError(t, err)

	return port
}
