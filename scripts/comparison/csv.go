package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

type SunSchedule struct {
	Date    string
	Dawn    time.Time
	Sunrise time.Time
	Transit time.Time
	Sunset  time.Time
	Dusk    time.Time
}

func parseSunCSV(srcPath string, tz *time.Location) ([]SunSchedule, error) {
	// Open file
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse CSV file
	var records [][]string
	dateCount := make(map[string]int64)
	csvReader := csv.NewReader(f)

	for {
		// Read the record
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		// Skip empty record
		if len(record) != 20 {
			continue
		}

		// Skip header
		if record[0] == "Date" {
			continue
		}

		// Save data
		date := record[0]
		line, _ := strconv.ParseInt(record[1], 10, 64)
		records = append(records, record)
		if line > dateCount[date] {
			dateCount[date] = line
		}
	}

	// Convert records to schedules
	schedules := make([]SunSchedule, nDays)
	for _, record := range records {
		// Get record data
		strDate := record[0]
		date, _ := parseDate(strDate, tz)
		dawn, _ := parseTime(strDate, record[2], tz)
		dusk, _ := parseTime(strDate, record[3], tz)
		sunrise, _ := parseTime(strDate, record[8], tz)
		sunset, _ := parseTime(strDate, record[9], tz)
		transit, _ := parseTime(strDate, record[17], tz)

		// Prepare index for saving schedule
		idx := date.YearDay() - 1
		if transit.IsZero() {
			transit = schedules[idx].Transit
		}

		// Save the schedule
		schedules[idx].Date = strDate
		schedules[idx].Transit = transit

		if !dawn.IsZero() {
			dawnIdx := prepareSunIndex(idx, dawn, transit, true)
			schedules[dawnIdx].Dawn = dawn
		}

		if !dusk.IsZero() {
			duskIdx := prepareSunIndex(idx, dusk, transit, false)
			schedules[duskIdx].Dusk = dusk
		}

		if !sunrise.IsZero() {
			sunriseIdx := prepareSunIndex(idx, sunrise, transit, true)
			schedules[sunriseIdx].Sunrise = sunrise
		}

		if !sunset.IsZero() {
			sunsetIdx := prepareSunIndex(idx, sunset, transit, false)
			schedules[sunsetIdx].Sunset = sunset
		}
	}

	return schedules, nil
}

func prepareSunIndex(currentIdx int, currentTime, transitTime time.Time, isBeforeTransit bool) int {
	if (isBeforeTransit && currentTime.After(transitTime)) ||
		(!isBeforeTransit && currentTime.Before(transitTime)) {
		currentIdx--
	}

	if currentIdx < 0 {
		currentIdx = nDays - currentIdx
	}

	return currentIdx
}

func parseDate(s string, tz *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, tz)
}

func parseTime(strDate, strTime string, tz *time.Location) (time.Time, error) {
	str := strDate + " " + strTime
	return time.ParseInLocation("2006-01-02 15:04:05", str, tz)
}
