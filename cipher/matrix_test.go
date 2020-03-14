package cipher

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
			data := make([]int, test.dataSize)
			if _, err := NewMatrix(test.order, data); err == nil {
				t.Errorf("NewMatrix(%v, %v) did not fail, it should have.", test.order, data)
			}
		})
	}
}

// TestNewMatrix_CorrectMatrixCreation verify matrices are created correctly
func TestNewMatrix_CorrectMatrixCreation(t *testing.T) {
	tests := []struct {
		name       string
		order      int
		input      []int
		wantMatrix *Matrix
	}{
		{
			name:       "order 1",
			order:      1,
			input:      []int{-99999},
			wantMatrix: &Matrix{order: 1, data: [][]int{{-99999}}},
		},
		{
			name:  "order 2",
			order: 2,
			input: []int{12, 34, 56, 78},
			wantMatrix: &Matrix{
				order: 2,
				data: [][]int{
					{12, 34},
					{56, 78},
				},
			},
		},
		{
			name:  "order 10",
			order: 10,
			input: []int{
				52, 37, 38, 88, 89, 9, 23, 95, 99, 16,
				59, 23, 35, 36, 43, 13, 26, 46, 47, 85,
				7, 23, 84, 24, 83, 100, 30, 72, 86, 93,
				54, 94, 77, 59, 50, 29, 94, 64, 43, 37,
				68, 17, 65, 23, 19, 43, 68, 78, 15, 73,
				93, 96, 30, 86, 52, 55, 37, 58, 31, 22,
				58, 41, 85, 35, 18, 54, 26, 96, 43, 73,
				41, 88, 52, 36, 42, 6, 69, 12, 32, 3,
				72, 57, 9, 15, 78, 90, 63, 77, 17, 1,
				80, 49, 18, 67, 47, 22, 86, 13, 2, 33,
			},
			wantMatrix: &Matrix{
				order: 10,
				data: [][]int{
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
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotMatrix, err := NewMatrix(test.order, test.input)
			if err != nil {
				t.Errorf("NewMatrix(%v, %v) returned unexpected error; %v.", test.order, test.input, err)
			}
			if diff := cmp.Diff(test.wantMatrix, gotMatrix, unxOpt); diff != "" {
				t.Errorf("NewMatrix(%v, %v) = %v; want %v; diff: want -> got %s", test.order, test.input, gotMatrix, test.wantMatrix, diff)
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
			matrix: &Matrix{order: 1},
			p:      0, q: 0,
			wantMatrix: &Matrix{},
		},
		{
			name: "order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 3},
					{9, -1},
				},
			},
			p: 0, q: 0,
			wantMatrix: &Matrix{order: 1, data: [][]int{{-1}}},
		},
		{
			name: "order 3 middle cross",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			p: 1, q: 1,
			wantMatrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 3},
					{7, 9},
				},
			},
		},
		{
			name: "order 3 edge cross",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			p: 0, q: 0,
			wantMatrix: &Matrix{
				order: 2,
				data: [][]int{
					{5, 6},
					{8, 9},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotMatrix, err := Minor(test.matrix, test.p, test.q)
			if err != nil {
				t.Fatalf("Minor(%v, %d, %d) returned unexpected error; %v", test.matrix, test.p, test.q, err)
			}
			if diff := cmp.Diff(test.wantMatrix, gotMatrix, unxOpt); diff != "" {
				t.Errorf("Minor(%v, %d, %d) = %v; want %v; diff: want -> got %s", test.matrix, test.p, test.q, gotMatrix, test.wantMatrix, diff)
			}
		})
	}
}

// TestMinor_ThrowsErrorOnInvalidIndexes verifies row and column indices are in valid range
func TestMinor_ThrowsErrorOnInvalidIndexes(t *testing.T) {
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
			data := make([]int, test.matrixOrder*test.matrixOrder)
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

// TestString_ReturnsCorrectValue verifies returned string is the expected
func TestString_ReturnsCorrectValue(t *testing.T) {
	tests := []struct {
		name    string
		matrix  *Matrix
		wantRep string
	}{
		{
			name: "empty matrix",
			matrix: &Matrix{
				order: 0,
				data:  [][]int{},
			},
			wantRep: "",
		},
		{
			name: "order 1",
			matrix: &Matrix{
				order: 1,
				data:  [][]int{{1}},
			},
			wantRep: "|\t1\t|\n",
		},
		{
			name: "order 3",
			matrix: &Matrix{
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
		wantDet int
	}{
		{
			name: "order 1",
			matrix: &Matrix{
				order: 1,
				data:  [][]int{{1}},
			},
			wantDet: 1.0,
		},
		{
			name: "order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 3},
					{9, -1},
				},
			},
			wantDet: -28,
		},
		{
			name: "order 3 (without inverse)",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
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
				order: 10,
				data: [][]int{
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
				t.Fatalf("Determinant(%v) returned an unexpected error; %v", test.matrix, err)
			}
			if gotDet != test.wantDet {
				t.Errorf("got incorrect determinant %df want %df", gotDet, test.wantDet)
			}
		})
	}
}

// TestDeterminant_InvalidOrder verifies fails when invalid order is sent
func TestDeterminant_InvalidOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
	}{
		{name: "order 0", matrix: &Matrix{order: 0, data: [][]int{}}},
		{name: "negative order", matrix: &Matrix{order: -13, data: [][]int{}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := test.matrix.Determinant(); err == nil {
				t.Errorf("Determinant() of matrix did not fail, it should have\n%v.", test.name)
			}
		})
	}
}

// TestIsInvertibleMod
func TestIsInvertibleMod(t *testing.T) {
	tests := []struct {
		name         string
		matrix       *Matrix
		mod          int
		isInvertible bool
	}{
		{name: "order 0 not invertible mod 10", matrix: &Matrix{}, isInvertible: false, mod: 10},
		{
			name:         "order 1 not invertible mod less than 2",
			matrix:       &Matrix{order: 1, data: [][]int{{1}}},
			isInvertible: false,
			mod:          1,
		},
		{
			name:         "order 1 invertible mod 10",
			matrix:       &Matrix{order: 1, data: [][]int{{1}}},
			isInvertible: true,
			mod:          10,
		},
		{
			name:         "order 1 not invertible mod 10",
			matrix:       &Matrix{order: 1, data: [][]int{{11}}},
			isInvertible: false,
			mod:          10,
		},
		{
			name: "order 3 invertible mod 26",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{6, 24, 1},
					{13, 16, 10},
					{20, 17, 15},
				},
			},
			isInvertible: true,
			mod:          26,
		},
		{
			name: "order 3 invertible mod 27",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			isInvertible: true,
			mod:          27,
		},
		{
			name: "order 4 invertible mod 27",
			matrix: &Matrix{
				order: 4,
				data: [][]int{
					{6, 24, 1, 15},
					{13, 16, 10, 23},
					{20, 17, 15, 23},
					{1, 2, 9, 13},
				},
			},
			isInvertible: true,
			mod:          27,
		},
		{
			name: "order 5 invertible mod 49",
			matrix: &Matrix{
				order: 5,
				data: [][]int{
					{6, 24, 44, 1, 15},
					{13, 16, 48, 10, 23},
					{20, 20, 17, 15, 23},
					{1, 2, 9, 13, 0},
					{48, 47, 46, 45, 44},
				},
			},
			isInvertible: true,
			mod:          49,
		},
		{
			name: "same order 5 not invertible mod 50",
			matrix: &Matrix{
				order: 5,
				data: [][]int{
					{6, 24, 44, 1, 15},
					{13, 16, 48, 10, 23},
					{20, 20, 17, 15, 23},
					{1, 2, 9, 13, 0},
					{48, 47, 46, 45, 44},
				},
			},
			isInvertible: false,
			mod:          50,
		},
		{
			name: "same order 5 but invertible mod 51",
			matrix: &Matrix{
				order: 5,
				data: [][]int{
					{6, 24, 44, 1, 15},
					{13, 16, 48, 10, 23},
					{20, 20, 17, 15, 23},
					{1, 2, 9, 13, 0},
					{48, 47, 46, 45, 44},
				},
			},
			isInvertible: true,
			mod:          51,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if inv := test.matrix.IsInvertibleMod(test.mod); inv != test.isInvertible {
				t.Errorf("IsInvertibleMod(%d) = %v, want %v for matrix\n%s", test.mod, inv, test.isInvertible, test.matrix)
			}
		})
	}
}

// TestCofactor verifies correct definition of cofactor
func TestCofactor(t *testing.T) {
	tests := []struct {
		name            string
		matrix, wantCof *Matrix
	}{
		{
			name: "order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 3},
					{9, 1},
				},
			},
			wantCof: &Matrix{
				order: 2,
				data: [][]int{
					{1, -9},
					{-3, 1},
				},
			},
		},
		{
			name: "another order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{-23, 71},
					{92, 86},
				},
			},
			wantCof: &Matrix{
				order: 2,
				data: [][]int{
					{86, -92},
					{-71, -23},
				},
			},
		},
		{
			name: "order 3",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{0, 9, 3},
					{2, 0, 4},
					{3, 7, 0},
				},
			},
			wantCof: &Matrix{
				order: 3,
				data: [][]int{
					{-28, 12, 14},
					{21, -9, 27},
					{36, 6, -18},
				},
			},
		},
		{
			name: "another order 3",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			wantCof: &Matrix{
				order: 3,
				data: [][]int{
					{-286, 44, 520},
					{468, -72, -70},
					{165, 305, -300},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cof, err := test.matrix.Cofactor()
			if err != nil {
				t.Fatalf("Cofactor(m) returned unexpected error; %v\n%s", err, test.matrix)
			}
			if diff := cmp.Diff(cof, test.wantCof, unxOpt); diff != "" {
				t.Errorf("Cofactor(\n%s) =\n%s, want\n%s; diff want -> got:\n%s", test.matrix, cof, test.wantCof, diff)
			}
		})
	}
}

// TestCofactor_Error verifies error check in cofactor
func TestCofactor_Error(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
	}{
		{
			name:   "order 1",
			matrix: &Matrix{order: 1, data: [][]int{{1}}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := test.matrix.Cofactor(); err == nil {
				t.Fatalf("Cofactor(\n%s) returned non-nil error, want error; ", test.matrix)

			}
		})
	}
}

// TestTranspose verify method implementation
func TestTranspose(t *testing.T) {
	tests := []struct {
		name              string
		matrix, wantTrans *Matrix
	}{
		{
			name:      "order 1",
			matrix:    &Matrix{order: 1, data: [][]int{{1}}},
			wantTrans: &Matrix{order: 1, data: [][]int{{1}}},
		},
		{
			name: "order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			wantTrans: &Matrix{
				order: 2,
				data: [][]int{
					{1, 3},
					{2, 4},
				},
			},
		},
		{
			name: "order 3",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			wantTrans: &Matrix{
				order: 3,
				data: [][]int{
					{5, 20, 4},
					{15, 0, 26},
					{18, 11, 0},
				},
			},
		},
		{
			name: "another order 3",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{-286, 44, 520},
					{468, -72, -70},
					{165, 305, -300},
				},
			},
			wantTrans: &Matrix{
				order: 3,
				data: [][]int{
					{-286, 468, 165},
					{44, -72, 305},
					{520, -70, -300},
				},
			},
		},
		{
			name: "order 4",
			matrix: &Matrix{
				order: 4,
				data: [][]int{
					{5, 15, 18, 1},
					{20, 0, 11, 2},
					{4, 26, 0, 3},
					{4, 5, 6, 7},
				},
			},
			wantTrans: &Matrix{
				order: 4,
				data: [][]int{
					{5, 20, 4, 4},
					{15, 0, 26, 5},
					{18, 11, 0, 6},
					{1, 2, 3, 7},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTrans := test.matrix.Transpose()
			if diff := cmp.Diff(gotTrans, test.wantTrans, unxOpt); diff != "" {
				t.Errorf("Transpose(\n%s) =\n%s, want\n%s; diff want -> got:\n%s", test.matrix, gotTrans, test.wantTrans, diff)
			}
		})
	}
}

// TestAdjoint verify method implementation
func TestAdjoint(t *testing.T) {
	tests := []struct {
		name               string
		matrix, wantMatrix *Matrix
	}{
		{
			name: "order 2",
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			wantMatrix: &Matrix{
				order: 2,
				data: [][]int{
					{4, -2},
					{-3, 1},
				},
			},
		},
		{
			name: "order 3",
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			wantMatrix: &Matrix{
				order: 3,
				data: [][]int{
					{-286, 468, 165},
					{44, -72, 305},
					{520, -70, -300},
				},
			},
		},
		{
			name: "order 4",
			matrix: &Matrix{
				order: 4,
				data: [][]int{
					{5, 15, 18, 1},
					{20, 0, 11, 2},
					{4, 26, 0, 3},
					{4, 5, 6, 7},
				},
			},
			wantMatrix: &Matrix{
				order: 4,
				data: [][]int{
					{-1525, 3120, 1100, -1145},
					{488, -354, 1975, -815},
					{3172, -511, -1930, 520},
					{-2196, -1092, -385, 8590},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotAdj, err := test.matrix.Adjoint()
			if err != nil {
				t.Fatalf("Adjoint(m) returned unexpected error; %v\n%s", err, test.matrix)
			}
			if diff := cmp.Diff(gotAdj, test.wantMatrix, unxOpt); diff != "" {
				t.Errorf("Adjoint(\n%s) =\n%s, want\n%s; diff want -> got:\n%s", test.matrix, gotAdj, test.wantMatrix, diff)
			}
		})
	}
}

// TestAdjoint_Error verifies error check in Adjoint
func TestAdjoint_Error(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
	}{
		{
			name:   "order 1",
			matrix: &Matrix{order: 1, data: [][]int{{1}}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := test.matrix.Adjoint(); err == nil {
				t.Fatalf("Adjoint(\n%s) returned non-nil error, want error; ", test.matrix)

			}
		})
	}
}

// TestInverseMod verify method implementation
func TestInverseMod(t *testing.T) {
	tests := []struct {
		name               string
		mod                int
		matrix, wantMatrix *Matrix
	}{
		{
			name: "order 2 mod 12",
			mod:  12,
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 5},
					{3, 4},
				},
			},
			wantMatrix: &Matrix{
				order: 2,
				data: [][]int{
					{4, 7},
					{9, 1},
				},
			},
		},
		{
			name: "order 3 mod 26",
			mod:  26,
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{6, 24, 1},
					{13, 16, 10},
					{20, 17, 15},
				},
			},
			wantMatrix: &Matrix{
				order: 3,
				data: [][]int{
					{8, 5, 10},
					{21, 8, 21},
					{21, 12, 8},
				},
			},
		},
		{
			name: "order 3 mod 27",
			mod:  27,
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			wantMatrix: &Matrix{
				order: 3,
				data: [][]int{
					{23, 9, 21},
					{11, 9, 2},
					{22, 23, 6},
				},
			},
		},
	}
	unxOpt := cmp.AllowUnexported(Matrix{})
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotInverse, err := test.matrix.InverseMod(test.mod)
			if err != nil {
				t.Fatalf("InverseMod(\n%s\n, %d) returned unexpected error; %v", test.matrix, test.mod, err)
			}
			if diff := cmp.Diff(gotInverse, test.wantMatrix, unxOpt); diff != "" {
				t.Errorf("InverseMod(\n%s\n, %d) =\n%s, want\n%s; diff want -> got:\n%s", test.matrix, test.mod, gotInverse, test.wantMatrix, diff)
			}
		})
	}
}

// TestInverseMod_Error verifies error check in InverseMod
func TestInverseMod_Error(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		mod    int
	}{
		{
			name:   "mod less than 2",
			matrix: &Matrix{order: 1, data: [][]int{{1}}},
			mod:    1,
		},
		{
			name:   "order 1",
			matrix: &Matrix{order: 1, data: [][]int{{1}}},
			mod:    12,
		},
		{
			name: "data out of bound",
			mod:  12,
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{12, 10},
				},
			},
		},
		{
			name: "order 2 without inverse",
			mod:  12,
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{3, 4},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := test.matrix.InverseMod(test.mod); err == nil {
				t.Fatalf("InverseMod(\n%s\n, %d) returned non-nil error, want error; ", test.matrix, test.mod)

			}
		})
	}
}

// TestVectorProductMod verifies implementation of matrix-vector multiplication
func TestVectorProductMod(t *testing.T) {
	tests := []struct {
		name                   string
		matrix                 *Matrix
		mod                    int
		multVector, wantVector []int
	}{
		{
			name: "order 2 mult 0",
			mod:  12,
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			multVector: []int{0, 0},
			wantVector: []int{0, 0},
		},
		{
			name: "order 3",
			mod:  27,
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			multVector: []int{2, 15, 13},
			wantVector: []int{10, 21, 20},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotVector, err := test.matrix.VectorProductMod(test.mod, test.multVector...)
			if err != nil {
				t.Fatalf("VectorProductMod(%d, %v) return unexpected error; %v", test.mod, test.multVector, err)
			}
			if diff := cmp.Diff(test.wantVector, gotVector); diff != "" {
				t.Errorf("VectorProductMod(%d, %v) = %v, want %v: diff want -> got %s", test.mod, test.multVector, gotVector, test.wantVector, diff)
			}
		})
	}
}

// TestVectorProductMod_Error verifies validations
func TestVectorProductMod_Error(t *testing.T) {
	tests := []struct {
		name       string
		matrix     *Matrix
		mod        int
		multVector []int
	}{
		{
			name: "vector different size than matrix order",
			mod:  2,
			matrix: &Matrix{
				order: 2,
				data: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			multVector: []int{0, 1, 2},
		},
		{
			name: "mod less than 2",
			mod:  1,
			matrix: &Matrix{
				order: 3,
				data: [][]int{
					{5, 15, 18},
					{20, 0, 11},
					{4, 26, 0},
				},
			},
			multVector: []int{2, 15, 13},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.matrix.VectorProductMod(test.mod, test.multVector...)
			if err == nil {
				t.Fatalf("VectorProductMod(%d, %v) return non-nil error", test.mod, test.multVector)
			}
		})
	}
}

// TestMatrixOrder verify order definition
func TestMatrixOrder(t *testing.T) {
	tests := []struct {
		matrix    *Matrix
		wantOrder int
	}{
		{matrix: &Matrix{order: 1}, wantOrder: 1},
		{matrix: &Matrix{order: 2}, wantOrder: 2},
		{matrix: &Matrix{order: 3}, wantOrder: 3},
		{matrix: &Matrix{order: 100}, wantOrder: 100},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("order %d", test.wantOrder), func(t *testing.T) {
			if test.matrix.Order() != test.wantOrder {
				t.Errorf("Order() = %d, want %d", test.matrix.Order(), test.wantOrder)
			}
		})
	}
}
