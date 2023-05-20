package main

import (
	"fmt"
	"log"
)

const (
	year      = 2023
	nDays     = 365
	maxDiff   = 60
	enableLog = false
)

func main() {
	log.SetFlags(0)

	fmt.Printf("# Comparison Results\n\n")

	fmt.Printf("## Sun Events\n\n")
	compareSunSchedules(LordHoweIsland)
	compareSunSchedules(Maputo)
	compareSunSchedules(Amsterdam)
	compareSunSchedules(Oslo)
	compareSunSchedules(Philipsburg)
	compareSunSchedules(NewYork)

	fmt.Printf("## Moon Events\n\n")
	compareMoonSchedules(LordHoweIsland)
	compareMoonSchedules(Maputo)
	compareMoonSchedules(Amsterdam)
	compareMoonSchedules(Oslo)
	compareMoonSchedules(Philipsburg)
	compareMoonSchedules(NewYork)
}
