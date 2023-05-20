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
	var dawn18Diffs []int
	var dawn12Diffs []int
	var dawn6Diffs []int
	var sunriseDiffs []int
	var transitDiffs []int
	var sunsetDiffs []int
	var dusk6Diffs []int
	var dusk12Diffs []int
	var dusk18Diffs []int
	var sunriseAzimuthDiffs []float64
	var sunsetAzimuthDiffs []float64
	var transitAltitudeDiffs []float64

	dawn18Title := location.Name + ", dawn18"
	dawn12Title := location.Name + ", dawn12"
	dawn6Title := location.Name + ", dawn6"
	sunriseTitle := location.Name + ", sunrise"
	transitTitle := location.Name + ", transit"
	sunsetTitle := location.Name + ", sunset"
	dusk6Title := location.Name + ", dusk6"
	dusk12Title := location.Name + ", dusk12"
	dusk18Title := location.Name + ", dusk18"

	for i, exp := range expectedSchedules {
		res := calculatedSchedules[i]
		dawn18Diff := compareTime(dawn18Title, res.Date, exp.Dawn18, res.Dawn18)
		dawn12Diff := compareTime(dawn12Title, res.Date, exp.Dawn12, res.Dawn12)
		dawn6Diff := compareTime(dawn6Title, res.Date, exp.Dawn6, res.Dawn6)
		sunriseDiff := compareTime(sunriseTitle, res.Date, exp.Sunrise, res.Sunrise)
		transitDiff := compareTime(transitTitle, res.Date, exp.Transit, res.Transit)
		sunsetDiff := compareTime(sunsetTitle, res.Date, exp.Sunset, res.Sunset)
		dusk6Diff := compareTime(dusk6Title, res.Date, exp.Dusk6, res.Dusk6)
		dusk12Diff := compareTime(dusk12Title, res.Date, exp.Dusk12, res.Dusk12)
		dusk18Diff := compareTime(dusk18Title, res.Date, exp.Dusk18, res.Dusk18)
		sunriseAzimuthDiff := math.Abs(res.SunriseAzimuth - exp.SunriseAzimuth)
		sunsetAzimuthDiff := math.Abs(res.SunsetAzimuth - exp.SunsetAzimuth)
		transitAltitudeDiff := math.Abs(res.TransitAltitude - exp.TransitAltitude)

		dawn18Diffs = append(dawn18Diffs, dawn18Diff)
		dawn12Diffs = append(dawn12Diffs, dawn12Diff)
		dawn6Diffs = append(dawn6Diffs, dawn6Diff)
		sunriseDiffs = append(sunriseDiffs, sunriseDiff)
		transitDiffs = append(transitDiffs, transitDiff)
		sunsetDiffs = append(sunsetDiffs, sunsetDiff)
		dusk6Diffs = append(dusk6Diffs, dusk6Diff)
		dusk12Diffs = append(dusk12Diffs, dusk12Diff)
		dusk18Diffs = append(dusk18Diffs, dusk18Diff)
		sunriseAzimuthDiffs = append(sunriseAzimuthDiffs, sunriseAzimuthDiff)
		sunsetAzimuthDiffs = append(sunsetAzimuthDiffs, sunsetAzimuthDiff)
		transitAltitudeDiffs = append(transitAltitudeDiffs, transitAltitudeDiff)
	}

	// Print diff table
	var t Table
	t.AddRow("Name", "Max", "Mode", "Avg")
	t.AddRow(timeDiffRow("Dawn 18", dawn18Diffs)...)
	t.AddRow(timeDiffRow("Dawn 12", dawn12Diffs)...)
	t.AddRow(timeDiffRow("Dawn 6", dawn6Diffs)...)
	t.AddRow(timeDiffRow("Sunrise", sunriseDiffs)...)
	t.AddRow(timeDiffRow("Transit", transitDiffs)...)
	t.AddRow(timeDiffRow("Sunset", sunsetDiffs)...)
	t.AddRow(timeDiffRow("Dusk 6", dusk6Diffs)...)
	t.AddRow(timeDiffRow("Dusk 12", dusk12Diffs)...)
	t.AddRow(timeDiffRow("Dusk 18", dusk18Diffs)...)
	t.AddRow(floatDiffRow("Sunrise azimuth", sunriseAzimuthDiffs)...)
	t.AddRow(floatDiffRow("Sunset azimuth", sunsetAzimuthDiffs)...)
	t.AddRow(floatDiffRow("Transit altitude", transitAltitudeDiffs)...)

	fmt.Printf("### Sun in %s\n\n", location.Name)
	fmt.Println(t.String())
	fmt.Println()
	return nil
}

func compareMoonSchedules(location Location) error {
	// Prepare expected schedules from CSV file
	expectedSchedules, err := parseMoonCSV(location.CsvMoon, location.Timezone)
	if err != nil {
		return err
	}

	// Calculate schedules using SAMPA
	calculatedSchedules, err := calculateMoonEvents(sampa.Location{
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
	}, location.Timezone)
	if err != nil {
		return err
	}

	// Compare schedules
	var moonriseDiffs []int
	var transitDiffs []int
	var moonsetDiffs []int
	var moonriseAzimuthDiffs []float64
	var moonsetAzimuthDiffs []float64
	var transitAltitudeDiffs []float64
	var illuminationDiffs []float64

	moonriseTitle := location.Name + ", moonrise"
	transitTitle := location.Name + ", transit"
	moonsetTitle := location.Name + ", moonset"

	for i, exp := range expectedSchedules {
		res := calculatedSchedules[i]
		moonriseDiff := compareTime(moonriseTitle, res.Date, exp.Moonrise, res.Moonrise)
		transitDiff := compareTime(transitTitle, res.Date, exp.Transit, res.Transit)
		moonsetDiff := compareTime(moonsetTitle, res.Date, exp.Moonset, res.Moonset)
		moonriseAzimuthDiff := math.Abs(res.MoonriseAzimuth - exp.MoonriseAzimuth)
		moonsetAzimuthDiff := math.Abs(res.MoonsetAzimuth - exp.MoonsetAzimuth)
		transitAltitudeDiff := math.Abs(res.TransitAltitude - exp.TransitAltitude)
		illuminationDiff := math.Abs(res.Illumination - exp.Illumination)

		moonriseDiffs = append(moonriseDiffs, moonriseDiff)
		transitDiffs = append(transitDiffs, transitDiff)
		moonsetDiffs = append(moonsetDiffs, moonsetDiff)
		moonriseAzimuthDiffs = append(moonriseAzimuthDiffs, moonriseAzimuthDiff)
		moonsetAzimuthDiffs = append(moonsetAzimuthDiffs, moonsetAzimuthDiff)
		transitAltitudeDiffs = append(transitAltitudeDiffs, transitAltitudeDiff)
		illuminationDiffs = append(illuminationDiffs, illuminationDiff)
	}

	// Print diff table
	var t Table
	t.AddRow("Name", "Max", "Mode", "Avg")
	t.AddRow(timeDiffRow("Moonrise", moonriseDiffs)...)
	t.AddRow(timeDiffRow("Transit", transitDiffs)...)
	t.AddRow(timeDiffRow("Moonset", moonsetDiffs)...)
	t.AddRow(floatDiffRow("Moonrise azimuth", moonriseAzimuthDiffs)...)
	t.AddRow(floatDiffRow("Moonset azimuth", moonsetAzimuthDiffs)...)
	t.AddRow(floatDiffRow("Transit altitude", transitAltitudeDiffs)...)
	t.AddRow(floatDiffRow("Illumination", illuminationDiffs)...)

	// Print diff stat
	fmt.Printf("### Moon in %s\n\n", location.Name)
	fmt.Println(t.String())
	fmt.Println()
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

func diffStat[T int | float64](diffs []T) (max T, mode T, avg float64) {
	nDiff := len(diffs)
	if nDiff == 0 {
		return
	}

	// Calculate various stat
	var sum T
	diffCount := make(map[T]int)
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

func timeDiffRow(title string, diffs []int) []string {
	max, mode, avg := diffStat(diffs)
	return []string{
		title,
		fmt.Sprintf("%ds", max),
		fmt.Sprintf("%ds", mode),
		fmt.Sprintf("%.2fs", avg),
	}
}

func floatDiffRow(title string, diffs []float64) []string {
	max, mode, avg := diffStat(diffs)
	return []string{
		title,
		fmt.Sprintf("%.2f", max),
		fmt.Sprintf("%.2f", mode),
		fmt.Sprintf("%.2f", avg),
	}
}
