package main

import (
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	location := sampa.Location{
		Latitude:    -6.21138888888889,
		Longitude:   106.845277777778,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("WIB", 7*60*60)
	dt := time.Date(2022, 4, 1, 12, 0, 0, 0, tz)
	limit := dt.AddDate(0, 1, 0)

	for dt.Before(limit) {
		e, _ := sampa.GetMoonEvents(dt, location, nil)
		if e.Transit.DateTime.IsZero() {
			fmt.Printf("%s DOESNT PASS MERIDIAN\n", dt.Format("2006-01-02"))
			dt = dt.AddDate(0, 0, 1)
			continue
		}

		// sun, _ := sampa.GetSunPosition(e.Transit.DateTime, location, nil)
		moon, _ := sampa.GetMoonPosition(e.Transit.DateTime, location, nil)
		fmt.Printf("%s, E=%f, k=%f\n",
			dt.Format("2006-01-02"), moon.Elongation,
			moon.PercentIlluminated)
		dt = dt.AddDate(0, 0, 1)
	}
}
