package port

import (
	"context"
	"testing"

	"github.com/gfelixc/maritime/port/internal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreatePort(t *testing.T) {
	t.Run("Should create port in repository", func(t *testing.T) {
		repository := internal.NewMockRepository(t)
		repository.EXPECT().Save(
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("*internal.Port"),
		).Return(nil)

		usecase := CreateUpdatePort{
			repository: repository,
		}

		err := usecase.Create(
			context.TODO(),
			stubCreatePortCommand(),
		)

		require.NoError(t, err)
	})

	t.Run("Should return the expected DomainErrorCode", func(t *testing.T) {
		usecase := CreateUpdatePort{
			repository: internal.NewMockRepository(t),
		}

		t.Run("InvalidUNLOC", func(t *testing.T) {
			var expectedError *internal.DomainError
			test_case := stubCreatePortCommand()
			test_case.UNLOC = ""
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidUNLOC, expectedError.Code())
		})

		t.Run("InvalidName", func(t *testing.T) {
			var expectedError *internal.DomainError

			test_case := stubCreatePortCommand()
			test_case.Name = ""
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidName, expectedError.Code())
		})

		t.Run("InvalidCity", func(t *testing.T) {
			var expectedError *internal.DomainError

			test_case := stubCreatePortCommand()
			test_case.City = ""
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidCity, expectedError.Code())
		})

		t.Run("InvalidCountry", func(t *testing.T) {
			var expectedError *internal.DomainError

			test_case := stubCreatePortCommand()
			test_case.Country = ""
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidCountry, expectedError.Code())
		})

		t.Run("InvalidCoordinates", func(t *testing.T) {
			var expectedError *internal.DomainError

			test_case := stubCreatePortCommand()
			test_case.CoordinatesLatitude = -200
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidCoordinates, expectedError.Code())
		})

		t.Run("InvalidProvince", func(t *testing.T) {
			var expectedError *internal.DomainError

			province := ""
			test_case := stubCreatePortCommand()
			test_case.Province = &province
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidProvince, expectedError.Code())
		})

		t.Run("InvalidTimezone", func(t *testing.T) {
			var expectedError *internal.DomainError

			tz := "Dubai"
			test_case := stubCreatePortCommand()
			test_case.Timezone = &tz
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidTimezone, expectedError.Code())
		})

		t.Run("InvalidPortConsistency", func(t *testing.T) {
			var expectedError *internal.DomainError

			test_case := stubCreatePortCommand()
			test_case.UNLOCS = []string{}
			err := usecase.Create(context.TODO(), test_case)

			require.ErrorAs(t, err, &expectedError)
			require.Equal(t, internal.InvalidPortConsistency, expectedError.Code())
		})

	})
}

func stubCreatePortCommand() CreateUpdatePortCommand {
	timezone := "Asia/Dubai"
	province := "Ajman"
	code := "52000"

	return CreateUpdatePortCommand{
		"AEAJM",
		"Ajman",
		"Ajman",
		"United Arab Emirates",
		[]string{},
		[]string{},
		55.5136433,
		25.4052165,
		&province,
		&timezone,
		[]string{
			"AEAJM",
		},
		&code,
	}
}
