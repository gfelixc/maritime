package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCountry(t *testing.T) {
	t.Run("Should fail as it is empty", func(t *testing.T) {
		_, err := NewCountry("")

		require.Error(t, err, "[InvalidCountry] empty country")
	})

	t.Run("Should return a Country with no errors", func(t *testing.T) {
		country, err := NewCountry("United Arab Emirates")

		require.NoError(t, err)
		require.Equal(t, "United Arab Emirates", country.String())
	})
}
