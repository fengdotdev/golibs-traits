package comparableex

import (
	"reflect"

	"github.com/fengdotdev/golibs-traits/trait"
)

type ComparablePet = trait.Comparable

var _ trait.Comparable = &Pet{}

type Pet struct {
	name      string
	typeOfPet string
}

// constructors

func NewPet(name string, typeOfPet string) *Pet {
	return &Pet{
		name:      name,
		typeOfPet: typeOfPet,
	}
}

func NewDog(name string) *Pet {
	return &Pet{
		name:      name,
		typeOfPet: "Dog",
	}
}

func NewCat(name string) *Pet {
	return &Pet{
		name:      name,
		typeOfPet: "Cat",
	}
}

// test constructors
func NewScobbyDoo() *Pet {
	return &Pet{
		name:      "Scooby Doo",
		typeOfPet: "Dog",
	}
}

func NewGarfield() *Pet {
	return &Pet{
		name:      "Garfield",
		typeOfPet: "Cat",
	}
}

// methods comparable interface

// AreEqual implements trait.Comparable.
func (p *Pet) AreEqual(other any) bool {
	otherValue := reflect.ValueOf(other)
	petValue := reflect.ValueOf(p)

	if otherValue.Type() != petValue.Type() {
		return false
	}

	if otherValue.Kind() == reflect.Ptr {
		otherValue = otherValue.Elem()
	}
	if petValue.Kind() == reflect.Ptr {
		petValue = petValue.Elem()
	}

	if otherValue.Kind() != reflect.Struct {
		return reflect.DeepEqual(p, other) // Para tipos no struct, usar DeepEqual
	}

	for i := 0; i < petValue.NumField(); i++ {
		field := petValue.Type().Field(i).Name
		pValue := petValue.FieldByName(field)
		oValue := otherValue.FieldByName(field)

		if !reflect.DeepEqual(pValue.Interface(), oValue.Interface()) {
			return false
		}
	}

	return true
}

// AreNotEqual implements trait.Comparable.
func (p *Pet) AreNotEqual(other any) bool {
	if _, ok := other.(Pet); !ok {
		return true
	}
	var otherPet = other.(Pet)
	if p.name == otherPet.name && p.typeOfPet == otherPet.typeOfPet {
		return false
	}
	return true
}

// AreNotSameType implements trait.Comparable.
func (p *Pet) AreNotSameType(other any) bool {
	if _, ok := other.(Pet); !ok {
		return true
	}
	return false
}

// AreSameType implements trait.Comparable.
func (p *Pet) AreSameType(other any) bool {
	if _, ok := other.(Pet); ok {
		return true
	}
	return false
}

// methods
func (p *Pet) Name() string {
	return p.name
}
func (p *Pet) TypeOfPet() string {
	return p.typeOfPet
}
