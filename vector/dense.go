package vector

import (
	"errors"
)

type DenseVector struct {
	valueSeq  []float64
	undefined bool
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

func NewZeroDense(size int) *DenseVector {
	v := &DenseVector{
		valueSeq: make([]float64, size),
	}

	return v
}

func (v *DenseVector) Size() int {
	return len(v.valueSeq)
}

func (v *DenseVector) Undefined() bool {
	return v.undefined
}

func (v *DenseVector) Get(index int) (float64, error) {
	if index < 0 || v.Size() <= index {
		// TODO: Write the error message.
		return 0, errors.New("")
	}

	return v.valueSeq[index], nil
}

func (v *DenseVector) Add(vector Vector) {
	if v.Size() != vector.Size() || v.Undefined() || vector.Undefined() {
		v.valueSeq = []float64{}
		v.undefined = true
		return
	}

	iter := vector.NonZeros()
	for iter.HasNext() {
		index, value := iter.Get()
		v.valueSeq[index] += value
	}
}

func (v *DenseVector) Scalar(scalar float64) {
	if v.Undefined() {
		return
	}

	iter := v.NonZeros()
	for iter.HasNext() {
		index, value := iter.Get()
		v.valueSeq[index] = value * scalar
	}
}

func (v *DenseVector) Dot(vector Vector) (float64, error) {
	if v.Size() != vector.Size() || v.Undefined() || vector.Undefined() {
		// TODO: Write the error message.
		return 0.0, errors.New("")
	}

	product := 0.0

	iter := vector.NonZeros()
	for iter.HasNext() {
		index, value := iter.Get()
		product += v.valueSeq[index] * value
	}

	return product, nil
}

func (v *DenseVector) Resize(size int) {
	valueSeq := make([]float64, size)

	var limit int
	if size < v.Size() {
		limit = size
	} else {
		limit = v.Size()
	}

	for offset := 0; offset < limit; offset++ {
		valueSeq[offset] = v.valueSeq[offset]
	}

	v.valueSeq = valueSeq
}

func (v *DenseVector) NonZeros() Iter {
	iter := &denseNonZeroIter{
		vector: v,
		index:  -1,
	}

	return iter
}

type denseNonZeroIter struct {
	vector *DenseVector
	index  int
	value  float64
}

func (iter *denseNonZeroIter) HasNext() bool {
	for index := iter.index + 1; index < iter.vector.Size(); index++ {
		value, _ := iter.vector.Get(index)
		if value == 0 {
			continue
		}

		iter.index = index
		iter.value = value

		return true
	}

	return false
}

func (iter *denseNonZeroIter) Get() (index int, value float64) {
	return iter.index, iter.value
}
