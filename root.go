package main

import (
	"fmt"
	"math"
	"time"

	"github.com/hablullah/go-juliandays"
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
	deltaT := 67

	latitudeRad := degreeToRad(latitude)
	surfaceSlopeRad := degreeToRad(surfaceSlope)

	// 1. CALCULATE THE JULIAN AND JULIAN EPHEMERIS DAY, CENTURY, AND MILLENNIUM
	// 1.1. Calculate the Julian Day
	JD, _ := juliandays.FromTime(dt)
	printf("JD: %f\n", JD)

	// 1.2. Calculate the Julian Ephemeris Day (JDE)
	JDE := JD + float64(deltaT)/86_400
	printf("JDE: %f\n", JDE)

	// 1.3. Calculate the Julian century (JC) and the Julian Ephemeris Century (JCE)
	// for the 2000 standard epoch
	JC := (JD - 2_451_545) / 36_525
	JCE := (JDE - 2_451_545) / 36_525
	printf("JC: %f\n", JC)
	printf("JCE: %f\n", JCE)

	// 1.4. Calculate the Julian Ephemeris Millennium (JME) for the 2000 standard epoch
	JME := JCE / 10
	printf("JME: %f\n", JME)

	// 2. CALCULATE THE EARTH HELIOCENTRIC LONGITUDE, LATITUDE, AND RADIUS VECTOR (L, B, R)
	// 2.1. Calculate the Earth heliocentric longitude
	L0 := getEarthPeriodicTermSum("L0", JME)
	L1 := getEarthPeriodicTermSum("L1", JME)
	L2 := getEarthPeriodicTermSum("L2", JME)
	L3 := getEarthPeriodicTermSum("L3", JME)
	L4 := getEarthPeriodicTermSum("L4", JME)
	L5 := getEarthPeriodicTermSum("L5", JME)

	L := (L0 + L1*JME +
		L2*math.Pow(JME, 2) +
		L3*math.Pow(JME, 3) +
		L4*math.Pow(JME, 4) +
		L5*math.Pow(JME, 5)) /
		math.Pow10(8)
	L = radToDegree(L)
	L = limitDegrees(L)

	printf("L0: %f\n", L0)
	printf("L1: %f\n", L1)
	printf("L2: %f\n", L2)
	printf("L3: %f\n", L3)
	printf("L4: %f\n", L4)
	printf("L5: %f\n", L5)
	printf("L: %f\n", L)

	// 2.2. Calculate the Earth heliocentric latitude
	B0 := getEarthPeriodicTermSum("B0", JME)
	B1 := getEarthPeriodicTermSum("B1", JME)

	B := (B0 + B1*JME) / math.Pow10(8)
	B = radToDegree(B)

	printf("B0: %f\n", B0)
	printf("B1: %f\n", B1)
	printf("B: %f\n", B)

	// 2.3. Calculate the Earth radius vector
	R0 := getEarthPeriodicTermSum("R0", JME)
	R1 := getEarthPeriodicTermSum("R1", JME)
	R2 := getEarthPeriodicTermSum("R2", JME)
	R3 := getEarthPeriodicTermSum("R3", JME)
	R4 := getEarthPeriodicTermSum("R4", JME)

	R := (R0 + R1*JME +
		R2*math.Pow(JME, 2) +
		R3*math.Pow(JME, 3) +
		R4*math.Pow(JME, 4)) /
		math.Pow10(8)

	printf("R0: %f\n", R0)
	printf("R1: %f\n", R1)
	printf("R2: %f\n", R2)
	printf("R3: %f\n", R3)
	printf("R4: %f\n", R4)
	printf("R: %f\n", R)

	// 3. CALCULATE THE GEOCENTRIC LONGITUDE AND LATITUDE
	// 3.1. Calculate the geocentric longitude
	theta := L + 180
	theta = limitDegrees(theta)
	printf("THETA: %f\n", theta)

	// 3.2. Calculate the geocentric latitude
	beta := -B
	betaRad := degreeToRad(beta)
	printf("BETA: %f\n", beta)

	// 4. CALCULATE THE NUTATION IN LONGITUDE AND OBLIQUITY
	X := make([]float64, 5)

	// 4.1. Calculate the mean elongation of the moon from the sun, X0 (in degrees)
	X[0] = 297.85036 +
		445267.11148*JCE -
		0.0019142*math.Pow(JCE, 2) +
		math.Pow(JCE, 3)/189474
	printf("X0: %f\n", X[0])

	// 4.2. Calculate the mean anomaly of the sun (Earth), X1 (in degrees)
	X[1] = 357.52772 +
		35999.050340*JCE -
		0.0001603*math.Pow(JCE, 2) -
		math.Pow(JCE, 3)/300000
	printf("X1: %f\n", X[1])

	// 4.3. Calculate the mean anomaly of the moon, X2 (in degrees)
	X[2] = 134.96298 +
		477198.867398*JCE +
		0.0086972*math.Pow(JCE, 2) +
		math.Pow(JCE, 3)/56250
	printf("X2: %f\n", X[2])

	// 4.4. Calculate the moon’s argument of latitude, X3 (in degrees)
	X[3] = 93.27191 +
		483202.017538*JCE -
		0.0036825*math.Pow(JCE, 2) +
		math.Pow(JCE, 3)/327270
	printf("X3: %f\n", X[3])

	// 4.5. Calculate the longitude of the ascending node of the moon's
	// mean orbit on the ecliptic, measured from the mean equinox of the
	// date, X4 (in degrees)
	X[4] = 125.04452 -
		1934.136261*JCE +
		0.0020708*math.Pow(JCE, 2) +
		math.Pow(JCE, 3)/450000
	printf("X4: %f\n", X[4])

	// 4.6. Calculate the nutation in longitude and obliquity (in degrees)
	sumDeltaPsi, sumDeltaEpsilon := getNutationPeriodicTermSum(X, JCE)
	deltaPsi := sumDeltaPsi / 36_000_000
	deltaEpsilon := sumDeltaEpsilon / 36_000_000

	printf("DELTA PSI: %f\n", deltaPsi)
	printf("DELTA EPSILON: %f\n", deltaEpsilon)

	// 5. CALCULATE THE TRUE OBLIQUITY OF THE ECLIPTIC (IN DEGREES)
	// 5.1. Calculate the mean obliquity of the ecliptic (in arc seconds)
	U := JME / 10
	epsilonZero := 84381.448 - 4680.93*U -
		1.55*math.Pow(U, 2) +
		1999.25*math.Pow(U, 3) -
		51.38*math.Pow(U, 4) -
		249.67*math.Pow(U, 5) -
		39.05*math.Pow(U, 6) +
		7.12*math.Pow(U, 7) +
		27.87*math.Pow(U, 8) +
		5.79*math.Pow(U, 9) +
		2.45*math.Pow(U, 10)
	printf("EPSILON ZERO: %f\n", epsilonZero)

	// 5.2. Calculate the true obliquity of the ecliptic (in degrees)
	epsilon := (epsilonZero / 3600) + deltaEpsilon
	epsilonRad := degreeToRad(epsilon)
	printf("EPSILON: %f\n", epsilon)

	// 6. CALCULATE THE ABERRATION CORRECTION (IN DEGREES)
	deltaTau := -20.4898 / (3600 * R)
	printf("DELTA TAU: %f\n", deltaTau)

	// 7. CALCULATE THE APPARENT SUN LONGITUDE (IN DEGREES)
	lambda := theta + deltaPsi + deltaTau
	lambdaRad := degreeToRad(lambda)
	printf("LAMBDA: %f\n", lambda)

	// 8. CALCULATE THE APPARENT SIDEREAL TIME AT GREENWICH AT ANY GIVEN TIME (IN DEGREES)
	// 8.1. Calculate the mean sidereal time at Greenwich (in degrees)
	nu0 := 280.46061837 +
		360.98564736629*(JD-2451545) +
		0.000387933*math.Pow(JC, 2) -
		math.Pow(JC, 3)/38710000
	nu0 = limitDegrees(nu0)
	printf("nu0: %f\n", nu0)

	// 8.2. Calculate the apparent sidereal time at Greenwich (in degrees)
	nu := nu0 + deltaPsi*math.Cos(epsilonRad)
	printf("nu: %f\n", nu)

	// 9. CALCULATE THE GEOCENTRIC SUN RIGHT ASCENSION (IN DEGREES)
	alpha := math.Atan2(
		math.Sin(lambdaRad)*math.Cos(epsilonRad)-math.Tan(betaRad)*math.Sin(epsilonRad),
		math.Cos(lambdaRad))
	alpha = radToDegree(alpha)
	alpha = limitDegrees(alpha)
	printf("alpha: %f\n", alpha)

	// 10. CALCULATE THE GEOCENTRIC SUN DECLINATION (IN DEGREES)
	delta := math.Asin(math.Sin(betaRad)*math.Cos(epsilonRad) +
		math.Cos(betaRad)*math.Sin(epsilonRad)*math.Sin(lambdaRad))
	deltaRad := delta
	delta = radToDegree(delta)
	printf("delta: %f\n", delta)

	// 11. CALCULATE THE OBSERVER LOCAL HOUR ANGLE (IN DEGREES)
	H := nu + longitude - alpha
	H = limitDegrees(H)
	HRad := degreeToRad(H)
	printf("H: %f\n", H)

	// 12. CALCULATE THE TOPOCENTRIC SUN RIGHT ASCENSION (IN DEGREES)
	// 12.1. Calculate the equatorial horizontal parallax of the sun (in degrees)
	xi := 8.794 / (3600 * R)
	xiRad := degreeToRad(xi)
	printf("xi: %f\n", xi)

	// 12.2. Calculate the term u (in radians)
	u := math.Atan(0.99664719 * math.Tan(latitudeRad))

	// 12.3. Calculate the term x
	x := math.Cos(u) + (elevation/6378140)*math.Cos(latitudeRad)

	// 12.4. Calculate the term y
	y := 0.99664719*math.Sin(u) + (elevation/6378140)*math.Sin(latitudeRad)

	// 12.5. Calculate the parallax in the sun right ascension (in degrees)
	deltaAlpha := math.Atan2(
		-x*math.Sin(xiRad)*math.Sin(HRad),
		math.Cos(deltaRad)-x*math.Sin(xiRad)*math.Cos(HRad))
	deltaAlphaRad := deltaAlpha
	deltaAlpha = radToDegree(deltaAlpha)
	printf("DELTA ALPHA: %f\n", deltaAlpha)

	// 12.6. Calculate the topocentric sun right ascension (in degrees)
	alphaPrime := alpha + deltaAlpha
	printf("ALPHA PRIME: %f\n", alphaPrime)

	// 12.7. Calculate the topocentric sun declination (in degrees)
	deltaPrime := math.Atan2(
		(math.Sin(deltaRad)-y*math.Sin(xiRad))*math.Cos(deltaAlphaRad),
		math.Cos(deltaRad)-x*math.Sin(xiRad)*math.Cos(HRad))
	deltaPrimeRad := deltaPrime
	deltaPrime = radToDegree(deltaPrime)
	printf("DELTA PRIME: %f\n", deltaPrime)

	// 13. CALCULATE THE TOPOCENTRIC LOCAL HOUR ANGLE (IN DEGREES)
	HPrime := H - deltaAlpha
	HPrimeRad := degreeToRad(HPrime)
	printf("H PRIME: %f\n", HPrime)

	// 14. CALCULATE THE TOPOCENTRIC ZENITH ANGLE (IN DEGREES)
	// 14.1. Calculate the topocentric elevation angle without atmospheric
	// refraction correction (in degrees)
	e0 := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) +
		math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
	e0 = radToDegree(e0)
	printf("e0: %f\n", e0)

	// 14.2. Calculate the atmospheric atmRefraction correction (in degrees)
	atmRefraction := (pressure / 1010) *
		(283 / (273 + temperature)) *
		(1.02 / (60 * math.Tan(degreeToRad(e0+10.3/(e0+5.11)))))
	printf("ATM REFRACTION: %f\n", atmRefraction)

	// 14.3. Calculate the topocentric elevation angle (in degrees)
	e := e0 + atmRefraction
	printf("e: %f\n", e)

	// 14.4. Calculate the topocentric zenith angle (in degrees)
	zenith := 90 - e
	zenithRad := degreeToRad(zenith)
	printf("ZENITH: %f\n", zenith)

	// 15. CALCULATE THE TOPOCENTRIC AZIMUTH ANGLE (IN DEGREES)
	// 15.1. Calculate the topocentric astronomers azimuth angle (in degrees)
	astroAzimuth := math.Atan2(math.Sin(HPrimeRad),
		math.Cos(HPrimeRad)*math.Sin(latitudeRad)-math.Tan(deltaPrimeRad)*math.Cos(latitudeRad))
	astroAzimuth = radToDegree(astroAzimuth)
	astroAzimuth = limitDegrees(astroAzimuth)
	printf("ASTRO AZIMUTH: %f\n", astroAzimuth)

	// 15.2. Calculate the topocentric azimuth angle for navigators and solar
	// radiation users (in degrees)
	azimuth := astroAzimuth + 180
	azimuth = limitDegrees(azimuth)
	printf("AZIMUTH: %f\n", azimuth)

	// 16. CALCULATE THE INCIDENCE ANGLE FOR A SURFACE ORIENTED IN ANY DIRECTION (IN DEGREES)
	incidenceAngle := math.Acos(math.Cos(zenithRad)*math.Cos(surfaceSlopeRad) +
		math.Sin(surfaceSlopeRad)*math.Sin(zenithRad)*math.Cos(degreeToRad(astroAzimuth-surfaceAzimuthRotation)))
	incidenceAngle = radToDegree(incidenceAngle)
	printf("INCIDENCE ANGLE: %f\n", incidenceAngle)

	// 17. CALCULATE EQUATION OF TIME (IN DEGREES)
	// 17.1. Calculate sun's mean longitude (in degrees)
	M := 280.4664567 +
		360007.6982779*JME +
		0.03032028*math.Pow(JME, 2) +
		math.Pow(JME, 3)/49931 -
		math.Pow(JME, 4)/15300 -
		math.Pow(JME, 5)/2000000

	// 17.2. Calculate equation of time (in degrees)
	EoT := M - 0.0057183 - alpha + deltaPsi*math.Cos(epsilonRad)
	EoT = limitValues(1440, 4*EoT)
	printf("EoT: %f\n", EoT)
}

func getEarthPeriodicTermSum(key string, JME float64) float64 {
	var sum float64
	for _, term := range _EarthPeriodicTerms[key] {
		sum += term.A * math.Cos(term.B+term.C*JME)
	}
	return sum
}

func getNutationPeriodicTermSum(X []float64, JCE float64) (float64, float64) {
	var sumDeltaPsi float64
	var sumDeltaEpsilon float64

	for _, term := range _NutationPeriodicTerms {
		var sumXY float64
		for j := 0; j <= 4; j++ {
			sumXY += X[j] * term.Y[j]
		}

		sumXY = degreeToRad(sumXY)
		sumDeltaPsi += (term.a + term.b*JCE) * math.Sin(sumXY)
		sumDeltaEpsilon += (term.c + term.d*JCE) * math.Cos(sumXY)
	}

	return sumDeltaPsi, sumDeltaEpsilon
}

func radToDegree(rad float64) float64 {
	return rad * 180 / math.Pi
}

func degreeToRad(deg float64) float64 {
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

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
