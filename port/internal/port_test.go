package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPort(t *testing.T) {
	t.Run("Should fail as UNLOC is not in the UNLOCS list", func(t *testing.T) {
		unloc, _ := NewUNLOC("AEAJM")
		name, _ := NewName("Ajman")
		city, _ := NewCity("Ajman")
		country, _ := NewCountry("United Arab Emirates")
		alias := EmptyAliasCollection()
		regions := EmptyRegionCollection()
		coordinates, _ := NewCoordinates(
			55.5136433,
			25.4052165,
		)
		province, _ := NewProvince("Ajman")
		timezone, _ := NewTimezone("Asia/Dubai")
		unlocs := EmptyUNLOCCollection()
		code := NoCode()

		_, err := NewPort(
			unloc,
			name,
			city,
			country,
			alias,
			regions,
			coordinates,
			province,
			timezone,
			unlocs,
			code,
		)

		require.Error(t, err, "[InvalidPort] UNLOC is not present in the UNLOCS collection")
	})

	t.Run("Should return a Port with no errors", func(t *testing.T) {
		unloc, _ := NewUNLOC("AEAJM")
		name, _ := NewName("Ajman")
		city, _ := NewCity("Ajman")
		country, _ := NewCountry("United Arab Emirates")
		alias := EmptyAliasCollection()
		regions := EmptyRegionCollection()
		coordinates, _ := NewCoordinates(
			55.5136433,
			25.4052165,
		)
		province, _ := NewProvince("Ajman")
		timezone, _ := NewTimezone("Asia/Dubai")
		unlocs := NewUNLOCCollection(unloc)
		code := NoCode()

		_, err := NewPort(
			unloc,
			name,
			city,
			country,
			alias,
			regions,
			coordinates,
			province,
			timezone,
			unlocs,
			code,
		)

		require.NoError(t, err)
	})
}
