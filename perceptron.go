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

	if instance.Label() != inference.Label() {
		classifier.Weight().Add(inference.Update())
	}

	return nil
}

type Learner interface {
	Learn(classifier *Classifier, iter InstanceIter) error
}

type Instance interface {
	Label() int
	Feature() Vector
	Update() Vector
}

type InstanceIter interface {
	HasNext() bool
	Get() Instance
	Error() error
	Init() error
}

type Vector interface {
	Add(matrix Vector)
	Sub(matrix Vector)
}
