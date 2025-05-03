package validableex

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

var (
	ErrorEmailNotValidatedYet = errors.New("not validated yet")
	ErrorInvalidEmail         = errors.New("invalid email")
	ErrorEmptyEmail           = errors.New("empty email")
	ErrorLengthEmail          = errors.New("email length is invalid") // 5 - 254
)

func EmailValidator(email string) (bool, error) {

	if IsEmailEmpty(email) {
		return false, ErrorEmptyEmail
	}

	if IsEmailInRange(email, 5, 254) {
		return false, ErrorLengthEmail
	}

	if EmailHaveForbiddenChars(email) {
		return false, ErrorInvalidEmail
	}

	if EmailContainArrobas(email) {
		return false, ErrorInvalidEmail
	}

	if EmailContainDomain(email) {
		return false, ErrorInvalidEmail
	}
	// other validations

	// Use Go's built-in email parsing for more robust validation
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, ErrorInvalidEmail
	}

	return true, nil
}

func IsEmailInRange(email string, lower int, upper int) bool {
	emailLength := len(email)
	return emailLength >= lower && emailLength <= upper
}

func IsEmailEmpty(email string) bool {
	return strings.TrimSpace(email) == ""
}

func EmailContainArrobas(email string) bool {
	return strings.Count(email, "@") == 1
}

func EmailContainDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[1] == "" {
		return false
	}
	domain := parts[1]
	// This is a basic check; more sophisticated domain validation might be needed
	return strings.Contains(domain, ".")
}

func EmailHaveForbiddenChars(email string) bool {
	// Regular expression to match characters that are generally not allowed
	// in email addresses (excluding those handled by mail.ParseAddress)
	forbidden := regexp.MustCompile(`[\s<>()\[\]\\,;:"']`)
	return forbidden.MatchString(email)
}
