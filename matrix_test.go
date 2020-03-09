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
			t.Errorf("NewMatrix(%v, %v) did not fail, should've failed.", test.order, data)
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
			p:     0, q: 0,
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
		{
			input: &Matrix{
				Order: 3,
				Data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			p: 0, q: 0,
			wantMatrix: &Matrix{
				Order: 2,
				Data: [][]float64{
					{5, 6},
					{8, 9},
				},
			},
		},
	}
	for _, test := range tests {
		gotMatrix, err := Minor(test.input, test.p, test.q)
		if err != nil {
			t.Fatalf("Minor(%v, %d, %d) returned unexpected error; %v", test.input, test.p, test.q, err)
		}
		if diff := cmp.Diff(test.wantMatrix, gotMatrix); diff != "" {
			t.Errorf("Minor(%v, %d, %d) got %v; want %v; got -> want %s", test.input, test.p, test.q, gotMatrix, test.wantMatrix, diff)
		}
	}
}

// TestMinut_ThrowsErrorOnInvalidIndexes verifies row and column indices are in valid range
func TestMinut_ThrowsErrorOnInvalidIndexes(t *testing.T) {
	tests := []struct {
		matrixOrder, p, q int
	}{
		{matrixOrder: 3, p: 4, q: 1},
		{matrixOrder: 3, p: 1, q: 4},
		{matrixOrder: 0, p: 1, q: 4},
		{matrixOrder: 5, p: 2, q: -1},
	}
	for _, test := range tests {
		data := make([]float64, test.matrixOrder*test.matrixOrder)
		m, err := NewMatrix(test.matrixOrder, data)
		if err != nil {
			t.Fatalf("NewMatrix(%d, %v) returned unexpected error; %v", test.matrixOrder, data, err)
		}
		if _, err := Minor(m, test.p, test.q); err == nil {
			t.Errorf("Minor(%v, %d, %d) did not fail, should've failed.", m, test.p, test.q)
		}
	}
}
