package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCity(t *testing.T) {
	t.Run("Should fail as it is empty", func(t *testing.T) {
		_, err := NewCity("")

		require.Error(t, err, "[InvalidCity] empty city")
	})

	t.Run("Should return a City with no errors", func(t *testing.T) {
		city, err := NewCity("Ajman")

		require.NoError(t, err)
		require.Equal(t, "Ajman", city.String())
	})
}
