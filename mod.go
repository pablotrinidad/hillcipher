package hillcipher

import "math"

// residue returns the residue of a modulo m
func residue(a, m int) int {
	reminder := int(math.Abs(float64(a % m)))
	if a >= 0 {
		return reminder
	} else if a < 0 && reminder != 0 {
		return m - reminder
	}
	return 0
}

// egcd computes ax + by = gcd(a, b) for any given a+b integers. Assumes a <= b.
func egcd(a, b int) (x, y, g int) {
	x1, y1, q := 1, 0, 0
	y = 1
	for a != 0 {
		q, b, a = b/a, a, b%a
		y, y1 = y1, y-q*y1
		x, x1 = x1, x-q*x1
	}
	return x, y, b
}
