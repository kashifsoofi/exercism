// Package difference of squares implements a simple library that implements
// methods to calculates square of sum, sum of squrares and their difference
package diffsquares

// SquareOfSum returns square of sum of first n natural numbers, if successful
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns sum of squares of first n natural numbers, if successful
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference returns difference between square of sums and sum of squares of first n natural numbers, if successful
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
