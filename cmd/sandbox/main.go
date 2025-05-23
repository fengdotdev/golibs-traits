package main

import (
	"time"
)

type Stream[T any] interface {
	GetError() chan error
	GetValue() chan T
	Get() (chan T, chan error)
	Pause()
	Resume()
	Stop()
}

var _ Stream[int] = (*IntStream)(nil)

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

	// clock at 1 second

	go func() {
		i := 0
		paused := false
		for {
			select {
			case p, ok := <-stream.pause:
				if !ok {
					return
				}
				paused = p
			default:
				if paused {
					continue
				}
				select {
				case stream.value <- i:
					i++
				default:
					// evitar bloqueo si nadie está leyendo
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return stream
}
func TimeUP[T any](Stream Stream[T], timer time.Duration) {
	go func() {
		time.Sleep(timer)
		Stream.Stop()
	}()
}

func IntermitenPause[T any](stream Stream[T], timer time.Duration) {
	go func() {
		for {

			stream.Pause()
			time.Sleep(timer)
			stream.Resume()

		}
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

func main() {
	stream := NumberGenerator()
	TimeUP(stream, 30*time.Second)
	IntermitenPause(stream, 5*time.Second)
	StreamReader(stream)

}
