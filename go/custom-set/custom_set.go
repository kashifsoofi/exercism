// Package stringset implemnts custom Set
package stringset

import (
	"fmt"
	"strings"
)

// Set holds custom set
type Set struct {
	items []string
}

// New returns new Set
func New() Set {
	return Set{
		items: []string{},
	}
}

// NewFromSlice returns new Set with items
func NewFromSlice(items []string) Set {
	s := New()
	for _, i := range items {
		s.Add(i)
	}

	return s
}

// String returns string representation of Set
func (s Set) String() string {
	output := strings.Builder{}
	output.WriteString("{")
	for i, item := range s.items {
		output.WriteString(fmt.Sprintf("%q", item))
		if i+1 < len(s.items) {
			output.WriteString(", ")
		}
	}
	output.WriteString("}")
	return output.String()
}

// IsEmpty returns true if Set is empty
func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

// Has returns true if element is in Set
func (s Set) Has(i string) bool {
	for _, item := range s.items {
		if item == i {
			return true
		}
	}
	return false
}

// Add adds a new item to Set
func (s *Set) Add(i string) {
	if s.Has(i) {
		return
	}

	s.items = append(s.items, i)
}

// Subset return true if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for _, i := range s1.items {
		if !s2.Has(i) {
			return false
		}
	}
	return true
}

// Disjoint returns true if no element of s1 is in s2
func Disjoint(s1, s2 Set) bool {
	for _, i := range s1.items {
		if s2.Has(i) {
			return false
		}
	}
	return true
}

// Equal retuns true if both sets are equal
func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}

	for _, i := range s1.items {
		if !s2.Has(i) {
			return false
		}
	}

	return true
}

// Intersection returns new set of elements from s1 that are also in s2
func Intersection(s1, s2 Set) Set {
	s3 := New()
	for _, i := range s1.items {
		if s2.Has(i) {
			s3.Add(i)
		}
	}
	return s3
}

// Difference returns new set of elements from s1 that are not in s2
func Difference(s1, s2 Set) Set {
	s3 := New()
	for _, i := range s1.items {
		if !s2.Has(i) {
			s3.Add(i)
		}
	}
	return s3
}

// Union returns a new set containing elements of both sets
func Union(s1, s2 Set) Set {
	s3 := New()
	for _, i := range s1.items {
		s3.Add(i)
	}

	for _, i := range s2.items {
		s3.Add(i)
	}
	return s3
}
