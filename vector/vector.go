package vector

type Vector interface {
	Size() int
	Get(index int) (float64, error)
	Add(vector Vector)
	Dot(vector Vector) (float64, error)
	Resize(size int)
	NonZeros() Iter
}

type Iter interface {
	HasNext() bool
	Get() (index int, value float64)
}
