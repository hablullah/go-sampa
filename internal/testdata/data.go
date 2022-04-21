package testdata

import "time"

type TestLocation struct {
	Latitude    float64
	Longitude   float64
	Elevation   float64
	Temperature float64
	Pressure    float64
}

type TestTime struct {
	Date    string
	Rise    string
	Transit string
	Set     string
}

type TestData struct {
	Name     string
	Z        *time.Location
	Location TestLocation
	Times    []TestTime
}
