package sampa

// Location is coordinate and data for the location where Sun and Moon
// position will be calculated.
type Location struct {
	Latitude    float64
	Longitude   float64
	Elevation   float64
	Temperature float64
	Pressure    float64
}

func setDefaultLocation(l Location) Location {
	if l.Elevation < 0 {
		l.Elevation = 0
	}

	if l.Pressure == 0 {
		l.Pressure = 1013.25
	}

	if l.Temperature == 0 {
		l.Temperature = 10
	}

	return l
}
