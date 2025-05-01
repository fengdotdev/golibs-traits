package validableex

import (
	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.Validable = &PhoneNumber{}

type PhoneNumber struct {
	number  string
	isValid bool
	reason  error
}

// by default, the phone number is not validated
func NewPhoneNumber(number string) *PhoneNumber {
	return &PhoneNumber{
		number:  number,
		isValid: false,
		reason:  ErrorPhoneNotValidatedYet,
	}
}

// retuns a new phone number with the validation
func NewValidatedPhoneNumber(number string) *PhoneNumber {
	p := NewPhoneNumber(number)
	p.Validate()
	return p
}


func (p *PhoneNumber) Validate() {
	if p.isValid && p.reason == nil {
		return
	}

	result, err := PhoneNumberValidator(p.number)
	p.isValid = result
	p.reason = err
}

// IsInvalid implements trait.Validable.
func (p *PhoneNumber) IsInvalid() bool {
	return !p.isValid
}

// IsValid implements trait.Validable.
func (p *PhoneNumber) IsValid() bool {
	return p.isValid
}

func (p *PhoneNumber) Reason() error {
		// protect against zero values
	if p.reason == nil  && !p.isValid {
		return ErrorPhoneNotValidatedYet
	}
	return p.reason
}