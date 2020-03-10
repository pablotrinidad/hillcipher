package hillcipher

import "math"

// Residue returns the residue of a modulo m
func Residue(a, m int) int {
	reminder := int(math.Abs(float64(a % m)))
	if a >= 0 {
		return reminder
	} else if a < 0 && reminder != 0 {
		return m - reminder
	}
	return 0
}

// EGCD computes the Bezouts identity using the Extended Euclidean Algorithm.
// Assumes a <= b and returns x,y,g such that g=gcd(a,b) and ax + by = g.
func EGCD(a, b int) (x, y, g int) {
	x1, y1, q := 1, 0, 0
	x, y, g = 0, 1, 0
	for a != 0 {
		q, b, a = b/a, a, b%a
		y, y1 = y1, y-q*y1
		x, x1 = x1, x-q*x1
	}
	return x, y, b
}

// IsModUnit returns a boolean indicating whether an element of Zn is a unit in Zn.
// Assumes a in Zn, that is, a in [0, n-1]
func IsModUnit(a, n int) bool {
	_, _, gcd := EGCD(a, n)
	return gcd == 1
}
