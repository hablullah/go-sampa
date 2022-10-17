package main

import (
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	location, dt, limit := wellington()

	for dt.Before(limit) {
		// dt = time.
		e, _ := sampa.GetSunEvents(dt, location, nil)

		fmt.Printf("%s\t%s\t%s\t%s\n",
			dt.Format("2006-01-02"),
			strTime(e.Sunrise.DateTime),
			strTime(e.Transit.DateTime),
			strTime(e.Sunset.DateTime))
		dt = dt.AddDate(0, 0, 1)
	}
}

// tromso is sample for location in north frigid zone.
func tromso() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    69.682778,
		Longitude:   18.942778,
		Elevation:   0,
		Temperature: 10,
		Pressure:    1010,
	}

	tz := time.FixedZone("CST", 1*60*60)
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

// london is sample for location in north temperate zone.
func london() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    51.507222,
		Longitude:   -0.1275,
		Elevation:   11,
		Temperature: 10,
	}

	tz := time.UTC
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

// jakarta is sample for location in torrid zone.
func jakarta() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    -6.175,
		Longitude:   106.825,
		Elevation:   8,
		Temperature: 10,
	}

	tz := time.FixedZone("WIB", 7*60*60)
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

// wellington is sample for location in south temperate zone.
func wellington() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    -41.288889,
		Longitude:   174.777222,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("NZST", 12*60*60)
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

func strTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("15:04:05")
}
