package matrix

type RVector struct {
	row       int
	undefined bool
	matrix    Matrix
}

func NewRVector(matrix Matrix, row int) *RVector {
	rows, _ := matrix.Shape()

	v := &RVector{
		row:       row,
		undefined: row >= rows,
		matrix:    matrix,
	}

	return v
}

func (v *RVector) Matrix() Matrix {
	return v.matrix
}

func (v *RVector) Size() int {
	_, column := v.matrix.Shape()
	return column
}

func (v *RVector) IsUndefined() bool {
	return v.undefined || v.matrix.IsUndefined()
}

func (v *RVector) Get(index int) (float64, error) {
	return v.matrix.Get(v.row, index)
}

func (v *RVector) Update(index int) *RVector {
	v.matrix.Update(v.row, index)

	return v
}
