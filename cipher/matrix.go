package cipher

import (
	"fmt"
	"strings"
)

// Matrix represents square matrices
type Matrix struct {
	order int
	data  [][]int
}

// String implements Stringer interface
func (m Matrix) String() string {
	var b strings.Builder
	for _, row := range m.data {
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
	m := &Matrix{order: order}
	m.data = make([][]int, order)
	for i := 0; i < order; i++ {
		row := make([]int, order)
		for j := 0; j < order; j++ {
			row[j] = data[(i*order)+j]
		}
		m.data[i] = row
	}
	return m, nil
}

// Determinant returns the matrix determinant. This algorithm is extremely slow O(n!) since it
// builds on the naive approach. Implement LU decomposition for better performance O(n^3).
func (m *Matrix) Determinant() (int, error) {
	if m.order < 1 {
		return 0, fmt.Errorf("determinant is undefined for order < 1")
	}
	if m.order == 1 {
		return m.data[0][0], nil
	}
	sign := 1
	var det int
	for i := 0; i < m.order; i++ {
		subM, _ := Minor(m, 0, i)
		minor, _ := subM.Determinant()
		det += sign * m.data[0][i] * minor
		sign *= -1
	}
	return det, nil
}

// IsInvertibleMod returns whether the matrix is invertible mod n. A matrix A with entries in Zn is
// invertible modulo n if and only if the residue of det(A) modulo n has an inverse modulo m. Also,
// A is invertible modulo n if m and the residue of det(A) modulo n have no common prime factors.
func (m *Matrix) IsInvertibleMod(n int) bool {
	for _, col := range m.data {
		for _, x := range col {
			if x < 0 || x >= n {
				return false
			}
		}
	}
	det, err := m.Determinant()
	if err != nil {
		return false
	}
	res := Residue(det, n)
	if _, err := ModularInverse(res, n); err != nil {
		return false
	}

	return true
}

// InverseMod returns a the inverted square matrix mod n.
func (m *Matrix) InverseMod(n int) (*Matrix, error) {
	if !m.IsInvertibleMod(n) {
		return nil, fmt.Errorf("matrix is not invertible mod %d", n)
	}
	det, _ := m.Determinant() // Neglect error since its checked by IsInvertibleMod
	res := Residue(det, n)
	inverse, _ := ModularInverse(res, n) // Neglect error since its checked by IsInvertibleMod
	adj, err := m.Adjoint()
	if err != nil {
		return nil, fmt.Errorf("failed to compute Adj(\n%s\n); %v", m, err)
	}
	for i := range adj.data {
		for j := range adj.data {
			adj.data[i][j] = Residue(adj.data[i][j]*inverse, n)
		}
	}
	return adj, nil
}

// Adjoint returns the adjoint matrix
func (m *Matrix) Adjoint() (*Matrix, error) {
	cof, err := m.Cofactor()
	if err != nil {
		return nil, fmt.Errorf("failed to compute cofactor matrix for \n%s; %v", m, err)
	}
	return cof.Transpose(), nil
}

// Cofactor returns the cofactor matrix
func (m *Matrix) Cofactor() (*Matrix, error) {
	cof := &Matrix{order: m.order, data: make([][]int, m.order)}
	for i := 0; i < m.order; i++ {
		row := make([]int, m.order)
		for j := 0; j < m.order; j++ {
			minor, _ := Minor(m, i, j) // Error is neglected since row & col are always in bound
			detM, err := minor.Determinant()
			if err != nil {
				return nil, fmt.Errorf("failed to compute det(m) for minor at row:%d col:%d\n%s;%v", i, j, m, err)
			}
			if (i+j)%2 == 0 {
				row[j] = detM
			} else {
				row[j] = detM * -1
			}
		}
		cof.data[i] = row
	}
	return cof, nil
}

// Transpose returns the transposed matrix
func (m *Matrix) Transpose() *Matrix {
	t := &Matrix{order: m.order, data: make([][]int, m.order)}
	for i := 0; i < m.order; i++ {
		t.data[i] = make([]int, m.order)
		for j := 0; j < m.order; j++ {
			t.data[i][j] = m.data[j][i]
		}
	}
	return t
}

// Minor returns the co-factor matrix of the given matrix at p, q (row, col).
func Minor(m *Matrix, p, q int) (*Matrix, error) {
	if 0 > p || p >= m.order || 0 > q || q >= m.order {
		return nil, fmt.Errorf("received row and/or col out of bound")
	}
	if m.order <= 1 {
		return &Matrix{}, nil
	}
	r := &Matrix{order: m.order - 1}
	r.data = make([][]int, 0, r.order)
	for row := 0; row < m.order; row++ {
		tmp := make([]int, 0, r.order)
		for col := 0; col < m.order; col++ {
			if row != p && col != q {
				tmp = append(tmp, m.data[row][col])
			}
		}
		if row != p {
			r.data = append(r.data, tmp)
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
