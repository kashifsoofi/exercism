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

	var sum = 0
	for i := len(runes) - 1; i >= 0; i -= 2 {
		if !unicode.IsDigit(runes[i]) {
			return false
		}

		n1, _ := strconv.Atoi(string(runes[i]))
		sum += n1

		if i > 0 {
			if !unicode.IsDigit(runes[i-1]) {
				return false
			}

			n2, _ := strconv.Atoi(string(runes[i-1]))
			var double = n2 * 2
			if double > 9 {
				double -= 9
			}
			sum += double
		}
	}

	return sum%10 == 0
}
