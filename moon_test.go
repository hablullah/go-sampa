package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetMoonEvents(t *testing.T) {
	testMoonEvents(t, testdata.Tromso)     // North Frigid
	testMoonEvents(t, testdata.London)     // North Temperate
	testMoonEvents(t, testdata.Jakarta)    // Torrid
	testMoonEvents(t, testdata.Wellington) // South Temperate
}

func testMoonEvents(t *testing.T, td testdata.TestData) {
	location := sampa.Location{
		Latitude:  td.Latitude,
		Longitude: td.Longitude,
	}

	for _, tt := range td.MoonEvents {
		dt, _ := time.ParseInLocation("2006-01-02", tt.Date, td.Timezone)

		e, err := sampa.GetMoonEvents(dt, location, nil, testdata.MoonEvents...)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertMoonEvents(t, td.Name, dt, tt, e)
	}
}
