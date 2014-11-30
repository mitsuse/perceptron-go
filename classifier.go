package perceptron

import "github.com/mitsuse/perceptron-go/vector"

type Classifier struct {
	weight    *vector.DenseVector
	extractor Extractor
}

func NewClassifier(extractor Extractor) *Classifier {
	c := &Classifier{
		weight:    vector.NewZeroDense(0),
		extractor: extractor,
	}

	return c
}

func (c *Classifier) Weight() vector.Vector {
	return c.weight
}

func (c *Classifier) Infer(instance Instance) (example, inference Instance, err error) {
	example = instance.Clone()
	inference = instance.Clone()

	feature, err := c.extractor.Extract(instance)
	if err != nil {
		return nil, nil, err
	}

	example.SetFeature(feature)
	inference.SetFeature(feature)

	if c.Weight().Size() < feature.Size() {
		c.Weight().Resize(feature.Size())
	}

	score, err := c.Weight().Dot(feature)
	if err != nil {
		return nil, nil, err
	}

	example.SetLabel(instance.Label())
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

type Extractor interface {
	Extract(instance Instance) (vector.Vector, error)
}
