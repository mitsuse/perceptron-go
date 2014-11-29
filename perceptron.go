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
	for iter.HasNext() {
		instance := iter.Get()

		if err := p.learnInstance(model, instance); err != nil {
			return err
		}
	}

	if err := iter.Error(); err != nil {
		return err
	}

	return nil
}

func (p *Perceptron) learnInstance(model Model, instance Instance) error {
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
	HasNext() bool
	Get() Instance
	Error() error
}
