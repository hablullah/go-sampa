package main

import (
	"encoding/csv"
	"io"
	"os"
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

type MoonSchedule struct {
	Date     string
	Moonrise time.Time
	Transit  time.Time
	Moonset  time.Time
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
		records = append(records, record)
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
			dawnIdx := prepareTimeIndex(idx, dawn, transit, true)
			if dawnIdx >= 0 && dawnIdx < nDays {
				schedules[dawnIdx].Dawn = dawn
			}
		}

		if !dusk.IsZero() {
			duskIdx := prepareTimeIndex(idx, dusk, transit, false)
			if duskIdx >= 0 && duskIdx < nDays {
				schedules[duskIdx].Dusk = dusk
			}
		}

		if !sunrise.IsZero() {
			sunriseIdx := prepareTimeIndex(idx, sunrise, transit, true)
			if sunriseIdx >= 0 && sunriseIdx < nDays {
				schedules[sunriseIdx].Sunrise = sunrise
			}
		}

		if !sunset.IsZero() {
			sunsetIdx := prepareTimeIndex(idx, sunset, transit, false)
			if sunsetIdx >= 0 && sunsetIdx < nDays {
				schedules[sunsetIdx].Sunset = sunset
			}
		}
	}

	return schedules, nil
}

func parseMoonCSV(srcPath string, tz *time.Location) ([]MoonSchedule, error) {
	// Open file
	f, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse CSV file
	var records [][]string
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
		if len(record) != 13 {
			continue
		}

		// Skip header
		if record[0] == "Date" {
			continue
		}

		// Save data
		records = append(records, record)
	}

	// Convert records to schedules
	schedules := make([]MoonSchedule, nDays)
	for _, record := range records {
		// Get record data
		strDate := record[0]
		date, _ := parseDate(strDate, tz)
		moonrise, _ := parseTime(strDate, record[2], tz)
		moonset, _ := parseTime(strDate, record[3], tz)
		transit, _ := parseTime(strDate, record[6], tz)

		// Prepare index for saving schedule
		idx := date.YearDay() - 1
		if transit.IsZero() {
			transit = schedules[idx].Transit
		}

		// Save the schedule
		schedules[idx].Date = strDate
		schedules[idx].Transit = transit

		if !moonrise.IsZero() {
			moonriseIdx := prepareTimeIndex(idx, moonrise, transit, true)
			if moonriseIdx >= 0 && moonriseIdx < nDays {
				schedules[moonriseIdx].Moonrise = moonrise
			}
		}

		if !moonset.IsZero() {
			moonsetIdx := prepareTimeIndex(idx, moonset, transit, false)
			if moonsetIdx >= 0 && moonsetIdx < nDays {
				schedules[moonsetIdx].Moonset = moonset
			}
		}
	}

	return schedules, nil
}

func prepareTimeIndex(currentIdx int, currentTime, transitTime time.Time, isBeforeTransit bool) int {
	if transitTime.IsZero() {
		if isBeforeTransit {
			return currentIdx + 1
		} else {
			return currentIdx - 1
		}
	}

	if isBeforeTransit && currentTime.After(transitTime) {
		currentIdx++
	} else if !isBeforeTransit && currentTime.Before(transitTime) {
		currentIdx--
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
