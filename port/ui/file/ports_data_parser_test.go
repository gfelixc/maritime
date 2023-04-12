package file

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPortsDataParser(t *testing.T) {
	reader := strings.NewReader(inMemoryPortsData)

	parser := NewPortsDataParser(reader)

	t.Run("Should parse AEAJM properly", func(t *testing.T) {
		port, err := parser.NextPort()

		require.NoError(t, err)
		require.Equal(t, "AEAJM", port.FieldName)
		require.Equal(t, "Ajman", port.Name)
		require.Equal(t, "Ajman", port.City)
		require.Equal(t, "United Arab Emirates", port.Country)
		require.Equal(t, []string{}, port.Alias)
		require.Equal(t, []string{}, port.Regions)
		require.Equal(t, []float64{55.5136433, 25.4052165}, port.Coordinates)
		require.Equal(t, "Ajman", *port.Province)
		require.Equal(t, "Asia/Dubai", port.Timezone)
		require.Equal(t, []string{"AEAJM"}, port.Unlocs)
		require.Equal(t, "52000", *port.Code)
	})

	t.Run("Should parse missing Code from AEAUH as nil", func(t *testing.T) {
		port, err := parser.NextPort()

		require.NoError(t, err)
		require.Equal(t, "AEAUH", port.FieldName)
		require.Equal(t, (*string)(nil), port.Code)
	})

	t.Run("Should return io.EOF", func(t *testing.T) {
		_, err := parser.NextPort()

		require.ErrorIs(t, err, io.EOF)
	})
}

var inMemoryPortsData = `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ]
  }
}`
