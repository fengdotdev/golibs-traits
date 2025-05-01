package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/validableex"
	"github.com/fengdotdev/golibs-traits/trait"
)



func TestValidable(t *testing.T) {

	t.Run("Validable with PhoneNumber", func(t *testing.T) {

		type ValidablePhoneNumber = trait.Validable

		var _ ValidablePhoneNumber = &validableex.PhoneNumber{}

		rawNumber := "+56 9 1234 5678"
		p := validableex.NewPhoneNumber(rawNumber)

		t.Run("Interface", func(t *testing.T) {
			assert.TypeWithMessage(t, p, ValidablePhoneNumber(nil), "PhoneNumber should implement Validable[any]")
		})

		t.Run("Validable getters", func(t *testing.T) {
			assert.EqualWithMessage(t, false, p.IsValid(), "PhoneNumber should not be valid")
			assert.EqualWithMessage(t, true, p.IsInvalid(), "PhoneNumber should be invalid")
			assert.EqualWithMessage(t, rawNumber, p.Reason().Error(), "PhoneNumber should have a reason")
		})

		t.Run("Validable setters", func(t *testing.T) {
			p.Validate()
			assert.EqualWithMessage(t, true, p.IsValid(), "PhoneNumber should be valid")
			assert.EqualWithMessage(t, false, p.IsInvalid(), "PhoneNumber should not be invalid")
			assert.EqualWithMessage(t, nil, p.Reason(), "PhoneNumber should have no reason")
		})

	})

	t.Run("Validable with Email", func(t *testing.T) {
		
	})

}
