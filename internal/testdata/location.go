package testdata

import "time"

var (
	WIB = time.FixedZone("WIB", 7*60*60)
	PST = time.FixedZone("PST", -8*60*60)
	CET = time.FixedZone("CET", 1*60*60)
)

var (
	jakarta = TestLocation{
		Latitude:    -6.21138888888889,
		Longitude:   106.845277777778,
		Elevation:   0,
		Temperature: 10,
	}

	losAngeles = TestLocation{
		Latitude:    34.16667,
		Longitude:   -118.46667,
		Elevation:   0,
		Temperature: 10,
	}

	tromso = TestLocation{
		Latitude:    69.1044,
		Longitude:   18.0594,
		Elevation:   0,
		Temperature: 10,
	}
)
