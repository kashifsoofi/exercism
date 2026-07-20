// Package triangle implements a simple library that
// determins the type of a triangle given its sides
package triangle

import (
	"math"
)

// Kind of triangle
type Kind int

// Kinds of triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides return type of triangle given the sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if (a+b) < c || (a+c) < b || (b+c) < a {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
