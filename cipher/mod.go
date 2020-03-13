package cipher

import (
	"fmt"
	"math"
)

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
	x1, y1, q := 0, 1, 0
	x, y, g = 1, 0, 0
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

// ModularInverse returns the modular invers of a over m, returns error if a is not unit of Zn
func ModularInverse(a, m int) (int, error) {
	if m == 1 {
		return 0, nil
	}
	if _, _, g := EGCD(a, m); g != 1 {
		return -1, fmt.Errorf("%d and %d are not coprimes", a, m)
	}
	m0, x, y := m, 1, 0
	for a > 1 {
		q := a / m
		m, a = a%m, m
		y, x = x-q*y, y
	}
	if x < 0 {
		x += m0
	}
	return x, nil
}
