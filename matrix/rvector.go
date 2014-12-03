package matrix

type RVector struct {
	row    int
	matrix Matrix
}

func (v *RVector) Matrix() Matrix {
	return v.matrix
}

func (v *RVector) Size() int {
	_, column := v.matrix
	return column
}
