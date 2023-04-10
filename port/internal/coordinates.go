package internal

type Coordinates struct {
	longitude float64
	latitude  float64
}

func NewCoordinates(longitude float64, latitude float64) (Coordinates, error) {
	if longitude < -180 || longitude > 180 {
		return Coordinates{}, NewDomainError(InvalidCoordinates, "longitude must be between [-180,180] both inclusive")
	}

	if latitude < -90 || latitude > 90 {
		return Coordinates{}, NewDomainError(InvalidCoordinates, "latitude must be between [-90,90] both inclusive")
	}

	return Coordinates{
		longitude: longitude,
		latitude:  latitude,
	}, nil
}

func (c Coordinates) Longitude() float64 {
	return c.longitude
}

func (c Coordinates) Latitude() float64 {
	return c.latitude
}
