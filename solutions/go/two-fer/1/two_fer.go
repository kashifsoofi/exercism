// Package twofer implements a simple library for sharing quote.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"strings"
)

// ShareWith returns quote for name, returns default if name is not provided.
func ShareWith(name string) string {
	var twoFer strings.Builder
	twoFer.WriteString("One for ")
	if (len(name) == 0) {
		twoFer.WriteString("you")
	} else {
		twoFer.WriteString(name)
	}
	twoFer.WriteString(", one for me.")

	return twoFer.String()
}
