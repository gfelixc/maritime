package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTimezone(t *testing.T) {
	t.Run("Should fail as it is empty", func(t *testing.T) {
		_, err := NewTimezone("")

		require.Error(t, err, "[InvalidTimezone] empty timezone")
	})

	t.Run("Should fail as it doesn't exists in the IANA Database", func(t *testing.T) {
		_, err := NewTimezone("Dubai")

		require.Error(t, err, "[InvalidTimezone] timezone provided doesn't exists in IANA timezone database")
		require.Error(t, errors.Unwrap(err))
	})

	t.Run("Should return a Timezone with no errors", func(t *testing.T) {
		country, err := NewTimezone("Asia/Dubai")

		require.NoError(t, err)
		require.Equal(t, "Asia/Dubai", country.String())
	})
}
