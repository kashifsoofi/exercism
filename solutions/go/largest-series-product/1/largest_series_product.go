// Package lsproduct implements utility routine to find largest serries prodcut from a string of digits
package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct returns largest series product
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
		return 0, errors.New("invalid argument")
	}

	numbers := make([]int64, len(digits))
	for i, d := range digits {
		if !unicode.IsDigit(d) {
			return 0, errors.New("invalid argument")
		}
		numbers[i] = int64(d - '0')
	}

	if span == 0 {
		return 1, nil
	}

	var lsp int64
	for i := 0; i < len(numbers)-span+1; i++ {
		p := numbers[i]
		for j := 1; j < span; j++ {
			p *= numbers[j+i]
		}

		if lsp < p {
			lsp = p
		}
	}
	return lsp, nil
}
