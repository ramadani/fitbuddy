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
