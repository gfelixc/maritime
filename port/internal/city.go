package internal

type City struct {
	value string
}

func NewCity(value string) (City, error) {
	if value == "" {
		return City{}, NewDomainError(InvalidCity, "empty city")
	}

	return City{
		value: value,
	}, nil
}

func (vo City) String() string {
	return vo.value
}
