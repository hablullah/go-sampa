package main

import (
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	// Prepare location
	tz, _ := time.LoadLocation("Asia/Jakarta")
	jakarta := sampa.Location{Latitude: -6.14, Longitude: 106.81}

	// Fetch Sun position in Jakarta at 2023-05-20 18:15:00
	dt := time.Date(2023, 5, 20, 18, 15, 0, 0, tz)
	sunPosition, _ := sampa.GetSunPosition(dt, jakarta, nil)
	doSomething(sunPosition)

	// Fetch Moon position in Jakarta at 2023-05-20 20:15:00
	dt = time.Date(2023, 5, 20, 20, 15, 0, 0, tz)
	moonPosition, _ := sampa.GetMoonPosition(dt, jakarta, nil)
	doSomething(moonPosition)

	// Fetch Sun and Moon events in Jakarta at 2023-11-20
	dt = time.Date(2023, 11, 20, 0, 0, 0, 0, tz)
	sunEvents, _ := sampa.GetSunEvents(dt, jakarta, nil)
	doSomething(sunEvents)

	moonEvents, _ := sampa.GetMoonEvents(dt, jakarta, nil)
	doSomething(moonEvents)

	// Fetch Moon phases in Jakarta around 2023-03-05
	dt = time.Date(2023, 3, 5, 0, 0, 0, 0, tz)
	moonPhases := sampa.GetMoonPhases(dt, nil)
	doSomething(moonPhases)

	// To handle DST, just load the timezone properly.
	// For example, here we calculate Sun in Oslo that use DST.
	cet, _ := time.LoadLocation("CET")
	oslo := sampa.Location{Latitude: 59.91, Longitude: 10.74}
	dt = time.Date(2023, 5, 20, 18, 15, 0, 0, cet)
	sunPositionInOslo, _ := sampa.GetSunPosition(dt, oslo, nil)
	doSomething(sunPositionInOslo)
}

func doSomething(v any) {}
