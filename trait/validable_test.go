package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/validableex"
	"github.com/fengdotdev/golibs-traits/trait"
)

func TestValidable(t *testing.T) {

	t.Run("Validable with PhoneNumber", func(t *testing.T) {

		type ValidablePhoneNumber = trait.Validable[string]

		var _ ValidablePhoneNumber = &validableex.PhoneNumber{}

		rawNumber := "+56 9 1234 5678"
		p := validableex.NewPhoneNumber(rawNumber) //by default, it is invalid

		t.Run("Interface", func(t *testing.T) {
			assert.TypeWithMessage(t, p, ValidablePhoneNumber(nil), "PhoneNumber should implement Validable[any]")
		})

		t.Run("Validable getters", func(t *testing.T) {
			assert.FalseWithMessage(t, p.IsValid(), "PhoneNumber should not be valid")
			assert.TrueWithMessage(t, p.IsInvalid(), "PhoneNumber should be invalid")
			assert.ErrorWithMessage(t, p.Reason(), "PhoneNumber should have a reason")
		})

		t.Run("Validable setters", func(t *testing.T) {
			p.Validate()

			t.Log(p.IsValid())
			t.Log(p.String())
			t.Log(p.IsInvalid())
			t.Log(p.Reason())
			assert.TrueWithMessage(t, p.IsValid(), "PhoneNumber should be valid")
			assert.FalseWithMessage(t, p.IsInvalid(), "PhoneNumber should not be invalid")
			assert.NoErrorWithMessage(t, p.Reason(), "PhoneNumber should not have a reason")
			assert.EqualWithMessage(t, p.String(), rawNumber, "PhoneNumber should be equal to the raw number")
		})

	})

	t.Run("Validable with Email", func(t *testing.T) {
		type ValidableEmail = trait.Validable[string]

		var _ ValidableEmail = &validableex.Email{}

		rawEmail := "foo@gmail.com"
		p := validableex.NewEmail(rawEmail) //by default, it is invalid

		t.Run("Interface", func(t *testing.T) {
			assert.TypeWithMessage(t, p, ValidableEmail(nil), "Email should implement Validable[any]")
		})

		t.Run("Validable getters", func(t *testing.T) {

			assert.FalseWithMessage(t, p.IsValid(), "Email should not be valid")
			assert.TrueWithMessage(t, p.IsInvalid(), "Email should be invalid")
			assert.ErrorWithMessage(t, p.Reason(), "Email should have a reason")
			assert.EqualWithMessage(t, p.Value(), rawEmail, "Email should be equal to the raw number")
		})

		t.Run("Validable setters", func(t *testing.T) {
			p.Validate()
			assert.TrueWithMessage(t, p.IsValid(), "Email should be valid")
			t.Log(p.IsValid())
			t.Log(p.String())
			t.Log(p.IsInvalid())
			t.Log(p.Reason())
			assert.FalseWithMessage(t, p.IsInvalid(), "Email should not be invalid")
			t.Log(p.Reason())
			assert.NoErrorWithMessage(t, p.Reason(), "Email should not have a reason")

		})

	})

}
