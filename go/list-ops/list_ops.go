// Package listops implements a simple library that
// implements methods to perform basic list operations
package listops

// IntList type for list of integers
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Foldr given a list, func and initial value
// reduces each item into initial value from right to left using fn
// and returns final result
func (l IntList) Foldr(fn binFunc, n int) int {
	r := n
	for i := l.Length() - 1; i >= 0; i-- {
		r = fn(l[i], r)
	}
	return r
}

// Foldl given a list, func and initial value
// reduces each item into initial value from left to right using fn
// and returns final result
func (l IntList) Foldl(fn binFunc, n int) int {
	result := n
	for _, v := range l {
		result = fn(result, v)
	}
	return result
}

// Filter method filters IntList and return new list
func (l IntList) Filter(fn predFunc) IntList {
	fl := make(IntList, 0, l.Length())
	i := 0
	for _, v := range l {
		if fn(v) {
			fl = fl[:i+1]
			fl[i] = v
			i++
		}
	}
	return fl
}

// Length method returns length of IntList
func (l IntList) Length() (count int) {
	for range l {
		count++
	}
	return count
}

// Map method performs operation on each element of list and return new list
func (l IntList) Map(fn unaryFunc) IntList {
	lm := make(IntList, l.Length())
	for i, v := range l {
		lm[i] = fn(v)
	}
	return lm
}

// Append takes a list, appends it to current list and returns new list
func (l IntList) Append(l1 IntList) IntList {
	lc := l.Length()
	al := make(IntList, lc+l1.Length())
	for i, v := range l {
		al[i] = v
	}

	for i, v := range l1 {
		al[lc+i] = v
	}
	return al
}

// Reverse reverses the order of list and return new list
func (l IntList) Reverse() IntList {
	lc := l.Length()
	rl := make(IntList, lc)
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
