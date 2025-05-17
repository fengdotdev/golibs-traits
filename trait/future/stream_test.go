package future_test

import (
	"time"

	"github.com/fengdotdev/golibs-traits/trait/future"
)

type Stream[T any] interface {
	GetError() chan error
	GetValue() chan T
	Get() (chan T, chan error)
	Pause()
	Resume()
	Stop()
}

var _ future.Stream[int] = (*IntStream)(nil)

type IntStream struct {
	value chan int
	err   chan error
	pause chan bool
}

func NewIntStream() *IntStream {
	return &IntStream{
		value: make(chan int, 1),
		err:   make(chan error, 1),
		pause: make(chan bool),
	}
}

// Get implements future.Stream.
func (i *IntStream) Get() (chan int, chan error) {
	return i.value, i.err
}

// GetError implements future.Stream.
func (i *IntStream) GetError() chan error {
	return i.err
}

// GetValue implements future.Stream.
func (i *IntStream) GetValue() chan int {
	return i.value
}

// Pause implements future.Stream.
func (i *IntStream) Pause() {
	i.pause <- true
}

// Resume implements future.Stream.
func (i *IntStream) Resume() {
	i.pause <- false
}

// Stop implements future.Stream.
func (i *IntStream) Stop() {
	close(i.value)
	close(i.err)
}

func NumberGenerator() *IntStream {

	stream := NewIntStream()
	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)
			select {
			case <-stream.pause:

			default:

				stream.value <- i

			}

		}

	}()
	return stream
}

func TimeUP[T any](timer time.Duration, Stream Stream[T]) {
	go func() {
		time.Sleep(timer)
		Stream.Stop()
	}()
}

func StreamReader(stream *IntStream) {
	for {
		select {
		case v, ok := <-stream.GetValue():
			if !ok {
				return
			}
			println(v)
		case err := <-stream.GetError():
			println(err.Error())
		}
	}
}
