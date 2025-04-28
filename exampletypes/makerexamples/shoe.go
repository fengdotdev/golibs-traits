package makerexamples

import (
	"github.com/fengdotdev/golibs-traits/exampletypes/handyfuncs"
	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.Maker[Shoe] = (*Shoe)(nil)

type Shoe struct {
	Self  string
	Id    string
	Size  int
	Color string
}

func NewShoeBlueprint() Shoe {
	return Shoe{
		Self:  "Shoe",
		Id:    "",
		Size:  0,
		Color: "",
	}
}

func (s Shoe) Make(args ...any) Shoe {
	return Shoe{
		Self:  s.Self,
		Id:    handyfuncs.IdGen(),
		Size:  41,
		Color: "black",
	}
}
