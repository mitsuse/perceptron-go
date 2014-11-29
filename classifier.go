package perceptron

type Classifier struct {
	weight    Vector
	extractor Extractor
}

func NewClassifier(extractor Extractor) *Classifier {
	// TODO: Initialize weight with zero-vector.
	c := &Classifier{
		weight:    nil,
		extractor: extractor,
	}

	return c
}

func (c *Classifier) Weight() Vector {
	return c.weight
}

func (c *Classifier) Classify(instance Instance) (Instance, error) {
	inference := instance.Clone()

	feature, err := c.extractor.Extract(instance)
	if err != nil {
		return nil, err
	}
	inference.SetFeature(feature)

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
	Extract(instance Instance) (Vector, error)
}
