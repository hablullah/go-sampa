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

var timeFormat = "2006-01-02 15:04:05 -0700"

func assertSunEvents(t *testing.T, name string, dt time.Time, expected testdata.CelestialEvent, got sampa.SunEvents) {
	// Calculate diff
	diffDawn := diffTestTime(dt, expected.Dawn, got.Others["Dawn"].DateTime)
	diffSunrise := diffTestTime(dt, expected.Rise, got.Sunrise.DateTime)
	diffTransit := diffTestTime(dt, expected.Transit, got.Transit.DateTime)
	diffSunset := diffTestTime(dt, expected.Set, got.Sunset.DateTime)
	diffDusk := diffTestTime(dt, expected.Dusk, got.Others["Dusk"].DateTime)

	// Prepare log message
	strDate := dt.Format("2006-01-02")
	strDawn := got.Others["Dawn"].DateTime.Format(timeFormat)
	strSunrise := got.Sunrise.DateTime.Format(timeFormat)
	strTransit := got.Transit.DateTime.Format(timeFormat)
	strSunset := got.Sunset.DateTime.Format(timeFormat)
	strDusk := got.Others["Dusk"].DateTime.Format(timeFormat)

	msgFormat := "%s, %s => want %s got %s"
	dawnMsg := fmt.Sprintf("Dawn, "+msgFormat, name, strDate, expected.Dawn, strDawn)
	sunriseMsg := fmt.Sprintf("Sunrise, "+msgFormat, name, strDate, expected.Rise, strSunrise)
	transitMsg := fmt.Sprintf("Transit, "+msgFormat, name, strDate, expected.Transit, strTransit)
	sunsetMsg := fmt.Sprintf("Sunset, "+msgFormat, name, strDate, expected.Set, strSunset)
	duskMsg := fmt.Sprintf("Dusk, "+msgFormat, name, strDate, expected.Dusk, strDusk)

	// Diff only allowed up to 10 seconds
	assert.LessOrEqual(t, diffDawn, 10, dawnMsg)
	assert.LessOrEqual(t, diffSunrise, 10, sunriseMsg)
	assert.LessOrEqual(t, diffTransit, 10, transitMsg)
	assert.LessOrEqual(t, diffSunset, 10, sunsetMsg)
	assert.LessOrEqual(t, diffDusk, 10, duskMsg)
}

func assertMoonEvents(t *testing.T, name string, dt time.Time, expected testdata.CelestialEvent, got sampa.MoonEvents) {
	// Calculate diff
	diffDawn := diffTestTime(dt, expected.Dawn, got.Others["Dawn"].DateTime)
	diffMoonrise := diffTestTime(dt, expected.Rise, got.Moonrise.DateTime)
	diffTransit := diffTestTime(dt, expected.Transit, got.Transit.DateTime)
	diffMoonset := diffTestTime(dt, expected.Set, got.Moonset.DateTime)
	diffDusk := diffTestTime(dt, expected.Dusk, got.Others["Dusk"].DateTime)

	// Prepare log message
	strDate := dt.Format("2006-01-02")
	strDawn := got.Others["Dawn"].DateTime.Format(timeFormat)
	strMoonrise := got.Moonrise.DateTime.Format("15:04:05")
	strTransit := got.Transit.DateTime.Format("15:04:05")
	strMoonset := got.Moonset.DateTime.Format("15:04:05")
	strDusk := got.Others["Dusk"].DateTime.Format(timeFormat)

	msgFormat := "%s, %s => want %s got %s"
	dawnMsg := fmt.Sprintf("Dawn, "+msgFormat, name, strDate, expected.Dawn, strDawn)
	moonriseMsg := fmt.Sprintf("Moonrise, "+msgFormat, name, strDate, expected.Rise, strMoonrise)
	transitMsg := fmt.Sprintf("Transit, "+msgFormat, name, strDate, expected.Transit, strTransit)
	moonsetMsg := fmt.Sprintf("Moonset, "+msgFormat, name, strDate, expected.Set, strMoonset)
	duskMsg := fmt.Sprintf("Dusk, "+msgFormat, name, strDate, expected.Dusk, strDusk)

	// Diff only allowed up to 10 seconds
	assert.LessOrEqual(t, diffDawn, 10, dawnMsg)
	assert.LessOrEqual(t, diffTransit, 10, transitMsg)
	assert.LessOrEqual(t, diffMoonrise, 10, moonriseMsg)
	assert.LessOrEqual(t, diffMoonset, 10, moonsetMsg)
	assert.LessOrEqual(t, diffDusk, 10, duskMsg)
}

func diffTestTime(dt time.Time, expected string, got time.Time) int {
	// Parse expected time
	var expectedTime time.Time
	if expected != "" {
		expectedTime, _ = time.ParseInLocation("2006-01-02 15:04:05 -0700", expected, got.Location())
	}

	diff := math.Round(math.Abs(expectedTime.Sub(got).Seconds()))
	return int(diff)
}
