package hillcipher

import (
	"fmt"
	"strings"
)

// Matrix represents square matrices
type Matrix struct {
	Order int
	Data  [][]int
}

// String implements Stringer interface
func (m Matrix) String() string {
	var b strings.Builder
	for _, row := range m.Data {
		for i, item := range row {
			if i == 0 {
				b.WriteString("|")
			}
			b.WriteString("\t" + fmt.Sprintf("%d", item) + "\t|")
		}
		b.WriteString("\n")
	}
	return b.String()
}

// NewMatrix returns a new square matrix of the given order loaded with the given data.
// The size of the input data must be exactly order squared (order^2).
func NewMatrix(order int, data []int) (*Matrix, error) {
	if len(data) != order*order {
		return nil, fmt.Errorf("failed to build square matrix, got invalid data size %d, wantMatrix %d", len(data), order*order)
	}
	m := &Matrix{Order: order}
	m.Data = make([][]int, order)
	for i := 0; i < order; i++ {
		row := make([]int, order)
		for j := 0; j < order; j++ {
			row[j] = data[(i*order)+j]
		}
		m.Data[i] = row
	}
	return m, nil
}

// Determinant returns the matrix determinant. This algorithm is extremely slow O(n!) since it
// builds on the naive approach. Implement LU decomposition for better performance O(n^3).
func (m *Matrix) Determinant() (int, error) {
	if m.Order < 1 {
		return 0, fmt.Errorf("determinant is undefined for order < 1")
	}
	if m.Order == 1 {
		return m.Data[0][0], nil
	}
	sign := 1
	var det int
	for i := 0; i < m.Order; i++ {
		cofactor, _ := Minor(m, 0, i)
		minor, _ := cofactor.Determinant()
		det += sign * m.Data[0][i] * minor
		sign *= -1
	}
	return det, nil
}

// IsInvertibleMod returns whether the matrix is invertible mod n. A matrix A with entries in Zn is
// invertible modulo n if and only if the residue of det(A) modulo n has an inverse modulo m. Also,
// A is invertible modulo n if m and the residue of det(A) modulo n have no common prime factors.
func (m *Matrix) IsInvertibleMod(n int) bool {
	for _, col := range m.Data {
		for _, x := range col {
			if x < 0 || x >= (n) {
				return false
			}
		}
	}

	return true
}

// Minor returns the co-factor matrix of the given matrix at p, q (row, col).
func Minor(m *Matrix, p, q int) (*Matrix, error) {
	if 0 > p || p >= m.Order || 0 > q || q >= m.Order {
		return nil, fmt.Errorf("received row and/or col out of bound")
	}
	if m.Order <= 1 {
		return &Matrix{}, nil
	}
	r := &Matrix{Order: m.Order - 1}
	r.Data = make([][]int, 0, r.Order)
	for row := 0; row < m.Order; row++ {
		tmp := make([]int, 0, r.Order)
		for col := 0; col < m.Order; col++ {
			if row != p && col != q {
				tmp = append(tmp, m.Data[row][col])
			}
		}
		if row != p {
			r.Data = append(r.Data, tmp)
		}
	}
	return r, nil
}

// By that same reasoning, since the cipher used would have to be a
// matrix modulo 26, that inverse cipher would have to be the inverse
// of that matrix modulo 26.

// If m is a positive interger, then a square matrix A with entries in Zm is said to be
// "INVERTIBLE MODULO M" if there is a matrix B with entries in Zm such that
// AB = BA = I (mod m)
