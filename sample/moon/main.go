package main

import (
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	location := sampa.Location{
		Latitude:    24.61167,
		Longitude:   143.36167,
		Elevation:   0,
		Temperature: 11,
	}

	opts := sampa.Options{
		DeltaT: 66.4,
	}

	dt := time.Date(2009, 7, 22, 1, 33, 0, 0, time.UTC)
	data, err := sampa.GetMoonPosition(dt, location, &opts)
	if err != nil {
		panic(err)
	}
	data = data

	// bt, _ := json.MarshalIndent(&data, "", "  ")
	// fmt.Println(string(bt))
}
