package math

import "testing"

// TestNewMatrix_ThrowsErrorOnInvalidInput verify invalid inputs fail matrix creation
func TestNewMatrix_ThrowsErrorOnInvalidInput(t *testing.T) {
	tests := []struct{
		order int
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
	tests := []struct{
		input *Matrix
		p, q int
		want *Matrix
	}{
		{
			input: &Matrix{order: 1},
			p: 1, q: 1,
			want: &Matrix{},
		},
		{
			input: &Matrix{
				order: 3,
				data: [][]float64{
					{1,2,3},
					{4,5,6},
					{7,8,9},
				},
			},
			p:1,q:1,
			want: &Matrix{
				order: 2,
				data: [][]float64{
					{1,3},
					{7,9},
				},
			},
		},
	}
	for _, test := range tests {
		if
	}
}