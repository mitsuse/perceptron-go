package perceptron

import (
	"errors"

	"github.com/mitsuse/perceptron-go/matrix"
)

type Classifier struct {
	model *Model
}

func NewClassifier(size int, indexer Indexer) *Classifier {
	c := &Classifier{
		model: &Model{
			weight:  matrix.NewZeroDense(size, 0),
			indexer: indexer,
		},
	}

	return c
}

func (c *Classifier) Update(learner Learner, instance Instance) (int, error) {
	feature := c.model.Extract(instance, true)

	score := c.model.Score(feature)
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
		return label, learner.Learn(c.model, example, inference)
	}

	return label, nil
}

func (c *Classifier) getUpdate(label int, feature matrix.Matrix) matrix.Matrix {
	return nil
}

func (c *Classifier) Classify(instance Instance) (int, error) {
	feature := c.model.Extract(instance, false)

	score := c.model.Score(feature)
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

type Indexer interface {
	Size() int
	Index(identifier []int32, indexed bool) int
}
