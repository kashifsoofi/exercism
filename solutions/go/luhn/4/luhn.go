// Package luhn implements a simple library that implements
// method to validate number using luhn formula
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns square of sum of first n natural numbers, if successful
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) < 2 {
		return false
	}

	var sum int
	double := len(input)%2 == 0
	for _, r := range input {
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
