// Package acronym implements a utility routine to generate acronym
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns acronym of passed string
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.Is(unicode.Quotation_Mark, r)
	})
	a := ""
	for _, w := range words {
		a += strings.ToUpper(w[0:1])
	}
	return a
}
