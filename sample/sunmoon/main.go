package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	location := sampa.Location{
		Latitude:    -6.21138888888889,
		Longitude:   106.845277777778,
		Elevation:   0,
		Temperature: 10,
	}

	tz := time.FixedZone("WIB", 7*60*60)
	dt := time.Date(2022, 4, 15, 23, 3, 0, 0, tz)

	sun, _ := sampa.GetSunPosition(dt, location, nil)
	moon, _ := sampa.GetMoonPosition(dt, location, nil)

	sunZenith := degToRad(sun.TopocentricZenithAngle)
	sunAzimuth := degToRad(sun.TopocentricAzimuthAngle)
	moonZenith := degToRad(moon.TopocentricZenithAngle)
	moonAzimuth := degToRad(moon.TopocentricAzimuthAngle)
	E := radToDeg(math.Acos(math.Cos(sunZenith)*math.Cos(moonZenith) +
		math.Sin(sunZenith)*math.Sin(moonZenith)*math.Cos(sunAzimuth-moonAzimuth)))
	E = limitDegrees(E)

	Mm := degToRad(moon.MeanAnomaly)
	PA := 180 - E - 0.1468*((1-0.0549*math.Sin(Mm))/(1-0.0167*math.Sin(Mm)))*math.Sin(degToRad(E))
	K := 100 * ((1 + math.Cos(degToRad(PA))) / 2)

	fmt.Println("ELONGATION:", E)
	fmt.Println("PHASE ANGLE:", PA)
	fmt.Println("K:", K)
}

func radToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func limitDegrees(deg float64) float64 {
	return limitValues(360, deg)
}

func limitValues(max, val float64) float64 {
	f := val / max
	val = val - math.Trunc(f)*max
	if val < 0 {
		val = max + val
	}
	return val
}
