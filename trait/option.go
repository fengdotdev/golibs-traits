package trait

// NOT READY FOR PRODUCTION
type Option[T any] interface {
	IsSome() bool // IsSome returns true if the Option is Some
	IsNone() bool // IsNone returns true if the Option is None
	Unwrap() T // Unwrap returns the value if the Option is Some, panics if it is None
	UnwrapOr(defaultValue T) T // UnwrapOr returns the value if the Option is Some, otherwise returns the default value
	UnwrapOrElse(f func() T) T // UnwrapOrElse returns the value if the Option is Some, otherwise calls the function and returns its result
	UnwrapOrDefault() T // UnwrapOrDefault returns the value if the Option is Some, otherwise returns the zero value of the type
}
