package future

type Stream[T any] interface {
	GetError() chan error
	GetValue() chan T
	Get() (chan T, chan error)
	Pause()
	Resume()
	Stop()
}
