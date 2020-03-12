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
			g: 30, x: 1, y: -1,
		},
		{
			a: 18, b: 348,
			g: 6, x: 1, y: -19,
		},
		{
			a: 24, b: 60,
			g: 12, x: 1, y: -2,
		},
		{
			a: 148, b: 772,
			g: 4, x: 14, y: -73,
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

// TestModularInverse verify the definition of modular inverse
func TestModularInverse(t *testing.T) {
	tests := []struct {
		input map[int]int // map[unit]inverse
		mod   int
	}{
		{
			input: map[int]int{
				1:  1,
				3:  9,
				5:  21,
				7:  15,
				9:  3,
				11: 19,
				15: 7,
				17: 23,
				19: 11,
				21: 5,
				23: 17,
				25: 25,
			},
			mod: 26,
		},
		{
			input: map[int]int{
				1:  1,
				3:  27,
				7:  23,
				9:  9,
				11: 51,
				13: 37,
				17: 33,
				19: 59,
				21: 61,
				23: 7,
				27: 3,
				29: 69,
			},
			mod: 80,
		},
		{
			input: map[int]int{
				1:  1,
				5:  5,
				7:  7,
				11: 11,
			},
			mod: 12,
		},
		{
			input: map[int]int{
				1:  1,
				2:  14,
				4:  7,
				5:  11,
				7:  4,
				8:  17,
				10: 19,
				11: 5,
				13: 25,
				14: 2,
				16: 22,
				17: 8,
				19: 10,
				20: 23,
				22: 16,
				23: 20,
				25: 13,
				26: 26,
			},
			mod: 27,
		},
		{input: map[int]int{1: 0}, mod: 1},
	}
	for _, test := range tests {
		for u, i := range test.input {
			name := fmt.Sprintf("%x^-1 (mod %d)", u, test.mod)
			t.Run(name, func(t *testing.T) {
				inverse, err := ModularInverse(u, test.mod)
				if err != nil {
					t.Fatalf("ModularInverse(%d, %d) returned unexpected error; %v", u, test.mod, err)
				}
				if inverse != i {
					t.Errorf("ModularInverse(%d, %d) = %d, want %d", u, test.mod, inverse, i)
				}
			})
		}
	}
}

// TestModularInverse_Error verify definition rejects pairs that aren't coprimes
func TestModularInverse_Error(t *testing.T) {
	tests := []struct {
		factors []int
		mod     int
	}{
		{
			factors: []int{2, 4, 6, 8, 10, 12, 13, 14, 16, 18, 20, 22},
			mod:     26,
		},
		{
			factors: []int{2, 3, 4, 6, 8, 9, 10},
			mod:     12,
		},
		{
			factors: []int{3, 6, 9, 12, 15, 18, 21, 24},
			mod:     27,
		},
	}
	for _, test := range tests {
		for _, f := range test.factors {
			name := fmt.Sprintf("%d^-1 (mod %d)", f, test.mod)
			t.Run(name, func(t *testing.T) {
				if _, err := ModularInverse(f, test.mod); err == nil {
					t.Fatalf("ModularInverse(%d, %d) did not fail, should've failed", f, test.mod)
				}
			})
		}
	}
}
