package sampa_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func newTestLocation(l testdata.TestLocation) sampa.Location {
	return sampa.Location{
		Latitude:    l.Latitude,
		Longitude:   l.Longitude,
		Elevation:   l.Elevation,
		Temperature: l.Temperature,
		Pressure:    l.Pressure,
	}
}

func assertSunEvents(t *testing.T, name string, dt time.Time, expected testdata.TestTime, got sampa.SunEvents) {
	// Calculate diff
	diffSunrise := diffTestTime(dt, expected.Rise, got.Sunrise.DateTime, false)
	diffTransit := diffTestTime(dt, expected.Transit, got.Transit.DateTime, false)
	diffSunset := diffTestTime(dt, expected.Set, got.Sunset.DateTime, false)

	// Prepare log message
	strDate := dt.Format("2006-01-02")
	strSunrise := got.Sunrise.DateTime.Format("15:04:05")
	strTransit := got.Transit.DateTime.Format("15:04:05")
	strSunset := got.Sunset.DateTime.Format("15:04:05")

	msgFormat := "%s, %s => want %s got %s"
	sunriseMsg := fmt.Sprintf("Sunrise, "+msgFormat, name, strDate, expected.Rise, strSunrise)
	transitMsg := fmt.Sprintf("Transit, "+msgFormat, name, strDate, expected.Transit, strTransit)
	sunsetMsg := fmt.Sprintf("Sunset, "+msgFormat, name, strDate, expected.Set, strSunset)

	// For transit, diff only allowed up to 10 seconds
	assert.LessOrEqual(t, diffTransit, 10, transitMsg)

	// For sunrise and sunset, diff is allowed up to 60 seconds
	assert.LessOrEqual(t, diffSunrise, 60, sunriseMsg)
	assert.LessOrEqual(t, diffSunset, 60, sunsetMsg)
}

func assertMoonEvents(t *testing.T, name string, dt time.Time, expected testdata.TestTime, got sampa.MoonEvents) {
	// Calculate diff
	diffMoonrise := diffTestTime(dt, expected.Rise, got.Moonrise.DateTime, false)
	diffTransit := diffTestTime(dt, expected.Transit, got.Transit.DateTime, false)
	diffMoonset := diffTestTime(dt, expected.Set, got.Moonset.DateTime, false)

	// Prepare log message
	strDate := dt.Format("2006-01-02")
	strMoonrise := got.Moonrise.DateTime.Format("15:04:05")
	strTransit := got.Transit.DateTime.Format("15:04:05")
	strMoonset := got.Moonset.DateTime.Format("15:04:05")

	msgFormat := "%s, %s => want %s got %s"
	sunriseMsg := fmt.Sprintf("Moonrise, "+msgFormat, name, strDate, expected.Rise, strMoonrise)
	transitMsg := fmt.Sprintf("Transit, "+msgFormat, name, strDate, expected.Transit, strTransit)
	sunsetMsg := fmt.Sprintf("Moonset, "+msgFormat, name, strDate, expected.Set, strMoonset)

	// For transit, diff only allowed up to 10 seconds
	assert.LessOrEqual(t, diffTransit, 10, transitMsg)

	// For moonrise and moonset, diff is allowed up to 60 seconds
	assert.LessOrEqual(t, diffMoonrise, 60, sunriseMsg)
	assert.LessOrEqual(t, diffMoonset, 60, sunsetMsg)
}

func diffTestTime(dt time.Time, expected string, got time.Time, strictEmpty bool) int {
	if !strictEmpty && (expected == "" || got.IsZero()) {
		return -99
	}

	var expectedTime time.Time
	if expected != "" {
		tmp, _ := time.Parse("15:04:05", expected)
		expectedTime = time.Date(dt.Year(), dt.Month(), dt.Day(),
			tmp.Hour(), tmp.Minute(), tmp.Second(), 0,
			dt.Location())
	}

	diff := math.Round(math.Abs(expectedTime.Sub(got).Seconds()))
	return int(diff)
}
