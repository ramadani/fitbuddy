package fitbuddy_test

import (
	"github.com/ramadani/fitbuddy"
	"testing"
)

func TestDistanceUnitType_String(t *testing.T) {
	tests := []struct {
		name   string
		unit   fitbuddy.DistanceUnitType
		result string
	}{
		{
			name:   "Kilometer",
			unit:   fitbuddy.Kilometer,
			result: "km",
		},
		{
			name:   "Mile",
			unit:   fitbuddy.Mile,
			result: "mi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.unit.String()

			if test.result != actual {
				t.Errorf("pace label should be %v not %v", test.result, actual)
			}
		})
	}
}

func TestDistance_String(t *testing.T) {
	tests := []struct {
		name     string
		distance fitbuddy.Distance
		result   string
	}{
		{
			name: "5 km",
			distance: fitbuddy.Distance{
				Value: 5,
				Unit:  fitbuddy.Kilometer,
			},
			result: "5 km",
		},
		{
			name: "21.1 km",
			distance: fitbuddy.Distance{
				Value: 21.1,
				Unit:  fitbuddy.Kilometer,
			},
			result: "21.1 km",
		},
		{
			name: "3.107 mi",
			distance: fitbuddy.Distance{
				Value: 3.107,
				Unit:  fitbuddy.Mile,
			},
			result: "3.107 mi",
		},
		{
			name: "10 mi",
			distance: fitbuddy.Distance{
				Value: 10,
				Unit:  fitbuddy.Mile,
			},
			result: "10 mi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.distance.String()

			if test.result != actual {
				t.Errorf("distance label should be %v not %v", test.result, actual)
			}
		})
	}
}

func TestDistanceConverter(t *testing.T) {
	tests := []struct {
		name     string
		distance fitbuddy.Distance
		destUnit fitbuddy.DistanceUnitType
		result   fitbuddy.Distance
		err      error
	}{
		{
			name:     "5 km to 3.11 mi",
			distance: fitbuddy.Distance{Value: 5, Unit: fitbuddy.Kilometer},
			destUnit: fitbuddy.Mile,
			result:   fitbuddy.Distance{Value: 3.11, Unit: fitbuddy.Mile},
			err:      nil,
		},
		{
			name:     "13.11 mi to 21.1 km",
			distance: fitbuddy.Distance{Value: 13.11, Unit: fitbuddy.Mile},
			destUnit: fitbuddy.Kilometer,
			result:   fitbuddy.Distance{Value: 21.1, Unit: fitbuddy.Kilometer},
			err:      nil,
		},
		{
			name:     "10 km to 10 km",
			distance: fitbuddy.Distance{Value: 10, Unit: fitbuddy.Kilometer},
			destUnit: fitbuddy.Kilometer,
			result:   fitbuddy.Distance{Value: 10, Unit: fitbuddy.Kilometer},
			err:      nil,
		},
		{
			name:     "Destination Unit Error",
			distance: fitbuddy.Distance{Value: 10, Unit: fitbuddy.Kilometer},
			destUnit: 2,
			result:   fitbuddy.Distance{Value: 0, Unit: fitbuddy.DistanceUnitType(0)},
			err:      fitbuddy.ErrUnsupportedDistanceUnit,
		},
		{
			name:     "Distance Unit Error",
			distance: fitbuddy.Distance{Value: 10, Unit: fitbuddy.DistanceUnitType(-1)},
			destUnit: 2,
			result:   fitbuddy.Distance{Value: 0, Unit: fitbuddy.DistanceUnitType(0)},
			err:      fitbuddy.ErrUnsupportedDistanceUnit,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := fitbuddy.DistanceConverter(test.distance, test.destUnit)

			if test.err != err {
				t.Errorf("error should be equal to %v not %v", test.err, err)
			}

			if test.result != actual {
				t.Errorf("distance label should be %v not %v", test.result, actual)
			}
		})
	}
}
