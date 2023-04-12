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
		province, err := NewProvince("Ajman")
		require.NoError(t, err)

		s, _ := province.String()
		require.Equal(t, "Ajman", s)
	})

	t.Run("Should return a ProvinceNotProvided error", func(t *testing.T) {
		province := NoProvince()
		_, err := province.String()

		var expectedErr *DomainError
		require.ErrorAs(t, err, &expectedErr)
		require.Equal(t, ProvinceNotProvided, expectedErr.Code())
	})
}
