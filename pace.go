package fitbuddy

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// Pace refers to the rate at which a runner covers a certain distance,
// usually measured in minutes per mile or kilometers per hour.
type Pace struct {
	Value time.Duration
	Unit  DistanceUnitType
}

// String get pace per distance unit
func (p Pace) String() string {
	var h, m int
	duration := p.Value
	items := make([]string, 0)

	if h = int(duration.Hours()); h > 0 {
		items = append(items, strconv.Itoa(h))
		duration -= time.Hour * time.Duration(h)
	}

	m = int(duration.Minutes())
	duration -= time.Minute * time.Duration(m)

	if h > 0 {
		items = append(items, fmt.Sprintf("%02d", m))
	} else {
		items = append(items, strconv.Itoa(m))
	}

	items = append(items, fmt.Sprintf("%02d", int(duration.Seconds())))

	return fmt.Sprintf("%s /%s", strings.Join(items, ":"), p.Unit.String())
}

// GetPace Run's pace refers to the speed at which a runner is running, usually measured in minutes per mile or
// kilometers per hour. It is a measure of how long it takes a runner to cover a certain distance,
// and is commonly used to help runners set goals and track their progress in training and racing.
func GetPace(distance Distance, duration time.Duration) Pace {
	paceValue := math.Round(duration.Seconds() / distance.Value)
	paceDuration := time.Second * time.Duration(paceValue)

	return Pace{
		Value: paceDuration,
		Unit:  distance.Unit,
	}
}
