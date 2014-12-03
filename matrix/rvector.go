package matrix

type RVector struct {
	row       int
	undefined bool
	matrix    Matrix
}

func (v *RVector) Matrix() Matrix {
	return v.matrix
}

func (v *RVector) Size() int {
	_, column := v.matrix
	return column
}

func (v *RVector) IsUndefined() bool {
	return v.undefined
}
