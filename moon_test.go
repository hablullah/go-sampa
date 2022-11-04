package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetMoonEvents(t *testing.T) {
	testMoonEvents(t, testdata.TromsoMoon)     // North Frigid
	testMoonEvents(t, testdata.LondonMoon)     // North Temperate
	testMoonEvents(t, testdata.JakartaMoon)    // Torrid
	testMoonEvents(t, testdata.WellingtonMoon) // South Temperate
}

func testMoonEvents(t *testing.T, td testdata.TestData) {
	for _, tt := range td.Times {
		location := newTestLocation(td.Location)
		dt, _ := time.ParseInLocation("2006-01-02", tt.Date, td.Z)

		e, err := sampa.GetMoonEvents(dt, location, nil)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertMoonEvents(t, td.Name, dt, tt, e)
	}
}
