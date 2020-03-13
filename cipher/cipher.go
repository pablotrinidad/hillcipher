// Package cipher implements the Hill Cipher algorithm. the Hill cipher
// is a polygraphic substitution cipher that encrypts messages through
// matrix transformations. Matrix operations and modular arithmetic
// definitions are implemented and exported as part of the package.
package cipher

import (
	"fmt"
	"math"
)

// Key represents a Hill Cipher key matrix
type Key Matrix

// Alphabet is the set of symbols valid through a cipher
type Alphabet []rune

// Cipher is an instance of the Hill Cipher on an specific alphabet
type Cipher struct {
	mod      int
	alphabet Alphabet
}

// String makes Key implement Stringer
func (k Key) String() string {
	return Matrix(k).String()
}

// NewKey initializes a Hill Cipher in an specific modulo
func NewKey(k []int, mod int) (*Key, error) {
	if mod < 2 {
		return nil, fmt.Errorf("cannot create key for mod %d < 2", mod)
	}
	sqr := math.Sqrt(float64(len(k)))
	if sqr-math.Floor(sqr) != 0 {
		return nil, fmt.Errorf("key size must be a square number, got %d", len(k))
	}
	if int(sqr) < 2 {
		return nil, fmt.Errorf("cannot create key of order %d < 2", int(sqr))
	}
	m, _ := NewMatrix(int(sqr), k) // Error is neglected since order is square
	if !m.IsInvertibleMod(mod) {
		return nil, fmt.Errorf("key is not invertible modulo %d", mod)
	}
	key := Key(*m)
	return &key, nil
}
