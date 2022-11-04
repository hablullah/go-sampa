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

var (
	// Tromso (Denmark) is representation for location in North Frigid area
	tromso = TestLocation{
		Latitude:    69.682778,
		Longitude:   18.942778,
		Elevation:   0,
		Temperature: 10,
		Pressure:    1010,
	}

	// London (UK) is representation for location in North Temperate area
	london = TestLocation{
		Latitude:    51.507222,
		Longitude:   -0.1275,
		Elevation:   11,
		Temperature: 10,
		Pressure:    1010,
	}

	// Jakarta (Indonesia) is representation for location in Torrid area
	jakarta = TestLocation{
		Latitude:    -6.175,
		Longitude:   106.825,
		Elevation:   8,
		Temperature: 10,
		Pressure:    1010,
	}

	// Wellington (New Zealand) is representation for location in South Temperate area
	wellington = TestLocation{
		Latitude:    -41.288889,
		Longitude:   174.777222,
		Elevation:   0,
		Temperature: 10,
		Pressure:    1010,
	}
)

var (
	CET  = time.FixedZone("CET", 1*60*60)   // Tromso
	UTC  = time.UTC                         // London
	WIB  = time.FixedZone("WIB", 7*60*60)   // Jakarta
	NZST = time.FixedZone("NZST", 12*60*60) // Jakarta
)
