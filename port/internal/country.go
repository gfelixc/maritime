package internal

type Country struct {
	value string
}

func NewCountry(value string) (Country, error) {
	if value == "" {
		return Country{}, NewDomainError(InvalidCountry, "empty country")
	}

	return Country{
		value: value,
	}, nil
}

func (vo Country) String() string {
	return vo.value
}
