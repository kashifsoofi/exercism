// Package strain implements utility routines
// to filter collection using a predicate
package strain

// Ints list of integers
type Ints []int

// Lists list of integers of integers
type Lists [][]int

// Strings list of strings
type Strings []string

// Keep returns collection of int elements for which predicate is true
func (col Ints) Keep(predicate func(int) bool) (res Ints) {
	for _, e := range col {
		if predicate(e) {
			res = append(res, e)
		}
	}
	return
}

// Discard returns collection of elements for which predicate is false
func (col Ints) Discard(predicate func(int) bool) (res Ints) {
	res = col.Keep(func(e int) bool {
		return !predicate(e)
	})
	return
}

// Keep returns collection of string elements for which predicate is true
func (col Strings) Keep(predicate func(string) bool) (res Strings) {
	for _, e := range col {
		if predicate(e) {
			res = append(res, e)
		}
	}
	return
}

// Keep returns collection of string elements for which predicate is true
func (col Lists) Keep(predicate func([]int) bool) (res Lists) {
	for _, e := range col {
		if predicate(e) {
			res = append(res, e)
		}
	}
	return
}
