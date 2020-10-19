// Package summultiples implemnts utility method to calculate sum of multiples
package summultiples

// SumMultiples sums of all the unique multiples of factors up to but not including limit.
func SumMultiples(limit int, factors ...int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		for _, f := range factors {
			if f == 0 {
				continue
			}
			if i%f == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
