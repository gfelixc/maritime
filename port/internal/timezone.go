package internal

import "time"

type Timezone struct {
	value *time.Location
}

func NewTimezone(value string) (Timezone, error) {
	if value == "" {
		return Timezone{}, NewDomainError(
			InvalidTimezone,
			"empty timezone",
		)
	}

	location, err := time.LoadLocation(value)
	if err != nil {
		return Timezone{}, WrapWithDomainError(
			InvalidTimezone,
			"timezone provided doesn't exists in IANA timezone database",
			err,
		)
	}

	return Timezone{
		value: location,
	}, nil
}

func (vo Timezone) String() string {
	return vo.value.String()
}
