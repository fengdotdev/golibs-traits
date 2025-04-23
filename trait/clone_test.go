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

// TestShoeClonability verifies that Shoe correctly implements the Cloneable interface
func TestShoeClonability(t *testing.T) {
	var _ CloneableShoe = &Shoe{}
	var cloneableShoe trait.Cloneable[Shoe] = &Shoe{}
	assert.NotNil(t, cloneableShoe)
	assert.NotNil(t, &Shoe{})
}

// TestShoeClone tests the cloning functionality for the Shoe type
func TestShoeClone(t *testing.T) {
	// Arrange
	original := NewSneakers()

	// Act
	clone := original.Clone()

	// Assert - Verify content equality
	assert.EqualWithMessage(t, original, clone, "expected cloned object to have the same content")

	// Verify that they are different objects in memory
	assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

	// Verify that the fields were copied correctly
	assert.EqualWithMessage(t, original.size, clone.size, fmt.Sprintf("expected same size for %v and %v", original.size, clone.size))
	assert.EqualWithMessage(t, original.model, clone.model, fmt.Sprintf("expected same model for %v and %v", original.model, clone.model))

	// Act - Modify the clone
	clone.size = 44
	clone.model = "ClonedSneakers"

	// Assert - Verify that the original hasn't changed
	assert.NotEqualWithMessage(t, original.size, clone.size, fmt.Sprintf("expected different size after modification for %v and %v", original.size, clone.size))
	assert.NotEqualWithMessage(t, original.model, clone.model, fmt.Sprintf("expected different model after modification for %v and %v", original.model, clone.model))
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

// TestBottleClonability verifies that Bottle correctly implements the Cloneable interface
func TestBottleClonability(t *testing.T) {
	bottle := NewBottleOfWater()
	var _ CloneableBottleOfWater = &bottle
	var cloneableBottle trait.Cloneable[Bottle[string]] = &bottle
	assert.NotNil(t, cloneableBottle)
}

// TestBottleClone tests the cloning functionality for the Bottle type
func TestBottleClone(t *testing.T) {
	// Arrange
	original := NewBottleOfWater()
	cloneableBottle := trait.Cloneable[Bottle[string]](&original)

	// Act
	clone := cloneableBottle.Clone()

	// Assert - Verify content equality
	assert.EqualWithMessage(t, original, clone, fmt.Sprintf("expected cloned object to have the same content for %v and %v", original, clone))

	// Verify that they are different objects in memory
	assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

	// Verify that the fields were copied correctly
	assert.EqualWithMessage(t, original.volume, clone.volume, fmt.Sprintf("expected same volume for %v and %v", original.volume, clone.volume))
	assert.EqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected same content for %v and %v", original.content, clone.content))

	// Act - Modify the clone
	clone.volume = 1000
	clone.content = "juice"

	// Assert - Verify that the original hasn't changed
	assert.NotEqualWithMessage(t, original.volume, clone.volume, fmt.Sprintf("expected different volume after modification for %v and %v", original.volume, clone.volume))
	assert.NotEqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected different content after modification for %v and %v", original.content, clone.content))
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

// TestBoxOfWaterBottleClone tests the cloning functionality for the Box type with Bottle content
func TestBoxOfWaterBottleClone(t *testing.T) {
	// Arrange
	original := NewBoxOfWaterBottle()
	cloneableBox := trait.Cloneable[Box[Bottle[string]]](&original)
	assert.NotNil(t, cloneableBox)

	// Act
	clone := cloneableBox.Clone()

	// Assert - Verify content equality
	assert.EqualWithMessage(t, original, clone, fmt.Sprintf("expected cloned object to have the same content for %v and %v", original, clone))

	// Verify that they are different objects in memory
	assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

	// Verify that the fields were copied correctly
	assert.EqualWithMessage(t, original.width, clone.width, fmt.Sprintf("expected same width for %v and %v", original.width, clone.width))
	assert.EqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected same content for %v and %v", original.content, clone.content))

	// Act - Modify the clone
	clone.width = 1000
	clone.content = NewBottleOfJuice()

	// Assert - Verify that the original hasn't changed
	assert.NotEqualWithMessage(t, original.width, clone.width, fmt.Sprintf("expected different width after modification for %v and %v", original.width, clone.width))
	assert.NotEqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected different content after modification for %v and %v", original.content, clone.content))
}

// TestBoxOfShoeClone tests the cloning functionality for the Box type with Shoe content
func TestBoxOfShoeClone(t *testing.T) {
	// Arrange
	original := NewBoxOfShoe()
	cloneableBox := trait.Cloneable[Box[Shoe]](&original)
	assert.NotNil(t, cloneableBox)

	// Act
	clone := cloneableBox.Clone()

	// Assert - Verify content equality
	assert.EqualWithMessage(t, original, clone, fmt.Sprintf("expected cloned object to have the same content for %v and %v", original, clone))

	// Verify that they are different objects in memory
	assert.NotEqualWithMessage(t, &original, &clone, fmt.Sprintf("expected different memory addresses for %v and %v", &original, &clone))

	// Verify that the fields were copied correctly
	assert.EqualWithMessage(t, original.width, clone.width, fmt.Sprintf("expected same width for %v and %v", original.width, clone.width))
	assert.EqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected same content for %v and %v", original.content, clone.content))

	// Act - Modify the clone
	clone.width = 1000
	clone.content = NewOxfords()

	// Assert - Verify that the original hasn't changed
	assert.NotEqualWithMessage(t, original.width, clone.width, fmt.Sprintf("expected different width after modification for %v and %v", original.width, clone.width))
	assert.NotEqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected different content after modification for %v and %v", original.content, clone.content))
}

// TestDeepClone verifies that cloning works properly in nested structures
func TestDeepClone(t *testing.T) {
	// Arrange
	original := NewBoxOfWaterBottle()
	originalBottle := original.content
	cloneableBox := trait.Cloneable[Box[Bottle[string]]](&original)

	// Act
	clone := cloneableBox.Clone()

	// Assert - Verify that the internal content is also a separate object
	assert.EqualWithMessage(t, original.content, clone.content, fmt.Sprintf("expected clone content to have the same values for %v and %v", original.content, clone.content))

	// Modify the content of the clone
	clone.content.volume = 1500
	clone.content.content = "sparkling water"

	// Verify that the original content hasn't changed
	assert.EqualWithMessage(t, originalBottle.volume, original.content.volume, fmt.Sprintf("expected original bottle volume to remain unchanged for %v and %v", originalBottle.volume, original.content.volume))
	assert.EqualWithMessage(t, originalBottle.content, original.content.content, fmt.Sprintf("expected original bottle content to remain unchanged for %v and %v", originalBottle.content, original.content.content))
	assert.NotEqualWithMessage(t, original.content.volume, clone.content.volume, fmt.Sprintf("expected different volume in cloned bottle for %v and %v", original.content.volume, clone.content.volume))
	assert.NotEqualWithMessage(t, original.content.content, clone.content.content, fmt.Sprintf("expected different content in cloned bottle for %v and %v", original.content.content, clone.content.content))
}
