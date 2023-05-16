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

func limitValue(value, limit float64) float64 {
	f := value / limit
	f = math.Abs(f - math.Trunc(f))

	switch {
	case value > 0:
		return limit * f
	case value < 0:
		return limit - limit*f
	default:
		return value
	}
}

func limit180Degrees(val float64) float64 {
	val = limitValue(val, 360)
	if val <= -180 {
		val += 360
	} else if val >= 180 {
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
	dt = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	seconds := time.Duration(math.Round(f*24*60*60)) * time.Second
	return dt.Add(seconds)
}

func fractionDiff(f1, f2 float64) int {
	f1s := int(math.Round(f1 * 24 * 60 * 60))
	f2s := int(math.Round(f2 * 24 * 60 * 60))
	return f2s - f1s
}
