package cipher

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestNewKey correct behavior
func TestNewKey(t *testing.T) {
	tests := []struct {
		name    string
		mod     int
		data    []int
		wantKey *Key
	}{
		{
			name: "order 2 mod 2",
			mod:  2, data: []int{1, 0, 1, 1},
			wantKey: &Key{Order: 2, Data: [][]int{{1, 0}, {1, 1}}},
		},
		{
			name: "order 3 mod 27",
			mod:  27,
			data: []int{5, 15, 18, 20, 0, 11, 4, 26, 0},
			wantKey: &Key{
				Order: 3,
				Data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotKey, err := NewKey(test.data, test.mod)
			if err != nil {
				t.Fatalf("NewKey(%v, %d) returned unexpected error; %v", test.data, test.mod, err)
			}
			if diff := cmp.Diff(test.wantKey, gotKey); diff != "" {
				t.Errorf("NewKey(%v, %d)=\n%s, want \n%s: diff want -> got\n%s", test.data, test.mod, gotKey, test.wantKey, diff)
			}
		})
	}
}

// TestNewKey_Error verify validations are applied
func TestNewKey_Error(t *testing.T) {
	tests := []struct {
		name string
		mod  int
		data []int
	}{
		{name: "mod under 2", mod: 1},
		{name: "non-square number", mod: 2, data: []int{1, 2}},
		{name: "another non-square number", mod: 2, data: []int{1, 2, 3}},
		{name: "square number, order 1", mod: 2, data: []int{1}},
		{
			name: "order 5 non-invertible",
			mod:  50,
			data: []int{6, 24, 44, 1, 15, 13, 16, 48, 10, 23, 20, 20, 17, 15, 23, 1, 2, 9, 13, 0, 48, 47, 46, 45, 44},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := NewKey(test.data, test.mod); err == nil {
				t.Fatalf("NewKey(%v, %d) returned non-nil error, want non-nil", test.data, test.mod)
			}
		})
	}
}

// TestKeyString verify Key mirrors matrix string
func TestKeyString(t *testing.T) {
	tests := []struct {
		name    string
		key     *Key
		wantRep string
	}{
		{
			name: "order 3",
			key: &Key{
				Order: 3,
				Data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantRep: "|	1	|	2	|	3	|\n|	4	|	5	|	6	|\n|	7	|	8	|	9	|\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotRep := fmt.Sprintf("%s", test.key)
			if gotRep != test.wantRep {
				t.Errorf("String(m) = %s, want %s", gotRep, test.wantRep)
			}
		})
	}
}

// TestNewAlphabet verifies correct initialization of alphaber
func TestNewAlphabet(t *testing.T) {
	tests := []struct {
		name, data   string
		wantAlphabet *Alphabet
	}{
		{
			name:         "binary alphabet",
			data:         "01",
			wantAlphabet: &Alphabet{Symbols: []rune{'0', '1'}},
		},
		{
			name:         "english alphabet",
			data:         "abcdefghijklmnopqrstuvwxyz",
			wantAlphabet: &Alphabet{Symbols: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}},
		},
		{
			name:         "spanish alphabet without diacritics",
			data:         "abcdefghijklmnñopqrstuvwxyz",
			wantAlphabet: &Alphabet{Symbols: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'ñ', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}},
		},
		{
			name:         "spanish alphabet with diacritics",
			data:         "aábcdeéfghiíjklmnñoópqrstuúüvwxyz",
			wantAlphabet: &Alphabet{Symbols: []rune{'a', 'á', 'b', 'c', 'd', 'e', 'é', 'f', 'g', 'h', 'i', 'í', 'j', 'k', 'l', 'm', 'n', 'ñ', 'o', 'ó', 'p', 'q', 'r', 's', 't', 'u', 'ú', 'ü', 'v', 'w', 'x', 'y', 'z'}},
		},
		{
			name:         "spanish upper case without diacritics",
			data:         "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			wantAlphabet: &Alphabet{Symbols: []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'Ñ', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotAlphabet := NewAlphabet(test.data)
			if diff := cmp.Diff(test.wantAlphabet, gotAlphabet); diff != "" {
				t.Errorf("NewAlphabet(%q) = %v, want %v: diff want -> got \n%s", test.data, *gotAlphabet, *test.wantAlphabet, diff)
			}
		})
	}
}
