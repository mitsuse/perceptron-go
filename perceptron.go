package perceptron

import (
	"errors"

	"github.com/mitsuse/perceptron-go/vector"
)

type Perceptron struct {
	iteration int
}

func New(iteration byte) *Perceptron {
	p := &Perceptron{
		iteration: int(iteration),
	}

	return p
}

func (p *Perceptron) Learn(classifier *Classifier, iter InstanceIter) error {
	for iteration := 1; iteration <= p.iteration; iteration++ {
		for iter.HasNext() {
			instance := iter.Get()

			if err := p.learnInstance(classifier, instance); err != nil {
				return err
			}
		}

		if err := iter.Error(); err != nil {
			return err
		}

		if err := iter.Init(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Perceptron) learnInstance(classifier *Classifier, instance Instance) error {
	inference, err := classifier.Classify(instance)
	if err != nil {
		return err
	}

	if instance.Label() == inference.Label() {
		classifier.Weight().Add(inference.Update())

		if classifier.Weight().Undefined() {
			// TODO: Write the error message.
			return errors.New("")
		}
	}

	return nil
}

type Learner interface {
	Learn(classifier *Classifier, iter InstanceIter) error
}

type Instance interface {
	Label() int
	Score() float64
	Feature() vector.Vector

	SetLabel(label int)
	SetScore(score float64)
	SetFeature(vector vector.Vector)

	Update() vector.Vector
	Clone() Instance
}

type InstanceIter interface {
	HasNext() bool
	Get() Instance
	Error() error
	Init() error
}
