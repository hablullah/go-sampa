package testdata

import "time"
import "github.com/hablullah/go-sampa"

type CelestialEvent struct {
	Date    string
	Dawn    string
	Rise    string
	Transit string
	Set     string
	Dusk    string
}

type TestData struct {
	Name       string
	Latitude   float64
	Longitude  float64
	Timezone   *time.Location
	SunEvents  []CelestialEvent
	MoonEvents []CelestialEvent
}

var SunEvents = []sampa.CustomSunEvent{{
	Name:          "Dawn",
	BeforeTransit: true,
	Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
}, {
	Name:          "Dusk",
	BeforeTransit: false,
	Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
}}

var MoonEvents = []sampa.CustomMoonEvent{{
	Name:          "Dawn",
	BeforeTransit: true,
	Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
}, {
	Name:          "Dusk",
	BeforeTransit: false,
	Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
}}
