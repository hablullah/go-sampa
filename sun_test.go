package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSunEvents(t *testing.T) {
	testSunEvents(t, testdata.Tromso)     // North Frigid
	testSunEvents(t, testdata.London)     // North Temperate
	testSunEvents(t, testdata.Jakarta)    // Torrid
	testSunEvents(t, testdata.Wellington) // South Temperate
}

func testSunEvents(t *testing.T, td testdata.TestData) {
	location := sampa.Location{
		Latitude:  td.Latitude,
		Longitude: td.Longitude,
	}

	for _, tt := range td.SunEvents {
		dt, _ := time.ParseInLocation("2006-01-02", tt.Date, td.Timezone)

		e, err := sampa.GetSunEvents(dt, location, nil, testdata.SunEvents...)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertSunEvents(t, td.Name, dt, tt, e)
	}
}
