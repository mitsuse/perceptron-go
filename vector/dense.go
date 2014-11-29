package vector

type DenseVector struct {
	valueSeq []float64
}

func NewDense(valueSeq ...float64) *DenseVector {
	v := &DenseVector{
		valueSeq: make([]float64, len(valueSeq)),
	}

	for offset, value := range valueSeq {
		v.valueSeq[offset] = value
	}

	return v
}

func NewZeroDens(size int) *DenseVector {
	v := &DenseVector{
		valueSeq: make([]float64, size),
	}

	return v
}

func (v *DenseVector) Size() int {
	return len(v.valueSeq)
}

func (v *DenseVector) Add(vector Vector) {
	// TODO: Implement this.
}

func (v *DenseVector) Dot(vector Vector) (float64, error) {
	// TODO: Implement this.
	return 0, nil
}

func (v *DenseVector) Extend(size int) {
	// TODO: Implement this.
}
