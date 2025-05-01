package validableex

import "errors"

var (
	ErrorEmailNotValidatedYet = errors.New("not validated yet")
	ErrorInvalidEmail         = errors.New("invalid email")
	ErrorEmptyEmail           = errors.New("empty email")
)

func EmailValidator(email string) (bool, error) {

	return false, ErrorInvalidEmail
}
