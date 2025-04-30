package makerex

import (
	"errors"

	"github.com/fengdotdev/golibs-traits/trait"
)

// todo btn to  html btn
// title to html h1
func Factory(widget Widget) (trait.TypeStringer, error) {
	return nil, errors.New("make not implemented")
}
