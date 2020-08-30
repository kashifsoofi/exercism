// Package collatzconjecture implements utility routine for calculation steps required to reach 1
package collatzconjecture

import "errors"

// CollatzConjecture returns steps to reach 1 if successful
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid number")
	}

	steps := 0
	for ; n != 1; steps++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}

	return steps, nil
}
