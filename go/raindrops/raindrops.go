// Package raindrops implements a simple library that convert
// a number into a string that contains raindrop sounds
package raindrops

import (
	"strconv"
)

// Convert returns string that contains raindrop sounds, if successful
func Convert(n int) string {
	sounds := ""
	if n%3 == 0 {
		sounds += "Pling"
	}
	if n%5 == 0 {
		sounds += "Plang"
	}
	if n%7 == 0 {
		sounds += "Plong"
	}

	if sounds == "" {
		sounds += strconv.Itoa(n)
	}

	return sounds
}
