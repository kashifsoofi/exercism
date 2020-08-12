// Package isogram implements a simple library that
// determines if a word or phrase is an isogram
package isogram

import (
	"unicode"
)

// IsIsogram returns true if given word is an isogram, if successful
func IsIsogram(word string) bool {
	letterVisitMap := make(map[rune]bool)
	for _, letter := range word {
		lowerCaseLetter := unicode.ToLower(letter)

		if lowerCaseLetter == ' ' || lowerCaseLetter == '-' {
			continue
		}

		letterVisited := letterVisitMap[lowerCaseLetter]
		if letterVisited {
			return false
		}

		letterVisitMap[lowerCaseLetter] = true
	}

	return true
}
