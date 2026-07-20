// Package grains implements a simple library for calculating
// number of grains that can be places on a square of chess board.
package grains

import (
	"errors"
)

// Square calculats the number of grains in the given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("number should be between 1 and 64")
	}

	return uint64(1) << uint(n-1), nil
}

// Total calculats the toal number of grains on a chess board
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}

	return total
}
