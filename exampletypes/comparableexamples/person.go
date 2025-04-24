package comparableexamples

import "github.com/fengdotdev/golibs-traits/trait"

type ComparablePerson = trait.Comparable

var _ trait.Comparable = &Person{}

type Person struct {
	name string
	age  int
}

// AreEqual implements trait.Comparable.
func (p *Person) AreEqual(other any) bool {

	if _, ok := other.(Person); !ok {
		return false
	}
	var otherPerson = other.(Person)
	if p.age == otherPerson.age && p.name == otherPerson.name {
		return true
	}
	return false

}

// AreNotEqual implements trait.Comparable.
func (p *Person) AreNotEqual(other any) bool {
	if _, ok := other.(Person); !ok {
		return true
	}
	var otherPerson = other.(Person)
	if p.age == otherPerson.age && p.name == otherPerson.name {
		return false
	}
	return true
}

// AreNotSameType implements trait.Comparable.
func (p *Person) AreNotSameType(other any) bool {
	if _, ok := other.(Person); !ok {
		return true
	}
	return false

}

// AreSameType implements trait.Comparable.
func (p *Person) AreSameType(other any) bool {
	if _, ok := other.(Person); ok {
		return true
	}
	return false
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) Name() string {
	return p.name
}
func (p *Person) Age() int {
	return p.age
}
