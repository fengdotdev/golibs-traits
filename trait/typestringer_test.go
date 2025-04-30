package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/typestringerex"
	"github.com/fengdotdev/golibs-traits/trait"
)

type TypeStringer = trait.TypeStringer

type NotToml struct {
	content string
}

func TestTypeStringer_Toml(t *testing.T) {

	t.Run("interface TypeStringer TOML", func(t *testing.T) {
		nottommy := &NotToml{"hello toml"}
		tommy := typestringerex.NewToML("hello toml")

		assert.TypeWithMessage(t, tommy, TypeStringer(nil), "tommy should be of type TypeStringer")

		assert.Fail(t, func(t *testing.T) {
			assert.TypeWithMessage(t, nottommy, TypeStringer(nil), "nottommy should not be of type TypeStringer")
		})

	})

	t.Run("TypeStringer TOML getters", func(t *testing.T) {
		tommy := typestringerex.NewToML("hello toml")

		assert.EqualWithMessage(t, "toml", tommy.TypeName(), "TypeName should be toml and not %s", tommy.TypeName())
		assert.EqualWithMessage(t, "hello toml", tommy.ValueString(), "ValueString should be hello toml and not %s", tommy.ValueString())
	})

}

func TestTypeStringer_blob(t *testing.T) {

	t.Run("interface TypeStringer binaryblob", func(t *testing.T) {
		notblob := &NotToml{"hello toml"}
		blob := typestringerex.NewBinaryBlob([]byte{1, 2, 3, 4})

		assert.TypeWithMessage(t, blob, TypeStringer(nil), "blob should be of type TypeStringer")

		assert.Fail(t, func(t *testing.T) {
			assert.TypeWithMessage(t, notblob, TypeStringer(nil), "notblob should not be of type TypeStringer")
		})
	})

	t.Run("TypeStringer binaryblob getters", func(t *testing.T) {

		data := []byte("Hello")

		blob := typestringerex.NewBinaryBlob(data) // "Hello"

		// 64base  SGVsbG8=

		assert.EqualWithMessage(t, "binaryblob", blob.TypeName(), "TypeName should be blob and not %s", blob.TypeName())
		assert.EqualWithMessage(t, "SGVsbG8=", blob.ValueString(), "ValueString should be SGVsbG8= and not %s", blob.ValueString())
	})

	t.Run("TypeStringer binaryblob from", func(t *testing.T) {

		data := []byte("Hello")

		blob := typestringerex.NewBinaryBlob(data) // "Hello"

		data64 := blob.ValueString()

		blob2, err := typestringerex.NewBinaryBlobFromString(data64)
		assert.NotErrorWithMessage(t, err, "should not return error")

		blob2data := blob2.Content()
		assert.EqualWithMessage(t, data, blob2data, "blob2data should be %s and not %s", data, blob2data)

	})

}

func TestTypeStringer_mockDB(t *testing.T) {

	type TV struct {
		Type  string
		Value string
	}

	db := make(map[string]TV)

	tommy := typestringerex.NewToML("hello toml")

	bloby := typestringerex.NewBinaryBlob([]byte("Hello"))

	db["tommy"] = TV{
		Type:  tommy.TypeName(),
		Value: tommy.ValueString(),
	}

	db["bloby"] = TV{
		Type:  bloby.TypeName(),
		Value: bloby.ValueString(),
	}

	// lookup tommy

	t.Run("lookup tommy TypeStringer", func(t *testing.T) {

		tommyTV, ok := db["tommy"]
		assert.TrueWithMessage(t, ok, "should be true")

		newTommy, err := typestringerex.Builder[trait.TypeStringer](tommyTV.Type, tommyTV.Value)
		assert.NotErrorWithMessage(t, err, "should not return error")

		assert.EqualWithMessage(t, tommy.TypeName(), newTommy.TypeName(), "TypeName should be %s and not %s", tommy.TypeName(), newTommy.TypeName())
		assert.EqualWithMessage(t, tommy.ValueString(), newTommy.ValueString(), "ValueString should be %s and not %s", tommy.ValueString(), newTommy.ValueString())
	})

	t.Run("lookup tommy * TOML", func(t *testing.T) {

		tommyTV, ok := db["tommy"]
		assert.TrueWithMessage(t, ok, "should be true")

		newTommy, err := typestringerex.Builder[*typestringerex.TOML](tommyTV.Type, tommyTV.Value)
		assert.NotErrorWithMessage(t, err, "should not return error")

		assert.EqualWithMessage(t, tommy.TypeName(), newTommy.TypeName(), "TypeName should be %s and not %s", tommy.TypeName(), newTommy.TypeName())
		assert.EqualWithMessage(t, tommy.ValueString(), newTommy.ValueString(), "ValueString should be %s and not %s", tommy.ValueString(), newTommy.ValueString())
	})

	t.Run("lookup bloby TypeStringer", func(t *testing.T) {
		// lookup bloby
		blobyTV, ok := db["bloby"]
		assert.TrueWithMessage(t, ok, "should be true")

		newBloby, err := typestringerex.Builder[trait.TypeStringer](blobyTV.Type, blobyTV.Value)
		assert.NotErrorWithMessage(t, err, "should not return error")
		assert.EqualWithMessage(t, bloby.TypeName(), newBloby.TypeName(), "TypeName should be %s and not %s", bloby.TypeName(), newBloby.TypeName())
		assert.EqualWithMessage(t, bloby.ValueString(), newBloby.ValueString(), "ValueString should be %s and not %s", bloby.ValueString(), newBloby.ValueString())
	})

	t.Run("lookup bloby * BinaryBlob", func(t *testing.T) {
		// lookup bloby
		blobyTV, ok := db["bloby"]
		assert.TrueWithMessage(t, ok, "should be true")

		newBloby, err := typestringerex.Builder[*typestringerex.BinaryBlob](blobyTV.Type, blobyTV.Value)
		assert.NotErrorWithMessage(t, err, "should not return error")
		assert.EqualWithMessage(t, bloby.TypeName(), newBloby.TypeName(), "TypeName should be %s and not %s", bloby.TypeName(), newBloby.TypeName())
		assert.EqualWithMessage(t, bloby.ValueString(), newBloby.ValueString(), "ValueString should be %s and not %s", bloby.ValueString(), newBloby.ValueString())
	})

}
