package perceptron

import (
	"github.com/mitsuse/perceptron-go/matrix"
)

type Scorer struct {
	weight  *matrix.DenseMatrix
	indexer *Indexer
}

func NewScorer(size int) *Scorer {
	s := &Scorer{
		weight:  matrix.ZeroDense(size, 0),
		indexer: NewIndexer(),
	}

	return s
}

func (s *Scorer) Weight() matrix.Matrix {
	return s.weight
}

func (s *Scorer) Extract(instance Instance, indexed bool) matrix.Matrix {
	return instance.Extract(s.indexer, indexed)
}

func (s *Scorer) Score(feature matrix.Matrix) matrix.Matrix {
	weightRows, weightColumns := s.Weight().Shape()
	featuresRows, _ := feature.Shape()

	if weightColumns < featuresRows {
		s.Weight().Resize(weightRows, featuresRows)
	}

	return s.Weight().Mul(feature)
}
