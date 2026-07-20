// Package isogram implements a simple library that
// determines if a word or phrase is an isogram
package isogram

import (
	"strings"
)

// IsIsogram returns true if given word is an isogram, if successful
func IsIsogram(word string) bool {
	letterCountMap := make(map[rune]int)
	for _, letter := range strings.ToLower(word) {
		if letter == ' ' || letter == '-' {
			continue
		}

		letterCount := letterCountMap[letter]
		if letterCount > 0 {
			return false
		}

		letterCountMap[letter] = 1
	}

	return true
}
