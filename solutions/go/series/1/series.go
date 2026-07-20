// Package series implements a simple library that
// implements methods to calculate age on different planets
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (series []string) {
	for i, j := 0, n; j <= len(s); i, j = i+1, j+1 {
		first := UnsafeFirst(n, s[i:])
		series = append(series, first)
		// if first, ok := UnsafeFirst(n, s); ok {
		// 	s = s[1:]
		// 	series = append(series, first)
		// }
	}
	return series
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	ok = n <= len(s)
	if ok {
		first = UnsafeFirst(n, s)
	}
	return first, ok
}
