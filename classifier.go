package perceptron

type Classifier struct {
	weight Matrix
}

func (c *Classifier) Weight() Matrix {
	return c.weight
}

func (c *Classifier) Classify(instance Instance) (int, error) {
	// TODO: Implement this.
	return nil, nil
}
