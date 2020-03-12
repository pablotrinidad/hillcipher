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
		subM, _ := Minor(m, 0, i)
		minor, _ := subM.Determinant()
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
	det, err := m.Determinant()
	if err != nil {
		return nil, fmt.Errorf("failed to compute det(\n%s\n); %v", m, err)
	}
	res := Residue(det, n)
	inverse, err := ModularInverse(res, n)
	if err != nil {
		return nil, fmt.Errorf("failed to compute ModularInverse(\n%d\n, %d); %v", res, n, err)
	}
	adj, err := m.Adjoint()
	if err != nil {
		return nil, fmt.Errorf("failed to compute Adj(\n%s\n); %v", m, err)
	}
	for i := range adj.Data {
		for j := range adj.Data {
			adj.Data[i][j] = Residue(adj.Data[i][j]*inverse, n)
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
	cof := &Matrix{Order: m.Order, Data: make([][]int, m.Order)}
	for i := 0; i < m.Order; i++ {
		row := make([]int, m.Order)
		for j := 0; j < m.Order; j++ {
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
		cof.Data[i] = row
	}
	return cof, nil
}

// Transpose returns the transposed matrix
func (m *Matrix) Transpose() *Matrix {
	t := &Matrix{Order: m.Order, Data: make([][]int, m.Order)}
	for i := 0; i < m.Order; i++ {
		t.Data[i] = make([]int, m.Order)
		for j := 0; j < m.Order; j++ {
			t.Data[i][j] = m.Data[j][i]
		}
	}
	return t
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
