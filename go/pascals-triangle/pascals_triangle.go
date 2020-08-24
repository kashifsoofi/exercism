// Package pascal implements a simple library that implements
// method to generate n-depth Pascal's Triangle
package pascal

// Triangle given a number n returns n-depth Pascal's Triangle
func Triangle(n int) [][]int {
	t := make([][]int, n)
	t[0] = []int{1}
	for i := 1; i < n; i++ {
		t[i] = getRow(t[i-1])
	}
	return t
}

func getRow(p []int) []int {
	l := len(p) + 1
	r := make([]int, l)
	r[0], r[l-1] = 1, 1
	for i := 1; i < l-1; i++ {
		r[i] = p[i-1] + p[i]
	}
	return r
}
