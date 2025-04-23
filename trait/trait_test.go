package trait_test

import (
	"fmt"
	"math/rand"
)

// types and helper funcs for testing

func CoinFlip() bool {
	// Simulate a coin flip
	flip := rand.Intn(2) // 0 or 1
	return flip == 0     // true for heads, false for tails
}

// Bottle

type Bottle[T any] struct {
	volume  int
	content T
}

func NewBottleOfWater() Bottle[string] {
	return Bottle[string]{volume: 500, content: "water"}
}

func NewBottleOfJuice() Bottle[string] {
	return Bottle[string]{volume: 300, content: "juice"}
}

// Box

type Box[T any] struct {
	width   int
	height  int
	depth   int
	content T
}

func NewBoxOf[T any](width, height, depth int, content T) Box[T] {
	return Box[T]{width: width, height: height, depth: depth, content: content}
}

func NewBoxOfWaterBottle() Box[Bottle[string]] {
	return NewBoxOf(10, 20, 30, NewBottleOfWater())
}
func NewBoxOfJuiceBottle() Box[Bottle[string]] {
	return NewBoxOf(10, 20, 30, NewBottleOfJuice())
}

func NewBoxOfShoe() Box[Shoe] {
	return NewBoxOf(10, 20, 30, Shoe{model: "Sneakers", size: 42})
}

func NewBoxOfShoeWithContent(shoe Shoe) Box[Shoe] {
	return NewBoxOf(10, 20, 30, shoe)
}




// Shoe
type Shoe struct {
	model string
	size  int // 42
}

func NewOxfords() Shoe {
	return Shoe{model: "Oxford", size: 43}
}
func NewSneakers() Shoe {
	return Shoe{model: "Sneakers", size: 42}
}

func (s Shoe) String() string {
	return fmt.Sprintf("Shoe model: %s, size: %d", s.model, s.size)
}
