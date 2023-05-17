package sampa

import (
	"math"
	"time"
)

type celestialPosition struct {
	ApparentSiderealTime     float64
	GeocentricRightAscension float64
	GeocentricDeclination    float64
}

type celestialArgs struct {
	date      time.Time
	location  Location
	deltaT    float64
	today     celestialPosition
	yesterday celestialPosition
	tomorrow  celestialPosition
}

func toCelestial(data interface{}) celestialPosition {
	switch v := data.(type) {
	case SunPosition:
		return celestialPosition{
			ApparentSiderealTime:     v.ApparentSiderealTime,
			GeocentricRightAscension: v.GeocentricRightAscension,
			GeocentricDeclination:    v.GeocentricDeclination,
		}
	case MoonPosition:
		return celestialPosition{
			ApparentSiderealTime:     v.ApparentSiderealTime,
			GeocentricRightAscension: v.GeocentricRightAscension,
			GeocentricDeclination:    v.GeocentricDeclination,
		}
	}
	return celestialPosition{}
}

func getCelestialTransit(args celestialArgs, approx float64) (float64, time.Time) {
	// Calculate at most 5 iterations
	for i := 1; i <= 5; i++ {
		// Calculate the sidereal time at Greenwich, in degrees, for the transit
		nu := args.today.ApparentSiderealTime + 360.985647*approx
		nu = limitValue(nu, 360)

		// Interpolate right ascension α` (in degrees)
		n := approx + args.deltaT/86400
		alphaPrime := interpolate(n,
			args.today.GeocentricRightAscension,
			args.yesterday.GeocentricRightAscension,
			args.tomorrow.GeocentricRightAscension)

		// Calculate the local hour angle for the sun transit
		// TODO: in Meeus HPrime = nu - loc.Longitude - alphaPrime
		// Here we use the one from Reda in his SPA paper.
		HPrime := nu + args.location.Longitude - alphaPrime
		HPrime = limit180Degrees(HPrime)

		// Calculate new approximate transit time in fraction of day
		newApprox := approx - HPrime/360

		// If the new approximation is similar with the previous, stop
		if fractionDiff(approx, newApprox) == 0 {
			break
		}

		approx = newApprox
	}

	// Make sure the transit occured within the day, not next nor previous day
	T := approx
	if T > 1 || T < 0 {
		return 0, time.Time{}
	}

	return T, dayFractionToTime(args.date, T)
}

func getCelestialAtElevation(args celestialArgs, transit, celestialElevation float64, isBeforeTransit bool) time.Time {
	// Calculate the approximate local hour angle
	H := getLocalHourAngle(celestialElevation, args.location.Latitude, args.today.GeocentricDeclination)
	if math.IsNaN(H) {
		return time.Time{}
	}

	// Calculate the approximate time in fraction of day
	approx := transit
	if isBeforeTransit {
		approx -= H / 360
	} else {
		approx += H / 360
	}

	// TODO: in original paper, the approximate time must be limited to 0 and 1.
	// However, there are case where the event occured in next or previous day,
	// so here we don't limit it.
	// approx = limitValue(approx, 1)

	// Calculate at most 5 iterations
	for i := 1; i <= 5; i++ {
		// Calculate the sidereal time at Greenwich, in degrees
		nu := args.today.ApparentSiderealTime + 360.985647*approx
		nu = limitValue(nu, 360)

		// Interpolate right ascension and declination (α` δ` in degrees)
		n := approx + args.deltaT/86400
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
		// Here we use the one from Reda in his SPA paper.
		HPrime := nu + args.location.Longitude - alphaPrime
		HPrime = limit180Degrees(HPrime)

		// Calculate the celestial altitude
		HPrimeRad := degToRad(HPrime)
		latitudeRad := degToRad(args.location.Latitude)
		deltaPrimeRad := degToRad(deltaPrime)
		h := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) +
			math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
		h = radToDeg(h)

		// Calculate the new approximate time in fraction of day
		newApprox := approx + ((h - celestialElevation) /
			(360 * math.Cos(deltaPrimeRad) * math.Cos(latitudeRad) * math.Sin(HPrimeRad)))

		// Make sure the new approximation is correctly ordered related to transit
		if (isBeforeTransit && newApprox > transit) || (!isBeforeTransit && newApprox < transit) {
			break
		}

		// If the new approximation is similar with the previous, stop
		if fractionDiff(approx, newApprox) == 0 {
			break
		}

		approx = newApprox
	}

	return dayFractionToTime(args.date, approx)
}
