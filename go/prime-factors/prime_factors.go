// Package prime implements a utility method to calculate all prime factors of a given number
package prime

// Factors returns all prime factors of a given numbers
func Factors(n int64) []int64 {
	f := []int64{}
	for i := int64(2); n > 1; i++ {
		for ; n%i == 0; n = n / i {
			f = append(f, i)
		}
	}
	return f
}
