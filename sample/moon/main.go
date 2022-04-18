package main

import (
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	location, dt, limit := jakarta()

	for dt.Before(limit) {
		e, _ := sampa.GetMoonEvents(dt, location, nil)

		fmt.Printf("%s\t%s\t%s\t%s\n",
			dt.Format("2006-01-02"),
			e.Moonrise.DateTime.Format("15:04:05"),
			e.Transit.DateTime.Format("15:04:05"),
			e.Moonset.DateTime.Format("15:04:05"))
		dt = dt.AddDate(0, 0, 1)
	}
}

func jakarta() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    -6.21138888888889,
		Longitude:   106.845277777778,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("WIB", 7*60*60)
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

func losAngeles() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    34.16667,
		Longitude:   -118.46667,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("EST", -8*60*60)
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

func laPaz() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    -16.5,
		Longitude:   -68.2,
		Elevation:   3812,
		Temperature: 10,
	}

	tz := time.FixedZone("BOT", -4*60*60)
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

func mexicoCity() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    19.4333,
		Longitude:   -99.0667,
		Elevation:   2216,
		Temperature: 10,
	}

	tz := time.FixedZone("CDT", -6*60*60)
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}

func tromso() (sampa.Location, time.Time, time.Time) {
	location := sampa.Location{
		Latitude:    69.1044,
		Longitude:   18.0594,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("CST", 1*60*60)
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	return location, start, limit
}
