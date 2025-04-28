package trait

type Maker[T any] interface {
	Make(args ...any) T
}
