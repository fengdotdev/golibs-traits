package crudex_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/crudex"
)

func TestMap(t *testing.T) {

	t.Run("new", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		assert.NotNil(t, m)
	})

	t.Run("Exists", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		result, err := m.Exists("1")
		assert.False(t, result)
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
	})

	t.Run("Create", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		result, _ := m.Exists("1")
		assert.False(t, result)

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		result, _ = m.Exists("1")
		assert.True(t, result)
	})

	t.Run("Create AlreadyExists", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("1", "item2")
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrAlreadyExists.Error(), err.Error())
	})

	t.Run("Create InvalidID", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		assert.NotNil(t, m)
		//err := m.Create(nil, "item1")
		//assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrInvalidID.Error(), err.Error())
	})

	t.Run("Create id diferent type", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		assert.NotNil(t, m)
		err := m.Create("1", "item1")

		assert.NoError(t, err)
		//err = m.Create(1, "item2")
		//assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrInvalidID.Error(), err.Error())
	})
}
