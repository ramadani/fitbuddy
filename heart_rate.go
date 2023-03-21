package fitbuddy

import (
	"github.com/ramadani/fitbuddy/percent"
	"time"
)

type HRZone int

const (
	HRZone1 HRZone = iota
	HRZone2
	HRZone3
	HRZone4
	HRZone5
	UnknownHRZone
)

var hrZoneRanges = []struct {
	zone  HRZone
	start float64
	end   float64
}{
	{zone: HRZone1, start: 50, end: 60},
	{zone: HRZone2, start: 60, end: 70},
	{zone: HRZone3, start: 70, end: 80},
	{zone: HRZone4, start: 80, end: 90},
	{zone: HRZone5, start: 90, end: 100},
}

// GetMaxHeartRate Maximum heart rate (Max HR) is the highest number of times a person's heart
// can beat in one minute during physical activity.
// The traditional formula for calculating max HR is 220 minus the person's age.
func GetMaxHeartRate(dob time.Time) int {
	days := time.Since(dob).Hours() / 24
	age := days / 365

	return 220 - int(age)
}

// GetHeartRateZone Heart rate zone refers to a range of heart rate levels that are used to determine
// the intensity of a person's workout or exercise. The heart rate zone is typically calculated
// based on a percentage of an individual's maximum heart rate,
// which is the highest number of times a person's heart can beat in a minute.
func GetHeartRateZone(hr, maxHr int) HRZone {
	for _, z := range hrZoneRanges {
		if isHrZoneInRange(percent.PercentOf(hr, maxHr), z.start, z.end) {
			return z.zone
		}
	}

	return UnknownHRZone
}

func isHrZoneInRange(val, start, end float64) bool {
	return val > start && val <= end
}
