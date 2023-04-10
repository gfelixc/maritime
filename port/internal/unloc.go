package internal

type UNLOC struct {
	value string
}

func NewUNLOC(unloc string) (UNLOC, error) {
	if len(unloc) != 5 {
		return UNLOC{}, NewDomainError(InvalidUNLOC, "unloc must have five characters")
	}

	return UNLOC{
		value: unloc,
	}, nil
}

func (vo UNLOC) String() string {
	return vo.value
}

type UNLOCCollection []UNLOC
