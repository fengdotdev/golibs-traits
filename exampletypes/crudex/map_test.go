package crudex_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/crudex"
)

func TestMap_StringString(t *testing.T) {

	t.Run("new", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		assert.NotNil(t, m)
	})

	t.Run("Exists", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		result, err := m.Exists("1")
		assert.False(t, result)
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())

		err = m.Create("1", "item1")
		assert.NoError(t, err)
		result2, err2 := m.Exists("1")
		assert.True(t, result2)
		assert.NoError(t, err2)
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

	t.Run("Read", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		item, err := m.Read("1")
		assert.NoError(t, err)
		assert.Equal(t, "item1", item)
	})

	t.Run("Read NotFound", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		item, err := m.Read("1")
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
		assert.Equal(t, "", item)
	})

	t.Run("Read InvalidID", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		assert.NotNil(t, m)
		//err := m.Read(nil)
		//assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrInvalidID.Error(), err.Error())
	})

	t.Run("Update", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Update("1", "item2")
		assert.NoError(t, err)

		item, err := m.Read("1")
		assert.NoError(t, err)
		assert.Equal(t, "item2", item)
	})

	t.Run("Update NotFound", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Update("1", "item2")
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
	})

	t.Run("Delete", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Delete("1")
		assert.NoError(t, err)

		result, err := m.Exists("1")
		assert.False(t, result)
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
	})
	t.Run("Delete NotFound", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Delete("1")
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
	})

	t.Run("Len", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		assert.Equal(t, 2, m.Len())
	})

	t.Run("Keys", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		keys := m.Keys()
		assert.Equal(t, 2, len(keys))
	})

	t.Run("Values", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		values := m.Values()
		assert.Equal(t, 2, len(values))
	})

	t.Run("All", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		all := m.All()
		assert.Equal(t, 2, len(all))
	})

	t.Run("All empty", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		all := m.All()
		assert.Equal(t, 0, len(all))
	})

	t.Run("All with nil", func(t *testing.T) {
		m := crudex.NewMap[string, any]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", nil)
		assert.NoError(t, err)

		all := m.All()
		assert.Equal(t, 2, len(all))
	})

	t.Run("iterate read", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		err := m.Create("1", "item1")
		assert.NoError(t, err)
		err = m.Create("2", "item2")
		assert.NoError(t, err)
		err = m.Create("3", "item3")
		assert.NoError(t, err)
		err = m.Create("4", "item4")
		assert.NoError(t, err)
		err = m.Create("5", "item5")
		assert.NoError(t, err)
		err = m.Create("6", "item6")
		assert.NoError(t, err)

		listKeys := make([]string, 0)

		listValues := make([]string, 0)

		m.Iterate(func(k string, v string) error {
			listKeys = append(listKeys, k)
			listValues = append(listValues, v)
			return nil
		})

		assert.Equal(t, 6, len(listKeys))
		assert.Equal(t, 6, len(listValues))
	})

	t.Run("iterate read with error", func(t *testing.T) {
		m := crudex.NewMap[string, string]()
		err := m.Create("1", "item1")
		assert.NoError(t, err)
		err = m.Create("2", "item2")
		assert.NoError(t, err)
		err = m.Create("3", "item3")
		assert.NoError(t, err)
		err = m.Create("4", "item4")
		assert.NoError(t, err)
		err = m.Create("5", "item5")
		assert.NoError(t, err)
		err = m.Create("6", "item6")
		assert.NoError(t, err)

		listKeys := make([]string, 0)
		listValues := make([]string, 0)
		err = m.Iterate(func(k string, v string)error  {
			listKeys = append(listKeys, k)
			listValues = append(listValues, v)
			if k == "3" {
				return crudex.ErrNotFound
			}
			return nil
		})

		assert.Equal(t, 3, len(listKeys))
		assert.Equal(t, 3, len(listValues))
		assert.ErrorWithMessage(t, err, "Expected error to be %s, but was %s", crudex.ErrNotFound.Error(), err.Error())
	})

	t.Run("clean", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		m.Clean()
		assert.Equal(t, 0, m.Len())
	})

	t.Run("Populate", func(t *testing.T) {
		m := crudex.NewMap[string, string]()

		err := m.Create("1", "item1")
		assert.NoError(t, err)

		err = m.Create("2", "item2")
		assert.NoError(t, err)

		m.Populate(map[string]string{"3": "item3", "4": "item4"})
		assert.Equal(t, 4, m.Len())
	})
}
