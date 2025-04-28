package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/makerexamples"
)

func TestMaker(t *testing.T) {

	type Shoe = makerexamples.Shoe
	type Shorts = makerexamples.Shorts
	type Tshirt = makerexamples.Tshirt

	factory := makerexamples.NewFactory(Shoe{}, Shorts{}, Tshirt{})

	tshirtRaw, err := factory.Makeit("makerexamples.Tshirt")
	assert.NoError(t, err)

	tshirt, ok := tshirtRaw.(Tshirt)
	assert.True(t, ok)
	assert.Equal(t, tshirt.Self, "Tshirt")
}
