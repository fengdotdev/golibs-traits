package validableex_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/validableex"
)

var validPhones = []string{
	"+56 (45) 259 2200", // Chile
	"56 2 2391 9997",    // Chile
	"+1 (206) 266-1000", // USA
	"+1 (650) 253-0000", // USA
	"+1-416-979-3033",   // Canada
}

var invalidPhones = []string{
	"1234567890",   // Invalid format
	"123.456.7890", // Invalid format
}

func TestPhoneNumber(t *testing.T) {
	t.Run("Valid Phone Numbers", func(t *testing.T) {

		for _, phone := range validPhones {
			t.Run(phone, func(t *testing.T) {
				p := validableex.NewPhoneNumber(phone)
				assert.FalseWithMessage(t, p.IsValid(), "Phone number %s should not be valid", phone)
				assert.TrueWithMessage(t, p.IsInvalid(), "Phone number %s should be invalid", phone)
				assert.ErrorWithMessage(t, p.Reason(), "Phone number %s should have a reason", phone)
				assert.EqualWithMessage(t, p.String(), phone, "Phone number %s should be equal to the raw number", phone)

				p.Validate()
				assert.TrueWithMessage(t, p.IsValid(), "Phone number %s should be valid", phone)
				assert.FalseWithMessage(t, p.IsInvalid(), "Phone number %s should not be invalid", phone)
				t.Log(p.Reason())
				assert.NotErrorWithMessage(t, p.Reason(), "Phone number %s should not have a reason", phone)
			})
		}

	})

	t.Run("Invalid Phone Numbers", func(t *testing.T) {
		for _, phone := range invalidPhones {
			t.Run(phone, func(t *testing.T) {
				p := validableex.NewPhoneNumber(phone)
				assert.FalseWithMessage(t, p.IsValid(), "Phone number %s should not be valid", phone)
				assert.TrueWithMessage(t, p.IsInvalid(), "Phone number %s should be invalid", phone)
				assert.ErrorWithMessage(t, p.Reason(), "Phone number %s should have a reason", phone)
				assert.EqualWithMessage(t, p.String(), phone, "Phone number %s should be equal to the raw number", phone)

				p.Validate()
				assert.FalseWithMessage(t, p.IsValid(), "Phone number %s should not be valid", phone)
				assert.TrueWithMessage(t, p.IsInvalid(), "Phone number %s should be invalid", phone)
				assert.ErrorWithMessage(t, p.Reason(), "Phone number %s should have a reason", phone)
			})
		}

	})
}
