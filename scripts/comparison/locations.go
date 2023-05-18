package main

import (
	"math"
	"path/filepath"
	"time"
)

var dataDir = "./scripts/comparison/data/"

type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
	Timezone  *time.Location
	CsvSun    string
	CsvMoon   string
}

var (
	LordHoweIsland = Location{
		Name:      "Lord Howe Island, Australia",
		Latitude:  coordinate(false, 31, 33),
		Longitude: coordinate(true, 159, 5),
		Timezone:  mustLoadLocation("Australia/Lord_Howe"),
		CsvSun:    filepath.Join(dataDir, "sun-australia-lord-howe-island.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-australia-lord-howe-island.csv"),
	}

	Maputo = Location{
		Name:      "Maputo, Mozambique",
		Latitude:  coordinate(false, 25, 58),
		Longitude: coordinate(true, 32, 34),
		Timezone:  mustLoadLocation("Africa/Maputo"),
		CsvSun:    filepath.Join(dataDir, "sun-mozambique-maputo.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-mozambique-maputo.csv"),
	}

	Amsterdam = Location{
		Name:      "Amsterdam, Netherlands",
		Latitude:  coordinate(true, 52, 22),
		Longitude: coordinate(true, 4, 54),
		Timezone:  mustLoadLocation("CET"),
		CsvSun:    filepath.Join(dataDir, "sun-netherlands-amsterdam.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-netherlands-amsterdam.csv"),
	}

	Oslo = Location{
		Name:      "Oslo, Norway",
		Latitude:  coordinate(true, 59, 55),
		Longitude: coordinate(true, 10, 44),
		Timezone:  mustLoadLocation("CET"),
		CsvSun:    filepath.Join(dataDir, "sun-norway-oslo.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-norway-oslo.csv"),
	}

	Philipsburg = Location{
		Name:      "Philipsburg, Sint Maarten",
		Latitude:  coordinate(true, 18, 2),
		Longitude: coordinate(false, 63, 3),
		Timezone:  time.FixedZone("AST", -4*60*60),
		CsvSun:    filepath.Join(dataDir, "sun-sint-maarten-philipsburg.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-sint-maarten-philipsburg.csv"),
	}

	NewYork = Location{
		Name:      "New York, USA",
		Latitude:  coordinate(true, 40, 43),
		Longitude: coordinate(false, 74, 1),
		Timezone:  mustLoadLocation("America/New_York"),
		CsvSun:    filepath.Join(dataDir, "sun-us-new-york.csv"),
		CsvMoon:   filepath.Join(dataDir, "moon-us-new-york.csv"),
	}
)

func coordinate(positive bool, degrees, minutes float64) float64 {
	fl := math.Abs(degrees) + math.Abs(minutes)/60
	if !positive {
		fl = -fl
	}
	return fl
}

func mustLoadLocation(name string) *time.Location {
	tz, _ := time.LoadLocation(name)
	return tz
}
