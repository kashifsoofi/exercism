// Package listops implements a simple library that
// implements methods to perform basic list operations
package listops

// IntList type for list of integers
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

/*
func binFunc(i, j int) bool {

}
*/

// Foldr method folds list values from right to left with initial value and returns the final result
func (l IntList) Foldr(fn binFunc, initial int) int {
	r := initial
	rl := l.Reverse()
	for _, v := range rl {
		r = fn(v, r)
	}
	return r
}

// Foldl method folds list values from left to right with initial value and returns the final result
func (l IntList) Foldl(fn binFunc, initial int) int {
	result := initial
	for _, v := range l {
		result = fn(result, v)
	}
	return result
}

// Filter method filters IntList and return new list
func (l IntList) Filter(fn predFunc) IntList {
	fl := make([]int, 0)
	for _, v := range l {
		if fn(v) {
			fl = append(fl, v)
		}
	}
	return fl
}

// Length method returns length of IntList
func (l IntList) Length() int {
	return len(l)
}

// Map method applies performs operation on each value of list and return new list
func (l IntList) Map(fn unaryFunc) IntList {
	lm := make([]int, len(l))
	for i, v := range l {
		lm[i] = fn(v)
	}
	return lm
}

// Append appends two IntLists and returns new list
func (l IntList) Append(l1 IntList) IntList {
	lc := len(l)
	al := make([]int, lc+len(l1))
	for i, v := range l {
		al[i] = v
	}

	for i, v := range l1 {
		al[lc+i] = v
	}
	return al
}

// Reverse reverses a list and return new list
func (l IntList) Reverse() IntList {
	lc := len(l)
	rl := make([]int, lc)
	for i, v := range l {
		rl[lc-i-1] = v
	}
	return rl
}

// Concat concatenates lists and returns new list
func (l IntList) Concat(lists []IntList) IntList {
	cl := IntList{}.Append(l)
	for _, l1 := range lists {
		cl = cl.Append(l1)
	}
	return cl
}
