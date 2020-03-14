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
			name: "order 2 mod 2 1011",
			mod:  2, data: []int{1, 0, 1, 1},
			wantKey: &Key{order: 2, data: [][]int{{1, 0}, {1, 1}}},
		},
		{
			name: "order 3 mod 27 FORTALEZA",
			mod:  27,
			data: []int{5, 15, 18, 20, 0, 11, 4, 26, 0},
			wantKey: &Key{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
		},
		{
			name: "order 4 mod 27 UNAMFCIENCIASCYS",
			mod:  27,
			data: []int{21, 13, 0, 12, 5, 2, 8, 4, 13, 2, 8, 0, 19, 2, 25, 19},
			wantKey: &Key{
				order: 4,
				data: [][]int{
					{21, 13, 0, 12},
					{5, 2, 8, 4},
					{13, 2, 8, 0},
					{19, 2, 25, 19},
				},
			},
		},
		{
			name: "order 5 mod 27 ÑOMEGUSTALCORONAVIRUSHELP",
			mod:  27,
			data: []int{14, 15, 12, 4, 6, 21, 19, 20, 0, 11, 2, 15, 18, 15, 13, 0, 22, 8, 18, 21, 19, 7, 4, 11, 16},
			wantKey: &Key{
				order: 5,
				data: [][]int{
					{14, 15, 12, 4, 6},
					{21, 19, 20, 0, 11},
					{2, 15, 18, 15, 13},
					{0, 22, 8, 18, 21},
					{19, 7, 4, 11, 16},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Key{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotKey, err := NewKey(test.data, test.mod)
			if err != nil {
				t.Fatalf("NewKey(%v, %d) returned unexpected error; %v", test.data, test.mod, err)
			}
			if diff := cmp.Diff(test.wantKey, gotKey, unxOpt); diff != "" {
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
				order: 3,
				data: [][]int{
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
			name: "binary alphabet",
			data: "01",
			wantAlphabet: &Alphabet{
				symbols:     []rune{'0', '1'},
				intIndex:    map[int]rune{0: '0', 1: '1'},
				symbolIndex: map[rune]int{'0': 0, '1': 1},
			},
		},
		{
			name: "english alphabet",
			data: "abcdefghijklmnopqrstuvwxyz",
			wantAlphabet: &Alphabet{
				symbols:     []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
				intIndex:    map[int]rune{0: 'a', 1: 'b', 2: 'c', 3: 'd', 4: 'e', 5: 'f', 6: 'g', 7: 'h', 8: 'i', 9: 'j', 10: 'k', 11: 'l', 12: 'm', 13: 'n', 14: 'o', 15: 'p', 16: 'q', 17: 'r', 18: 's', 19: 't', 20: 'u', 21: 'v', 22: 'w', 23: 'x', 24: 'y', 25: 'z'},
				symbolIndex: map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25},
			},
		},
		{
			name: "spanish upper case alphabet without diacritics",
			data: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			wantAlphabet: &Alphabet{
				symbols:     []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'Ñ', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
				intIndex:    map[int]rune{0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G', 7: 'H', 8: 'I', 9: 'J', 10: 'K', 11: 'L', 12: 'M', 13: 'N', 14: 'Ñ', 15: 'O', 16: 'P', 17: 'Q', 18: 'R', 19: 'S', 20: 'T', 21: 'U', 22: 'V', 23: 'W', 24: 'X', 25: 'Y', 26: 'Z'},
				symbolIndex: map[rune]int{'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'Ñ': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Alphabet{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotAlphabet := NewAlphabet(test.data)
			if diff := cmp.Diff(test.wantAlphabet, gotAlphabet, unxOpt); diff != "" {
				t.Errorf("NewAlphabet(%q) = %v, want %v: diff want -> got \n%s", test.data, *gotAlphabet, *test.wantAlphabet, diff)
			}
		})
	}
}

// TestContains verify implementation of alphabet contains
func TestContains(t *testing.T) {
	tests := []struct {
		alphabet  *Alphabet
		r         rune
		contained bool
	}{
		{alphabet: NewAlphabet("01"), r: '1', contained: true},
		{alphabet: NewAlphabet("01"), r: '2', contained: false},
		{alphabet: NewAlphabet("0123456789abcdef"), r: 'f', contained: true},
		{alphabet: NewAlphabet("0123456789abcdef"), r: 'g', contained: false},
		{alphabet: NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), r: 'a', contained: false},
		{alphabet: NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), r: 'A', contained: true},
	}
	for _, test := range tests {
		in := "in"
		if !test.contained {
			in = "not in"
		}
		name := fmt.Sprintf("%q %s %s", test.r, in, test.alphabet)
		t.Run(name, func(t *testing.T) {
			if test.alphabet.Contains(test.r) != test.contained {
				t.Errorf("Contains(%q) = %v, want %v", test.r, !test.contained, test.contained)
			}
		})
	}
}

// TestEncryption verify corrrectness of algorithm
func TestEncryption(t *testing.T) {
	tests := []struct {
		name, alphabet, msg, key, wantCipherText string
	}{
		{
			name:           "spanish alphabet key size 3",
			alphabet:       "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			msg:            "CONSUL",
			key:            "FORTALEZA",
			wantCipherText: "KUTÑOB",
		},
		{
			name:           "spanish alphabet key size 4",
			alphabet:       "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			msg:            "IBOMBATOMICALLYSOCRATESPHILOSOPHIESANDHYPOTHESESCANTDEFINEHOWIBEDROPPINGTHESEMOCKERY",
			key:            "UNAMFCIENCIASCYS",
			wantCipherText: "BUKBMJLUFLZXICÑCQHSKPAOGZKGHDLAGELUÑRLOMUBTVSEÑIMFVÑVÑGRSAQNÑTZÑDSGPZIEDKGJRUKHAVFMP",
		},
		{
			name:           "spanish alphabet key size 5",
			alphabet:       "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			msg:            "LYRICALLYPERFORMARMEDROBBERYFLEEWITHLOTTERYPOSSIBLYTHEYSPOTTEDMEBATTLESCARREDSHOGUNEXPLOSIONWHENMYPENHITSTREMENDOUSULTRAVIOLETSHINEBLINDFORENSICS",
			key:            "ÑOMEGUSTALCORONAVIRUSHELP",
			wantCipherText: "GIDPJHLBJEÑVQEAXINMKPÑHOAKUBZTEDYZKVMKIAXLCLPOOECLJXXNHVIBKTXÑRMPÑÑNRAQÑFQXOLTLENWOIMROSIVENNFOUSDZWKSMFOVVTZPLCMRZOXAXYBNDLQDLLAAPHXÑROPYQJKEDZJ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alphabet := NewAlphabet(test.alphabet)
			cipher, err := NewCipher(alphabet)
			if err != nil {
				t.Fatalf("NewCipher(%q) returned unexpected error; %v", alphabet, err)
			}
			gotCipherText, err := cipher.Encrypt(test.msg, test.key)
			if err != nil {
				t.Fatalf("Encrypt(msg:%q, key:%q) returned unexpected error; %v", test.msg, test.key, err)
			}
			if gotCipherText != test.wantCipherText {
				t.Errorf("Encrypt(msg:%q, key:%q) = %q, want %q", test.msg, test.key, gotCipherText, test.wantCipherText)
			}
		})
	}
}
