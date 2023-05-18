package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type SunData struct {
	Date    string
	Dawn18  time.Time
	Dawn12  time.Time
	Dawn6   time.Time
	Sunrise time.Time
	Transit time.Time
	Sunset  time.Time
	Dusk6   time.Time
	Dusk12  time.Time
	Dusk18  time.Time

	SunriseAzimuth  float64
	SunsetAzimuth   float64
	TransitAltitude float64
}

type MoonSchedule struct {
	Date     string
	Moonrise time.Time
	Transit  time.Time
	Moonset  time.Time

	MoonriseAzimuth float64
	MoonsetAzimuth  float64
	TransitAltitude float64
	Illumination    float64
}

func parseSunCSV(srcPath string, tz *time.Location) ([]SunData, error) {
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
		if len(record) != 19 {
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
	schedules := make([]SunData, nDays)
	for _, record := range records {
		// Get record data
		strDate := record[0]
		date, _ := parseDate(strDate, tz)
		dawn18, _ := parseTime(strDate, record[1], tz)
		dusk18, _ := parseTime(strDate, record[2], tz)
		dawn12, _ := parseTime(strDate, record[3], tz)
		dusk12, _ := parseTime(strDate, record[4], tz)
		dawn6, _ := parseTime(strDate, record[5], tz)
		dusk6, _ := parseTime(strDate, record[6], tz)
		sunrise, _ := parseTime(strDate, record[7], tz)
		sunset, _ := parseTime(strDate, record[8], tz)
		transit, _ := parseTime(strDate, record[16], tz)
		sunriseAzimuth, _ := parseFloat(record[9])
		sunsetAzimuth, _ := parseFloat(record[10])
		transitAltitude, _ := parseFloat(record[17])

		// Prepare index for saving schedule
		idx := date.YearDay() - 1

		// Save the schedule
		schedules[idx].Date = strDate
		schedules[idx].Transit = transit
		schedules[idx].TransitAltitude = transitAltitude

		if !dawn18.IsZero() {
			newIdx := prepareTimeIndex(idx, dawn18, transit, true)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dawn18 = dawn18
			}
		}

		if !dusk18.IsZero() {
			newIdx := prepareTimeIndex(idx, dusk18, transit, false)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dusk18 = dusk18
			}
		}

		if !dawn12.IsZero() {
			newIdx := prepareTimeIndex(idx, dawn12, transit, true)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dawn12 = dawn12
			}
		}

		if !dusk12.IsZero() {
			newIdx := prepareTimeIndex(idx, dusk12, transit, false)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dusk12 = dusk12
			}
		}

		if !dawn6.IsZero() {
			newIdx := prepareTimeIndex(idx, dawn6, transit, true)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dawn6 = dawn6
			}
		}

		if !dusk6.IsZero() {
			newIdx := prepareTimeIndex(idx, dusk6, transit, false)
			if newIdx >= 0 && newIdx < nDays {
				schedules[newIdx].Dusk6 = dusk6
			}
		}

		if !sunrise.IsZero() {
			sunriseIdx := prepareTimeIndex(idx, sunrise, transit, true)
			if sunriseIdx >= 0 && sunriseIdx < nDays {
				schedules[sunriseIdx].Sunrise = sunrise
				schedules[sunriseIdx].SunriseAzimuth = sunriseAzimuth
			}
		}

		if !sunset.IsZero() {
			sunsetIdx := prepareTimeIndex(idx, sunset, transit, false)
			if sunsetIdx >= 0 && sunsetIdx < nDays {
				schedules[sunsetIdx].Sunset = sunset
				schedules[sunsetIdx].SunsetAzimuth = sunsetAzimuth
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
		if len(record) != 12 {
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
		moonrise, _ := parseTime(strDate, record[1], tz)
		moonset, _ := parseTime(strDate, record[2], tz)
		transit, _ := parseTime(strDate, record[5], tz)
		moonriseAzimuth, _ := parseFloat(record[3])
		moonsetAzimuth, _ := parseFloat(record[4])
		transitAltitude, _ := parseFloat(record[6])
		illumination, _ := parseFloat(strings.TrimSuffix(record[8], "%"))

		// Prepare index for saving schedule
		idx := date.YearDay() - 1
		if transit.IsZero() {
			transit = schedules[idx].Transit
		}

		// Save the schedule
		schedules[idx].Date = strDate
		schedules[idx].Transit = transit
		schedules[idx].TransitAltitude = transitAltitude
		schedules[idx].Illumination = illumination

		if !moonrise.IsZero() {
			moonriseIdx := prepareTimeIndex(idx, moonrise, transit, true)
			if moonriseIdx >= 0 && moonriseIdx < nDays {
				schedules[moonriseIdx].Moonrise = moonrise
				schedules[moonriseIdx].MoonriseAzimuth = moonriseAzimuth
			}
		}

		if !moonset.IsZero() {
			moonsetIdx := prepareTimeIndex(idx, moonset, transit, false)
			if moonsetIdx >= 0 && moonsetIdx < nDays {
				schedules[moonsetIdx].Moonset = moonset
				schedules[moonsetIdx].MoonsetAzimuth = moonsetAzimuth
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

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
