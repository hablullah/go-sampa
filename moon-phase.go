package sampa

import (
	"math"
	"time"

	"github.com/hablullah/go-juliandays"
)

// MoonPhases is times when Moon reach its phases.
type MoonPhases struct {
	NewMoon      time.Time
	FirstQuarter time.Time
	FullMoon     time.Time
	LastQuarter  time.Time
	NextNewMoon  time.Time
	MonthLength  float64
}

// GetMoonPhases calculate the time for Moon phases around the specified date time.
func GetMoonPhases(dt time.Time, opts *Options) MoonPhases {
	// Set default value
	opts = setDefaultOptions(opts)

	// Calculate year fraction
	year := dt.Year()
	days := dt.YearDay()
	daysInYear := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC).YearDay()
	yearFraction := float64(year) + (float64(days) / float64(daysInYear))

	// Calculate base k value
	baseK := math.Floor((yearFraction - 2000) * 12.3685)

	// Calculate time for each phases
	tz := dt.Location()
	newMoon := getMoonPhaseTime(_NewMoon, baseK, tz, opts)
	firstQuarter := getMoonPhaseTime(_FirstQuarter, baseK, tz, opts)
	fullMoon := getMoonPhaseTime(_FullMoon, baseK, tz, opts)
	lastQuarter := getMoonPhaseTime(_LastQuarter, baseK, tz, opts)
	nextNewMoon := getMoonPhaseTime(_NextNewMoon, baseK, tz, opts)

	// Calculate month's length
	monthLength := nextNewMoon.Sub(newMoon).Hours() / 24

	return MoonPhases{
		NewMoon:      newMoon,
		FirstQuarter: firstQuarter,
		FullMoon:     fullMoon,
		LastQuarter:  lastQuarter,
		NextNewMoon:  nextNewMoon,
		MonthLength:  monthLength,
	}
}

func getMoonPhaseTime(phase _MoonPhase, baseK float64, tz *time.Location, opts *Options) time.Time {
	// Calculate k for the specified phase
	k := baseK + float64(phase)*0.25

	// Calculate term T
	T := k / 1236.85

	// Calculate Earth orbit eccentricity (E)
	E := 1 - 0.002516*T - 0.0000074*math.Pow(T, 2)

	// Calculate Sun's mean anomaly (M in degrees, convert to rad)
	M := 2.5534 + 29.1053567*k -
		0.0000014*math.Pow(T, 2) -
		0.00000011*math.Pow(T, 3)
	M = degToRad(limitDegrees(M))

	// Calculate Moon's mean anomaly (M' in degrees, convert to rad)
	MPrime := 201.5643 + 385.81693528*k +
		0.0107582*math.Pow(T, 2) +
		0.00001238*math.Pow(T, 3) -
		0.000000058*math.Pow(T, 4)
	MPrime = degToRad(limitDegrees(MPrime))

	// Calculate Moon's argument of latitude (F in degrees, convert to rad)
	F := 160.7108 + 390.67050284*k -
		0.0016118*math.Pow(T, 2) -
		0.00000227*math.Pow(T, 3) +
		0.000000011*math.Pow(T, 4)
	F = degToRad(limitDegrees(F))

	// Calculate longitude of ascending node of the lunar orbit (Ω in degrees, convert to rad)
	omega := 124.7746 - 1.56375588*k + 0.0020672*math.Pow(T, 2) + 0.00000215*math.Pow(T, 3)
	omega = degToRad(limitDegrees(omega))

	// Calculate planetary arguments (degrees, convert to rad)
	A1 := degToRad(299.77 + 0.107408*k - 0.009173*math.Pow(T, 2))
	A2 := degToRad(251.88 + 0.016321*k)
	A3 := degToRad(251.83 + 26.651886*k)
	A4 := degToRad(349.42 + 36.412478*k)
	A5 := degToRad(84.66 + 18.206239*k)
	A6 := degToRad(141.74 + 53.303771*k)
	A7 := degToRad(207.14 + 2.453732*k)
	A8 := degToRad(154.84 + 7.30686*k)
	A9 := degToRad(34.52 + 27.261239*k)
	A10 := degToRad(207.19 + 0.121824*k)
	A11 := degToRad(291.34 + 1.844379*k)
	A12 := degToRad(161.72 + 24.198154*k)
	A13 := degToRad(239.56 + 25.513099*k)
	A14 := degToRad(331.55 + 3.592518*k)

	// Calculate approximate JDE
	c := 2_451_550.09766 + 29.530588861*k
	JDE := polynomial(T, c, 0, 0.00015437, -0.00000015, 0.00000000073)

	// Calculate planetary arguments correction
	ACorrection := (0 +
		325*math.Sin(A1) +
		165*math.Sin(A2) +
		164*math.Sin(A3) +
		126*math.Sin(A4) +
		110*math.Sin(A5) +
		62*math.Sin(A6) +
		60*math.Sin(A7) +
		56*math.Sin(A8) +
		47*math.Sin(A9) +
		42*math.Sin(A10) +
		40*math.Sin(A11) +
		37*math.Sin(A12) +
		35*math.Sin(A13) +
		23*math.Sin(A14)) / 1000000

	// Calculate phase correction
	phaseCorrection := getPhaseCorrection(phase, E, M, MPrime, F, omega)

	// Calculate final JD
	JD := JDE + ACorrection + phaseCorrection - (opts.DeltaT / 86_400)
	return juliandays.ToTime(JD).In(tz)
}

func getPhaseCorrection(phase _MoonPhase, E, M, MPrime, F, omega float64) float64 {
	switch phase {
	case _NewMoon, _NextNewMoon:
		return (-40720*math.Sin(MPrime) +
			17241*E*math.Sin(M) +
			1608*math.Sin(2*MPrime) +
			1039*math.Sin(2*F) +
			739*E*math.Sin(MPrime-M) -
			514*E*math.Sin(MPrime+M) +
			208*E*E*math.Sin(2*M) -
			111*math.Sin(MPrime-2*F) -
			57*math.Sin(MPrime+2*F) +
			56*E*math.Sin(2*MPrime+M) -
			42*math.Sin(3*MPrime) +
			42*E*math.Sin(M+2*F) +
			38*E*math.Sin(M-2*F) -
			24*E*math.Sin(2*MPrime-M) -
			17*math.Sin(omega) -
			7*math.Sin(MPrime+2*M) +
			4*math.Sin(2*(MPrime-F)) +
			4*math.Sin(3*M) +
			3*math.Sin(MPrime+M-2*F) +
			3*math.Sin(2*(MPrime+F)) -
			3*math.Sin(MPrime+M+2*F) +
			3*math.Sin(MPrime-M+2*F) -
			2*math.Sin(MPrime-M-2*F) -
			2*math.Sin(3*MPrime+M) +
			2*math.Sin(4*MPrime)) / 100000
	case _FullMoon:
		return (-40614*math.Sin(MPrime) +
			17302*E*math.Sin(M) +
			1614*math.Sin(2*MPrime) +
			1043*math.Sin(2*F) +
			734*E*math.Sin(MPrime-M) -
			515*E*math.Sin(MPrime+M) +
			209*E*E*math.Sin(2*M) -
			111*math.Sin(MPrime-2*F) -
			57*math.Sin(MPrime+2*F) +
			56*E*math.Sin(2*MPrime+M) -
			42*math.Sin(3*MPrime) +
			42*E*math.Sin(M+2*F) +
			38*E*math.Sin(M-2*F) -
			24*E*math.Sin(2*MPrime-M) -
			17*math.Sin(omega) -
			7*math.Sin(MPrime+2*M) +
			4*math.Sin(2*(MPrime-F)) +
			4*math.Sin(3*M) +
			3*math.Sin(MPrime+M-2*F) +
			3*math.Sin(2*(MPrime+F)) -
			3*math.Sin(MPrime+M+2*F) +
			3*math.Sin(MPrime-M+2*F) -
			2*math.Sin(MPrime-M-2*F) -
			2*math.Sin(3*MPrime+M) +
			2*math.Sin(4*MPrime)) / 100000
	case _FirstQuarter, _LastQuarter:
		base := (-62801*math.Sin(MPrime) +
			17172*E*math.Sin(M) -
			1183*E*math.Sin(MPrime+M) +
			862*math.Sin(2*MPrime) +
			804*math.Sin(2*F) +
			454*E*math.Sin(MPrime-M) +
			204*E*E*math.Sin(2*M) -
			180*math.Sin(MPrime-2*F) -
			70*math.Sin(MPrime+2*F) -
			40*math.Sin(3*MPrime) -
			34*E*math.Sin(2*MPrime-M) +
			32*E*math.Sin(M+2*F) +
			32*E*math.Sin(M-2*F) -
			28*E*E*math.Sin(MPrime+2*M) +
			27*E*math.Sin(2*MPrime+M) -
			17*math.Sin(omega) -
			5*math.Sin(MPrime-M-2*F) +
			4*math.Sin(2*(MPrime+F)) -
			4*math.Sin(MPrime+M+2*F) +
			4*math.Sin(MPrime-2*M) +
			3*math.Sin(3*M) +
			3*math.Sin(MPrime+M-2*F) +
			2*math.Sin(2*(MPrime-F)) +
			2*math.Sin(MPrime-M+2*F) -
			2*math.Sin(3*MPrime+M)) / 100000
		W := (306 -
			38*E*math.Cos(M) +
			26*math.Cos(MPrime) -
			2*math.Cos(MPrime-M) +
			2*math.Cos(MPrime+M) +
			2*math.Cos(2*F)) / 100000
		if phase == _FirstQuarter {
			return base + W
		} else {
			return base - W
		}
	}

	return 0
}

type _MoonPhase uint8

const (
	_NewMoon _MoonPhase = iota
	_FirstQuarter
	_FullMoon
	_LastQuarter
	_NextNewMoon
)
