package makerex

import (
	"github.com/fengdotdev/golibs-traits/exampletypes/handyfuncs"
	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.Maker[Tshirt] = (*Tshirt)(nil)

type Tshirt struct {
	Self  string
	Id    string
	Size  int
	Color string
}

func NewTshirtBlueprint() Tshirt {
	return Tshirt{
		Self:  "Tshirt",
		Id:    "",
		Size:  0,
		Color: "",
	}
}
func (s *Tshirt) Make(args ...any) Tshirt {
	return Tshirt{
		Self:  s.Self,
		Id:    handyfuncs.IdGen(),
		Size:  42,
		Color: "red",
	}
}
