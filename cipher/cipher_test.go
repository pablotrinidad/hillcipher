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

// TestBelongs verify implementation of alphabet Belongs
func TestBelongs(t *testing.T) {
	tests := []struct {
		alphabet  *Alphabet
		substring string
		contained bool
	}{
		{alphabet: NewAlphabet("01"), substring: "100101101", contained: true},
		{alphabet: NewAlphabet("01"), substring: "2", contained: false},
		{alphabet: NewAlphabet("01"), substring: "0", contained: true},
		{alphabet: NewAlphabet("0123456789abcdef"), substring: "f14e92a", contained: true},
		{alphabet: NewAlphabet("0123456789abcdef"), substring: "asd92ssa0239asdgkq", contained: false},
		{alphabet: NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), substring: "abcdefghiu", contained: false},
		{alphabet: NewAlphabet("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), substring: "HOLAMUNDO", contained: true},
	}
	for _, test := range tests {
		in := "in"
		if !test.contained {
			in = "not in"
		}
		name := fmt.Sprintf("%q %s %s", test.substring, in, test.alphabet)
		t.Run(name, func(t *testing.T) {
			if test.alphabet.Belongs(test.substring) != test.contained {
				t.Errorf("Belongs(%q) = %v, want %v", test.substring, !test.contained, test.contained)
			}
		})
	}
}

// TestStoi verify correct mapping between symbols and numbers
func TestStoi(t *testing.T) {
	tests := []struct {
		alphabet *Alphabet
		mapping  map[rune]int
	}{
		{
			alphabet: NewAlphabet("01"),
			mapping:  map[rune]int{'0': 0, '1': 1},
		},
		{
			alphabet: NewAlphabet("0123456789ABCDEF"),
			mapping:  map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, 'F': 15},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("alphabet %q", test.alphabet.String())
		t.Run(name, func(t *testing.T) {
			for r, want := range test.mapping {
				got, err := test.alphabet.Stoi(r)
				if err != nil {
					t.Errorf("Stoi(%q) got unexpected error; %v", r, err)
				}
				if got != want {
					t.Errorf("Stoi(%q) = %d, want %d", r, got, want)
				}
			}
		})
	}
}

// TestStoi_Error verify correct validations
func TestStoi_Error(t *testing.T) {
	tests := []struct {
		alphabet       *Alphabet
		invalidSymbols []rune
	}{
		{
			alphabet:       NewAlphabet("01"),
			invalidSymbols: []rune{'2'},
		},
		{
			alphabet:       NewAlphabet("0123456789ABCDEF"),
			invalidSymbols: []rune{'G', 'H', 'I', 'J'},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("alphabet %q", test.alphabet.String())
		t.Run(name, func(t *testing.T) {
			for _, r := range test.invalidSymbols {
				_, err := test.alphabet.Stoi(r)
				if err == nil {
					t.Errorf("Stoi(%q) got non-nil error, expected error", r)
				}
			}
		})
	}
}

// TestItos verify correct mapping between symbols and numbers
func TestItos(t *testing.T) {
	tests := []struct {
		alphabet *Alphabet
		mapping  map[int]rune
	}{
		{
			alphabet: NewAlphabet("01"),
			mapping:  map[int]rune{0: '0', 1: '1'},
		},
		{
			alphabet: NewAlphabet("0123456789ABCDEF"),
			mapping:  map[int]rune{0: '0', 1: '1', 2: '2', 3: '3', 15: 'F'},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("alphabet %q", test.alphabet.String())
		t.Run(name, func(t *testing.T) {
			for i, want := range test.mapping {
				got, err := test.alphabet.Itos(i)
				if err != nil {
					t.Errorf("Itos(%d) got unexpected error; %v", i, err)
				}
				if got != want {
					t.Errorf("Itos(%d) = %d, want %d", i, got, want)
				}
			}
		})
	}
}

// TestItos_Error verify correct validations
func TestItos_Error(t *testing.T) {
	tests := []struct {
		alphabet       *Alphabet
		invalidSymbols []int
	}{
		{
			alphabet:       NewAlphabet("01"),
			invalidSymbols: []int{2},
		},
		{
			alphabet:       NewAlphabet("0123456789ABCDEF"),
			invalidSymbols: []int{-1, 16, 100},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("alphabet %q", test.alphabet.String())
		t.Run(name, func(t *testing.T) {
			for _, i := range test.invalidSymbols {
				_, err := test.alphabet.Itos(i)
				if err == nil {
					t.Errorf("Itos(%q) got non-nil error, expected error", i)
				}
			}
		})
	}
}

// TestNewCipher verify correct creation of cipher
func TestNewCipher(t *testing.T) {
	tests := []struct {
		name       string
		alphabet   *Alphabet
		wantCipher *Cipher
	}{
		{
			name:       "english alphabet",
			alphabet:   NewAlphabet("abcdefghijklmnopqrstuvwxyz"),
			wantCipher: &Cipher{mod: 26},
		},
		{
			name:       "hex alphabet",
			alphabet:   NewAlphabet("0123456789ABCDEF"),
			wantCipher: &Cipher{mod: 16},
		},
		{
			name:       "binary alphabet",
			alphabet:   NewAlphabet("01"),
			wantCipher: &Cipher{mod: 2},
		},
	}
	unxOpt := cmp.AllowUnexported(Cipher{}, Alphabet{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.wantCipher.alphabet = *test.alphabet
			gotCipher, err := NewCipher(test.alphabet)
			if err != nil {
				t.Fatalf("NewCipher(%s) returned unexpected error; %v", test.alphabet, err)
			}
			if diff := cmp.Diff(test.wantCipher, gotCipher, unxOpt); diff != "" {
				t.Errorf("NewCipher(%s) = %v, want %d; diff want -> got %s", test.alphabet, gotCipher, test.wantCipher, diff)
			}
		})
	}
}

// TestNewCipher_Error verify method validates input data
func TestNewCipher_Error(t *testing.T) {
	tests := []struct {
		name     string
		alphabet *Alphabet
	}{
		{
			name:     "one-symbol alphabet",
			alphabet: NewAlphabet("A"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewCipher(test.alphabet)
			if err == nil {
				t.Fatalf("NewCipher(%s) got non-nil error, expected error", test.alphabet)
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
		{
			name:           "hex alphabet k size 10",
			alphabet:       "0123456789ABCDEF",
			msg:            "F60DFBF641AA19A49D881A85A3F503EFA720097BB62489809A064D7B445A295BEE0665A83A9D686F6F8733736E28859D7C395D57C8A201F2ADCC3DDB58B341785EC2343E3BC667105F3BEA8AB63A03A8DC5C1CBE0E441618C90C19B20A434F3C07B666D4FC8F329D43898303B7579A3E6D94AB3635A683722C7B2309B4A2BCA49BF449D41FC18F764CD42AA0BA41785AD10277C982237F47B6A94C423DAE43DF65CE4E875C2356844C6B36E53C05B41A4B57BA5DA5A4F36A28ED1E574AF9A050F9CE5E1A530FF07F0B7DB27087AFE18B405DB86E80FD708CE6CB3D767E10C4A85C25C40C4AAAD1D59042CEA7EA8444FBC227E2EDE7E930821ACA014D6B44824D189DE2A84DF59A128A596C1DCA6FBEDD4BBA26643D6B854F81F682FAD1C4CCF41E15AC800EC70110B1559EDBF258E79748FBE744A99504877B77ECB08FE402C339865B896D420B64B7B7DF8139ECA3857AE34B4EA61218E9AB1AF0AFC4AAF13E838D8EA7A900560DE0A405C0EBE531F373BE39735339E4BD07E29EFBCD6532D88531D9AECF6771E5D7C9BF10E28EABE64A0B40F15BB0B7A17EFE171D6B46BFD4B49EF010337B65EAFFF1D99CC262B48DB7ED1543B05A737BB4CA48E9219B3659E45A5F0C1C8CD9F075E2BE09465AB95B81496E86688BDBB42741E749BAF5A8CE77F96552E7DFD2585A76772CF661B436936F1CC6",
			key:            "06095889993EA8DE0BBA646B61268E5E23E5A8C83BD43F8BF3F0220BC8480947875E6624CC812791A228171F1C62D14024F1",
			wantCipherText: "7E6DE032F2CDDDEAC5C438A10E5D3460D100ADAB36F4183AEDF0809FC2D68E3DD8AB4947D1279E667C64EDAECF282ACF1BA14BBA0B696D2570CBE7D522472A5477C2A0389F3FE871D5D681E4A8B84DAFDC0FFDF0E30646FA4FEFEE89D71C18D57C597708A9DE8180BAC22817C6C191F884668AFBBEE37A0954488725BF08EEB07E16F4B9C3655E5E0925065C8FBE0FA05DEE645F0C1F115681AFEFAAF3EA1F8E84A87410AB566C3075A64810ABD0B249451AF36399DD534742FA75014137CE5F27D672DE41DCB534D191CAD3875F86C93101214F75DEB94A9751E12D1E85336B3E0C00FA5666CA5EE74E54AB06B61A5E3DB3EC3A7E606E2A0455FE604B995DFB57AD29ED43717EF06A623EBE6E87FD01FED586B14E038D07472A088C45B66458963201A8C02CF187EB66F3C25CF20C68914FB293C361342C5BB02B6F54E7BBE099D44A873162BBBE099DB6E9B010A6B51D83EA0598A8A4A0ECE7F63AB59DFEF0DCE26BE0BD201594DC37A59A1FE9E8E30FFFD5678684863E168916A1F3BB34590854545BAE537461F83ECB1A8181C556447B1A01FF42CD6494D550CFC8F52AA161637BBA3ECDA4D7AAEE18A90CFE7E2B08868F0D0640E7EF28899984DF495AE88A1ED3EDD3BCA8A19C1BE1199A19B9056AF06A865925E6BC06C564593545CB858F4DBC0C87C36751CF1C1E85F33A13B3271615CA",
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

// TestEncryption_Error verify validations of input data
func TestEncryption_Error(t *testing.T) {
	tests := []struct {
		name, alphabet, msg, key string
	}{
		{
			name:     "key does not belong to alphabet",
			alphabet: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			key:      "fortaleza",
		},
		{
			name:     "message does not belong to alphabet",
			alphabet: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			msg:      "sup",
		},
		{
			name:     "message length is not multiple of key's length",
			alphabet: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			msg:      "SUPWORLD",
			key:      "FORTALEZA",
		},
		{
			name:     "key is not invertible in dict",
			alphabet: "0123456789ABCDEF",
			key:      "92666C703E4135B097C7D2EA9C699C274C4F9442F13D38013F28C1765D3461A52E82261E74EAB8C35D6BA6457DF68830B0E0",
		},
		{
			name:     "key length is less than 2",
			alphabet: "0123456789ABCDEF",
			key:      "A",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alphabet := NewAlphabet(test.alphabet)
			cipher, err := NewCipher(alphabet)
			if err != nil {
				t.Fatalf("NewCipher(%q) returned unexpected error; %v", alphabet, err)
			}
			_, err = cipher.Encrypt(test.msg, test.key)
			if err == nil {
				t.Fatalf("Encrypt(msg:%q, key:%q) returned non-nil error, expected error", test.msg, test.key)
			}
		})
	}
}

// TestDecryption verify corrrectness of algorithm
func TestDecryption(t *testing.T) {
	tests := []struct {
		name, alphabet, wantPlainText, key, cipherText string
	}{
		{
			name:          "spanish alphabet key size 3",
			alphabet:      "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			wantPlainText: "CONSUL",
			key:           "FORTALEZA",
			cipherText:    "KUTÑOB",
		},
		{
			name:          "spanish alphabet key size 4",
			alphabet:      "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			wantPlainText: "IBOMBATOMICALLYSOCRATESPHILOSOPHIESANDHYPOTHESESCANTDEFINEHOWIBEDROPPINGTHESEMOCKERY",
			key:           "UNAMFCIENCIASCYS",
			cipherText:    "BUKBMJLUFLZXICÑCQHSKPAOGZKGHDLAGELUÑRLOMUBTVSEÑIMFVÑVÑGRSAQNÑTZÑDSGPZIEDKGJRUKHAVFMP",
		},
		{
			name:          "spanish alphabet key size 5",
			alphabet:      "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			wantPlainText: "LYRICALLYPERFORMARMEDROBBERYFLEEWITHLOTTERYPOSSIBLYTHEYSPOTTEDMEBATTLESCARREDSHOGUNEXPLOSIONWHENMYPENHITSTREMENDOUSULTRAVIOLETSHINEBLINDFORENSICS",
			key:           "ÑOMEGUSTALCORONAVIRUSHELP",
			cipherText:    "GIDPJHLBJEÑVQEAXINMKPÑHOAKUBZTEDYZKVMKIAXLCLPOOECLJXXNHVIBKTXÑRMPÑÑNRAQÑFQXOLTLENWOIMROSIVENNFOUSDZWKSMFOVVTZPLCMRZOXAXYBNDLQDLLAAPHXÑROPYQJKEDZJ",
		},
		{
			name:          "hex alphabet k size 10",
			alphabet:      "0123456789ABCDEF",
			wantPlainText: "F60DFBF641AA19A49D881A85A3F503EFA720097BB62489809A064D7B445A295BEE0665A83A9D686F6F8733736E28859D7C395D57C8A201F2ADCC3DDB58B341785EC2343E3BC667105F3BEA8AB63A03A8DC5C1CBE0E441618C90C19B20A434F3C07B666D4FC8F329D43898303B7579A3E6D94AB3635A683722C7B2309B4A2BCA49BF449D41FC18F764CD42AA0BA41785AD10277C982237F47B6A94C423DAE43DF65CE4E875C2356844C6B36E53C05B41A4B57BA5DA5A4F36A28ED1E574AF9A050F9CE5E1A530FF07F0B7DB27087AFE18B405DB86E80FD708CE6CB3D767E10C4A85C25C40C4AAAD1D59042CEA7EA8444FBC227E2EDE7E930821ACA014D6B44824D189DE2A84DF59A128A596C1DCA6FBEDD4BBA26643D6B854F81F682FAD1C4CCF41E15AC800EC70110B1559EDBF258E79748FBE744A99504877B77ECB08FE402C339865B896D420B64B7B7DF8139ECA3857AE34B4EA61218E9AB1AF0AFC4AAF13E838D8EA7A900560DE0A405C0EBE531F373BE39735339E4BD07E29EFBCD6532D88531D9AECF6771E5D7C9BF10E28EABE64A0B40F15BB0B7A17EFE171D6B46BFD4B49EF010337B65EAFFF1D99CC262B48DB7ED1543B05A737BB4CA48E9219B3659E45A5F0C1C8CD9F075E2BE09465AB95B81496E86688BDBB42741E749BAF5A8CE77F96552E7DFD2585A76772CF661B436936F1CC6",
			key:           "06095889993EA8DE0BBA646B61268E5E23E5A8C83BD43F8BF3F0220BC8480947875E6624CC812791A228171F1C62D14024F1",
			cipherText:    "7E6DE032F2CDDDEAC5C438A10E5D3460D100ADAB36F4183AEDF0809FC2D68E3DD8AB4947D1279E667C64EDAECF282ACF1BA14BBA0B696D2570CBE7D522472A5477C2A0389F3FE871D5D681E4A8B84DAFDC0FFDF0E30646FA4FEFEE89D71C18D57C597708A9DE8180BAC22817C6C191F884668AFBBEE37A0954488725BF08EEB07E16F4B9C3655E5E0925065C8FBE0FA05DEE645F0C1F115681AFEFAAF3EA1F8E84A87410AB566C3075A64810ABD0B249451AF36399DD534742FA75014137CE5F27D672DE41DCB534D191CAD3875F86C93101214F75DEB94A9751E12D1E85336B3E0C00FA5666CA5EE74E54AB06B61A5E3DB3EC3A7E606E2A0455FE604B995DFB57AD29ED43717EF06A623EBE6E87FD01FED586B14E038D07472A088C45B66458963201A8C02CF187EB66F3C25CF20C68914FB293C361342C5BB02B6F54E7BBE099D44A873162BBBE099DB6E9B010A6B51D83EA0598A8A4A0ECE7F63AB59DFEF0DCE26BE0BD201594DC37A59A1FE9E8E30FFFD5678684863E168916A1F3BB34590854545BAE537461F83ECB1A8181C556447B1A01FF42CD6494D550CFC8F52AA161637BBA3ECDA4D7AAEE18A90CFE7E2B08868F0D0640E7EF28899984DF495AE88A1ED3EDD3BCA8A19C1BE1199A19B9056AF06A865925E6BC06C564593545CB858F4DBC0C87C36751CF1C1E85F33A13B3271615CA",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alphabet := NewAlphabet(test.alphabet)
			cipher, err := NewCipher(alphabet)
			if err != nil {
				t.Fatalf("NewCipher(%q) returned unexpected error; %v", alphabet, err)
			}
			gotPlainText, err := cipher.Decrypt(test.cipherText, test.key)
			if err != nil {
				t.Fatalf("Decrypt(msg:%q, key:%q) returned unexpected error; %v", test.cipherText, test.key, err)
			}
			if gotPlainText != test.wantPlainText {
				t.Errorf("Decrypt(msg:%q, key:%q) = %q, want %q", test.cipherText, test.key, gotPlainText, test.wantPlainText)
			}
		})
	}
}

// TestDecryption_Error verify validations of input data
func TestDecryption_Error(t *testing.T) {
	tests := []struct {
		name, alphabet, cipherText, key string
	}{
		{
			name:     "key does not belong to alphabet",
			alphabet: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			key:      "fortaleza",
		},
		{
			name:       "cipher text does not belong to alphabet",
			alphabet:   "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			cipherText: "sup",
		},
		{
			name:       "cipher text length is not multiple of key's length",
			alphabet:   "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
			cipherText: "SUPWORLD",
			key:        "FORTALEZA",
		},
		{
			name:     "key is not invertible in dict",
			alphabet: "0123456789ABCDEF",
			key:      "92666C703E4135B097C7D2EA9C699C274C4F9442F13D38013F28C1765D3461A52E82261E74EAB8C35D6BA6457DF68830B0E0",
		},
		{
			name:     "key length is less than 2",
			alphabet: "0123456789ABCDEF",
			key:      "A",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alphabet := NewAlphabet(test.alphabet)
			cipher, err := NewCipher(alphabet)
			if err != nil {
				t.Fatalf("NewCipher(%q) returned unexpected error; %v", alphabet, err)
			}
			_, err = cipher.Decrypt(test.cipherText, test.key)
			if err == nil {
				t.Fatalf("Decrypt(msg:%q, key:%q) returned non-nil error, expected error", test.cipherText, test.key)
			}
		})
	}
}
