package perceptron

import (
	"github.com/mitsuse/perceptron-go/vector"
)

type Model struct {
	weight  *vector.DenseVector
	indexer Indexer
}

func (m *Model) Weight() vector.Vector {
	return m.weight
}

func (m *Model) Extract(instance Instance, indexed bool) vector.Vector {
	return instance.Extract(m.indexer, indexed)
}

func (m *Model) Score(feature vector.Vector) (float64, error) {
	weight := m.Weight()
	if weight.Size() < feature.Size() {
		weight.Resize(feature.Size())
	}

	score, err := weight.Dot(feature)
	if err != nil {
		return 0, err
	}

	return score, nil
}
