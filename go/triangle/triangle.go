package triangle

import (
	"math"
)

// Kind represents the type of triangle
type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// KindFromSides determines what kind of triagele it is based on the lengths
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if a == b && b == c && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}
