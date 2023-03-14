package fitbuddy

import "time"

// GetMaxHeartRate 220 - age from dob
func GetMaxHeartRate(dob time.Time) int {
	days := time.Since(dob).Hours() / 24
	year := days / 365

	return 220 - int(year)
}
