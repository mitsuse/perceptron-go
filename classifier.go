package perceptron

type Classifier struct {
	weight Matrix
}

func (c *Classifier) Weight() Matrix {
	return c.weight
}

func (c *Classifier) Extract(instance Instance) (Vector, error) {
	// TODO: Implement this.
	return nil, nil
}

func (c *Classifier) Classify(instance Instance) (Instance, error) {
	inference := instance.Clone()

	feature, err := c.Extract(instance)
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
