// Package luhn implements a simple library that implements
// method to validate number using luhn formula
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns square of sum of first n natural numbers, if successful
func Valid(input string) bool {
	runes := []rune(strings.ReplaceAll(input, " ", ""))
	if len(runes) < 2 {
		return false
	}

	var sum, double = 0, len(runes)%2 == 0
	for _, r := range runes {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		double = !double

		sum += digit
	}

	return sum%10 == 0
}
