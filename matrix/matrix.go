package matrix

type Matrix interface {
	Shape() (rows, columns int)
	IsUndefined() bool
	Get(row, column int) (float64, error)
	Update(row, column int, value float64) Matrix
	Add(matrix Matrix) Matrix
	Sub(matrix Matrix) Matrix
	Mul(matrix Matrix) Matrix
	Scalar(scalar float64) Matrix
	Resize(rows, columns int) Matrix
	Clone() Matrix
	NonZeros() Iter
	Max() (row, column int, maxValue float64, err error)
}

type Iter interface {
	HasNext() bool
	Get() (row, column int, value float64)
}
