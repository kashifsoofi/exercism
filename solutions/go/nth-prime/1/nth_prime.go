// Package prime implements simple routine to find nth prime
package prime

// Nth returns nth prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	p := 2
	for i, c := 2, 1; c <= n; i++ {
		if isPrime(i) {
			p = i
			c++
		}
	}
	return p, true
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
