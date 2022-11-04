package sampa

// Options is the advanced configuration for calculating
// the Sun and Moon position.
type Options struct {
	SurfaceSlope           float64
	SurfaceAzimuthRotation float64
	DeltaT                 float64
}

func setDefaultOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{
			SurfaceSlope:           0,
			SurfaceAzimuthRotation: -180,
			DeltaT:                 66.9,
		}
	}

	if opts.DeltaT == 0 {
		opts.DeltaT = 66.9
	}

	return opts
}
