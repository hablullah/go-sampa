package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	tz := time.FixedZone("WIB", 7*60*60)
	dt := time.Date(1973, 12, 31, 0, 0, 0, 0, tz)

	opts := &sampa.Options{DeltaT: 72.8}
	phases := sampa.GetMoonPhases(dt, opts)
	bt, _ := json.MarshalIndent(&phases, "", "  ")
	fmt.Println(string(bt))
}
