package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewName(t *testing.T) {
	t.Run("Should fail as it is empty", func(t *testing.T) {
		_, err := NewName("")

		require.Error(t, err, "[InvalidName] empty name")
	})

	t.Run("Should return a Name with no errors", func(t *testing.T) {
		name, err := NewName("Ajman")

		require.NoError(t, err)
		require.Equal(t, "Ajman", name.String())
	})
}
