package perceptron

import (
	"github.com/mitsuse/perceptron-go/matrix"
)

type Model struct {
	weight  *matrix.DenseMatrix
	indexer *Indexer
}

func NewModel(size int) *Model {
	// TODO: Create indexer.
	m := &Model{
		weight: matrix.ZeroDense(size, 0),
	}

	return m
}

func (m *Model) Weight() matrix.Matrix {
	return m.weight
}

func (m *Model) Extract(instance Instance, indexed bool) matrix.Matrix {
	return instance.Extract(m.indexer, indexed)
}

func (m *Model) Score(feature matrix.Matrix) matrix.Matrix {
	weightRows, weightColumns := m.Weight().Shape()
	featuresRows, _ := feature.Shape()

	if weightColumns < featuresRows {
		m.Weight().Resize(weightRows, featuresRows)
	}

	return m.Weight().Mul(feature)
}

type Indexer struct {
}

func Size() int {
	// TODO: Implement this.
	return 0
}

func Index(identifier []int32, indexed bool) int {
	// TODO: Implement this.
	return 0
}
