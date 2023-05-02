package fitbuddy

import (
	"fmt"
	"math"
)

type DistanceUnitType int

const (
	// Kilometer is a unit of distance or length in the metric system, one kilometer is equal to 1,000 meters or
	// approximately 0.62 miles.
	Kilometer DistanceUnitType = iota
	// Mile refers to one mile is equivalent to 1.609 kilometers or 1,760 yards.
	Mile
)

const (
	Km float64 = 1
	Mi         = 1.60934
)

var (
	distanceUnitTypeLabels = map[DistanceUnitType]string{
		Kilometer: "km",
		Mile:      "mi",
	}

	distanceConversionUnitMap = map[DistanceUnitType]map[DistanceUnitType]float64{
		Kilometer: {
			Mile: 0.62137,
		},
		Mile: {
			Kilometer: 1.60934,
		},
	}
)

// String get label of distance unit
func (d DistanceUnitType) String() string {
	return distanceUnitTypeLabels[d]
}

// Distance refers to the length of the route or distance covered by a runner during a running activity.
// Running distance can be measured in various units of length such as kilometers, miles, or meters.
type Distance struct {
	Value float64
	Unit  DistanceUnitType
}

// String get distance label based on unit
func (d Distance) String() string {
	return fmt.Sprintf("%v %s", d.Value, d.Unit.String())
}

// DistanceConverter .
func DistanceConverter(distance Distance, destUnit DistanceUnitType) (Distance, error) {
	if distance.Unit == destUnit {
		return distance, nil
	}

	distUnits, ok := distanceConversionUnitMap[distance.Unit]
	if !ok {
		return Distance{}, ErrUnsupportedDistanceUnit
	}

	comp, ok := distUnits[destUnit]
	if !ok {
		return Distance{}, ErrUnsupportedDistanceUnit
	}

	val := distance.Value * comp

	return Distance{
		Value: math.Round(val*100) / 100,
		Unit:  destUnit,
	}, nil
}
