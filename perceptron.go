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

func (p *Perceptron) Learn(inferencer Inferencer, iter InstanceIter) error {
	for iteration := 1; iteration <= p.iteration; iteration++ {
		for iter.HasNext() {
			instance := iter.Get()

			if err := p.learnInstance(inferencer, instance); err != nil {
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

func (p *Perceptron) learnInstance(inferencer Inferencer, instance Instance) error {
	inference, err := inferencer.Infer(instance)
	if err != nil {
		return err
	}

	if instance.Label() != inference.Label() {
		inferencer.Weight().Add(instance.Update())
		inferencer.Weight().Sub(inference.Update())
	}

	return nil
}

type Learner interface {
	Learn(inferencer Inferencer, iter InstanceIter) error
}

type Inferencer interface {
	Infer(instance Instance) (Instance, error)
	Weight() Matrix
}

type Instance interface {
	Label() int
	Feature() Matrix
	Update() Matrix
}

type InstanceIter interface {
	HasNext() bool
	Get() Instance
	Error() error
	Init() error
}

type Matrix interface {
	Add(matrix Matrix)
	Sub(matrix Matrix)
}
