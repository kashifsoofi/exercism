// Package prime implements simple routine to find nth prime
package prime

// Nth returns nth prime
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	if n == 1 {
		return 2, true
	}

	p := 3
	for i, c := 3, 2; c <= n; i += 2 {
		if isPrime(i) {
			p = i
			c++
		}
	}
	return p, true
}

func isPrime(n int) bool {
	if n > 2 && n%2 == 0 {
		return false
	}

	l := n/2 + 1
	for i := 3; i < l; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
