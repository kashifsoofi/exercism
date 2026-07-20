// Package binarysearch implements routing to perform binary search
package binarysearch

// SearchInts return index of key in an array
func SearchInts(a []int, k int) int {
	l, r := 0, len(a)-1
	for l <= r {
		m := l + (r-l)/2
		if a[m] == k {
			return m
		}

		if a[m] < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
