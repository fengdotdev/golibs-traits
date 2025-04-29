package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/makerex"
)

func TestMaker(t *testing.T) {

	type Shoe = makerex.Shoe
	type Shorts = makerex.Shorts
	type Tshirt = makerex.Tshirt

	factory := makerex.NewFactory(Shoe{}, Shorts{}, Tshirt{})

	tshirtRaw, err := factory.Makeit("makerexamples.Tshirt")
	assert.NoError(t, err)

	tshirt, ok := tshirtRaw.(Tshirt)
	assert.True(t, ok)
	assert.Equal(t, tshirt.Self, "Tshirt")
}
