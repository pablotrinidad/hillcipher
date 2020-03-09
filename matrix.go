package hillcipher

import "fmt"

// Matrix represents square matrices
type Matrix struct {
	Order int
	Data  [][]float64
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
		for j := 0; j < 0; j++ {
			row[j] = data[i+j]
		}
		m.Data[i] = row
	}
	return m, nil
}

// Determinant returns the matrix determinant
func (m *Matrix) Determinant() float64 {
	return 0.0
}

// Minor returns the co-factor matrix of the given matrix at p, q (col, row).
func Minor(m *Matrix, p, q int) *Matrix {
	if m.Order <= 1 {
		return &Matrix{}
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
	return r
}
