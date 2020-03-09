package hillcipher

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

// TestNewMatrix_ThrowsErrorOnInvalidInput verify invalid inputs fail matrix creation
func TestNewMatrix_ThrowsErrorOnInvalidInput(t *testing.T) {
	tests := []struct {
		order    int
		dataSize int
	}{
		{order: 1, dataSize: 0},
		{order: 3, dataSize: 3},
		{order: 5, dataSize: 10},
	}
	for _, test := range tests {
		data := make([]float64, test.dataSize)
		if _, err := NewMatrix(test.order, data); err == nil {
			t.Errorf("NewMatrix(%v, %v) thwrew not nil error, should've failed.", test.order, data)
		}
	}
}

// TestMinor_ReturnsCorrect verifies correct minor matrix creation
func TestMinor_ReturnsCorrect(t *testing.T) {
	tests := []struct {
		input      *Matrix
		p, q       int
		wantMatrix *Matrix
	}{
		{
			input: &Matrix{Order: 1},
			p:     1, q: 1,
			wantMatrix: &Matrix{},
		},
		{
			input: &Matrix{
				Order: 3,
				Data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			p: 1, q: 1,
			wantMatrix: &Matrix{
				Order: 2,
				Data: [][]float64{
					{1, 3},
					{7, 9},
				},
			},
		},
	}
	for _, test := range tests {
		gotMatrix := Minor(test.input, test.p, test.q)
		if diff := cmp.Diff(test.wantMatrix, gotMatrix); diff != "" {
			t.Errorf("Minor(%v, %d, %d) got %v; want %v; got -> want %s", test.input, test.p, test.q, gotMatrix, test.wantMatrix, diff)
		}
	}
}
