package sampa

import (
	"math"
	"time"

	"github.com/hablullah/go-juliandays"
	fmt "github.com/hablullah/go-sampa/internal/fakefmt"
)

type MoonData struct {
	DateTime                     time.Time
	JulianDay                    float64
	JulianCentury                float64
	JulianEphemerisDay           float64
	JulianEphemerisCentury       float64
	JulianEphemerisMillenium     float64
	MeanAnomaly                  float64
	GeocentricLongitude          float64
	GeocentricLatitude           float64
	GeocentricDistance           float64
	HorizontalParallax           float64
	NutationLongitude            float64
	NutationObliquity            float64
	EclipticTrueObliquity        float64
	ApparentLongitude            float64
	MeanSiderealTime             float64
	ApparentSiderealTime         float64
	GeocentricRightAscension     float64
	GeocentricDeclination        float64
	ObserverLocalHourAngle       float64
	RightAscensionParallax       float64
	TopocentricRightAscension    float64
	TopocentricDeclination       float64
	TopocentricLocalHourAngle    float64
	TopocentricElevationAngle    float64
	TopocentricZenithAngle       float64
	TopocentricAstroAzimuthAngle float64
	TopocentricAzimuthAngle      float64
}

type CustomMoonEvent struct {
	Name          string
	BeforeTransit bool
	MoonElevation func(todayData MoonData) float64
}

type MoonEvents struct {
	Transit  MoonData
	Moonrise MoonData
	Moonset  MoonData
}

func GetMoonPosition(dt time.Time, loc Location, opts *Options) (MoonData, error) {
	// Make sure date time is not zero
	if dt.IsZero() {
		return MoonData{}, nil
	}

	// Set default value
	loc = setDefaultLocation(loc)
	opts = setDefaultOptions(opts)

	// 1. Calculate the Julian and Julian Ephemeris Day, Century, and Millennium
	JD, err := juliandays.FromTime(dt)
	if err != nil {
		return MoonData{}, err
	}

	JC := getJulianCentury(JD)
	JDE := getJulianEphemerisDays(JD, opts.DeltaT)
	JCE := getJulianEphemerisCentury(JDE)
	JME := getJulianEphemerisMillenium(JCE)
	fmt.Printf("JD: %f\n", JD)
	fmt.Printf("JC: %f\n", JC)

	// 2. Calculate the Moon Geocentric Longitude, Latitude, and Distance Between
	// the Centers of Earth and Moon
	MPrime := getMoonMeanAnomaly(JCE)
	fmt.Printf("M Prime: %f (%f rad)\n", MPrime, degToRad(MPrime))
	lambdaPrime, beta, dDelta := getMoonGeocentricPosition(JCE, MPrime)

	// 3. Calculate the Moon's Equatorial Horizontal Parallax (π in degrees)
	pi := math.Asin(6378.14 / dDelta)
	pi = radToDeg(pi)
	fmt.Printf("PI: %f\n", pi)

	// 4. Calculate the Nutation in Longitude and Obliquity (Δψ and Δε in degrees)
	deltaPsi, deltaEpsilon := getNutationLongitudeAndObliquity(JCE)
	fmt.Printf("DELTA PSI: %f\n", deltaPsi)
	fmt.Printf("DELTA EPSILON: %f\n", deltaEpsilon)

	// 5. Calculate the True Obliquity of the Ecliptic, ε (in degrees)
	epsilon := getEclipticTrueObliquity(JME, deltaEpsilon)
	fmt.Printf("EPSILON: %f (%f rad)\n", epsilon, degToRad(epsilon))

	// 6. Calculate the Apparent Moon Longitude, λ (in degrees)
	lambda := getApparentMoonLongitude(lambdaPrime, deltaPsi)
	fmt.Printf("LAMBDA: %f (%f rad)\n", lambda, degToRad(lambda))

	// 7. Calculate the Apparent Sidereal Time at Greenwich at any given time, ν (in degrees)
	nu0 := getMeanSiderealTime(JD, JC)
	nu := getApparentSiderealTime(deltaPsi, epsilon, nu0)
	nu = limitDegrees(nu)
	fmt.Printf("NU0: %f\n", nu0)
	fmt.Printf("NU: %f\n", nu)

	// 8. Calculate the Moon's Geocentric Right Ascension, α (in degrees)
	alpha := getGeocentricRightAscension(beta, epsilon, lambda)
	fmt.Printf("ALPHA: %f (%f rad)\n", alpha, degToRad(alpha))

	// 9. Calculate the Moon's Geocentric Declination, δ (in degrees)
	delta := getGeocentricDeclination(beta, epsilon, lambda)
	fmt.Printf("DELTA: %f (%f rad)\n", delta, degToRad(delta))

	// 10. Calculate the Observer Local Hour Angle, H (in degrees)
	H := getObserverLocalHourAngle(loc.Longitude, nu, alpha)

	// 11. Calculate the topocentric Moon right ascension α` and declination δ` (in degrees).
	// While on it also return the parallax in Moon right ascension Δα (in degrees).
	deltaAlpha, alphaPrime, deltaPrime := getEquatorialMoonCoordinates(loc.Latitude, loc.Elevation, pi, alpha, delta, H)

	// 12. Calculate the topocentric local hour angle (in degrees)
	HPrime := getTopocentricLocalHourAngle(H, deltaAlpha)

	// 13. Calculate the topocentric zenith angle (in degrees)
	zenith, moonElevation := getTopocentricZenithAngle(loc.Latitude, loc.Temperature, loc.Pressure, deltaPrime, HPrime)

	// 14. Calculate the topocentric azimuth angle (in degrees)
	astroAzimuth, azimuth := getTopocentricAzimuthAngle(loc.Latitude, deltaPrime, HPrime)

	return MoonData{
		DateTime:                     dt,
		JulianDay:                    JD,
		JulianCentury:                JC,
		JulianEphemerisDay:           JDE,
		JulianEphemerisCentury:       JCE,
		JulianEphemerisMillenium:     JME,
		MeanAnomaly:                  MPrime,
		GeocentricLongitude:          lambdaPrime,
		GeocentricLatitude:           beta,
		GeocentricDistance:           dDelta,
		HorizontalParallax:           pi,
		NutationLongitude:            deltaPsi,
		NutationObliquity:            deltaEpsilon,
		EclipticTrueObliquity:        epsilon,
		ApparentLongitude:            lambda,
		MeanSiderealTime:             nu0,
		ApparentSiderealTime:         nu,
		GeocentricRightAscension:     alpha,
		GeocentricDeclination:        delta,
		ObserverLocalHourAngle:       H,
		RightAscensionParallax:       deltaAlpha,
		TopocentricRightAscension:    alphaPrime,
		TopocentricDeclination:       deltaPrime,
		TopocentricLocalHourAngle:    HPrime,
		TopocentricElevationAngle:    moonElevation,
		TopocentricZenithAngle:       zenith,
		TopocentricAstroAzimuthAngle: astroAzimuth,
		TopocentricAzimuthAngle:      azimuth,
	}, nil
}

func GetMoonEvents(date time.Time, loc Location, opts *Options, customEvents ...CustomMoonEvent) (MoonEvents, error) {
	// Set default value
	loc = setDefaultLocation(loc)
	opts = setDefaultOptions(opts)

	// Change time to 0 LCT
	tz := date.Location()
	_, tzOffset := date.Zone()
	dt := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, tz)
	// dt = date

	// Set TT to zero
	ttZero := *opts
	ttZero.DeltaT = 0

	// Get data for current, previous and next day
	prevDate := dt.AddDate(0, 0, -1)
	yesterday, err := GetMoonPosition(prevDate, loc, &ttZero)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("yesterday moon error: %v", err)
	}

	nextDate := dt.AddDate(0, 0, 1)
	tomorrow, err := GetMoonPosition(nextDate, loc, &ttZero)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("tomorrow moon error: %v", err)
	}

	fmt.Printf("\nSTART HERE ========\n\n")
	today, err := GetMoonPosition(dt, loc, &ttZero)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("today moon error: %v", err)
	}

	// Prepare calculation args
	elevationAdjustment := 2.076 * math.Sqrt(loc.Elevation)
	h0 := 0.7275*today.HorizontalParallax - (34+elevationAdjustment)/60
	fmt.Printf("h0: %f (%f rad)\n", h0, degToRad(h0))

	args := celestialArgs{
		date:      dt,
		location:  loc,
		deltaT:    opts.DeltaT,
		today:     toCelestial(today),
		yesterday: toCelestial(yesterday),
		tomorrow:  toCelestial(tomorrow),
		tz:        tz,
		tzOffset:  tzOffset,
	}

	// Calculate the approximate moon transit time, st0, in fraction of day
	// Limit it to value between 0 and 1
	st0 := (today.GeocentricRightAscension - loc.Longitude - today.ApparentSiderealTime) / 360
	st0 = limitZeroOne(st0)
	fmt.Printf("ST0: %f\n", st0)

	// Calculate transit time
	mt := getCelestialTransit(args, st0)
	mtData, err := GetMoonPosition(mt, loc, opts)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("moon transit error: %v", err)
	}

	// Calculate moonrise
	mr := getCelestialAtElevation(args, st0, h0, true)
	mrData, err := GetMoonPosition(mr, loc, opts)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("moonrise error: %v", err)
	}

	// Calculate moonset
	ms := getCelestialAtElevation(args, st0, h0, false)
	msData, err := GetMoonPosition(ms, loc, opts)
	if err != nil {
		return MoonEvents{}, fmt.Errorf("moonset error: %v", err)
	}

	return MoonEvents{
		Transit:  mtData,
		Moonrise: mrData,
		Moonset:  msData,
	}, nil
}

func getMoonMeanAnomaly(JCE float64) float64 {
	MPrime := polynomial(JCE, 134.9633964, 477198.8675055, 0.0087414, 1/69699.0, -1/14712000.0)
	MPrime = limitDegrees(MPrime)
	return MPrime
}

func getMoonGeocentricPosition(JCE, MPrime float64) (float64, float64, float64) {
	// Calculate the Moon's Mean Longitude, L' (in degrees)
	LPrime := polynomial(JCE, 218.3164477, 481267.88123421, -0.0015786, 1/538841.0, -1/65194000.0)
	LPrime = limitDegrees(LPrime)
	fmt.Printf("L Prime: %f (%f rad)\n", LPrime, degToRad(LPrime))

	// Calculate the Mean Elongation of the Moon, D (in degrees)
	D := polynomial(JCE, 297.8501921, 445267.1114034, -0.0018819, 1/545868.0, -1/113065000.0)
	D = limitDegrees(D)
	fmt.Printf("D: %f (%f rad)\n", D, degToRad(D))

	// Calculate the Sun's Mean Anomaly, M (in degrees)
	M := polynomial(JCE, 357.5291092, 35999.0502909, -0.0001536, 1/24490000.0)
	M = limitDegrees(M)
	fmt.Printf("M: %f (%f rad)\n", M, degToRad(M))

	// Calculate the Moon's Argument of Latitude, F (in degrees)
	F := polynomial(JCE, 93.2720950, 483202.0175233, -0.0036539, -1/3526000.0, 1/863310000.0)
	F = limitDegrees(F)
	fmt.Printf("F: %f (%f rad)\n", F, degToRad(F))

	// Calculate term l (in 0.000001 degrees), r (in 0.001 kilometers), and
	// b (in 0.000001 degrees)
	E := 1.0 - 0.002516*JCE - 0.0000074*math.Pow(JCE, 2)
	fmt.Printf("E: %f (%f rad)\n", E, degToRad(E))
	l, r, b := getMoonPeriodicTermSum(E, D, M, MPrime, F)

	// Calculate term a
	a1 := limitDegrees(119.75 + 131.849*JCE)
	a2 := limitDegrees(53.09 + 479264.29*JCE)
	a3 := limitDegrees(313.45 + 481266.484*JCE)
	fmt.Printf("A1: %f (%f rad)\n", a1, degToRad(a1))
	fmt.Printf("A2: %f (%f rad)\n", a2, degToRad(a2))
	fmt.Printf("A3: %f (%f rad)\n", a3, degToRad(a3))

	// Calculate term Δl and Δb
	deltal := 3958*math.Sin(degToRad(a1)) +
		1962*math.Sin(degToRad(LPrime-F)) +
		318*math.Sin(degToRad(a2))
	fmt.Printf("deltaL: %f\n", deltal)

	deltab := -2235*math.Sin(degToRad(LPrime)) +
		382*math.Sin(degToRad(a3)) +
		175*math.Sin(degToRad(a1-F)) +
		175*math.Sin(degToRad(a1+F)) +
		127*math.Sin(degToRad(LPrime-MPrime)) -
		115*math.Sin(degToRad(LPrime+MPrime))
	fmt.Printf("deltaB: %f (%f rad)\n", deltab, degToRad(deltab))

	// Calculate the Moon's Longitude, λ' (in degrees), then limit it to 0 and 360
	lambdaPrime := LPrime + (l+deltal)/1_000_000
	lambdaPrime = limitDegrees(lambdaPrime)
	fmt.Printf("LAMBDA PRIME: %f (%f rad)\n", lambdaPrime, degToRad(lambdaPrime))

	// Calculate the Moon's Latitude, β (in degrees), then limit it to 0 and 360
	beta := (b + deltab) / 1_000_000
	beta = limitDegrees(beta)
	fmt.Printf("BETA: %f (%f rad)\n", beta, degToRad(beta))

	// Calculate the Moon's Distance From the Center of Earth, Δ (in kilometers)
	dDelta := 385000.56 + r/1000
	fmt.Printf("D DELTA: %f\n", dDelta)

	return lambdaPrime, beta, dDelta
}

func getMoonPeriodicTermSum(E, D, M, MPrime, F float64) (l, r, b float64) {
	l, r, b = 0, 0, 0

	for i := 0; i < 60; i++ {
		bTerm := _MoonPeriodicLatTerms[i]
		lrTerm := _MoonPeriodicDisLongTerms[i]
		bMultiplier := math.Pow(E, math.Abs(bTerm.m))
		lrMultiplier := math.Pow(E, math.Abs(lrTerm.m))
		bTrigArg := degToRad(bTerm.d*D + bTerm.m*M + bTerm.mP*MPrime + bTerm.f*F)
		lrTrigArg := degToRad(lrTerm.d*D + lrTerm.m*M + lrTerm.mP*MPrime + lrTerm.f*F)

		l += lrTerm.l * lrMultiplier * math.Sin(lrTrigArg)
		r += lrTerm.r * lrMultiplier * math.Cos(lrTrigArg)
		b += bTerm.b * bMultiplier * math.Sin(bTrigArg)
	}

	return l, r, b
}

func getApparentMoonLongitude(lambdaPrime, deltaPsi float64) float64 {
	lambda := lambdaPrime + deltaPsi
	return lambda
}

func getEquatorialMoonCoordinates(latitude, elevation, pi, alpha, delta, H float64) (float64, float64, float64) {
	latitudeRad := degToRad(latitude)
	deltaRad := degToRad(delta)
	piRad := degToRad(pi)
	HRad := degToRad(H)

	// Calculate the term u (in radians)
	u := math.Atan(0.99664719 * math.Tan(latitudeRad))

	// Calculate the term x
	x := math.Cos(u) + (elevation/6378140)*math.Cos(latitudeRad)

	// Calculate the term y
	y := 0.99664719*math.Sin(u) + (elevation/6378140)*math.Sin(latitudeRad)

	// Calculate the parallax in the moon right ascension (in degrees)
	deltaAlpha := math.Atan2(
		-x*math.Sin(piRad)*math.Sin(HRad),
		math.Cos(deltaRad)-x*math.Sin(piRad)*math.Cos(HRad))
	deltaAlphaRad := deltaAlpha
	deltaAlpha = radToDeg(deltaAlpha)

	// Calculate the topocentric moon right ascension (in degrees)
	alphaPrime := alpha + deltaAlpha

	// Calculate the topocentric moon declination (in degrees)
	deltaPrime := math.Atan2(
		(math.Sin(deltaRad)-y*math.Sin(piRad))*math.Cos(deltaAlphaRad),
		math.Cos(deltaRad)-x*math.Sin(piRad)*math.Cos(HRad))
	deltaPrime = radToDeg(deltaPrime)

	return deltaAlpha, alphaPrime, deltaPrime
}
