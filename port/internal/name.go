package internal

type Name struct {
	value string
}

func NewName(name string) (Name, error) {
	if name == "" {
		return Name{}, NewDomainError(InvalidName, "empty name")
	}

	return Name{
		value: name,
	}, nil
}

func (vo Name) String() string {
	return vo.value
}
