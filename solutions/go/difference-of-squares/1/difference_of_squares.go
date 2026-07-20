// Package difference of squares implements a simple library that implements
// methods to calculates square of sum, sum of squrares and their difference
package diffsquares

// SquareOfSum returns square of sum of first n natural numbers, if successful
func SquareOfSum(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns sum of squares of first n natural numbers, if successful
func SumOfSquares(n int) int {
	var sumOfSquares int = 0
	for i := 1; i <= n; i++ {
		sumOfSquares += (i * i)
	}
	return sumOfSquares
}

// Difference returns difference between square of sums and sum of squares of first n natural numbers, if successful
func Difference(n int) int {
	squareOfSum := SquareOfSum(n)
	sumOfSquares := SumOfSquares(n)
	return squareOfSum - sumOfSquares
}
