// Package leap implements utility routine for checking leap year.
package leap

// IsLeapYear returns, if given value is a leap year, true
func IsLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}

	return false
}
