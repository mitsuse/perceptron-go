package vector

type Vector interface {
	Size() int
	Add(vector Vector)
	Dot(vector Vector) (float64, error)
	Extend(size int)
}
