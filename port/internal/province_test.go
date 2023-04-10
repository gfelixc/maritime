package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProvince(t *testing.T) {
	t.Run("Should fail as it is empty", func(t *testing.T) {
		_, err := NewProvince("")

		require.Error(t, err, "[InvalidProvince] empty country")
	})

	t.Run("Should return a Province with no errors", func(t *testing.T) {
		country, err := NewProvince("Ajman")

		require.NoError(t, err)
		require.Equal(t, "Ajman", country.String())
	})
}
