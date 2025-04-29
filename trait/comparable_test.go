package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/comparableex"
	"github.com/fengdotdev/golibs-traits/trait"
)

type Person = comparableex.Person
type Pet = comparableex.Pet
type ComparablePerson = trait.Comparable
type ComparablePet = trait.Comparable

func TestComparable(t *testing.T) {

	var jhon22 ComparablePerson = comparableex.NewPerson("John", 22)

	var emily18 ComparablePerson = comparableex.NewPerson("Emily", 18)

	jhonUnknown := comparableex.NewPerson("John", 22)

	philip33 := comparableex.NewPerson("Philip", 33)

	kat60 := comparableex.NewPerson("Kat", 60)

	t.Run("Person", func(t *testing.T) {

		t.Run("AreEqual", func(t *testing.T) {
			assert.TrueWithMessage(t, jhon22.AreEqual(jhon22), "Expected jhon22 and jhon22 to be equal")

			assert.TrueWithMessage(t, jhon22.AreEqual(jhonUnknown), "Expected jhon22 and jhonUnknown to be equal")
			assert.TrueWithMessage(t, jhonUnknown.AreEqual(jhon22), "Expected jhon22 and jhon22 to be equal")

			assert.FalseWithMessage(t, jhon22.AreEqual(emily18), "Expected jhon22 and emily18 to be not equal")
			assert.FalseWithMessage(t, jhon22.AreEqual(philip33), "Expected jhon22 and philip33 to be not equal")
			assert.FalseWithMessage(t, jhon22.AreEqual(kat60), "Expected jhon22 and kat60 to be not equal")

		})

		t.Run("AreNotEqual", func(t *testing.T) {

			assert.FalseWithMessage(t, jhon22.AreNotEqual(jhon22), "Expected jhon22 and jhon22 to be not equal")

			assert.TrueWithMessage(t, jhon22.AreNotEqual(emily18), "Expected jhon22 and emily18 to be not equal")
			assert.TrueWithMessage(t, jhon22.AreNotEqual(philip33), "Expected jhon22 and philip33 to be not equal")
			assert.TrueWithMessage(t, jhon22.AreNotEqual(kat60), "Expected jhon22 and kat60 to be not equal")

		})

		t.Run("AreSameType", func(t *testing.T) {
			assert.TrueWithMessage(t, jhon22.AreSameType(emily18), "Expected jhon22 and emily18 to be same type")
			assert.TrueWithMessage(t, jhon22.AreSameType(philip33), "Expected jhon22 and philip33 to be same type")
			assert.TrueWithMessage(t, jhon22.AreSameType(kat60), "Expected jhon22 and kat60 to be same type")

		})
		t.Run("AreNotSameType", func(t *testing.T) {
			assert.FalseWithMessage(t, jhon22.AreNotSameType(emily18), "Expected jhon22 and emily18 be same type")
			assert.FalseWithMessage(t, jhon22.AreNotSameType(philip33), "Expected jhon22 and philip33 to be same type")
			assert.FalseWithMessage(t, jhon22.AreNotSameType(kat60), "Expected jhon22 and kat60 to be same type")

		})

	})

	// Pet
	var scooby ComparablePet = comparableex.NewScobbyDoo()
	var garfield1 ComparablePet = comparableex.NewGarfield()

	garfield2 := comparableex.NewGarfield()
	snowball := comparableex.NewPet("Snowball", "Cat")
	lassie := comparableex.NewPet("Lassie", "Dog")

	t.Run("Pet", func(t *testing.T) {
		t.Run("AreEqual", func(t *testing.T) {
			assert.TrueWithMessage(t, scooby.AreEqual(scooby), "Expected scooby and scooby to be equal")

			assert.TrueWithMessage(t, garfield1.AreEqual(garfield2), "Expected garfield1 and garfield2 to be equal")
			assert.TrueWithMessage(t, garfield2.AreEqual(garfield1), "Expected garfield2 and garfield1 to be equal")

			assert.FalseWithMessage(t, scooby.AreEqual(garfield1), "Expected scooby and garfield to be not equal")
			assert.FalseWithMessage(t, scooby.AreEqual(snowball), "Expected scooby and snowball to be not equal")
			assert.FalseWithMessage(t, scooby.AreEqual(lassie), "Expected scooby and lassie to be not equal")
		})

		t.Run("AreNotEqual", func(t *testing.T) {
			assert.FalseWithMessage(t, scooby.AreNotEqual(scooby), "Expected scooby and scooby to be not equal")
			assert.FalseWithMessage(t, garfield1.AreNotEqual(garfield2), "Expected garfield1 and garfield2 to be not equal")
			assert.FalseWithMessage(t, garfield2.AreNotEqual(garfield1), "Expected garfield2 and garfield1 to be not equal")

			assert.TrueWithMessage(t, scooby.AreNotEqual(garfield1), "Expected scooby and garfield to be not equal")
			assert.TrueWithMessage(t, scooby.AreNotEqual(snowball), "Expected scooby and snowball to be not equal")
			assert.TrueWithMessage(t, scooby.AreNotEqual(lassie), "Expected scooby and lassie to be not equal")

		})

		t.Run("AreSameType", func(t *testing.T) {
			assert.TrueWithMessage(t, scooby.AreSameType(garfield1), "Expected scooby and garfield to be same type")
			assert.TrueWithMessage(t, scooby.AreSameType(garfield2), "Expected scooby and garfield to be same type")
			assert.TrueWithMessage(t, garfield1.AreSameType(garfield2), "Expected garfield1 and garfield2 to be same type")
			assert.TrueWithMessage(t, garfield2.AreSameType(garfield1), "Expected garfield2 and garfield1 to be same type")
		})

		t.Run("AreNotSameType", func(t *testing.T) {
			assert.FalseWithMessage(t, scooby.AreNotSameType(garfield1), "Expected scooby and garfield to be same type")
			assert.FalseWithMessage(t, scooby.AreNotSameType(garfield2), "Expected scooby and garfield to be same type")
			assert.FalseWithMessage(t, garfield1.AreNotSameType(garfield2), "Expected garfield1 and garfield2 to be same type")
			assert.FalseWithMessage(t, garfield2.AreNotSameType(garfield1), "Expected garfield2 and garfield1 to be same type")

		})
	})

	t.Run("person and pet", func(t *testing.T) {

		t.Run("AreEqual", func(t *testing.T) {
			assert.FalseWithMessage(t, jhon22.AreEqual(scooby), "Expected jhon22 and scooby to be not equal")
			assert.FalseWithMessage(t, jhon22.AreEqual(garfield1), "Expected jhon22 and garfield to be not equal")
			assert.FalseWithMessage(t, jhon22.AreEqual(garfield2), "Expected jhon22 and garfield to be not equal")

		})

		t.Run("AreNotEqual", func(t *testing.T) {
			assert.TrueWithMessage(t, jhon22.AreNotEqual(scooby), "Expected jhon22 and scooby to be not equal")
			assert.TrueWithMessage(t, jhon22.AreNotEqual(garfield1), "Expected jhon22 and garfield to be not equal")
			assert.TrueWithMessage(t, jhon22.AreNotEqual(garfield2), "Expected jhon22 and garfield to be not equal")

		})

		t.Run("AreSameType", func(t *testing.T) {
			assert.FalseWithMessage(t, jhon22.AreSameType(scooby), "Expected jhon22 and scooby to be not same type")
			assert.FalseWithMessage(t, jhon22.AreSameType(garfield1), "Expected jhon22 and garfield to be not same type")
			assert.FalseWithMessage(t, jhon22.AreSameType(garfield2), "Expected jhon22 and garfield to be not same type")

		})
		t.Run("AreNotSameType", func(t *testing.T) {
			assert.TrueWithMessage(t, jhon22.AreNotSameType(scooby), "Expected jhon22 and scooby to be not same type")
			assert.TrueWithMessage(t, jhon22.AreNotSameType(garfield1), "Expected jhon22 and garfield to be not same type")
			assert.TrueWithMessage(t, jhon22.AreNotSameType(garfield2), "Expected jhon22 and garfield to be not same type")

		})

	})

	t.Run("pet or person vs any", func(t *testing.T) {

		unknown1 := "cat"
		unknown2 := 22
		unknown3 := 22.0
		unknown4 := true
		unknown5 := []int{1, 2, 3}
		unknown6 := map[string]string{"key": "value"}
		unknown7 := struct {
			Name string
			Age  int
		}{Name: "John", Age: 22}

		assert.FalseWithMessage(t, jhon22.AreEqual(unknown1), "Expected jhon22 and unknown1 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown2), "Expected jhon22 and unknown2 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown3), "Expected jhon22 and unknown3 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown4), "Expected jhon22 and unknown4 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown5), "Expected jhon22 and unknown5 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown6), "Expected jhon22 and unknown6 to be not equal")
		assert.FalseWithMessage(t, jhon22.AreEqual(unknown7), "Expected jhon22 and unknown7 to be not equal")

		assert.FalseWithMessage(t, jhon22.AreSameType(unknown1), "Expected jhon22 and unknown1 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown2), "Expected jhon22 and unknown2 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown3), "Expected jhon22 and unknown3 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown4), "Expected jhon22 and unknown4 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown5), "Expected jhon22 and unknown5 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown6), "Expected jhon22 and unknown6 to be not same type")
		assert.FalseWithMessage(t, jhon22.AreSameType(unknown7), "Expected jhon22 and unknown7 to be not same type")

		assert.FalseWithMessage(t, scooby.AreEqual(unknown1), "Expected scooby and unknown1 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown2), "Expected scooby and unknown2 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown3), "Expected scooby and unknown3 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown4), "Expected scooby and unknown4 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown5), "Expected scooby and unknown5 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown6), "Expected scooby and unknown6 to be not equal")
		assert.FalseWithMessage(t, scooby.AreEqual(unknown7), "Expected scooby and unknown7 to be not equal")

	})

}
