package matrix

type CVector struct {
	column int
	matrix Matrix
}

func (v *CVector) Matrix() Matrix {
	return v.matrix
}

func (v *RVector) Size() int {
	row, _ := v.matrix
	return row
}
