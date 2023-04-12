package internal

const provinceNotProvided = "provinceNotProvided"

type Province struct {
	value string
}

func NoProvince() Province {
	return Province{
		value: provinceNotProvided,
	}
}

func NewProvince(value string) (Province, error) {
	if value == "" {
		return Province{}, NewDomainError(InvalidProvince, "empty province")
	}

	return Province{
		value: value,
	}, nil
}

func (vo Province) String() (string, error) {
	if vo.value == provinceNotProvided {
		return "", NewDomainError(
			ProvinceNotProvided,
			"province not provided",
		)
	}

	return vo.value, nil
}
