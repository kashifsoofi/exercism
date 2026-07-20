// Package hamming implements a simple library for calculating hamming distance.
package hamming

import (
	"errors"
)

// Distance returns distance between two strands, if successful, error if distance cannot be calculated
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("Sequences are not of equal length")
	}

	var distance = 0
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}

	return distance, nil
}
