package main

import (
	"math"
	"time"

	"github.com/hablullah/go-juliandays"
)

type Options struct {
	Pressure               float64
	SurfaceSlope           float64
	SurfaceAzimuthRotation float64
	DeltaT                 float64
}

type Data struct {
	JulianDay                    float64
	JulianCentury                float64
	JulianEphemerisDay           float64
	JulianEphemerisCentury       float64
	JulianEphemerisMillenium     float64
	EarthHeliocentricLongitude   float64
	EarthHeliocentricLatitude    float64
	EarthRadiusVector            float64
	GeocentricLongitude          float64
	GeocentricLatitude           float64
	NutationLongitude            float64
	NutationObliquity            float64
	EclipticTrueObliquity        float64
	AbberationCorrection         float64
	ApparentSunLongitude         float64
	MeanSiderealTime             float64
	ApparentSiderealTime         float64
	GeocentricSunRightAscension  float64
	GeocentricSunDeclination     float64
	ObserverLocalHourAngle       float64
	SunRightAscensionParallax    float64
	TopocentricSunRightAscension float64
	TopocentricSunDeclination    float64
	TopocentricLocalHourAngle    float64
	TopocentricZenithAngle       float64
	TopocentricAstroAzimuthAngle float64
	TopocentricAzimuthAngle      float64
	SurfaceIncidenceAngle        float64
}

func getData(dt time.Time, latitude, longitude, elevation, temperature float64, opts *Options) (Data, error) {
	// Set default value
	opts = setDefaultOptions(opts)

	// 1. Calculate the Julian and Julian ephemeris day century and millennium
	JD, err := juliandays.FromTime(dt)
	if err != nil {
		return Data{}, err
	}

	JC := getJulianCentury(JD)
	JDE := getJulianEphemerisDays(JD, opts.DeltaT)
	JCE := getJulianEphemerisCentury(JDE)
	JME := getJulianEphemerisMillenium(JCE)

	// 2. Calculate the Earth heliocentric longitude latitude and radius vector
	L := getEarthHeliocentricLongitude(JME)
	B := getEarthHeliocentricLatitude(JME)
	R := getEarthRadiusVector(JME)

	// 3. Calculate the geocentric longitude and latitude
	theta := getGeocentricLongitude(L)
	beta := getGeocentricLatitude(B)

	// 4. Calculate the nutation in longitude and obliquity
	deltaPsi, deltaEpsilon := getNutationLongitudeAndObliquity(JCE)

	// 5. Calculate the true obliquity of the ecliptic (in degrees)
	epsilon := getEclipticTrueObliquity(JME, deltaEpsilon)

	// 6. Calculate the aberration correction (in degrees)
	deltaTau := getAbberationCorrection(R)

	// 7. Calculate the apparent sun longitude (in degrees)
	lambda := getApparentSunLongitude(theta, deltaPsi, deltaTau)

	// 8. Calculate the apparent sidereal time at greenwich at any given time (in degrees)
	nu0 := getMeanSiderealTime(JD, JC)
	nu := getApparentSiderealTime(deltaPsi, epsilon, nu0)

	// 9. Calculate the geocentric sun right ascension (in degrees)
	alpha := getGeocentricSunRightAscension(beta, epsilon, lambda)

	// 10. Calculate the geocentric sun declination (in degrees)
	delta := getGeocentricSunDeclination(beta, epsilon, lambda)

	// 11. Calculate the observer local hour angle (in degrees)
	H := getObserverLocalHourAngle(longitude, nu, alpha)

	// 12. Calculate the topocentric sun right ascension α` and declination δ` (in degrees).
	// While on it also return the parallax in sun right ascension Δα (in degrees).
	deltaAlpha, alphaPrime, deltaPrime := getEquatorialSunCoordinates(latitude, elevation, R, alpha, delta, H)

	// 13. Calculate the topocentric local hour angle (in degrees)
	HPrime := getTopocentricLocalHourAngle(H, deltaAlpha)

	// 14. Calculate the topocentric zenith angle (in degrees)
	zenith := getTopocentricZenithAngle(latitude, opts.Pressure, temperature, deltaPrime, HPrime)

	// 15. Calculate the topocentric azimuth angle (in degrees)
	astroAzimuth, azimuth := getTopocentricAzimuthAngle(latitude, deltaPrime, HPrime)

	// 16. Calculate the incidence angle for a surface oriented in any direction (in degrees)
	incidenceAngle := getSurfaceIncidenceAngle(opts.SurfaceSlope, opts.SurfaceAzimuthRotation, zenith, astroAzimuth)

	return Data{
		JulianDay:                    JD,
		JulianCentury:                JC,
		JulianEphemerisDay:           JDE,
		JulianEphemerisCentury:       JCE,
		JulianEphemerisMillenium:     JME,
		EarthHeliocentricLongitude:   L,
		EarthHeliocentricLatitude:    B,
		EarthRadiusVector:            R,
		GeocentricLongitude:          theta,
		GeocentricLatitude:           beta,
		NutationLongitude:            deltaPsi,
		NutationObliquity:            deltaEpsilon,
		EclipticTrueObliquity:        epsilon,
		AbberationCorrection:         deltaTau,
		ApparentSunLongitude:         lambda,
		MeanSiderealTime:             nu0,
		ApparentSiderealTime:         nu,
		GeocentricSunRightAscension:  alpha,
		GeocentricSunDeclination:     delta,
		ObserverLocalHourAngle:       H,
		SunRightAscensionParallax:    deltaAlpha,
		TopocentricSunRightAscension: alphaPrime,
		TopocentricSunDeclination:    deltaPrime,
		TopocentricLocalHourAngle:    HPrime,
		TopocentricZenithAngle:       zenith,
		TopocentricAstroAzimuthAngle: astroAzimuth,
		TopocentricAzimuthAngle:      azimuth,
		SurfaceIncidenceAngle:        incidenceAngle,
	}, nil
}

func setDefaultOptions(opts *Options) *Options {
	if opts == nil {
		opts = &Options{}
	}

	if opts.Pressure == 0 {
		opts.Pressure = 101325
	}

	return opts
}

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

func getEarthHeliocentricLongitude(JME float64) float64 {
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
	L = radToDeg(L)
	L = limitDegrees(L)
	return L
}

func getEarthHeliocentricLatitude(JME float64) float64 {
	B0 := getEarthPeriodicTermSum("B0", JME)
	B1 := getEarthPeriodicTermSum("B1", JME)

	B := (B0 + B1*JME) / math.Pow10(8)
	B = radToDeg(B)
	return B
}

func getEarthRadiusVector(JME float64) float64 {
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
	return R
}

func getEarthPeriodicTermSum(key string, JME float64) float64 {
	var sum float64
	for _, term := range _EarthPeriodicTerms[key] {
		sum += term.A * math.Cos(term.B+term.C*JME)
	}
	return sum
}

func getGeocentricLongitude(L float64) float64 {
	theta := L + 180
	theta = limitDegrees(theta)
	return theta
}

func getGeocentricLatitude(B float64) float64 {
	return -B
}

func getNutationLongitudeAndObliquity(JCE float64) (float64, float64) {
	X := make([]float64, 5)

	// Calculate the mean elongation of the moon from the sun, X0 (in degrees)
	X[0] = polynomial(JCE, 297.85036, 445267.11148, -0.0019142, 1/189474)

	// Calculate the mean anomaly of the sun (Earth), X1 (in degrees)
	X[1] = polynomial(JCE, 357.52772, 35999.050340, -0.0001603, -1/300000)

	// Calculate the mean anomaly of the moon, X2 (in degrees)
	X[2] = polynomial(JCE, 134.96298, 477198.867398, 0.0086972, 1/56250)

	// Calculate the moon's argument of latitude, X3 (in degrees)
	X[3] = polynomial(JCE, 93.27191, 483202.017538, -0.0036825, 1/327270)

	// Calculate the longitude of the ascending node of the moon's
	// mean orbit on the ecliptic, measured from the mean equinox of the
	// date, X4 (in degrees)
	X[4] = polynomial(JCE, 125.04452, -1934.136261, 0.0020708, 1/450000)

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

func getAbberationCorrection(R float64) float64 {
	return -20.4898 / (3600 * R)
}

func getApparentSunLongitude(theta, deltaPsi, deltaTau float64) float64 {
	lambda := theta + deltaPsi + deltaTau
	return lambda
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
	return nu0 + deltaPsi*math.Cos(epsilonRad)
}

func getGeocentricSunRightAscension(beta, epsilon, lambda float64) float64 {
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

func getGeocentricSunDeclination(beta, epsilon, lambda float64) float64 {
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

func getEquatorialSunCoordinates(latitude, elevation, R, alpha, delta, H float64) (float64, float64, float64) {
	latitudeRad := degToRad(latitude)
	deltaRad := degToRad(delta)
	HRad := degToRad(H)

	// Calculate the equatorial horizontal parallax of the sun (in degrees)
	xi := 8.794 / (3600 * R)
	xiRad := degToRad(xi)

	// Calculate the term u (in radians)
	u := math.Atan(0.99664719 * math.Tan(latitudeRad))

	// Calculate the term x
	x := math.Cos(u) + (elevation/6378140)*math.Cos(latitudeRad)

	// Calculate the term y
	y := 0.99664719*math.Sin(u) + (elevation/6378140)*math.Sin(latitudeRad)

	// Calculate the parallax in the sun right ascension (in degrees)
	deltaAlpha := math.Atan2(
		-x*math.Sin(xiRad)*math.Sin(HRad),
		math.Cos(deltaRad)-x*math.Sin(xiRad)*math.Cos(HRad))
	deltaAlphaRad := deltaAlpha
	deltaAlpha = radToDeg(deltaAlpha)

	// Calculate the topocentric sun right ascension (in degrees)
	alphaPrime := alpha + deltaAlpha

	// Calculate the topocentric sun declination (in degrees)
	deltaPrime := math.Atan2(
		(math.Sin(deltaRad)-y*math.Sin(xiRad))*math.Cos(deltaAlphaRad),
		math.Cos(deltaRad)-x*math.Sin(xiRad)*math.Cos(HRad))
	deltaPrime = radToDeg(deltaPrime)

	return deltaAlpha, alphaPrime, deltaPrime
}

func getTopocentricLocalHourAngle(H, deltaAlpha float64) float64 {
	HPrime := H - deltaAlpha
	return HPrime
}

func getTopocentricZenithAngle(latitude, pressure, temperature, deltaPrime, HPrime float64) float64 {
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
	return zenith
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

func getSurfaceIncidenceAngle(surfaceSlope, surfaceAzimuthRotation, zenith, astroAzimuth float64) float64 {
	zenithRad := degToRad(zenith)
	surfaceSlopeRad := degToRad(surfaceSlope)
	incidenceAngle := math.Acos(math.Cos(zenithRad)*math.Cos(surfaceSlopeRad) +
		math.Sin(surfaceSlopeRad)*math.Sin(zenithRad)*math.Cos(degToRad(astroAzimuth-surfaceAzimuthRotation)))
	incidenceAngle = radToDeg(incidenceAngle)
	return incidenceAngle
}
