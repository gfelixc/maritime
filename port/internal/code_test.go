package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCode(t *testing.T) {
	t.Run("Should fail representing Code as String as it was not provided", func(t *testing.T) {
		code := NoCode()

		_, err := code.String()
		require.Error(t, err, "[CodeNotProvided] code not provided")
	})

	t.Run("Should return a Code with no errors", func(t *testing.T) {
		code := NewCode("52000")

		s, err := code.String()
		require.NoError(t, err)
		require.Equal(t, "52000", s)
	})
}
