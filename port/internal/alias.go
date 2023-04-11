package internal

type Alias struct {
	value string
}

func NewAlias(value string) Alias {
	return Alias{
		value: value,
	}
}

func (vo Alias) String() string {
	return vo.value
}

type AliasCollection []Alias

func EmptyAliasCollection() AliasCollection {
	return []Alias{}
}

func NewAliasCollection(alias ...Alias) AliasCollection {
	return alias
}
