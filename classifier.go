package perceptron

import "github.com/mitsuse/perceptron-go/vector"

type Classifier struct {
	weight    *vector.DenseVector
	extractor Extractor
}

func NewClassifier(extractor Extractor) *Classifier {
	c := &Classifier{
		weight:    vector.NewZeroDens(0),
		extractor: extractor,
	}

	return c
}

func (c *Classifier) Weight() vector.Vector {
	return c.weight
}

func (c *Classifier) Classify(instance Instance) (Instance, error) {
	inference := instance.Clone()

	feature, err := c.extractor.Extract(instance)
	if err != nil {
		return nil, err
	}
	inference.SetFeature(feature)

	if c.Weight().Size() < feature.Size() {
		c.Weight().Resize(feature.Size())
	}

	score, err := c.Weight().Dot(feature)
	if err != nil {
		return nil, err
	}
	inference.SetScore(score)

	if score > 0 {
		inference.SetLabel(1)
	} else {
		inference.SetLabel(-1)
	}

	return inference, nil
}

type Extractor interface {
	Extract(instance Instance) (vector.Vector, error)
}
