package matrix

import (
	"errors"
	"math"
)

type denseMatrix struct {
	valueSeq  []float64
	rows      int
	columns   int
	undefined bool
}

func NewDense(rows, columns int) func(valueSeq ...float64) (*denseMatrix, error) {
	initialize := func(valueSeq ...float64) (*denseMatrix, error) {
		if rows*columns != len(valueSeq) {
			// TODO: Write the error message.
			return nil, errors.New("")
		}

		m := &denseMatrix{
			valueSeq: make([]float64, len(valueSeq)),
			rows:     rows,
			columns:  columns,
		}

		return m, nil
	}

	return initialize
}

func NewZeroDense(rows, columns int) *denseMatrix {
	m := &denseMatrix{
		valueSeq: make([]float64, rows*columns),
		rows:     rows,
		columns:  columns,
	}

	return m
}

func (m *denseMatrix) Shape() (rows, columns int) {
	return m.rows, m.columns
}

func (m *denseMatrix) IsUndefined() bool {
	return m.undefined
}

func (m *denseMatrix) invalidate(row, column int) bool {
	return row < 0 || m.rows <= row || column < 0 || m.columns <= column
}

func (m *denseMatrix) addable(matrix Matrix) bool {
	rows, columns := matrix.Shape()
	return m.rows == rows && m.columns == columns
}

func (m *denseMatrix) multipliable(matrix Matrix) bool {
	rows, _ := matrix.Shape()
	return m.columns == rows
}

func (m *denseMatrix) Get(row, column int) (float64, error) {
	if m.invalidate(row, column) || m.IsUndefined() {
		// TODO: Write the error message.
		return 0, errors.New("")
	}

	return m.valueSeq[m.columns*row+column], nil
}

func (m *denseMatrix) Update(row, column int, value float64) Matrix {
	if m.invalidate(row, column) || m.IsUndefined() {
		m.undefined = true
		return m
	}

	m.valueSeq[m.columns*row+column] = value

	return m
}

func (m *denseMatrix) Add(matrix Matrix) Matrix {
	if !m.addable(matrix) {
		m.undefined = true

		return m
	}

	iter := matrix.NonZeros()
	for iter.HasNext() {
		row, column, value := iter.Get()
		m.valueSeq[m.columns*row+column] += value
	}

	return m
}

func (m *denseMatrix) Sub(matrix Matrix) Matrix {
	if !m.addable(matrix) {
		m.undefined = true

		return m
	}

	iter := matrix.NonZeros()
	for iter.HasNext() {
		row, column, value := iter.Get()
		m.valueSeq[m.columns*row+column] -= value
	}

	return m
}

func (m *denseMatrix) Mul(matrix Matrix) Matrix {
	_, columns := matrix.Shape()
	n := NewZeroDense(m.rows, columns)

	if !m.multipliable(matrix) {
		n.undefined = true

		return n
	}

	iter := matrix.NonZeros()
	for iter.HasNext() {
		k, column, value := iter.Get()

		for row := 0; row < m.rows; row++ {
			n.valueSeq[row*columns+column] += m.valueSeq[row*m.columns+k] * value
		}
	}

	return n
}

func (m *denseMatrix) Scalar(scalar float64) Matrix {
	for i := 0; i < len(m.valueSeq); i++ {
		m.valueSeq[i] *= scalar
	}

	return m
}

func (m *denseMatrix) Resize(rows, columns int) Matrix {
	valueSeq := make([]float64, rows*columns)

	iter := m.NonZeros()
	for iter.HasNext() {
		row, column, value := iter.Get()
		if rows <= row || columns <= column {
			continue
		}

		valueSeq[columns*row+column] = value
	}

	m.valueSeq = valueSeq
	m.rows = rows
	m.columns = columns

	return m
}

func (m *denseMatrix) Clone() Matrix {
	matrix := &denseMatrix{
		valueSeq: make([]float64, len(m.valueSeq)),
		rows:     m.rows,
		columns:  m.columns,
	}

	copy(matrix.valueSeq, m.valueSeq)

	return matrix
}

func (m *denseMatrix) NonZeros() Iter {
	iter := &denseNonZeroIter{
		matrix: m,
		row:    0,
		column: 0,
		value:  0,
	}

	return iter
}

func (m *denseMatrix) Max() (row, column int, maxValue float64, err error) {
	if m.rows == 0 || m.columns == 0 || m.IsUndefined() {
		// TODO: Write the error message.
		err = errors.New("")
		return
	}

	maxValue = math.Inf(-1)

	for offset, value := range m.valueSeq {
		if value > maxValue {
			row = offset / m.columns
			column = offset % m.columns
			maxValue = value
		}
	}

	return
}

type denseNonZeroIter struct {
	matrix *denseMatrix
	row    int
	column int
	value  float64
}

func (iter *denseNonZeroIter) HasNext() bool {
	for row := iter.row; row < iter.matrix.rows; row++ {
		for column := iter.column; column < iter.matrix.columns; column++ {
			value, _ := iter.matrix.Get(row, column)
			if value == 0 {
				continue
			}

			iter.row = row
			iter.column = column
			iter.value = value

			return true
		}
	}

	return false
}

func (iter *denseNonZeroIter) Get() (row, column int, value float64) {
	return iter.row, iter.column, iter.value
}
