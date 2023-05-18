package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/hablullah/go-sampa"
)

func compareSunSchedules(location Location) error {
	// Prepare expected schedules from CSV file
	expectedSchedules, err := parseSunCSV(location.CsvSun, location.Timezone)
	if err != nil {
		return err
	}

	// Calculate schedules using SAMPA
	calculatedSchedules, err := calculateSunEvents(sampa.Location{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}, location.Timezone)
	if err != nil {
		return err
	}

	// Compare schedules
	var dawnDiffs []int
	var sunriseDiffs []int
	var transitDiffs []int
	var sunsetDiffs []int
	var duskDiffs []int

	dawnTitle := location.Name + ", dawn"
	sunriseTitle := location.Name + ", sunrise"
	transitTitle := location.Name + ", transit"
	sunsetTitle := location.Name + ", sunset"
	duskTitle := location.Name + ", dusk"

	for i, exp := range expectedSchedules {
		res := calculatedSchedules[i]
		dawnDiff := compareTime(dawnTitle, res.Date, exp.Dawn, res.Dawn)
		sunriseDiff := compareTime(sunriseTitle, res.Date, exp.Sunrise, res.Sunrise)
		transitDiff := compareTime(transitTitle, res.Date, exp.Transit, res.Transit)
		sunsetDiff := compareTime(sunsetTitle, res.Date, exp.Sunset, res.Sunset)
		duskDiff := compareTime(duskTitle, res.Date, exp.Dusk, res.Dusk)

		dawnDiffs = append(dawnDiffs, dawnDiff)
		sunriseDiffs = append(sunriseDiffs, sunriseDiff)
		transitDiffs = append(transitDiffs, transitDiff)
		sunsetDiffs = append(sunsetDiffs, sunsetDiff)
		duskDiffs = append(duskDiffs, duskDiff)
	}

	// Print diff stat
	fmt.Println("Sun event in", location.Name)
	printDiff("- Dawn   ", dawnDiffs)
	printDiff("- Sunrise", sunriseDiffs)
	printDiff("- Transit", transitDiffs)
	printDiff("- Sunset ", sunsetDiffs)
	printDiff("- Dusk   ", duskDiffs)
	return nil
}

func compareTime(logTitle, date string, expected, result time.Time) int {
	const dtFormat = "2006-01-02 15:04:05"

	// If expected and result is empty, everything is ok
	if expected.IsZero() && result.IsZero() {
		return 0
	}

	// If expected is empty but result exist, it's still ok
	if expected.IsZero() && !result.IsZero() {
		if enableLog {
			log.Printf("%s in %q, expect empty but got %q",
				logTitle, date, result.Format(dtFormat))
		}
		return 0
	}

	// If expected is not empty but result is, it's questionable
	if !expected.IsZero() && result.IsZero() {
		if enableLog {
			log.Printf("%s in %q, expect %q but got empty",
				logTitle, date, expected.Format(dtFormat))
		}
		return 0
	}

	// Calculate diff
	diff := int(math.Round(math.Abs(result.Sub(expected).Seconds())))
	if diff > maxDiff && enableLog {
		log.Printf("%s in %q, expect %q got %q with %d seconds diff",
			logTitle, date,
			expected.Format(dtFormat),
			result.Format(dtFormat),
			diff)
	}

	return diff
}

func diffStat(diffs []int) (max int, mode int, avg float64) {
	nDiff := len(diffs)
	if nDiff == 0 {
		return
	}

	// Calculate various stat
	var sum int
	diffCount := make(map[int]int)
	for _, diff := range diffs {
		sum += diff
		diffCount[diff]++
		if diff > max {
			max = diff
		}
	}

	// Calculate average
	avg = float64(sum) / float64(nDiff)

	// Calculate mode
	var modeCount int
	for diff, count := range diffCount {
		if count > modeCount {
			mode = diff
			modeCount = count
		}
	}

	return
}

func printDiff(title string, diffs []int) {
	max, mode, avg := diffStat(diffs)
	fmt.Printf("%s: diff max = %2ds, mode = %ds, avg = %.2fs\n", title, max, mode, avg)
}
