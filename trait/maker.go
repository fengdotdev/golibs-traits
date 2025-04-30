package trait


// NOT READY FOR PRODUCTION
type Maker[T any] interface {
	Make(args ...any) T
}
