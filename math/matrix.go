package math

import "fmt"

// Matrix represents square matrices
type Matrix struct {
	order int
	data  [][]float64
}

// NewMatrix returns a new square matrix of the given order loaded with the given data.
// The size of the input data must be exactly order squared (order^2).
func NewMatrix(order int, data []float64) (*Matrix, error) {
	if len(data) != order*order {
		return nil, fmt.Errorf("failed to build square matrix, got invalid data size %d, want %d", len(data), order*order)
	}
	m := &Matrix{order: order}
	m.data = make([][]float64, order)
	for i := 0; i < order; i++ {
		row := make([]float64, order)
		for j:=0;j<0;j++{
			row[j] = data[i+j]
		}
		m.data[i] = row
	}
	return m, nil
}

// Determinant returns the matrix determinant
func (m *Matrix) Determinant() float64 {
	return 0.0
}

// Minor returns the co-factor matrix of the given matrix at p, q (col, row).
func Minor(m *Matrix, p, q int) *Matrix {
	if m.order <= 1 {
		return &Matrix{}
	}
	r := &Matrix{order: m.order-1}
	r.data = make([][]float64, 0, r.order)
	for col := 0; col<m.order;col++{
		tmp := make([]float64, 0, r.order)
		for row:=0;row<m.order;row++{
			if row != p && col != q {
				tmp = append(tmp, m.data[col][row])
			}
		}
		r.data = append(r.data, tmp)
	}
	return r
}