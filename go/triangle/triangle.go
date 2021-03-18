// Package triangle checks the type of triangle from the length of its sides
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides returns the type of the triangle with given sides a, b, and c
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	// check sides are not NaN
	if math.IsNaN(a) || math.IsInf(a, 0) ||
		math.IsNaN(b) || math.IsInf(b, 0) ||
		math.IsNaN(c) || math.IsInf(c, 0) {
		k = NaT
		// check sides greater than zero
	} else if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
		// check triangle inequality holds
	} else if a > b+c || b > a+c || c > a+b {
		k = NaT
	} else {
		// check if equilateral
		if a == b && b == c {
			k = Equ
			// check if isosceles
		} else if a == b || b == c || c == a {
			k = Iso
			// must be scalene
		} else {
			k = Sca
		}

	}
	return k
}
