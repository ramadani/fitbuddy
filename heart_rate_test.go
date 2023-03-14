package fitbuddy_test

import (
	"github.com/ramadani/fitbuddy"
	"testing"
	"time"
)

func TestGetMaxHeartRate(t *testing.T) {
	t.Run("20YearsOld", func(t *testing.T) {
		year, month, day := time.Now().Date()
		dob := time.Date(year-20, month, day, 0, 0, 0, 0, time.Local)
		expected := 200

		actual := fitbuddy.GetMaxHeartRate(dob)

		if expected != actual {
			t.Errorf("max heart rate should be %v", expected)
		}
	})

	t.Run("27YearsOld", func(t *testing.T) {
		year, month, day := time.Now().Date()
		dob := time.Date(year-27, month+3, day, 0, 0, 0, 0, time.Local)
		expected := 194

		actual := fitbuddy.GetMaxHeartRate(dob)

		if expected != actual {
			t.Errorf("max heart rate should be %v not %v", expected, actual)
		}
	})
}
