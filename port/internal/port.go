package internal

type Port struct {
	unloc       UNLOC
	name        Name
	city        City
	country     Country
	alias       AliasCollection
	regions     RegionCollection
	coordinates Coordinates
	province    Province
	timezone    Timezone
	unlocs      UNLOCCollection
	code        Code
}

func NewPort(
	unloc UNLOC,
	name Name,
	city City,
	country Country,
	alias AliasCollection,
	regions RegionCollection,
	coordinates Coordinates,
	province Province,
	timezone Timezone,
	unlocs UNLOCCollection,
	code Code,
) (*Port, error) {
	if !unlocs.Contains(unloc) {
		return nil, NewDomainError(
			InvalidPortConsistency,
			"UNLOC is not present in the UNLOCS collection",
		)
	}

	return &Port{
		unloc,
		name,
		city,
		country,
		alias,
		regions,
		coordinates,
		province,
		timezone,
		unlocs,
		code,
	}, nil
}

func (p *Port) UNLOC() UNLOC {
	return p.unloc
}
