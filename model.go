package perceptron

import (
	"github.com/mitsuse/perceptron-go/matrix"
)

type Model struct {
	weight  matrix.Matrix
	indexer Indexer
}

func (m *Model) Weight() matrix.Matrix {
	return m.weight
}

func (m *Model) Extract(instance Instance, indexed bool) matrix.Matrix {
	return instance.Extract(m.indexer, indexed)
}

func (m *Model) Score(feature matrix.Matrix) matrix.Matrix {
	weightRows, weightColumns := m.Weight().Shape()
	featuresRows, _ := feature.Size()

	if weightColumns < featuresRows {
		m.Weight.Resize(weightRows, featuresRows)
	}

	return m.Mul(feature)
}
