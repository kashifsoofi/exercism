// Package armstrong implements a utility method to check if a number is armstrong number
package armstrong

import (
	"math"
)

// IsNumber returns true if number is an armstrong number
func IsNumber(n int) bool {
	digits := []int{}
	for i := n; i > 0; i /= 10 {
		digits = append(digits, i%10)
	}

	p := len(digits)
	sum := 0
	for _, d := range digits {
		sum += int(math.Pow(float64(d), float64(p)))
	}
	return n == sum
}
