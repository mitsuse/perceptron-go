package perceptron

import "github.com/mitsuse/perceptron-go/vector"

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
	return c.weight
}

func (c *Classifier) Infer(instance Instance) (example, inference Instance, err error) {
	example = instance.Clone()
	example.Extract(c.indexer)
	example.SetLabel(instance.Label())

	feature := example.Feature()

	if c.Weight().Size() < feature.Size() {
		c.Weight().Resize(feature.Size())
	}

	score, err := c.Weight().Dot(feature)
	if err != nil {
		return nil, nil, err
	}

	inference = instance.Clone()
	if score > 0 {
		inference.SetLabel(1)
	} else {
		inference.SetLabel(-1)
	}

	return example, inference, nil
}

func (c *Classifier) Classify(instance Instance) (label int, err error) {
	_, inference, err := c.Infer(instance)
	if err != nil {
		return 0, err
	}

	return inference.Label(), nil
}

type Indexer interface {
	Size() int
	Index(identifier []int32) int
}
