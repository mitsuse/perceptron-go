package perceptron

import (
	"errors"

	"github.com/mitsuse/perceptron-go/matrix"
)

type Perceptron struct {
}

func New() *Perceptron {
	p := &Perceptron{}

	return p
}

func (p *Perceptron) Learn(scorer *Scorer, example, infernce matrix.Matrix) error {
	scorer.Weight().Add(example).Sub(infernce)

	if scorer.Weight().IsUndefined() {
		// TODO: Write the error message.
		return errors.New("")
	}

	return nil
}

type Learner interface {
	Learn(scorer *Scorer, example, infernce matrix.Matrix) error
}

type Instance interface {
	Label() int
	Extract(indexer *Indexer, indexed bool) matrix.Matrix
}
