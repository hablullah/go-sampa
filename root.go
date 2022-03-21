package main

import (
	"fmt"
	"time"
)

func main() {
	// Get initial sun data
	location := Location{
		Latitude:    34.1667,
		Longitude:   -118.4667,
		Elevation:   0,
		Temperature: 10.0,
	}

	opts := &SunOptions{
		Pressure:               1010,
		SurfaceSlope:           30,
		SurfaceAzimuthRotation: -10,
		DeltaT:                 67,
	}

	tz := time.FixedZone("LST", -8*60*60)
	dt := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	// data, _ := GetSunAtTime(dt, location, opts)

	// bt, _ := json.MarshalIndent(&data, "", "\t")
	// fmt.Println(string(bt))

	transit, _ := GetSunTransit(dt, location, opts)
	fmt.Printf("TRANSIT: %s, elevation: %f\n", transit.DateTime, transit.TopocentricElevationAngle)

	sunrise, _ := GetSunAtElevation(-0.8333, transit, true, location, opts)
	fmt.Printf("SUNRISE: %s, elevation: %f\n", sunrise.DateTime, sunrise.TopocentricElevationAngle)

	sunset, _ := GetSunAtElevation(-0.8333, transit, false, location, opts)
	fmt.Printf("SUNSET: %s, elevation: %f\n", sunset.DateTime, sunset.TopocentricElevationAngle)
}
