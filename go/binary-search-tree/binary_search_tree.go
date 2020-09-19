// Package binarysearchtree implements simple binary search tree
package binarysearchtree

// Bst holds data and pointers to left and right nodes
// type Bst struct {
// }

// SearchTreeData holds data and pointers to left and right nodes
type SearchTreeData struct {
	data  int
	left  *SearchTreeData
	right *SearchTreeData
}

// Bst returns a new SearchTreeData
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{
		data: data,
	}
}

// Insert inserts new data in binary search tree
func (b *SearchTreeData) Insert(data int) {
	insert(b, data)
}

func insert(n *SearchTreeData, data int) *SearchTreeData {
	if n == nil {
		return Bst(data)
	}

	if data <= n.data {
		n.left = insert(n.left, data)
	} else if data > n.data {
		n.right = insert(n.right, data)
	}
	return n
}

// MapString returns data as array of strings
func (b *SearchTreeData) MapString(f func(int) string) []string {
	a := []string{}
	if b == nil {
		return a
	}

	a = append(a, b.left.MapString(f)...)
	a = append(a, f(b.data))
	a = append(a, b.right.MapString(f)...)
	return a
}

// MapInt returns data as array of integers
func (b *SearchTreeData) MapInt(f func(int) int) []int {
	a := []int{}
	if b == nil {
		return a
	}

	a = append(a, b.left.MapInt(f)...)
	a = append(a, f(b.data))
	a = append(a, b.right.MapInt(f)...)
	return a
}
