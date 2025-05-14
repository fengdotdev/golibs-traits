package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/crudex"
)

func TestCRUD_map(t *testing.T) {
	m := crudex.NewMap[int, string]()

	limit := 100

	for i := 0; i < limit; i++ {
		err := m.Create(i, "item")
		assert.NotErrorWithMessage(t, err, "Expected no error, but got %s", err.Error())
	}

	assert.EqualWithMessage(t, m.Len(), limit, "Expected length to be %d, but got %d", limit, m.Len())

}
