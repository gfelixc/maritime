package internal

type Province struct {
	value string
}

func NewProvince(value string) (Province, error) {
	if value == "" {
		return Province{}, NewDomainError(InvalidProvince, "empty province")
	}

	return Province{
		value: value,
	}, nil
}

func (vo Province) String() string {
	return vo.value
}
