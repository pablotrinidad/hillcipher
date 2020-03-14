// Package cipher implements the Hill Cipher algorithm. Hill Cipher
// is a polygraphic substitution cipher that encrypts messages through
// matrix transformations. Matrix operations and modular arithmetic
// definitions are implemented and exported as part of the package.
// Hill Cipher is broken by cipher-text only attacks. DON'T use this
// package for any purpose other than experimentation and educational
// purposes.
package cipher

import (
	"fmt"
	"math"
	"strings"
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
	return &Cipher{mod: n, alphabet: *alphabet}, nil
}

// verifyKeyTextPair makes sure key and text are usable in the current cipher.
// Returns Key and message if valid.
func (c *Cipher) verifyKeyTextPair(rawM, rawK string) (*Matrix, []rune, error) {
	if !c.alphabet.Belongs(rawM) {
		return nil, nil, fmt.Errorf("message %q does not belong to alphabet %q", rawM, c.alphabet)
	}
	if !c.alphabet.Belongs(rawK) {
		return nil, nil, fmt.Errorf("key %q does not belong to alphabet %q", rawK, c.alphabet)
	}
	msg, k := []rune(rawM), []rune(rawK)
	kInt := make([]int, len(k))
	for i, s := range k {
		kInt[i], _ = c.alphabet.Stoi(s) // Neglect error because key is permutation of alphabet
	}
	key, err := NewKey(kInt, c.mod)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create key for %q; %v", k, err)
	}
	if len(msg)%key.order != 0 {
		return nil, nil, fmt.Errorf("message length is not multiple of key's length, consider using EncryptWithPadding if encrypting message")
	}
	mKey := Matrix(*key)
	return &mKey, msg, nil
}

// performOperations apply cipher matrix operations on the given key and text. It assume all
// key and message validations were applied before. Returns the resulting string.
func (c *Cipher) performOperations(key *Matrix, msg []rune) string {
	// Use builder for optimum string creation
	var result strings.Builder

	for i := 0; i < len(msg); i += key.order {
		vector := make([]int, key.order)
		for j, r := range msg[i : i+key.order] {
			vector[j], _ = c.alphabet.Stoi(r) // Neglect error because message is permutation of alphabet.
		}
		prodVector, _ := key.VectorProductMod(c.mod, vector...) // Neglect error because size is exact
		for _, ri := range prodVector {
			r, _ := c.alphabet.Itos(ri) // Neglect error because mod operation
			result.WriteRune(r)
		}
	}
	return result.String()
}

// Encrypt plain text using given key. Returns an error if either key or message don't belong
// to the cipher's alphabet, if key is not invertible by cipher's modulo or if message length
// is not multiple of key's order (matrix order).
func (c *Cipher) Encrypt(rawM, rawK string) (string, error) {
	key, msg, err := c.verifyKeyTextPair(rawM, rawK)
	if err != nil {
		return "", err
	}
	return c.performOperations(key, msg), nil
}

// Decrypt cipher text using given key. Returns an error if either key or message don't belong
// to the cipher's alphabet, if key is not invertible by cipher's modulo or if message length
// is not multiple of key's order (matrix order).
func (c *Cipher) Decrypt(rawM, rawK string) (string, error) {
	key, plainText, err := c.verifyKeyTextPair(rawM, rawK)
	if err != nil {
		return "", err
	}

	invertedKey, err := key.InverseMod(c.mod)
	if err != nil {
		return "", fmt.Errorf("failed to invert key\n%s\n; %v", key, err)
	}

	return c.performOperations(invertedKey, plainText), nil
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
	symbols := []rune(s)
	n := len(symbols)
	a := &Alphabet{
		symbols:     symbols,
		intIndex:    make(map[int]rune, n),
		symbolIndex: make(map[rune]int, n),
	}
	for i, s := range symbols {
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
	for _, r := range []rune(s) {
		if !a.Contains(r) {
			return false
		}
	}
	return true
}
