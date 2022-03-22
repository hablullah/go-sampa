package main

import (
	"fmt"
	"time"
)

func main() {
	location := Location{
		Latitude:    34.1667,
		Longitude:   -118.4667,
		Elevation:   0,
		Temperature: 10.0,
	}

	opts := &SunOptions{
		Pressure:               1010,
		SurfaceSlope:           0,
		SurfaceAzimuthRotation: -0,
		DeltaT:                 69.8,
	}

	tz := time.FixedZone("LST", -8*60*60)
	dt := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := time.Date(2023, 1, 1, 0, 0, 0, 0, tz)

	for dt.Before(limit) {
		sunrise, _ := GetSunAtElevation(dt, -0.8333, true, location, opts)
		transit, _ := GetSunTransit(dt, location, opts)
		sunset, _ := GetSunAtElevation(dt, -0.8333, false, location, opts)
		fmt.Printf("%s\t%s\t%s\t%s\n",
			dt.Format("2006-01-02"),
			sunrise.DateTime.Format("15:04:05"),
			transit.DateTime.Format("15:04:05"),
			sunset.DateTime.Format("15:04:05"))
		dt = dt.AddDate(0, 0, 1)
	}
}
