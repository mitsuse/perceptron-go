package vector

type Vector interface {
	Size() int
	Undefined() bool
	Get(index int) (float64, error)
	Set(index int, value float64) error
	Add(vector Vector)
	Scalar(scalar float64)
	Dot(vector Vector) (float64, error)
	Resize(size int)
	NonZeros() Iter
	Clone() Vector
}

type Iter interface {
	HasNext() bool
	Get() (index int, value float64)
}
