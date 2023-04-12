package internal

import "time"

type Timezone struct {
	value *time.Location
}

func NoTimezone() Timezone {
	return Timezone{}
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

func (vo Timezone) String() (string, error) {
	if vo.value == nil {
		return "", NewDomainError(
			TimezoneNotProvided,
			"timezone not provided",
		)
	}
	return vo.value.String(), nil
}
