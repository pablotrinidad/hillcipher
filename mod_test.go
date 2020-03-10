package hillcipher

import (
	"fmt"
	"testing"
)

// TestResidue check definition of Residue function
func TestResidue(t *testing.T) {
	tests := []struct {
		a, m, r int
	}{
		{a: 87, m: 26, r: 9},
		{a: -38, m: 26, r: 14},
		{a: -26, m: 26, r: 0},
	}
	for _, test := range tests {
		name := fmt.Sprintf("Reminder(a:%d, m:%d)", test.a, test.m)
		t.Run(name, func(t *testing.T) {
			r := Residue(test.a, test.m)
			if r != test.r {
				t.Errorf("%s = %d, want %d", name, r, test.r)
			}
		})
	}
}

// TestEGCD verify correc implementation of Extended Euclidean Algorithm
func TestEGCD(t *testing.T) {
	tests := []struct {
		a, b, g, x, y int
	}{
		{
			a: 150, b: 180,
			g: 30, x: -1, y: 1,
		},
		{
			a: 18, b: 348,
			g: 6, x: -19, y: 1,
		},
		{
			a: 24, b: 60,
			g: 12, x: -2, y: 1,
		},
		{
			a: 148, b: 772,
			g: 4, x: -73, y: 14,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("EGCD(a:%d, b:%d)", test.a, test.b)
		t.Run(name, func(t *testing.T) {
			x, y, g := EGCD(test.a, test.b)
			if x != test.x || y != test.y || g != test.g {
				t.Errorf("%s = x:%d, y:%d, g:%d, want x:%d, y:%d, g:%d", name, x, y, g, test.x, test.y, test.g)
			}
		})
	}
}

// TestIsModUnit verify the definition of modular unit
func TestIsModUnit(t *testing.T) {
	tests := []struct {
		a, n   int
		isUnit bool
	}{
		{a: 14, n: 15, isUnit: true},
		{a: 3, n: 12, isUnit: false},
		{a: 1, n: 1, isUnit: true},
		{a: 51, n: 482, isUnit: true},
		{a: 52, n: 482, isUnit: false},
	}
	for _, test := range tests {
		name := fmt.Sprintf("IsModUnit(a:%d, n:%d)", test.a, test.n)
		t.Run(name, func(t *testing.T) {
			if r := IsModUnit(test.a, test.n); r != test.isUnit {
				t.Errorf("%s = %v, want %v", name, r, test.isUnit)
			}
		})
	}
}
