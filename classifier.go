package perceptron

import (
	"errors"

	"github.com/mitsuse/perceptron-go/matrix"
)

type Classifier struct {
	scorer *Scorer
}

func NewClassifier(size int) *Classifier {
	c := &Classifier{
		scorer: NewScorer(size),
	}

	return c
}

func (c *Classifier) Update(learner Learner, instance Instance) (int, error) {
	feature := c.scorer.Extract(instance, true)

	score := c.scorer.Score(feature)
	if score.IsUndefined() {
		// TODO: Write the error message.
		return 0, errors.New("")
	}

	label, _, _, err := score.Max()
	if err != nil {
		return 0, err
	}

	if label != instance.Label() {
		example := c.getUpdate(instance.Label(), feature)
		inference := c.getUpdate(label, feature)
		return label, learner.Learn(c.scorer, example, inference)
	}

	return label, nil
}

func (c *Classifier) getUpdate(label int, feature matrix.Matrix) matrix.Matrix {
	rows, columns := c.scorer.Weight().Shape()
	update := matrix.ZeroDense(rows, columns)

	iter := feature.NonZeros()
	for iter.HasNext() {
		id, _, value := iter.Get()
		update.Update(label, id, value)
	}

	return update
}

func (c *Classifier) Classify(instance Instance) (int, error) {
	feature := c.scorer.Extract(instance, false)

	score := c.scorer.Score(feature)
	if score.IsUndefined() {
		// TODO: Write the error message.
		return 0, errors.New("")
	}

	label, _, _, err := score.Max()
	if err != nil {
		return 0, err
	}

	return label, nil
}
