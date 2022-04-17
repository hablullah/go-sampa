package sampa

import (
	"math"
	"time"
)

type celestialData struct {
	ApparentSiderealTime     float64
	GeocentricRightAscension float64
	GeocentricDeclination    float64
}

type celestialArgs struct {
	date      time.Time
	location  Location
	deltaT    float64
	today     celestialData
	yesterday celestialData
	tomorrow  celestialData
	tz        *time.Location
	tzOffset  int
}

func toCelestial(data interface{}) celestialData {
	switch v := data.(type) {
	case SunData:
		return celestialData{
			ApparentSiderealTime:     v.ApparentSiderealTime,
			GeocentricRightAscension: v.GeocentricRightAscension,
			GeocentricDeclination:    v.GeocentricDeclination,
		}
	case MoonData:
		return celestialData{
			ApparentSiderealTime:     v.ApparentSiderealTime,
			GeocentricRightAscension: v.GeocentricRightAscension,
			GeocentricDeclination:    v.GeocentricDeclination,
		}
	}
	return celestialData{}
}

func getCelestialTransit(args celestialArgs, approx float64) time.Time {
	// Calculate the sidereal time at Greenwich, in degrees, for the transit
	nu := args.today.ApparentSiderealTime + 360.985647*approx

	// Interpolate right ascension α` (in degrees)
	n := approx + args.deltaT/86400
	alphaPrime := interpolate(n,
		args.today.GeocentricRightAscension,
		args.yesterday.GeocentricRightAscension,
		args.tomorrow.GeocentricRightAscension)

	// Calculate the local hour angle for the sun transit
	// TODO: in Meeus HPrime = nu - loc.Longitude - alphaPrime
	HPrime := nu + args.location.Longitude - alphaPrime
	HPrime = limit180Degrees(HPrime)

	// Calculate transit time in fraction of day
	T := approx - (HPrime / 360)
	T = limitZeroOne(T + float64(args.tzOffset)/(24*60*60))

	return dayFractionToTime(args.date, T, args.tz)
}

func getCelestialAtElevation(args celestialArgs, approxTransit, celestialElevation float64, beforeTransit bool) time.Time {
	// Calculate the approximate local hour angle
	H := getLocalHourAngle(celestialElevation, args.location.Latitude, args.today.GeocentricDeclination)
	if math.IsNaN(H) {
		return time.Time{}
	}

	// Calculate the approximate time in fraction of day
	m := approxTransit
	if beforeTransit {
		m -= H / 360
	} else {
		m += H / 360
	}

	// Calculate the sidereal time at Greenwich, in degrees
	nu := args.today.ApparentSiderealTime + 360.985647*m

	// Interpolate right ascension and declination (α` δ` in degrees)
	n := m + args.deltaT/86400
	alphaPrime := interpolate(n,
		args.today.GeocentricRightAscension,
		args.yesterday.GeocentricRightAscension,
		args.tomorrow.GeocentricRightAscension)
	deltaPrime := interpolate(n,
		args.today.GeocentricDeclination,
		args.yesterday.GeocentricDeclination,
		args.tomorrow.GeocentricDeclination)

	// Calculate the local hour angle
	// TODO: in Meeus HPrime = nu - loc.Longitude - alphaPrime
	HPrime := nu + args.location.Longitude - alphaPrime
	HPrime = limit180Degrees(HPrime)

	// Calculate the celestial altitude
	HPrimeRad := degToRad(HPrime)
	latitudeRad := degToRad(args.location.Latitude)
	deltaPrimeRad := degToRad(deltaPrime)
	h := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) +
		math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
	h = radToDeg(h)

	// Calculate the time in fraction of day
	T := m + ((h - celestialElevation) /
		(360 * math.Cos(deltaPrimeRad) * math.Cos(latitudeRad) * math.Sin(HPrimeRad)))
	T = limitZeroOne(T + float64(args.tzOffset)/(24*60*60))

	return dayFractionToTime(args.date, T, args.tz)
}
