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

func (vo UNLOC) Equals(v UNLOC) bool {
	return vo.value == v.value
}

type UNLOCCollection []UNLOC

func NewUNLOCCollection(unlocs ...UNLOC) UNLOCCollection {
	return unlocs
}

func EmptyUNLOCCollection() UNLOCCollection {
	return []UNLOC{}
}

func (c UNLOCCollection) Contains(unloc UNLOC) bool {
	if c == nil {
		return false
	}

	for _, v := range c {
		if unloc.Equals(v) {
			return true
		}
	}

	return false
}
