package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSunEvents(t *testing.T) {
	testSunEvents(t, testdata.TromsoSun)     // North Frigid
	testSunEvents(t, testdata.LondonSun)     // North Temperate
	testSunEvents(t, testdata.JakartaSun)    // Torrid
	testSunEvents(t, testdata.WellingtonSun) // South Temperate
}

func testSunEvents(t *testing.T, td testdata.TestData) {
	for _, tt := range td.Times {
		location := newTestLocation(td.Location)
		dt, _ := time.ParseInLocation("2006-01-02", tt.Date, td.Z)

		e, err := sampa.GetSunEvents(dt, location, nil)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertSunEvents(t, td.Name, dt, tt, e)
	}
}
