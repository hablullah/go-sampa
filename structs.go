package sampa

type Options struct {
	SurfaceSlope           float64
	SurfaceAzimuthRotation float64
	DeltaT                 float64
}

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

func setDefaultOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{
			SurfaceSlope:           0,
			SurfaceAzimuthRotation: -180,
			DeltaT:                 66.9,
		}
	}

	return opts
}
