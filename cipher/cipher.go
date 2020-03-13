// Package cipher implements the Hill Cipher algorithm. the Hill cipher
// is a polygraphic substitution cipher that encrypts messages through
// matrix transformations. Matrix operations and modular arithmetic
// definitions are implemented and exported as part of the package.
package cipher

import (
	"fmt"
	"math"
)

// Cipher is an instance of the Hill Cipher on an specific alphabet
type Cipher struct {
	mod      int
	alphabet Alphabet
}

// NewCipher initializes new cipher ready for given alphabet
func NewCipher(alphabet *Alphabet) (*Cipher, error) {
	n := len(alphabet.Symbols())
	if n < 2 {
		return nil, fmt.Errorf("alphabet must contain at least 2 symbols, got %d", n)
	}
	return &Cipher{mod: n - 1, alphabet: *alphabet}, nil
}

// Encrypt ciphers plain text using given key
func (c *Cipher) Encrypt(msg, k string) (string, error) {
	if !c.alphabet.Belongs(msg) {
		return "", fmt.Errorf("message %q does not belong to alphabet %q", msg, c.alphabet)
	}
	if !c.alphabet.Belongs(k) {
		return "", fmt.Errorf("key %q does not belong to alphabet %q", k, c.alphabet)
	}
	if len(msg) < len(k) || len(msg)%len(k) != 0 {
		return "", fmt.Errorf("message length is not multiple of key's length, consider using EncryptWithPadding")
	}

	kInt := make([]int, len(k))
	for i, s := range k {
		kInt[i], _ = c.alphabet.Stoi(s) // Neglect error because key is permutation of alphabet
	}
	_, err := NewKey(kInt, c.mod)
	if err != nil {
		return "", fmt.Errorf("failed to create key for %q; %v", k, err)
	}
	return "", nil
}

// Key represents a Hill Cipher key matrix
type Key Matrix

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

// Alphabet is the set of symbols valid through a cipher
type Alphabet struct {
	symbols     []rune
	symbolIndex map[rune]int
	intIndex    map[int]rune
}

// NewAlphabet initializes a new Hill Cipher Alphabet.
func NewAlphabet(s string) *Alphabet {
	n := len(s)
	a := &Alphabet{
		symbols:     make([]rune, n),
		intIndex:    make(map[int]rune, n),
		symbolIndex: make(map[rune]int, n),
	}
	for i, s := range s {
		a.symbols[i] = s
		a.intIndex[i] = s
		a.symbolIndex[s] = i
	}
	return a
}

// String makes Alphabet implement Stringer.
func (a Alphabet) String() string {
	return string(a.symbols)
}

// Symbols returns the alphabet's symbols.
func (a *Alphabet) Symbols() []rune {
	return a.symbols
}

// Contains returns wether r is defined in alphabet.
func (a *Alphabet) Contains(r rune) bool {
	_, found := a.symbolIndex[r]
	return found
}

// Stoi returns the int value of the given symbol s (Symbol To Int).
func (a *Alphabet) Stoi(s rune) (int, error) {
	if !a.Contains(s) {
		return -1, fmt.Errorf("symbols %q is not part of the alphabet", s)
	}
	return a.symbolIndex[s], nil
}

// Itos returns the symbol value of the given int i (Int To Symbol).
func (a *Alphabet) Itos(i int) (rune, error) {
	r, found := a.intIndex[i]
	if !found {
		return 'x', fmt.Errorf("%d cannot be mapped to symbol", i)
	}
	return r, nil
}

// Belongs returns whether a string belongs to the alphabet or not.
func (a *Alphabet) Belongs(s string) bool {
	for _, r := range s {
		if !a.Contains(r) {
			return false
		}
	}
	return true
}
