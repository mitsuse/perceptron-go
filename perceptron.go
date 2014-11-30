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

func (p *Perceptron) Learn(weight vector.Vector, example, inference Instance) error {
	if example.Label() != inference.Label() {
		weight.Add(example.Update())

		if weight.Undefined() {
			// TODO: Write the error message.
			return errors.New("")
		}
	}

	return nil
}

type Learner interface {
	Learn(weight vector.Vector, example, inference Instance) error
}

type Instance interface {
	Label() int
	Feature() vector.Vector

	SetLabel(label int)
	SetFeature(vector vector.Vector)

	Update() vector.Vector
	Clone() Instance
}
