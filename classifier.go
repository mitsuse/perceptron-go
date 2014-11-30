package perceptron

import (
	"github.com/mitsuse/perceptron-go/vector"
)

type Classifier struct {
	weight  *vector.DenseVector
	indexer Indexer
}

func NewClassifier(indexer Indexer) *Classifier {
	c := &Classifier{
		weight:  vector.NewZeroDense(0),
		indexer: indexer,
	}

	return c
}

func (c *Classifier) Weight() vector.Vector {
	if c.weight.Size() < c.indexer.Size() {
		c.weight.Resize(c.indexer.Size())
	}

	return c.weight
}

func (c *Classifier) Update(learner Learner, instance Instance) error {
	feature := instance.Extract(c.indexer, true)

	score, err := c.Weight().Dot(feature)
	if err != nil {
		return err
	}

	if score > 0 == (instance.Label() == 1) {
		return nil
	}

	if err := learner.Learn(c.Weight(), instance.Label(), feature); err != nil {
		return err
	}

	return nil
}

func (c *Classifier) Classify(instance Instance) (int, error) {
	feature := instance.Extract(c.indexer, false)

	score, err := c.Weight().Dot(feature)
	if err != nil {
		return 0, err
	}

	var label int
	if score > 0 {
		label = 1
	} else {
		label = -1
	}

	return label, nil
}

type Indexer interface {
	Size() int
	Index(identifier []int32, indexed bool) int
}
