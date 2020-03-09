package hillcipher

import (
	"fmt"
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

// TestNewMatrix_ThrowsErrorOnInvalidInput verify invalid inputs fail matrix creation
func TestNewMatrix_ThrowsErrorOnInvalidInput(t *testing.T) {
	tests := []struct {
		name     string
		order    int
		dataSize int
	}{
		{name: "order 1 size 0", order: 1, dataSize: 0},
		{name: "order 3 size 3 (equal)", order: 3, dataSize: 3},
		{name: "order 5 size 10 (double)", order: 5, dataSize: 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := make([]float64, test.dataSize)
			if _, err := NewMatrix(test.order, data); err == nil {
				t.Errorf("NewMatrix(%v, %v) did not fail, should've failed.", test.order, data)
			}
		})
	}
}

// TestMinor_ReturnsCorrectMatrix verifies correct minor matrix creation
func TestMinor_ReturnsCorrectMatrix(t *testing.T) {
	tests := []struct {
		name       string
		matrix     *Matrix
		p, q       int
		wantMatrix *Matrix
	}{
		{
			name:   "order 1",
			matrix: &Matrix{Order: 1},
			p:      0, q: 0,
			wantMatrix: &Matrix{},
		},
		{
			name: "order 2",
			matrix: &Matrix{
				Order: 2,
				Data: [][]float64{
					{1, 3},
					{9, -1},
				},
			},
			p: 0, q: 0,
			wantMatrix: &Matrix{Order: 1, Data: [][]float64{{-1}}},
		},
		{
			name: "order 3 middle cross",
			matrix: &Matrix{
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
			name: "order 3 edge cross",
			matrix: &Matrix{
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
		t.Run(test.name, func(t *testing.T) {
			gotMatrix, err := Minor(test.matrix, test.p, test.q)
			if err != nil {
				t.Fatalf("Minor(%v, %d, %d) returned unexpected error; %v", test.matrix, test.p, test.q, err)
			}
			if diff := cmp.Diff(test.wantMatrix, gotMatrix); diff != "" {
				t.Errorf("Minor(%v, %d, %d) = %v; want %v; diff: want -> got %s", test.matrix, test.p, test.q, gotMatrix, test.wantMatrix, diff)
			}
		})
	}
}

// TestMinut_ThrowsErrorOnInvalidIndexes verifies row and column indices are in valid range
func TestMinut_ThrowsErrorOnInvalidIndexes(t *testing.T) {
	tests := []struct {
		name              string
		matrixOrder, p, q int
	}{
		{name: "p exceeding", matrixOrder: 3, p: 4, q: 1},
		{name: "q exceeding", matrixOrder: 3, p: 1, q: 4},
		{name: "both exceeding", matrixOrder: 0, p: 1, q: 4},
		{name: "negative q", matrixOrder: 5, p: 2, q: -1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := make([]float64, test.matrixOrder*test.matrixOrder)
			m, err := NewMatrix(test.matrixOrder, data)
			if err != nil {
				t.Fatalf("NewMatrix(%d, %v) returned unexpected error; %v", test.matrixOrder, data, err)
			}
			if _, err := Minor(m, test.p, test.q); err == nil {
				t.Errorf("Minor(%v, %d, %d) did not fail, should've failed.", m, test.p, test.q)
			}
		})
	}
}

// TestString_RetursCorrectValue verifies returned string is the expected
func TestString_RetursCorrectValue(t *testing.T) {
	tests := []struct {
		name    string
		matrix  *Matrix
		wantRep string
	}{
		{
			name: "empty matrix",
			matrix: &Matrix{
				Order: 0,
				Data:  [][]float64{},
			},
			wantRep: "",
		},
		{
			name: "order 1",
			matrix: &Matrix{
				Order: 1,
				Data:  [][]float64{{1}},
			},
			wantRep: "|\t1.00\t|\n",
		},
		{
			name: "order 3",
			matrix: &Matrix{
				Order: 3,
				Data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantRep: "|	1.00	|	2.00	|	3.00	|\n|	4.00	|	5.00	|	6.00	|\n|	7.00	|	8.00	|	9.00	|\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotRep := fmt.Sprintf("%s", test.matrix)
			if gotRep != test.wantRep {
				t.Errorf("String(m) = %s, want %s", gotRep, test.wantRep)
			}
		})
	}
}

// TestDeterminant_CorrectResult verifies returns the correct result
func TestDeterminant_CorrectResult(t *testing.T) {
	tests := []struct {
		name    string
		matrix  *Matrix
		wantDet float64
	}{
		{
			name: "order 1",
			matrix: &Matrix{
				Order: 1,
				Data:  [][]float64{{1}},
			},
			wantDet: 1.0,
		},
		{
			name: "order 2",
			matrix: &Matrix{
				Order: 2,
				Data: [][]float64{
					{1, 3},
					{9, -1},
				},
			},
			wantDet: -28,
		},
		{
			name: "order 3 (without inverse)",
			matrix: &Matrix{
				Order: 3,
				Data: [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			wantDet: 0.0,
		},
		{
			name: "order 10",
			matrix: &Matrix{
				Order: 10,
				Data: [][]float64{
					{52, 37, 38, 88, 89, 9, 23, 95, 99, 16},
					{59, 23, 35, 36, 43, 13, 26, 46, 47, 85},
					{7, 23, 84, 24, 83, 100, 30, 72, 86, 93},
					{54, 94, 77, 59, 50, 29, 94, 64, 43, 37},
					{68, 17, 65, 23, 19, 43, 68, 78, 15, 73},
					{93, 96, 30, 86, 52, 55, 37, 58, 31, 22},
					{58, 41, 85, 35, 18, 54, 26, 96, 43, 73},
					{41, 88, 52, 36, 42, 6, 69, 12, 32, 3},
					{72, 57, 9, 15, 78, 90, 63, 77, 17, 1},
					{80, 49, 18, 67, 47, 22, 86, 13, 2, 33},
				},
			},
			wantDet: 165148033107009656.0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotDet, err := test.matrix.Determinant()
			if err != nil {
				t.Fatalf("Determinant(%v) returned an unexpered error; %v", test.matrix, err)
			}
			if gotDet != test.wantDet {
				t.Errorf("got incorrect determinant %.4f want %.4f", gotDet, test.wantDet)
			}
		})
	}
}
