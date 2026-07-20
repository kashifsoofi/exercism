// Package luhn implements a simple library that implements
// method to validate number using luhn formula
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid returns square of sum of first n natural numbers, if successful
func Valid(input string) bool {
	runes := []rune(strings.ReplaceAll(input, " ", ""))
	if len(runes) < 2 {
		return false
	}

	var sum, double = 0, false
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}

		digit, _ := strconv.Atoi(string(runes[i]))

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
