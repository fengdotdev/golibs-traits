package future

type Future[T any] interface {
	GetError() chan error
	GetValue() chan T
	Get() (T, error)
}
