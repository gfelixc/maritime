package internal

const notAvailable = "notAvailable"

type Code struct {
	value string
}

func NoCode() Code {
	return Code{
		value: notAvailable,
	}
}

func NewCode(value string) Code {
	return Code{
		value: value,
	}
}

func (vo Code) String() (string, error) {
	if vo.value == notAvailable {
		return "", NewDomainError(
			CodeNotProvided,
			"code not provided",
		)
	}

	return vo.value, nil
}
