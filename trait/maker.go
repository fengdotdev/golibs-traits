package trait

type Maker[T any] interface {
	Make() T
}
