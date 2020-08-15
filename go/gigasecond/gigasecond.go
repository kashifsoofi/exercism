// Package gigasecond implements a simple library that
// determine the moment that would be after a gigasecond has passed
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns next moment after a gigasecond has passed.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
