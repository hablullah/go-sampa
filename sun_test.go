package sampa_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
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
		assertNil(t, err, fmt.Sprintf("Sun in %s at %s has error", td.Name, tt.Date))
		assertSunEvents(t, td.Name, dt, tt, e)
	}
}
