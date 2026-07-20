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
	switch {
	case isNotATriangle(a, b, c):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}

func isNotATriangle(a, b, c float64) bool {
	isShortSide := (a+b) < c || (a+c) < b || (b+c) < a
	return isInvalidSide(a) || isInvalidSide(b) || isInvalidSide(c) || isShortSide
}

func isInvalidSide(side float64) bool {
	return math.IsNaN(side) || math.IsInf(side, 0) || side <= 0
}
