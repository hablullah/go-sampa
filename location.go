package sampa

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
		l.Pressure = 101325
	}

	if l.Temperature == 0 {
		l.Temperature = 10
	}

	return l
}
