package main

import (
	"time"

	"github.com/hablullah/go-sampa"
)

func calculateSunEvents(location sampa.Location, tz *time.Location) ([]SunSchedule, error) {
	sunEvents := []sampa.CustomSunEvent{{
		Name:          "dawn",
		BeforeTransit: true,
		Elevation:     func(sampa.SunPosition) float64 { return -18 },
	}, {
		Name:          "dusk",
		BeforeTransit: false,
		Elevation:     func(sampa.SunPosition) float64 { return -18 },
	}}

	start := time.Date(year, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	nDays := int(limit.Sub(start).Hours() / 24)
	schedules := make([]SunSchedule, nDays)

	var idx int
	for dt := start; dt.Before(limit); dt = dt.AddDate(0, 0, 1) {
		e, err := sampa.GetSunEvents(dt, location, nil, sunEvents...)
		if err != nil {
			return nil, err
		}

		schedules[idx] = SunSchedule{
			Date:    dt.Format("2006-01-02"),
			Dawn:    e.Others["dawn"].DateTime,
			Sunrise: e.Sunrise.DateTime,
			Transit: e.Transit.DateTime,
			Sunset:  e.Sunset.DateTime,
			Dusk:    e.Others["dusk"].DateTime,
		}

		idx++
	}

	return schedules, nil
}
