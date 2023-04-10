package internal

type Region struct {
	value string
}

func NewRegion(value string) Region {
	return Region{
		value: value,
	}
}

func (vo Region) String() string {
	return vo.value
}

type RegionCollection []Region

func EmptyRegionCollection() RegionCollection {
	return []Region{}
}
