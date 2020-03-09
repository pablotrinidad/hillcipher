package hillcipher

import (
	"fmt"
	"strings"
)

// Matrix represents square matrices
type Matrix struct {
	Order int
	Data  [][]float64
}

func (m Matrix) String() string {
	var b strings.Builder
	for _, row := range m.Data {
		for i, item := range row {
			if i == 0 {
				b.WriteString("|")
			}
			b.WriteString("\t" + fmt.Sprintf("%.2f", item) + "\t|")
		}
		b.WriteString("\n")
	}
	return b.String()
}

// NewMatrix returns a new square matrix of the given order loaded with the given data.
// The size of the input data must be exactly order squared (order^2).
func NewMatrix(order int, data []float64) (*Matrix, error) {
	if len(data) != order*order {
		return nil, fmt.Errorf("failed to build square matrix, got invalid data size %d, wantMatrix %d", len(data), order*order)
	}
	m := &Matrix{Order: order}
	m.Data = make([][]float64, order)
	for i := 0; i < order; i++ {
		row := make([]float64, order)
		for j := 0; j < order; j++ {
			row[j] = data[(i*order)+j]
		}
		m.Data[i] = row
	}
	return m, nil
}

// Determinant returns the matrix determinant
func (m *Matrix) Determinant() (float64, error) {
	if m.Order < 1 {
		return 0.0, fmt.Errorf("determinant is undefined for order < 1")
	}
	if m.Order == 1 {
		return m.Data[0][0], nil
	}
	sign := 1.0
	var det float64
	for i := 0; i < m.Order; i++ {
		cofactor, _ := Minor(m, 0, i)
		minor, _ := cofactor.Determinant()
		det += sign * m.Data[0][i] * minor
		sign *= -1
	}
	return det, nil
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
	r.Data = make([][]float64, 0, r.Order)
	for row := 0; row < m.Order; row++ {
		tmp := make([]float64, 0, r.Order)
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
