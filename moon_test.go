package sampa_test

import (
	"testing"
	"time"

	"github.com/hablullah/go-sampa"
	"github.com/hablullah/go-sampa/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGetMoonEvents(t *testing.T) {
	// Check tropical area
	testMoonEvents(t, testdata.MoonJakarta)

	// Check sub tropical area
	testMoonEvents(t, testdata.MoonLosAngeles)

	// Check area in extreme latitude
	// testMoonEvents(t, testdata.MoonTromso)
}

func testMoonEvents(t *testing.T, td testdata.TestData) {
	for _, tt := range td.Times {
		location := newTestLocation(td.Location)
		dt, _ := time.ParseInLocation("02/01/2006", tt.Date, td.Z)

		e, err := sampa.GetMoonEvents(dt, location, nil)
		assert.Nil(t, err, e, "%s %s error", td.Name, tt.Date)
		assertMoonEvents(t, td.Name, dt, tt, e)
	}
}
