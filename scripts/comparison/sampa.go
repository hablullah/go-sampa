package main

import (
	"math"
	"time"

	"github.com/hablullah/go-sampa"
)

func calculateSunEvents(location sampa.Location, tz *time.Location) ([]SunData, error) {
	sunEvents := []sampa.CustomSunEvent{{
		Name:          "dawn18",
		BeforeTransit: true,
		Elevation:     func(sampa.SunPosition) float64 { return -18 },
	}, {
		Name:          "dusk18",
		BeforeTransit: false,
		Elevation:     func(sampa.SunPosition) float64 { return -18 },
	}, {
		Name:          "dawn12",
		BeforeTransit: true,
		Elevation:     func(sampa.SunPosition) float64 { return -12 },
	}, {
		Name:          "dusk12",
		BeforeTransit: false,
		Elevation:     func(sampa.SunPosition) float64 { return -12 },
	}, {
		Name:          "dawn6",
		BeforeTransit: true,
		Elevation:     func(sampa.SunPosition) float64 { return -6 },
	}, {
		Name:          "dusk6",
		BeforeTransit: false,
		Elevation:     func(sampa.SunPosition) float64 { return -6 },
	}}

	start := time.Date(year, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	nDays := int(limit.Sub(start).Hours() / 24)
	schedules := make([]SunData, nDays)

	var idx int
	for dt := start; dt.Before(limit); dt = dt.AddDate(0, 0, 1) {
		e, err := sampa.GetSunEvents(dt, location, nil, sunEvents...)
		if err != nil {
			return nil, err
		}

		schedules[idx] = SunData{
			Date:    dt.Format("2006-01-02"),
			Dawn18:  e.Others["dawn18"].DateTime,
			Dawn12:  e.Others["dawn12"].DateTime,
			Dawn6:   e.Others["dawn6"].DateTime,
			Sunrise: e.Sunrise.DateTime,
			Transit: e.Transit.DateTime,
			Sunset:  e.Sunset.DateTime,
			Dusk6:   e.Others["dusk6"].DateTime,
			Dusk12:  e.Others["dusk12"].DateTime,
			Dusk18:  e.Others["dusk18"].DateTime,

			SunriseAzimuth:  round(e.Sunrise.TopocentricAzimuthAngle, 1),
			SunsetAzimuth:   round(e.Sunset.TopocentricAzimuthAngle, 1),
			TransitAltitude: round(e.Transit.TopocentricElevationAngle, 1),
		}

		idx++
	}

	return schedules, nil
}

func calculateMoonEvents(location sampa.Location, tz *time.Location) ([]MoonSchedule, error) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, tz)
	limit := start.AddDate(1, 0, 0)
	nDays := int(limit.Sub(start).Hours() / 24)
	schedules := make([]MoonSchedule, nDays)

	var idx int
	for dt := start; dt.Before(limit); dt = dt.AddDate(0, 0, 1) {
		e, err := sampa.GetMoonEvents(dt, location, nil)
		if err != nil {
			return nil, err
		}

		schedules[idx] = MoonSchedule{
			Date:     dt.Format("2006-01-02"),
			Moonrise: e.Moonrise.DateTime,
			Transit:  e.Transit.DateTime,
			Moonset:  e.Moonset.DateTime,

			MoonriseAzimuth: round(e.Moonrise.TopocentricAzimuthAngle, 1),
			MoonsetAzimuth:  round(e.Moonset.TopocentricAzimuthAngle, 1),
			TransitAltitude: round(e.Transit.TopocentricElevationAngle, 1),
			Illumination:    round(e.Transit.PercentIlluminated*100, 1),
		}

		idx++
	}

	return schedules, nil
}

func round(val float64, decimalPlace int) float64 {
	rounder := math.Pow10(decimalPlace)
	return math.Round(val*rounder) / rounder
}
