package main

import "log"

const (
	year      = 2023
	nDays     = 365
	maxDiff   = 30
	enableLog = false
)

func main() {
	log.SetFlags(0)

	compareSunSchedules(LordHoweIsland)
	compareSunSchedules(Maputo)
	compareSunSchedules(Amsterdam)
	compareSunSchedules(Oslo)
	compareSunSchedules(Philipsburg)
	compareSunSchedules(NewYork)
}
