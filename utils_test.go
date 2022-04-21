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
	diffSunrise := diffTestTime(dt, expected.Rise, got.Sunrise.DateTime)
	diffTransit := diffTestTime(dt, expected.Transit, got.Transit.DateTime)
	diffSunset := diffTestTime(dt, expected.Transit, got.Transit.DateTime)

	// Prepare log message
	strDate := dt.Format("2006-01-02")
	strSunrise := got.Sunrise.DateTime.Format("15:04:05")
	strTransit := got.Transit.DateTime.Format("15:04:05")
	strSunset := got.Sunset.DateTime.Format("15:04:05")

	msgFormat := "%s, %s => want %s got %s"
	sunriseMsg := fmt.Sprintf("Sunrise, "+msgFormat, name, strDate, expected.Rise, strSunrise)
	transitMsg := fmt.Sprintf("Transit, "+msgFormat, name, strDate, expected.Transit, strTransit)
	sunsetMsg := fmt.Sprintf("Sunset, "+msgFormat, name, strDate, expected.Set, strSunset)

	// For transit, diff only allowed up to 5 seconds
	assert.LessOrEqual(t, diffTransit, 5, transitMsg)

	// For sunrise and sunset, diff is allowed up to 60 seconds
	assert.LessOrEqual(t, diffSunrise, 60, sunriseMsg)
	assert.LessOrEqual(t, diffSunset, 60, sunsetMsg)
}

func diffTestTime(dt time.Time, expected string, got time.Time) int {
	if expected == "" || got.IsZero() {
		return -99
	}

	tmp, _ := time.Parse("15:04:05", expected)
	expectedTime := time.Date(dt.Year(), dt.Month(), dt.Day(),
		tmp.Hour(), tmp.Minute(), tmp.Second(), 0,
		dt.Location())

	diff := math.Round(math.Abs(expectedTime.Sub(got).Seconds()))
	return int(diff)
}
