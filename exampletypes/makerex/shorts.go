package makerex

import "github.com/fengdotdev/golibs-traits/exampletypes/handyfuncs"

type Shorts struct {
	Self  string
	Id    string
	Size  int
	Color string
}

func NewShortsBlueprint() Shorts {
	return Shorts{
		Self:  "Shorts",
		Id:    "",
		Size:  0,
		Color: "",
	}
}

func (s Shorts) Make(args ...any) Shorts {
	return Shorts{
		Self:  s.Self,
		Id:    handyfuncs.IdGen(),
		Size:  32,
		Color: "blue",
	}
}
