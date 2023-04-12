package internal

const codeNotProvided = "codeNotProvided"

type Code struct {
	value string
}

func NoCode() Code {
	return Code{
		value: codeNotProvided,
	}
}

func NewCode(value string) Code {
	return Code{
		value: value,
	}
}

func (vo Code) String() (string, error) {
	if vo.value == codeNotProvided {
		return "", NewDomainError(
			CodeNotProvided,
			"code not provided",
		)
	}

	return vo.value, nil
}
