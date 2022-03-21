package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Variables
	tz := time.FixedZone("LST", -7*60*60)
	dt := time.Date(2003, 10, 17, 12, 30, 30, 0, tz)
	latitude := 39.742476
	longitude := -105.1786
	pressure := 820.0
	elevation := 1830.14
	temperature := 11.0
	surfaceSlope := 30.0
	surfaceAzimuthRotation := -10.0
	deltaT := 67.0

	// Get initial sun data
	opts := &Options{
		Pressure:               pressure,
		SurfaceSlope:           surfaceSlope,
		SurfaceAzimuthRotation: surfaceAzimuthRotation,
		DeltaT:                 deltaT,
	}

	// data, _ := getData(dt, latitude, longitude, elevation, temperature, opts)
	// bt, _ := json.MarshalIndent(&data, "", "\t")
	// fmt.Println(string(bt))

	// Get transit time
	// getEarthHeliocentricLongitude(0.00379111)
	// return
	getSunTransitTime(dt, latitude, longitude, elevation, temperature, opts)
}

func getEquationOfTime(JME, deltaPsi, epsilon, alpha float64) float64 {
	epsilonRad := degToRad(epsilon)

	// Calculate sun's mean longitude (in degrees)
	M := polynomial(JME, 280.4664567, 360007.6982779, 0.03032028, 1/49931, -1/15300, -1/2000000)

	// Calculate equation of time (in degrees)
	EoT := M - 0.0057183 - alpha + deltaPsi*math.Cos(epsilonRad)
	EoT = limitValues(1440, 4*EoT)
	return EoT
}

func getSunTransitTime(dt time.Time, latitude, longitude, elevation, temperature float64, opts *Options) float64 {
	// Set default value
	opts = setDefaultOptions(opts)

	// Change time to 0 UT
	loc := dt.Location()
	dt = dt.UTC()
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())

	// Set TT to zero
	ttZero := *opts
	ttZero.DeltaT = 0

	// Get data for current, previous and next day
	nextDt := dt.AddDate(0, 0, 1)
	prevDt := dt.AddDate(0, 0, -1)
	today, _ := getData(dt, latitude, longitude, elevation, temperature, opts)
	tomorrow, _ := getData(nextDt, latitude, longitude, elevation, temperature, opts)
	yesterday, _ := getData(prevDt, latitude, longitude, elevation, temperature, opts)

	// printf("DATE: %s\n", dt)
	// printf("JD: %f\n", today.JulianDay)

	// Calculate the approximate sun transit time, m0, in fraction of day
	// Limit it to value between 0 and 1
	m0 := (today.GeocentricSunRightAscension - longitude - today.ApparentSiderealTime) / 360
	m0 = limitZeroOne(m0)
	// printf("beta: %f\n", today.GeocentricLatitude)
	// printf("epsilon: %f\n", today.EclipticTrueObliquity)
	// printf("deltaT: %.10f\n", opts.DeltaT)
	// printf("JC: %.10f\n", today.JulianCentury)
	// printf("JDE: %.10f\n", today.JulianEphemerisDay)
	// printf("JCE: %.10f\n", today.JulianEphemerisCentury)
	// printf("JME: %.10f\n", today.JulianEphemerisMillenium)
	// printf("L: %f\n", today.EarthHeliocentricLongitude)
	// printf("theta: %f\n", today.GeocentricLongitude)
	// printf("deltaPsi: %f\n", today.NutationLongitude)
	// printf("deltaTau: %f\n", today.AbberationCorrection)
	// printf("lambda: %f\n", today.ApparentSunLongitude)
	// printf("alpha: %f, %f, %f\n",
	// 	yesterday.GeocentricSunRightAscension,
	// 	today.GeocentricSunRightAscension,
	// 	tomorrow.GeocentricSunRightAscension)
	// printf("nu: %f\n", today.ApparentSiderealTime)
	// printf("m0: %f\n", m0)

	// Calculate the sidereal time at Greenwich, in degrees, for the sun transit
	nu := today.ApparentSiderealTime + 360.985647*m0
	// printf("nu_rts: %f\n", nu)

	// Calculate the terms n
	n := m0 + opts.DeltaT/86400

	// Calculate α` and δ` (in degrees)
	a := today.GeocentricSunRightAscension - yesterday.GeocentricSunRightAscension
	b := tomorrow.GeocentricSunRightAscension - today.GeocentricSunRightAscension
	c := b - a

	aPrime := today.GeocentricSunDeclination - yesterday.GeocentricSunDeclination
	bPrime := tomorrow.GeocentricSunDeclination - today.GeocentricSunDeclination
	cPrime := bPrime - aPrime

	a = limitAbsZeroOne(2, a)
	b = limitAbsZeroOne(2, b)
	aPrime = limitAbsZeroOne(2, aPrime)
	bPrime = limitAbsZeroOne(2, bPrime)

	alphaPrime := today.GeocentricSunRightAscension + (n*(a+b+c*n))/2
	deltaPrime := today.GeocentricSunDeclination + (n*(aPrime+bPrime+cPrime*n))/2
	// printf("alpha_prime: %f\n", alphaPrime)
	// printf("delta_prime: %f\n", deltaPrime)

	// Calculate the local hour angle for the sun transit
	HPrime := nu + longitude - alphaPrime
	// printf("h_prime: %f\n", HPrime)
	HPrime = limit180Degrees(HPrime)
	// printf("h_prime_limit: %f\n", HPrime)

	// Calculate sun transit time
	T := m0 - (HPrime / 360)
	// printf("t: %f\n", T)

	transitTime := dt.Add(time.Second * time.Duration(T*24*60*60)).In(loc)
	printf("transit: %s\n", transitTime)

	// Calculate sun altitude at transit
	HPrimeRad := degToRad(HPrime)
	latitudeRad := degToRad(latitude)
	deltaPrimeRad := degToRad(deltaPrime)
	h := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) + math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
	h = radToDeg(h)
	printf("altitude: %f\n", h)

	return 0
}

func getHourAngle(latitude, sunAltitude, geocentricSunDeclination float64) float64 {
	latitudeRad := degToRad(latitude)
	altitudeRad := degToRad(sunAltitude)
	deltaRad := degToRad(geocentricSunDeclination)

	ha := math.Acos((math.Sin(altitudeRad) - math.Sin(latitudeRad)*math.Sin(deltaRad)) /
		(math.Cos(latitudeRad) * math.Cos(deltaRad)))
	ha = radToDeg(ha)
	return ha
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

func limitZeroOne(val float64) float64 {
	val = val - math.Trunc(val)
	if val < 0 {
		val += 1
	}
	return val
}

func limitAbsZeroOne(abs, val float64) float64 {
	if math.Abs(val) >= abs {
		return limitZeroOne(val)
	}
	return val
}

func limit180Degrees(val float64) float64 {
	f := val / 360
	val = val - math.Trunc(f)*360
	if val < -180 {
		val += 360
	} else if val > 180 {
		val -= 360
	}
	return val
}

func polynomial(x float64, values ...float64) float64 {
	var sum float64
	for i, value := range values {
		sum += value * math.Pow(x, float64(i))
	}
	return sum
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
