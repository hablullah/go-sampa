package sampa

import (
	"math"
)

func getJulianEphemerisDays(JD, deltaT float64) float64 {
	return JD + deltaT/86_400
}

func getJulianCentury(JD float64) float64 {
	return (JD - 2_451_545) / 36_525
}

func getJulianEphemerisCentury(JDE float64) float64 {
	return (JDE - 2_451_545) / 36_525
}

func getJulianEphemerisMillenium(JCE float64) float64 {
	return JCE / 10
}

func getNutationLongitudeAndObliquity(JCE float64) (float64, float64) {
	X := make([]float64, 5)

	// Calculate the mean elongation of the moon from the sun, X0 (in degrees)
	X[0] = polynomial(JCE, 297.85036, 445267.11148, -0.0019142, 1/189474.0)

	// Calculate the mean anomaly of the sun (Earth), X1 (in degrees)
	X[1] = polynomial(JCE, 357.52772, 35999.050340, -0.0001603, -1/300000.0)

	// Calculate the mean anomaly of the moon, X2 (in degrees)
	X[2] = polynomial(JCE, 134.96298, 477198.867398, 0.0086972, 1/56250.0)

	// Calculate the moon's argument of latitude, X3 (in degrees)
	X[3] = polynomial(JCE, 93.27191, 483202.017538, -0.0036825, 1/327270.0)

	// Calculate the longitude of the ascending node of the moon's
	// mean orbit on the ecliptic, measured from the mean equinox of the
	// date, X4 (in degrees)
	X[4] = polynomial(JCE, 125.04452, -1934.136261, 0.0020708, 1.0/450000)

	// Calculate nutation periodic sum
	var sumDeltaPsi float64
	var sumDeltaEpsilon float64

	for _, term := range _NutationPeriodicTerms {
		var sumXY float64
		for j := 0; j <= 4; j++ {
			sumXY += X[j] * term.Y[j]
		}

		sumXY = degToRad(sumXY)
		sumDeltaPsi += (term.a + term.b*JCE) * math.Sin(sumXY)
		sumDeltaEpsilon += (term.c + term.d*JCE) * math.Cos(sumXY)
	}

	// Calculate the nutation in longitude and obliquity (in degrees)
	deltaPsi := sumDeltaPsi / 36_000_000
	deltaEpsilon := sumDeltaEpsilon / 36_000_000
	return deltaPsi, deltaEpsilon
}

func getEclipticTrueObliquity(JME, deltaEpsilon float64) float64 {
	// Calculate the mean obliquity of the ecliptic (in arc seconds)
	epsilonZero := polynomial(JME/10, 84381.448, -4680.93, -1.55, 1999.25,
		-51.38, -249.67, -39.05, 7.12, 27.87, 5.79, 2.45)

	// Calculate the true obliquity of the ecliptic (in degrees)
	epsilon := (epsilonZero / 3600) + deltaEpsilon
	return epsilon
}

func getMeanSiderealTime(JD, JC float64) float64 {
	nu0 := 280.46061837 +
		360.98564736629*(JD-2451545) +
		0.000387933*math.Pow(JC, 2) -
		math.Pow(JC, 3)/38710000
	nu0 = limitDegrees(nu0)
	return nu0
}

func getApparentSiderealTime(deltaPsi, epsilon, nu0 float64) float64 {
	epsilonRad := degToRad(epsilon)
	nu := nu0 + deltaPsi*math.Cos(epsilonRad)
	return nu
}

func getGeocentricRightAscension(beta, epsilon, lambda float64) float64 {
	betaRad := degToRad(beta)
	lambdaRad := degToRad(lambda)
	epsilonRad := degToRad(epsilon)

	alpha := math.Atan2(
		math.Sin(lambdaRad)*math.Cos(epsilonRad)-math.Tan(betaRad)*math.Sin(epsilonRad),
		math.Cos(lambdaRad))
	alpha = radToDeg(alpha)
	alpha = limitDegrees(alpha)
	return alpha
}

func getGeocentricDeclination(beta, epsilon, lambda float64) float64 {
	betaRad := degToRad(beta)
	lambdaRad := degToRad(lambda)
	epsilonRad := degToRad(epsilon)

	delta := math.Asin(math.Sin(betaRad)*math.Cos(epsilonRad) +
		math.Cos(betaRad)*math.Sin(epsilonRad)*math.Sin(lambdaRad))
	return radToDeg(delta)
}

func getObserverLocalHourAngle(longitude, nu, alpha float64) float64 {
	H := nu + longitude - alpha
	H = limitDegrees(H)
	return H
}

func getTopocentricLocalHourAngle(H, deltaAlpha float64) float64 {
	HPrime := H - deltaAlpha
	return HPrime
}

func getTopocentricZenithAngle(latitude, temperature, pressure, deltaPrime, HPrime float64) (float64, float64) {
	HPrimeRad := degToRad(HPrime)
	deltaPrimeRad := degToRad(deltaPrime)

	// Calculate the topocentric elevation angle without atmospheric
	// refraction correction (in degrees)
	latitudeRad := degToRad(latitude)
	e0 := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) +
		math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
	e0 = radToDeg(e0)

	// Calculate the atmospheric atmRefraction correction (in degrees)
	atmRefraction := (pressure / 1010) *
		(283 / (273 + temperature)) *
		(1.02 / (60 * math.Tan(degToRad(e0+10.3/(e0+5.11)))))

	// Calculate the topocentric elevation angle (in degrees)
	e := e0 + atmRefraction

	// Calculate the topocentric zenith angle (in degrees)
	zenith := 90 - e
	return zenith, e
}

func getTopocentricAzimuthAngle(latitude, deltaPrime, HPrime float64) (float64, float64) {
	HPrimeRad := degToRad(HPrime)
	deltaPrimeRad := degToRad(deltaPrime)

	// Calculate the topocentric astronomers azimuth angle (in degrees)
	latitudeRad := degToRad(latitude)
	astroAzimuth := math.Atan2(math.Sin(HPrimeRad),
		math.Cos(HPrimeRad)*math.Sin(latitudeRad)-math.Tan(deltaPrimeRad)*math.Cos(latitudeRad))
	astroAzimuth = radToDeg(astroAzimuth)
	astroAzimuth = limitDegrees(astroAzimuth)

	// Calculate the topocentric azimuth angle for navigators and solar radiation users (in degrees)
	azimuth := astroAzimuth + 180
	azimuth = limitDegrees(azimuth)
	return astroAzimuth, azimuth
}

func getLocalHourAngle(elevation, latitude, sunDeclination float64) float64 {
	deltaRad := degToRad(sunDeclination)
	latitudeRad := degToRad(latitude)
	elevationRad := degToRad(elevation)

	H := math.Acos((math.Sin(elevationRad) - math.Sin(latitudeRad)*math.Sin(deltaRad)) /
		(math.Cos(latitudeRad) * math.Cos(deltaRad)))
	H = radToDeg(H)
	H = limit180Degrees(H)
	return H
}

func interpolate(factor, d, dMin, dPlus float64) float64 {
	n := factor
	a := d - dMin
	b := dPlus - d

	// TODO: in SPA paper a & b values must be limited
	// between 0 and 1, i.e. a = limitAbsZeroOne(2, a)
	a = limitFullCircle(a)
	b = limitFullCircle(b)
	c := b - a
	d1 := d + n*(a+b+c*n)/2
	return limitDegrees(d1)
}
