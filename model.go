package perceptron

import (
	"github.com/mitsuse/perceptron-go/matrix"
	"github.com/mitsuse/perceptron-go/trie"
)

type Model struct {
	weight  *matrix.DenseMatrix
	indexer *Indexer
}

func NewModel(size int) *Model {
	m := &Model{
		weight:  matrix.ZeroDense(size, 0),
		indexer: NewIndexer(),
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
	trie *trie.SimpleTrie
	size int
}

func NewIndexer() *Indexer {
	i := &Indexer{
		trie: trie.NewSimpleTrie(),
		size: 0,
	}

	return i
}

func (i *Indexer) Size() int {
	// TODO: Implement this.
	return 0
}

func (i *Indexer) Index(identifier []int32, indexed bool) int {
	// TODO: Implement this.
	return 0
}
