package sampa

import (
	"fmt"
	"math"
	"time"

	"github.com/hablullah/go-juliandays"
)

type SunData struct {
	DateTime                     time.Time
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
	SurfaceIncidenceAngle        float64
}

type CustomSunEvent struct {
	Name          string
	BeforeTransit bool
	SunElevation  func(todayData SunData) float64
}

type SunEvents struct {
	Transit SunData
	Sunrise SunData
	Sunset  SunData
	Others  map[string]SunData
}

func GetSunPosition(dt time.Time, loc Location, opts *Options) (SunData, error) {
	// Make sure date time is not zero
	if dt.IsZero() {
		return SunData{}, nil
	}

	// Set default value
	loc = setDefaultLocation(loc)
	opts = setDefaultOptions(opts)

	// 1. Calculate the Julian and Julian ephemeris day century and millennium
	JD, err := juliandays.FromTime(dt)
	if err != nil {
		return SunData{}, err
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
	theta := getSunGeocentricLongitude(L)
	beta := getSunGeocentricLatitude(B)

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
	alpha := getGeocentricRightAscension(beta, epsilon, lambda)

	// 10. Calculate the geocentric sun declination (in degrees)
	delta := getGeocentricDeclination(beta, epsilon, lambda)

	// 11. Calculate the observer local hour angle (in degrees)
	H := getObserverLocalHourAngle(loc.Longitude, nu, alpha)

	// 12. Calculate the topocentric sun right ascension α` and declination δ` (in degrees).
	// While on it also return the parallax in sun right ascension Δα (in degrees).
	deltaAlpha, alphaPrime, deltaPrime := getEquatorialSunCoordinates(loc.Latitude, loc.Elevation, R, alpha, delta, H)

	// 13. Calculate the topocentric local hour angle (in degrees)
	HPrime := getTopocentricLocalHourAngle(H, deltaAlpha)

	// 14. Calculate the topocentric zenith angle (in degrees)
	zenith, sunElevation := getTopocentricZenithAngle(loc.Latitude, loc.Temperature, loc.Pressure, deltaPrime, HPrime)

	// 15. Calculate the topocentric azimuth angle (in degrees)
	astroAzimuth, azimuth := getTopocentricAzimuthAngle(loc.Latitude, deltaPrime, HPrime)

	// 16. Calculate the incidence angle for a surface oriented in any direction (in degrees)
	incidenceAngle := getSurfaceIncidenceAngle(opts.SurfaceSlope, opts.SurfaceAzimuthRotation, zenith, astroAzimuth)

	return SunData{
		DateTime:                     dt,
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
		TopocentricElevationAngle:    sunElevation,
		TopocentricZenithAngle:       zenith,
		TopocentricAstroAzimuthAngle: astroAzimuth,
		TopocentricAzimuthAngle:      azimuth,
		SurfaceIncidenceAngle:        incidenceAngle,
	}, nil
}

func GetSunEvents(date time.Time, loc Location, opts *Options, customEvents ...CustomSunEvent) (SunEvents, error) {
	// Set default value
	loc = setDefaultLocation(loc)
	opts = setDefaultOptions(opts)

	// Change time to 0 LCT
	dt := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	// Set TT to zero
	ttZero := *opts
	ttZero.DeltaT = 0

	// Get data for current, previous and next day
	today, err := GetSunPosition(dt, loc, &ttZero)
	if err != nil {
		return SunEvents{}, fmt.Errorf("today sun error: %v", err)
	}

	prevDate := dt.AddDate(0, 0, -1)
	yesterday, err := GetSunPosition(prevDate, loc, &ttZero)
	if err != nil {
		return SunEvents{}, fmt.Errorf("yesterday sun error: %v", err)
	}

	nextDate := dt.AddDate(0, 0, 1)
	tomorrow, err := GetSunPosition(nextDate, loc, &ttZero)
	if err != nil {
		return SunEvents{}, fmt.Errorf("tomorrow sun error: %v", err)
	}

	// Prepare calculation args
	elevationAdjustment := 2.076 * math.Sqrt(loc.Elevation)
	h0 := -(50 + elevationAdjustment) / 60.0

	args := celestialArgs{
		date:      dt,
		location:  loc,
		deltaT:    opts.DeltaT,
		today:     toCelestial(today),
		yesterday: toCelestial(yesterday),
		tomorrow:  toCelestial(tomorrow),
	}

	// Calculate the approximate sun transit time, st0, in fraction of day
	// Limit it to value between 0 and 1
	st0 := (today.GeocentricRightAscension - loc.Longitude - today.ApparentSiderealTime) / 360
	st0 = limitZeroOne(st0)

	// Calculate transit time
	st := getCelestialTransit(args, st0)
	stData, err := GetSunPosition(st, loc, opts)
	if err != nil {
		return SunEvents{}, fmt.Errorf("sun transit error: %v", err)
	}

	// Calculate sunrise and sunset
	sr := getCelestialAtElevation(args, st0, h0, true)
	srData, err := GetSunPosition(sr, loc, opts)
	if err != nil {
		return SunEvents{}, fmt.Errorf("sunrise error: %v", err)
	}

	ss := getCelestialAtElevation(args, st0, h0, false)
	ssData, err := GetSunPosition(ss, loc, opts)
	if err != nil {
		return SunEvents{}, fmt.Errorf("sunset error: %v", err)
	}

	// Calculate other events
	otherEvents := map[string]SunData{}
	for _, e := range customEvents {
		et := getCelestialAtElevation(args, st0, e.SunElevation(today), e.BeforeTransit)
		eData, err := GetSunPosition(et, loc, opts)
		if err != nil {
			return SunEvents{}, fmt.Errorf("event \"%s\" error: %v", e.Name, err)
		}
		otherEvents[e.Name] = eData
	}

	return SunEvents{
		Transit: stData,
		Sunrise: srData,
		Sunset:  ssData,
		Others:  otherEvents,
	}, nil
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

func getSunGeocentricLongitude(L float64) float64 {
	theta := L + 180
	theta = limitDegrees(theta)
	return theta
}

func getSunGeocentricLatitude(B float64) float64 {
	return -B
}

func getAbberationCorrection(R float64) float64 {
	return -20.4898 / (3600 * R)
}

func getApparentSunLongitude(theta, deltaPsi, deltaTau float64) float64 {
	lambda := theta + deltaPsi + deltaTau
	return lambda
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

func getSurfaceIncidenceAngle(surfaceSlope, surfaceAzimuthRotation, zenith, astroAzimuth float64) float64 {
	zenithRad := degToRad(zenith)
	surfaceSlopeRad := degToRad(surfaceSlope)
	incidenceAngle := math.Acos(math.Cos(zenithRad)*math.Cos(surfaceSlopeRad) +
		math.Sin(surfaceSlopeRad)*math.Sin(zenithRad)*math.Cos(degToRad(astroAzimuth-surfaceAzimuthRotation)))
	incidenceAngle = radToDeg(incidenceAngle)
	return incidenceAngle
}
