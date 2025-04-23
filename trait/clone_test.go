package trait_test

import (
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/trait"
)

func (s *Shoe) Clone() Shoe {
	return Shoe{
		model: s.model,
		size:  s.size,
	}
}

type CloneableShoe = trait.Cloneable[Shoe]

func TestClone_Shoe(t *testing.T) {


	var _ CloneableShoe = &Shoe{}
	var misterytrait trait.Cloneable[Shoe] = &Shoe{}
	assert.NotNil(t, misterytrait)
	assert.NotNil(t, &Shoe{})

	runningshoe := NewSneakers()

	// Clone the running shoe

	clonedShoe := runningshoe.Clone()

	// basic equality
	assert.EqualWithMessage(t, runningshoe, clonedShoe, "expected the same")

	// memory address
	assert.NotEqualWithMessage(t, &runningshoe, &clonedShoe, fmt.Sprintf("expected different memory address for %v and %v", &runningshoe, &clonedShoe))
	assert.EqualWithMessage(t, runningshoe.size, clonedShoe.size, fmt.Sprintf("expected same size for %v and %v", runningshoe, clonedShoe))

	assert.EqualWithMessage(t, runningshoe.model, clonedShoe.model, fmt.Sprintf("expected same model for %v and %v ", runningshoe.model, clonedShoe.model))

	// Change the size of the cloned shoe
	clonedShoe.size = 44

	assert.NotEqualWithMessage(t, runningshoe.size, clonedShoe.size, fmt.Sprintf("expected different size for %v and %v", runningshoe, clonedShoe))

	clonedShoe.model = "ClonedSneakers"

	assert.NotEqualWithMessage(t, runningshoe.model, clonedShoe.model, fmt.Sprintf("expected different model for %v and %v", runningshoe, clonedShoe))

}

func (b *Bottle[T]) Clone() Bottle[T] {
	return Bottle[T]{
		volume:  b.volume,
		content: b.content,
	}
}

type CloneableBottle[T any] = trait.Cloneable[Bottle[T]]
type CloneableBottleOfWater = trait.Cloneable[Bottle[string]]
type CloneableBottleOfJuice = trait.Cloneable[Bottle[string]]

func TestClone_Bottle(t *testing.T) {
	bottle := NewBottleOfWater()

	var _ CloneableBottleOfWater = &bottle

	var misterybottle trait.Cloneable[Bottle[string]] = &bottle
	assert.NotNil(t, misterybottle)

	// Clone the bottle
	clonedBottle := misterybottle.Clone()

	// basic equality
	assert.EqualWithMessage(t, bottle, clonedBottle, fmt.Sprintf("expected same volume for %v and %v", bottle, clonedBottle))

	// memory address
	assert.NotEqualWithMessage(t, &bottle, &clonedBottle, fmt.Sprintf("expected different memory address for %v and %v", &bottle, &clonedBottle))

	// basic equality inner
	assert.EqualWithMessage(t, bottle.volume, clonedBottle.volume, fmt.Sprintf("expected same volume for %v and %v", bottle.volume, clonedBottle.volume))

	// Change the volume of the cloned bottle

	clonedBottle.volume = 1000
	assert.NotEqualWithMessage(t, bottle.volume, clonedBottle.volume, fmt.Sprintf("expected different volume for %v and %v", bottle.volume, clonedBottle.volume))
	assert.EqualWithMessage(t, bottle.content, clonedBottle.content, fmt.Sprintf("expected same content for %v and %v", bottle.content, clonedBottle.content))

	// Change the content of the cloned bottle
	clonedBottle.content = "juice"

	assert.NotEqualWithMessage(t, bottle.content, clonedBottle.content, fmt.Sprintf("expected different content for %v and %v", bottle.content, clonedBottle.content))

}

func (b *Box[T]) Clone() Box[T] {
	return Box[T]{
		width:   b.width,
		height:  b.height,
		depth:   b.depth,
		content: b.content,
	}
}

type CloneableBox[T any] = trait.Cloneable[Box[T]]
type CloneableBoxOfWaterBottle = trait.Cloneable[Box[Bottle[string]]]
type CloneableBoxOfJuiceBottle = trait.Cloneable[Box[Bottle[string]]]

type CloneableBoxOfShoe = trait.Cloneable[Box[Shoe]]

func TestClone_Box(t *testing.T) {

	box := NewBoxOfWaterBottle()

	var misterybox trait.Cloneable[Box[Bottle[string]]] = &box
	assert.NotNil(t, misterybox)

	// Clone the box
	clonedBox := misterybox.Clone()

	// basic equality
	assert.EqualWithMessage(t, box, clonedBox, fmt.Sprintf("expected same volume for %v and %v", box, clonedBox))

	// memory address
	assert.NotEqualWithMessage(t, &box, &clonedBox, fmt.Sprintf("expected different memory address for %v and %v", &box, &clonedBox))

	// basic equality inner
	assert.EqualWithMessage(t, box.width, clonedBox.width, fmt.Sprintf("expected same width for %v and %v", box.width, clonedBox.width))

	// Change the width of the cloned box

	clonedBox.width = 1000
	assert.NotEqualWithMessage(t, box.width, clonedBox.width, fmt.Sprintf("expected different width for %v and %v", box.width, clonedBox.width))
	assert.EqualWithMessage(t, box.content, clonedBox.content, fmt.Sprintf("expected same content for %v and %v", box.content, clonedBox.content))

	// Change the content of the cloned box
	clonedBox.content = NewBottleOfJuice()

	assert.NotEqualWithMessage(t, box.content, clonedBox.content, fmt.Sprintf("expected different content for %v and %v", box.content, clonedBox.content))

}

func TestClone_BoxOfShoe(t *testing.T) {

	box := NewBoxOfShoe()

	var misterybox trait.Cloneable[Box[Shoe]] = &box

	assert.NotNil(t, misterybox)

	// Clone the box
	clonedBox := misterybox.Clone()

	// basic equality
	assert.EqualWithMessage(t, box, clonedBox, fmt.Sprintf("expected same volume for %v and %v", box, clonedBox))

	// memory address
	assert.NotEqualWithMessage(t, &box, &clonedBox, fmt.Sprintf("expected different memory address for %v and %v", &box, &clonedBox))

	// basic equality inner
	assert.EqualWithMessage(t, box.width, clonedBox.width, fmt.Sprintf("expected same width for %v and %v", box.width, clonedBox.width))

	// Change the width of the cloned box

	clonedBox.width = 1000
	assert.NotEqualWithMessage(t, box.width, clonedBox.width, fmt.Sprintf("expected different width for %v and %v", box.width, clonedBox.width))
	assert.EqualWithMessage(t, box.content, clonedBox.content, fmt.Sprintf("expected same content for %v and %v", box.content, clonedBox.content))

	// Change the content of the cloned box
	clonedBox.content = NewOxfords()

	assert.NotEqualWithMessage(t, box.content, clonedBox.content, fmt.Sprintf("expected different content for %v and %v", box.content, clonedBox.content))

}
