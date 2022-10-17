package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hablullah/go-sampa"
)

func main() {
	tz := time.FixedZone("CST", 1*60*60)
	dt := time.Date(2022, 2, 1, 0, 0, 0, 0, tz)

	phases := sampa.GetMoonPhases(dt, nil)
	bt, _ := json.MarshalIndent(&phases, "", "  ")
	fmt.Println(string(bt))
}
