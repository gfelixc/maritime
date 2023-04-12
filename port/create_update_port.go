package port

import (
	"context"

	"github.com/gfelixc/maritime/port/internal"
)

type CreateUpdatePort struct {
	repository internal.Repository
}

func NewCreateUpdatePort(repository internal.Repository) *CreateUpdatePort {
	return &CreateUpdatePort{repository: repository}
}

type CreateUpdatePortCommand struct {
	UNLOC                string
	Name                 string
	City                 string
	Country              string
	Alias                []string
	Regions              []string
	CoordinatesLongitude float64
	CoordinatesLatitude  float64
	Province             *string
	Timezone             *string
	UNLOCS               []string
	Code                 *string
}

func (uc CreateUpdatePort) Create(
	ctx context.Context,
	command CreateUpdatePortCommand,
) error {
	domainUnloc, err := internal.NewUNLOC(command.UNLOC)
	if err != nil {
		return err
	}

	domainName, err := internal.NewName(command.Name)
	if err != nil {
		return err
	}

	domainCity, err := internal.NewCity(command.City)
	if err != nil {
		return err
	}

	domainCountry, err := internal.NewCountry(command.Country)
	if err != nil {
		return err
	}

	var aliasSlice []internal.Alias

	for _, v := range command.Alias {
		aliasSlice = append(aliasSlice, internal.NewAlias(v))
	}

	domainAlias := internal.NewAliasCollection(aliasSlice...)

	var regionSlice []internal.Region

	for _, v := range command.Regions {
		regionSlice = append(regionSlice, internal.NewRegion(v))
	}

	domainRegions := internal.NewRegionCollection(regionSlice...)

	domainCoordinates, err := internal.NewCoordinates(
		command.CoordinatesLongitude,
		command.CoordinatesLatitude,
	)
	if err != nil {
		return err
	}

	domainProvince := internal.NoProvince()

	if command.Province != nil {
		domainProvince, err = internal.NewProvince(*command.Province)
		if err != nil {
			return err
		}
	}

	domainTimezone := internal.NoTimezone()
	if command.Timezone != nil {
		domainTimezone, err = internal.NewTimezone(*command.Timezone)
		if err != nil {
			return err
		}
	}

	var unlocsSlice []internal.UNLOC

	for _, v := range command.UNLOCS {
		u, err := internal.NewUNLOC(v)
		if err != nil {
			return err
		}
		unlocsSlice = append(unlocsSlice, u)
	}

	domainUnlocs := internal.NewUNLOCCollection(unlocsSlice...)

	domainCode := internal.NoCode()

	if command.Code != nil {
		domainCode = internal.NewCode(*command.Code)
	}

	port, err := internal.NewPort(
		domainUnloc,
		domainName,
		domainCity,
		domainCountry,
		domainAlias,
		domainRegions,
		domainCoordinates,
		domainProvince,
		domainTimezone,
		domainUnlocs,
		domainCode,
	)
	if err != nil {
		return err
	}

	return uc.repository.Save(ctx, port)
}
