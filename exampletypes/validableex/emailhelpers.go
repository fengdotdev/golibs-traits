package validableex

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

var (
	ErrorEmailNotValidatedYet   = errors.New("not validated yet")
	ErrorEmailInvalid           = errors.New("invalid email")
	ErrorEmailInvalidCharacters = errors.New("invalid characters in email")
	ErrorEmailInvalidFormat     = errors.New("invalid email format")
	ErrorEmailEmpty             = errors.New("empty email")
	ErrorEmailNotDomain         = errors.New("email domain is invalid")
	ErrorLengthEmail            = errors.New("email length is invalid") // 5 - 254
	ErrorEmailContainEspaces    = errors.New("email contains spaces")
)

func EmailValidator(email string) (bool, error) {

	if IsEmailEmpty(email) {
		return false, ErrorEmailEmpty
	}
	if EmailHaveSpaces(email) {
		return false, ErrorEmailContainEspaces
	}
	// a@c.c  4
	// 4 - 254
	if !IsEmailInRange(email, 4, 254) {
		return false, ErrorLengthEmail
	}

	if EmailHaveForbiddenChars(email) {
		return false, ErrorEmailInvalidCharacters
	}

	if !EmailContainArrobas(email) {
		return false, ErrorEmailInvalidFormat
	}

	if !EmailContainDomain(email) {
		return false, ErrorEmailNotDomain
	}

	if !EmailIsValidByGo(email) {
		return false, ErrorEmailInvalid
	}

	return true, nil
}

func EmailIsValidByGo(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsEmailInRange checks if the email length is within the specified range
// lower < x < upper -> true
func IsEmailInRange(email string, lower int, upper int) bool {
	emailLength := len(email)

	if emailLength < lower {
		return true
	}
	if emailLength > upper {
		return false
	}
	return true
}

// IsEmailEmpty checks if the email is empty or contains only whitespace
func IsEmailEmpty(email string) bool {
	return strings.TrimSpace(email) == ""
}

// EmailContainArrobas checks if the email contains exactly one '@' symbol
func EmailContainArrobas(email string) bool {
	return strings.Count(email, "@") == 1
}

// EmailContainDomain checks if the email contains a domain part after the '@' symbol
// and that the domain is valid (contains at least one dot)
func EmailContainDomain(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[1] == "" {
		return false
	}
	domain := parts[1]
	// This is a basic check; more sophisticated domain validation might be needed
	return strings.Contains(domain, ".")
}
// EmailHaveSpaces checks if the email contains spaces
func EmailHaveSpaces(email string) bool {
	return strings.Contains(email, " ")
}

// this regex is used to check for forbidden characters in the email
const regexForbiddenEmailChars = "[\\s\\t\\n\\r,;:\\\\\\/()\\[\\]{}<>|`^*!#$%&=?~]"

// EmailHaveForbiddenChars checks if the email contains forbidden characters
// allowed characters are: a-zA-Z0-9._%+-@
// not allowed characters are: <>()[]\,";:'`!#$%^&*|/{}~`
func EmailHaveForbiddenChars(email string) bool {
	// Regular expression to match characters that are generally not allowed
	// in email addresses (excluding those handled by mail.ParseAddress)
	forbidden := regexp.MustCompile(regexForbiddenEmailChars)
	return forbidden.MatchString(email)
}
