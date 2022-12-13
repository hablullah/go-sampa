package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Clean up dst dir
	dstDir := "internal/testdata"
	os.RemoveAll(dstDir)
	os.MkdirAll(dstDir, os.ModePerm)

	// Generate common files
	err := genCommonFiles(dstDir)
	checkError(err)

	// Generate test data for each location
	for _, location := range testLocations {
		err = genTestData(location, dstDir)
		checkError(err)
	}
}

func genCommonFiles(dstDir string) error {
	// Write package header and imports
	sb := bytes.NewBuffer(nil)
	sb.WriteString("package testdata\n")
	sb.WriteString("import \"time\"\n")
	sb.WriteString("import \"github.com/hablullah/go-sampa\"\n")

	// Put struct for celestial events
	sb.WriteString(`type CelestialEvent struct {
		Date    string
		Dawn    string
		Rise    string
		Transit string
		Set     string
		Dusk    string
	}` + "\n\n")

	// Put struct for test data
	sb.WriteString(`type TestData struct {
		Name       string
		Latitude   float64
		Longitude  float64
		Timezone   *time.Location
		SunEvents  []CelestialEvent
		MoonEvents []CelestialEvent
	}` + "\n")

	// Put variables for custom events
	sb.WriteString(`
	var SunEvents = []sampa.CustomSunEvent{{
		Name:          "Dawn",
		BeforeTransit: true,
		Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
	}, {
		Name:          "Dusk",
		BeforeTransit: false,
		Elevation:     func(_ sampa.SunPosition) float64 { return -18 },
	}}
	
	var MoonEvents = []sampa.CustomMoonEvent{{
		Name:          "Dawn",
		BeforeTransit: true,
		Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
	}, {
		Name:          "Dusk",
		BeforeTransit: false,
		Elevation:     func(_ sampa.MoonPosition) float64 { return -18 },
	}}
	`)

	// Format code
	bt, err := format.Source(sb.Bytes())
	if err != nil {
		return err
	}

	// Save to file
	dstPath := filepath.Join(dstDir, "common.go")
	return os.WriteFile(dstPath, bt, os.ModePerm)
}

func genTestData(loc Location, dstDir string) error {
	// Write package header and imports
	sb := bytes.NewBuffer(nil)
	sb.WriteString("package testdata\n")
	sb.WriteString("import \"time\"\n")

	// Put the variable for timezone
	sbWritef(sb,
		"var tz%s, _ = time.LoadLocation(%q)\n\n",
		loc.Name, loc.Timezone)

	// Put the variable for location
	sbWritef(sb, "var %s = TestData {\n"+
		"Name: %q,\n"+
		"Latitude: %f,\n"+
		"Longitude: %f,\n"+
		"Timezone: tz%s,\n",
		loc.Name, loc.Name, loc.Latitude, loc.Longitude, loc.Name)

	// Calculate and put sun events
	sb.WriteString("SunEvents: []CelestialEvent{\n")
	for _, e := range getSunEvents(loc) {
		sbWritef(sb, "{%q,%q,%q,%q,%q,%q},\n",
			e.Date,
			strTime(e.Others["Dawn"].DateTime),
			strTime(e.Sunrise.DateTime),
			strTime(e.Transit.DateTime),
			strTime(e.Sunset.DateTime),
			strTime(e.Others["Dusk"].DateTime))
	}
	sb.WriteString("},\n")

	// Calculate and put moon events
	sb.WriteString("MoonEvents: []CelestialEvent{\n")
	for _, e := range getMoonEvents(loc) {
		sbWritef(sb, "{%q,%q,%q,%q,%q,%q},\n",
			e.Date,
			strTime(e.Others["Dawn"].DateTime),
			strTime(e.Moonrise.DateTime),
			strTime(e.Transit.DateTime),
			strTime(e.Moonset.DateTime),
			strTime(e.Others["Dusk"].DateTime))
	}
	sb.WriteString("},\n")
	sb.WriteString("}\n")

	// Format code
	bt, err := format.Source(sb.Bytes())
	if err != nil {
		return err
	}

	// Save to file
	dstPath := filepath.Join(dstDir, strings.ToLower(loc.Name)+".go")
	return os.WriteFile(dstPath, bt, os.ModePerm)
}

func sbWritef(sb *bytes.Buffer, format string, args ...any) {
	sb.WriteString(fmt.Sprintf(format, args...))
}

func strTime(t time.Time) string {
	if !t.IsZero() {
		return t.Format("2006-01-02 15:04:05 -0700")
	} else {
		return strings.Repeat(" ", 25)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
