// Package sieve implements simple routine for Sieve of Eratosthenes
package sieve

// Sieve returns all prime numbers using Sieve of Eratosthenes
func Sieve(l int) []int {
	if l < 2 {
		return []int{}
	}

	m := make([]bool, l+1)
	p := []int{}
	for i := 2; i <= l; i++ {
		if !m[i] {
			p = append(p, i)
			for j := 2; i*j <= l; j++ {
				m[i*j] = true
			}
		}
	}

	return p
}
