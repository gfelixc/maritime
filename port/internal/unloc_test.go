package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUNLOC(t *testing.T) {
	t.Run("Should fail as it has more than 5 characters", func(t *testing.T) {
		_, err := NewUNLOC("ABCDEF")

		require.Error(t, err, "[InvalidUNLOC] unloc must have five characters")
	})

	t.Run("Should fail as it has less than 5 characters", func(t *testing.T) {
		_, err := NewUNLOC("ABCD")

		require.Error(t, err, "[InvalidUNLOC] unloc must have five characters")
	})

	t.Run("Should return a UNLOC with no errors", func(t *testing.T) {
		unloc, err := NewUNLOC("ABCDE")

		require.NoError(t, err)
		require.Equal(t, "ABCDE", unloc.String())
	})
}
