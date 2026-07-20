// Package bob implements a simple library that implements
// simple routine to calculate conversation responses for bob
package bob

import (
	"strings"
	"unicode"
)

// Hey returns converstaion response appropriate for remark, if successful.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(remark, "?")
	hasLetters := strings.IndexFunc(remark, unicode.IsLetter) >= 0
	isUppercase := strings.Compare(remark, strings.ToUpper(remark)) == 0
	isYelling := hasLetters && isUppercase

	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYelling:
		return "Whoa, chill out!"
	}

	return "Whatever."
}
