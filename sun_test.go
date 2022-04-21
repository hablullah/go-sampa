package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetSunEvents(t *testing.T) {
	// Check tropical area
	testSunEvents(t, testdata.SunJakarta)

	// Check sub tropical area
	testSunEvents(t, testdata.SunLosAngeles)

	// Check area in extreme latitude
	testSunEvents(t, testdata.SunTromso)
}

func testSunEvents(t *testing.T, td testdata.TestData) {
	for _, tt := range td.Times {
		location := newTestLocation(td.Location)
		dt, _ := time.ParseInLocation("02/01/2006", tt.Date, td.Z)

		e, err := sampa.GetSunEvents(dt, location, nil)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertSunEvents(t, td.Name, dt, tt, e)
	}
}
