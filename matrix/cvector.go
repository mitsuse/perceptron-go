package matrix

type CVector struct {
	column    int
	undefined bool
	matrix    Matrix
}

func NewCVector(matrix Matrix, column int) *CVector {
	_, columns := matrix.Shape()

	v := &CVector{
		column:    column,
		undefined: column >= columns,
		matrix:    matrix,
	}

	return v
}

func (v *CVector) Matrix() Matrix {
	return v.matrix
}

func (v *CVector) Size() int {
	row, _ := v.matrix.Shape()
	return row
}

func (v *CVector) IsUndefined() bool {
	return v.undefined || v.matrix.IsUndefined()
}

func (v *CVector) Get(index int) (float64, error) {
	return v.matrix.Get(index, v.column)
}

func (v *CVector) Update(index int, value float64) *CVector {
	v.matrix.Update(index, v.column, value)

	return v
}
