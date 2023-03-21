package fitbuddy_test

import (
	"github.com/ramadani/fitbuddy"
	"testing"
	"time"
)

func TestGetMaxHeartRate(t *testing.T) {
	year, month, day := time.Now().Date()

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "20YearsOld",
			dob:      time.Date(year-20, month, day, 0, 0, 0, 0, time.Local),
			expected: 200,
		},
		{
			name:     "27YearsOld",
			dob:      time.Date(year-27, month+3, day, 0, 0, 0, 0, time.Local),
			expected: 194,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := fitbuddy.GetMaxHeartRate(test.dob)

			if test.expected != actual {
				t.Errorf("max heart rate should be %v not %v", test.expected, actual)
			}
		})
	}
}

func TestGetHeartRateZone(t *testing.T) {
	tests := []struct {
		name     string
		hr       int
		maxHr    int
		expected fitbuddy.HRZone
	}{
		{
			name:     "Zone1",
			hr:       105,
			maxHr:    200,
			expected: fitbuddy.HRZone1,
		},
		{
			name:     "Zone3",
			hr:       138,
			maxHr:    185,
			expected: fitbuddy.HRZone3,
		},
		{
			name:     "Zone4",
			hr:       149,
			maxHr:    185,
			expected: fitbuddy.HRZone4,
		},
		{
			name:     "Zone2",
			hr:       124,
			maxHr:    195,
			expected: fitbuddy.HRZone2,
		},
		{
			name:     "Zone5",
			hr:       173,
			maxHr:    190,
			expected: fitbuddy.HRZone5,
		},
		{
			name:     "UnknownZone",
			hr:       75,
			maxHr:    165,
			expected: fitbuddy.UnknownHRZone,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := fitbuddy.GetHeartRateZone(test.hr, test.maxHr)

			if test.expected != actual {
				t.Errorf("heart rate zone should be %v not %v", test.expected, actual)
			}
		})
	}
}
