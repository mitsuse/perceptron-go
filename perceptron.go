package perceptron

import (
	"errors"

	"github.com/mitsuse/perceptron-go/vector"
)

type Perceptron struct {
}

func New() *Perceptron {
	p := &Perceptron{}

	return p
}

func (p *Perceptron) Learn(weight vector.Vector, label int, feature vector.Vector) error {
	update := feature.Clone()
	update.Scalar(float64(label))
	weight.Add(update)

	if weight.Undefined() {
		// TODO: Write the error message.
		return errors.New("")
	}

	return nil
}

type Learner interface {
	Learn(weight vector.Vector, label int, feature vector.Vector) error
}

type Instance interface {
	Label() int
	Extract(indexer Indexer, indexed bool) vector.Vector
}
