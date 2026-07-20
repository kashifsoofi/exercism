// Package sieve implements simple routine for Sieve of Eratosthenes
package sieve

// Sieve returns all prime numbers using Sieve of Eratosthenes
func Sieve(l int) []int {
	m := map[int]bool{}
	for i := 2; i <= l; i++ {
		m[i] = false
	}

	for i := 2; i*i <= l; i++ {
		if m[i] == false {
			for j := 2; j*i <= l; j++ {
				m[j*i] = true
			}
		}
	}

	p := []int{}
	for i := 2; i <= l; i++ {
		if m[i] == false {
			p = append(p, i)
		}
	}
	return p
}
