package sampa

import (
	"math"
	"time"

	fmt "github.com/hablullah/go-sampa/internal/fakefmt"
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
	fmt.Printf("APPROX 0: %f\n", approx)

	for i := 1; i <= 10; i++ {
		fmt.Printf("ITERATION: %d\n", i)

		// Calculate the sidereal time at Greenwich, in degrees, for the transit
		nu := limitDegrees(args.today.ApparentSiderealTime + 360.985647*approx)
		fmt.Printf("\tnu: %f\n", nu)

		// Interpolate right ascension α` (in degrees)
		n := approx + args.deltaT/86400
		alphaPrime := interpolate(n,
			args.today.GeocentricRightAscension,
			args.yesterday.GeocentricRightAscension,
			args.tomorrow.GeocentricRightAscension)
		fmt.Printf("\tn: %f\n", n)
		fmt.Printf("\talpha prime: %f\n", alphaPrime)

		// Calculate the local hour angle for the sun transit
		// TODO: in Meeus HPrime = nu - loc.Longitude - alphaPrime
		HPrime := nu + args.location.Longitude - alphaPrime
		HPrime = limit180Degrees(HPrime)
		fmt.Printf("\tH Prime: %f\n", HPrime)

		// Calculate transit time in fraction of day
		newApprox := approx - HPrime/360
		if fractionDiff(approx, newApprox) == 0 {
			break
		}

		approx = newApprox
		fmt.Printf("\tAPPROX %d: %f\n", i, approx)
	}

	T := approx
	if T > 1 || T < 0 {
		return time.Time{}
	}

	// T += float64(args.tzOffset) / (24 * 60 * 60)
	// T -= math.Trunc(T)
	// fmt.Printf("FINAL T: %f\n", T)
	fmt.Printf("RESULT: %s\n", dayFractionToTime(args.date, T, args.tz))
	return dayFractionToTime(args.date, T, args.tz)
}

func getCelestialAtElevation(args celestialArgs, approxTransit, celestialElevation float64, beforeTransit bool) time.Time {
	// Calculate the approximate local hour angle
	H := getLocalHourAngle(celestialElevation, args.location.Latitude, args.today.GeocentricDeclination)
	fmt.Printf("H: %f\n", H)
	if math.IsNaN(H) {
		return time.Time{}
	}

	// Calculate the approximate time in fraction of day
	approx := approxTransit
	if beforeTransit {
		approx -= H / 360
	} else {
		approx += H / 360
	}
	approx = limitZeroOne(approx)
	fmt.Printf("APPROX 0: %f\n", approx)

	for i := 1; i <= 10; i++ {
		fmt.Printf("ITERATION: %d\n", i)

		// Calculate the sidereal time at Greenwich, in degrees
		nu := limitDegrees(args.today.ApparentSiderealTime + 360.985647*approx)
		fmt.Printf("\tNU: %f\n", nu)

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
		fmt.Printf("\tn: %f\n", n)
		fmt.Printf("\tALPHA PRIME: %f\n", alphaPrime)
		fmt.Printf("\tDELTA PRIME: %f\n", deltaPrime)

		// Calculate the local hour angle
		// TODO: in Meeus HPrime = nu - loc.Longitude - alphaPrime
		HPrime := nu + args.location.Longitude - alphaPrime
		HPrime = limit180Degrees(HPrime)
		fmt.Printf("\tH PRIME: %f\n", HPrime)

		// Calculate the celestial altitude
		HPrimeRad := degToRad(HPrime)
		latitudeRad := degToRad(args.location.Latitude)
		deltaPrimeRad := degToRad(deltaPrime)
		h := math.Asin(math.Sin(latitudeRad)*math.Sin(deltaPrimeRad) +
			math.Cos(latitudeRad)*math.Cos(deltaPrimeRad)*math.Cos(HPrimeRad))
		h = radToDeg(h)
		fmt.Printf("\th: %f\n", h)

		// Calculate the time in fraction of day
		newApprox := approx + ((h - celestialElevation) /
			(360 * math.Cos(deltaPrimeRad) * math.Cos(latitudeRad) * math.Sin(HPrimeRad)))
		if fractionDiff(approx, newApprox) == 0 {
			break
		}

		approx = newApprox
		fmt.Printf("\tAPPROX %d: %f\n", i, approx)
	}

	T := approx
	if T > 1 || T < 0 {
		return time.Time{}
	}

	// T += float64(args.tzOffset) / (24 * 60 * 60)
	// T -= math.Trunc(T)
	// fmt.Printf("FINAL T: %f\n", T)
	fmt.Printf("RESULT: %s\n", dayFractionToTime(args.date, T, args.tz))
	return dayFractionToTime(args.date, T, args.tz)
}
