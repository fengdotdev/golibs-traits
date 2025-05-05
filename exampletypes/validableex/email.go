package validableex

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Validable[string] = &Email{}

type Email struct {
	email   string
	isValid bool
	reason  error
}

//constructors

// by default, the email is not validated
func NewEmail(email string) *Email {
	return &Email{
		email:   email,
		isValid: false,
		reason:  ErrorEmailNotValidatedYet,
	}
}

// retuns a new email with the validation
func NewValidatedEmail(email string) *Email {
	e := NewEmail(email)
	e.Validate()
	return e
}

// IsInvalid implements trait.Validable.
func (e *Email) IsInvalid() bool {
	return !e.isValid
}

// IsValid implements trait.Validable.
func (e *Email) IsValid() bool {
	return e.isValid
}

func (e *Email) Reason() error {
	// protect against zero values
	if e.reason == nil && !e.isValid {
		return ErrorEmailNotValidatedYet
	}
	return e.reason
}

func (e *Email) Validate() {
	// check if the email is already validated
	if e.isValid && e.reason == nil {
		return
	}

	result, err := EmailValidator(e.email)
	e.isValid = result
	e.reason = err
}

func (e *Email) String() string {
	return e.email
}

func (e *Email) Value() string {
	return e.email
}
