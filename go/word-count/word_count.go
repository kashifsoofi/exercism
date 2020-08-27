// Package wordcount implements a simple library for counting words in a phrase.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is frequency of words in phrase
type Frequency map[string]int

// WordCount returns a map of words and their count in a phrase.
func WordCount(phrase string) Frequency {
	words := strings.FieldsFunc(strings.ToLower(phrase), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	f := Frequency{}
	for _, w := range words {
		f[strings.Trim(w, "'")]++
	}
	return f
}
