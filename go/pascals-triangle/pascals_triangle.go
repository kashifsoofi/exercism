// Package pascal implements a simple library that implements
// method to generate n-depth Pascal's Triangle
package pascal

// Triangle given a number n returns n-depth Pascal's Triangle
func Triangle(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	t := Triangle(n - 1)
	r := row(n, t[n-2])
	t = append(t, r)
	return t
}

func row(n int, p []int) []int {
	l := len(p) + 1
	r := make([]int, l)
	r[0], r[l-1] = 1, 1
	for i := 1; i < l-1; i++ {
		r[i] = p[i-1] + p[i]
	}
	return r
}
