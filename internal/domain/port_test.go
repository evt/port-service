package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPort(t *testing.T) {
	t.Parallel()

	portID := "port id"
	portCode := "port code"
	portName := "port name"
	portCity := "port city"
	portCountry := "port country"

	t.Run("valid", func(t *testing.T) {
		port, err := NewPort(portID, portName, portCode, portCity, portCountry,
			nil, nil, nil, "", "", nil)
		require.NoError(t, err)

		require.Equal(t, portID, port.ID())
		require.Equal(t, portCode, port.Code())
		require.Equal(t, portName, port.Name())
		require.Equal(t, portCity, port.City())
		require.Equal(t, portCountry, port.Country())
	})

	t.Run("missing port ID", func(t *testing.T) {
		_, err := NewPort("", portName, portCode, portCity, portCountry,
			nil, nil, nil, "", "", nil)
		require.Error(t, err)
	})

	t.Run("missing port name", func(t *testing.T) {
		_, err := NewPort(portID, "", portCode, portCity, portCountry,
			nil, nil, nil, "", "", nil)
		require.Error(t, err)
	})

	t.Run("missing port city", func(t *testing.T) {
		_, err := NewPort(portID, portName, portCode, "", portCountry,
			nil, nil, nil, "", "", nil)
		require.Error(t, err)
	})

	t.Run("missing port country", func(t *testing.T) {
		_, err := NewPort(portID, portName, portCode, portCity, "",
			nil, nil, nil, "", "", nil)
		require.Error(t, err)
	})
}
