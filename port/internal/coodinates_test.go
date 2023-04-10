package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCoordinates(t *testing.T) {
	t.Run("Should fail as longitude is more than 180", func(t *testing.T) {
		_, err := NewCoordinates(180.01, 0)

		require.Error(t, err, "[InvalidCoordinates] longitude must be between [-180,180] both inclusive")
	})

	t.Run("Should fail as longitude is less than -180", func(t *testing.T) {
		_, err := NewCoordinates(-180.01, 0)

		require.Error(t, err, "[InvalidCoordinates] longitude must be between [-180,180] both inclusive")
	})

	t.Run("Should fail as latitude is more than 90", func(t *testing.T) {
		_, err := NewCoordinates(0, 90.01)

		require.Error(t, err, "[InvalidCoordinates] latitude must be between [-90,90] both inclusive")
	})

	t.Run("Should fail as latitude is less than -90", func(t *testing.T) {
		_, err := NewCoordinates(0, -90.01)

		require.Error(t, err, "[InvalidCoordinates] latitude must be between [-90,90] both inclusive")
	})

	t.Run("Should return a Coordinates with no errors", func(t *testing.T) {
		coordinates, err := NewCoordinates(
			55.5136433,
			25.4052165,
		)

		require.NoError(t, err)
		require.Equal(t, 55.5136433, coordinates.Longitude())
		require.Equal(t, 25.4052165, coordinates.Latitude())
	})
}
