package main

import (
	"time"

	"github.com/hablullah/go-sampa"
)

type SunEvents struct {
	sampa.SunEvents
	Date string
}

type MoonEvents struct {
	sampa.MoonEvents
	Date string
}

type Location struct {
	Name      string
	Timezone  string
	Latitude  float64
	Longitude float64
}

var testLocations = []Location{
	{ // Tromso (Norway) is representation for location in North Frigid area
		Name:      "Tromso",
		Timezone:  "CET",
		Latitude:  69.682778,
		Longitude: 18.942778,
	}, { // London (UK) is representation for location in North Temperate area
		Name:      "London",
		Timezone:  "Europe/London",
		Latitude:  51.507222,
		Longitude: -0.1275,
	}, { // Jakarta (Indonesia) is representation for location in Torrid area
		Name:      "Jakarta",
		Timezone:  "Asia/Jakarta",
		Latitude:  -6.175,
		Longitude: 106.825,
	}, { // Wellington (New Zealand) is representation for location in South Temperate area
		Name:      "Wellington",
		Timezone:  "Pacific/Auckland",
		Latitude:  -41.288889,
		Longitude: 174.777222,
	},
}

var customSunEvents = []sampa.CustomSunEvent{{
	Name:          "Dawn",
	BeforeTransit: true,
	Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
}, {
	Name:          "Dusk",
	BeforeTransit: false,
	Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
}}

var customMoonEvents = []sampa.CustomMoonEvent{{
	Name:          "Dawn",
	BeforeTransit: true,
	Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
}, {
	Name:          "Dusk",
	BeforeTransit: false,
	Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
}}

func getSunEvents(loc Location) []SunEvents {
	tz, _ := time.LoadLocation(loc.Timezone)
	location := sampa.Location{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}

	var events []SunEvents
	for dt := time.Date(2022, 1, 1, 0, 0, 0, 0, tz); dt.Year() == 2022; dt = dt.AddDate(0, 0, 1) {
		e, _ := sampa.GetSunEvents(dt, location, nil, customSunEvents...)
		events = append(events, SunEvents{
			SunEvents: e,
			Date:      dt.Format("2006-01-02"),
		})
	}

	return events
}

func getMoonEvents(loc Location) []MoonEvents {
	tz, _ := time.LoadLocation(loc.Timezone)
	location := sampa.Location{
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
	}

	var events []MoonEvents
	for dt := time.Date(2022, 1, 1, 0, 0, 0, 0, tz); dt.Year() == 2022; dt = dt.AddDate(0, 0, 1) {
		e, _ := sampa.GetMoonEvents(dt, location, nil, customMoonEvents...)
		events = append(events, MoonEvents{
			MoonEvents: e,
			Date:       dt.Format("2006-01-02"),
		})
	}

	return events
}
