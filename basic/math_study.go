package basic

import "math"

func MathStudy() {
	a, b, c := 2.0, 1.0, 0.0
	x, y := a/c, b/c     // infinity
	n := math.NaN()      // not a number
	m := math.Sqrt(-1.0) // not a number
	println(x == y, m == n)
}
