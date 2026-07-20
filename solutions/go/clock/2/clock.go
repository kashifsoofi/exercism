// Package clock implements a simple library that implements
// methods to handle times without dates
package clock

import (
	"fmt"
)

// Clock struct to keep time in minutes
type Clock struct {
	timeInMinutes int
}

// New returns a new instance of Clock
func New(hours, minutes int) Clock {
	maxMinutes := 24 * 60
	totalMinutes := hours*60 + minutes
	return Clock{
		timeInMinutes: (maxMinutes + totalMinutes%maxMinutes) % maxMinutes,
	}
}

// Add minutes to a given Clock and returns a new instance
func (c Clock) Add(minutesToAdd int) Clock {
	return New(c.hours(), c.minutes()+minutesToAdd)
}

// Subtract minutes from a given Clock and returns a new instance
func (c Clock) Subtract(minutesToSubtract int) Clock {
	return c.Add(-minutesToSubtract)
}

// String returns time in HH:mm format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours(), c.minutes())
}

func (c Clock) hours() int {
	return c.timeInMinutes / 60
}

func (c Clock) minutes() int {
	return c.timeInMinutes % 60
}
