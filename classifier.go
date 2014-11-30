package perceptron

import (
	"github.com/mitsuse/perceptron-go/vector"
)

type Classifier struct {
	model *Model
}

func NewClassifier(indexer Indexer) *Classifier {
	c := &Classifier{
		model: &Model{
			weight:  vector.NewZeroDense(0),
			indexer: indexer,
		},
	}

	return c
}

func (c *Classifier) Weight() vector.Vector {
	return c.model.Weight()
}

func (c *Classifier) Update(learner Learner, instance Instance) error {
	feature := c.model.Extract(instance, true)

	score, err := c.model.Score(feature)
	if err != nil {
		return err
	}

	if score > 0 == (instance.Label() == 1) {
		return nil
	}

	if err := learner.Learn(c.model, instance.Label(), feature); err != nil {
		return err
	}

	return nil
}

func (c *Classifier) Classify(instance Instance) (int, error) {
	feature := c.model.Extract(instance, false)

	score, err := c.model.Score(feature)
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
