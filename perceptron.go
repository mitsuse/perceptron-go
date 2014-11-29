package perceptron

type Perceptron struct {
	iteration int
}

func New(iteration byte) *Perceptron {
	p := &Perceptron{
		iteration: int(iteration),
	}

	return p
}

func (p *Perceptron) Learn(model Model, iter InstanceIter) error {
	// TODO: Implement this.
	return nil
}

type Learner interface {
	Learn(model Model, iter InstanceIter) error
}

type Model interface {
}

type Instance interface {
}

type InstanceIter interface {
}
