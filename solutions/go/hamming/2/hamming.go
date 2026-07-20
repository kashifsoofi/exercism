// Package hamming implements a simple library for calculating hamming distance.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance returns distance between two strands, if successful, error if distance cannot be calculated
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Sequences are not of equal length")
	}

	var distance = 0
	for i, aCell := range a {
		bCell, _ := utf8.DecodeRuneInString(b[i:])
		if aCell != bCell {
			distance++
		}
	}

	return distance, nil
}
