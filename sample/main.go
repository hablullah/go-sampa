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
	dt := time.Date(2022, 1, 1, 0, 0, 0, 0, tz)
	limit := dt.AddDate(1, 0, 0)

	prayerEvents := []sampa.CustomSunEvent{
		{
			Name:          "twilight rise",
			SunElevation:  func(_ sampa.SunData) float64 { return -18 },
			BeforeTransit: true,
		}, {
			Name:          "twilight set",
			SunElevation:  func(_ sampa.SunData) float64 { return -18 },
			BeforeTransit: false,
		},
	}

	for dt.Before(limit) {
		e, _ := sampa.GetSunEvents(dt, location, nil, prayerEvents...)

		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n",
			dt.Format("2006-01-02"),
			e.Others["twilight rise"].DateTime.Format("15:04:05"),
			e.Sunrise.DateTime.Format("15:04:05"),
			e.Transit.DateTime.Format("15:04:05"),
			e.Sunset.DateTime.Format("15:04:05"),
			e.Others["twilight rise"].DateTime.Format("15:04:05"))
		dt = dt.AddDate(0, 0, 1)
	}
}
