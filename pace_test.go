package fitbuddy_test

import (
	"github.com/ramadani/fitbuddy"
	"testing"
	"time"
)

func TestPace_String(t *testing.T) {
	tests := []struct {
		name   string
		pace   fitbuddy.Pace
		result string
	}{
		{
			name: "5:00 /km",
			pace: fitbuddy.Pace{
				Value: 300 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "5:00 /km",
		},
		{
			name: "5:05 /km",
			pace: fitbuddy.Pace{
				Value: 305 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "5:05 /km",
		},
		{
			name: "5:15 /km",
			pace: fitbuddy.Pace{
				Value: 315 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "5:15 /km",
		},
		{
			name: "15:15 /km",
			pace: fitbuddy.Pace{
				Value: 915 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "15:15 /km",
		},
		{
			name: "1:00:15 /km",
			pace: fitbuddy.Pace{
				Value: 3615 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "1:00:15 /km",
		},
		{
			name: "1:05:15 /km",
			pace: fitbuddy.Pace{
				Value: 3915 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "1:05:15 /km",
		},
		{
			name: "1:15:15 /km",
			pace: fitbuddy.Pace{
				Value: 4515 * time.Second,
				Unit:  fitbuddy.Kilometer,
			},
			result: "1:15:15 /km",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.pace.String()

			if test.result != actual {
				t.Errorf("pace label should be %v not %v", test.result, actual)
			}
		})
	}
}

func TestGetPace(t *testing.T) {
	thirtyMinutes, _ := time.ParseDuration("30m0s")
	sixPaceDur, _ := time.ParseDuration("6m0s")
	thirtyMinutesAndFortyFive, _ := time.ParseDuration("30m45s")
	nineMinutesAndFiftyFourDur, _ := time.ParseDuration("9m54s")
	oneHourAndThirtyMinutes, _ := time.ParseDuration("1h30m0s")
	sixMinutesAndFiftyTwoDur, _ := time.ParseDuration("6m52s")
	twoHoursTenMinutesAndFifteen, _ := time.ParseDuration("2h10m15s")
	threeMinutesAndFiveDur, _ := time.ParseDuration("3m5s")

	tests := []struct {
		name     string
		distance fitbuddy.Distance
		duration time.Duration
		result   fitbuddy.Pace
	}{
		{
			name: "5km in 30:00 is 6:00 /km",
			distance: fitbuddy.Distance{
				Value: 5 * fitbuddy.Km,
				Unit:  fitbuddy.Kilometer,
			},
			duration: thirtyMinutes,
			result: fitbuddy.Pace{
				Value: sixPaceDur,
				Unit:  fitbuddy.Kilometer,
			},
		},
		{
			name: "5km in 30:45 is 9:54 /mi",
			distance: fitbuddy.Distance{
				Value: 5 / fitbuddy.Mi,
				Unit:  fitbuddy.Mile,
			},
			duration: thirtyMinutesAndFortyFive,
			result: fitbuddy.Pace{
				Value: nineMinutesAndFiftyFourDur,
				Unit:  fitbuddy.Mile,
			},
		},
		{
			name: "half marathon in 1:30:00 is 6:52 /mi",
			distance: fitbuddy.Distance{
				Value: 13.1,
				Unit:  fitbuddy.Mile,
			},
			duration: oneHourAndThirtyMinutes,
			result: fitbuddy.Pace{
				Value: sixMinutesAndFiftyTwoDur,
				Unit:  fitbuddy.Mile,
			},
		},
		{
			name: "marathon in 2:10:15 is 3:05 /km",
			distance: fitbuddy.Distance{
				Value: 26.2 * fitbuddy.Mi,
				Unit:  fitbuddy.Kilometer,
			},
			duration: twoHoursTenMinutesAndFifteen,
			result: fitbuddy.Pace{
				Value: threeMinutesAndFiveDur,
				Unit:  fitbuddy.Kilometer,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := fitbuddy.GetPace(test.distance, test.duration)

			if test.result != actual {
				t.Errorf("pace should be %v not %v", test.result, actual)
			}
		})
	}
}
