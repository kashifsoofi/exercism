// Package meetup implements a simple library for calculating date of meetup.
package meetup

import "time"

// WeekSchedule is the.
type WeekSchedule int

const (
	// First day
	First WeekSchedule = iota
	// Second day
	Second
	// Third day
	Third
	// Fourth day
	Fourth
	// Fifth day
	Fifth
	// Last day
	Last
	// Teenth day
	Teenth
)

// Day returns scheduled day of month for the meetup
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	matchedDays := make([]int, 7)

	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	s := First
	for d.Month() == month {
		if d.Weekday() == weekday {
			day := d.Day()
			matchedDays[s] = day
			s++

			if day > 12 && day < 20 {
				matchedDays[Teenth] = day
			}

			d = d.AddDate(0, 0, 7)
			if d.Month() != month {
				matchedDays[Last] = day
			}
		} else {
			d = d.AddDate(0, 0, 1)
		}
	}

	return matchedDays[schedule]
}
