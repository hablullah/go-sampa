package sampa

import (
	"math"
	"time"
)

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

func limitFullCircle(val float64) float64 {
	if int(math.Round(math.Abs(val)/360)) >= 1 {
		val = limitDegrees(val)
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

func dayFractionToTime(dt time.Time, f float64) time.Time {
	fs := int(math.Round(f * 24 * 60 * 60))
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, fs, 0, dt.Location())
}

func fractionDiff(f1, f2 float64) int {
	f1s := int(math.Round(f1 * 24 * 60 * 60))
	f2s := int(math.Round(f2 * 24 * 60 * 60))
	return f2s - f1s
}
