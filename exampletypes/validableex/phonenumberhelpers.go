package validableex

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorPhoneNotValidatedYet       = errors.New("phone number not validated yet")
	ErrorIvalidPhoneNumber   = errors.New("invalid phone number")
	ErrorEmptyPhoneNumber    = errors.New("empty phone number")
	ErrorTooShortPhoneNumber = errors.New("phone number is too short") // 10
	ErrorTooLongPhoneNumber  = errors.New("phone number is too long")  // 15
	ErrorNoCountryCode       = errors.New("phone number has no country code")
	ErrorInvalidCountryCode  = errors.New("invalid country code")
)

type CountryCodes = map[string]string // code => country ex +56 Chile

var (
	// From Canada, USA, Mexico, Peru and Chile Only
	ValidCountryCodes = CountryCodes{
		"+1":  "Canada/USA",
		"+52": "Mexico",
		"+51": "Peru",
		"+56": "Chile",
	}
)

// PhoneNumberValidator validates a string phone number
func PhoneNumberValidator(fullPhoneNumber string) (bool, error) {

	normalizedPhoneNumber := NormalicePhoneNumber(fullPhoneNumber)

	err := ValidateLength(normalizedPhoneNumber)
	if err != nil {
		return false, err
	}

	err = ValidateCountryCode(normalizedPhoneNumber)
	if err != nil {
		return false, err
	}

	// other validations

	return true, nil
}

// ValidateCountryCode validates the country code of a phone number.
// a valid country code is a string that starts with a '+' and is followed by 1 to 3 digits.
// a valid country code will return nil error
func ValidateCountryCode(fullPhoneNumber string) error {
	if !strings.HasPrefix(fullPhoneNumber, "+") {
		return ErrorNoCountryCode
	}

	// Regular expression to match the country code part (e.g., +1, +52, +56)
	re := regexp.MustCompile(`^\+(\d{1,3})`)
	match := re.FindStringSubmatch(fullPhoneNumber)

	if len(match) < 2 {
		return ErrorInvalidCountryCode
	}

	countryCode := "+" + match[1]
	if _, ok := ValidCountryCodes[countryCode]; !ok {
		return ErrorInvalidCountryCode
	}

	return nil
}

// ValidateLength validates the length of a phone number.
// a valid phone number is a string that is between 10 and 15 characters long.
// a valid phone number will return nil error
func ValidateLength(fullPhoneNumber string) error {
	// check if the phone number is empty
	if fullPhoneNumber == "" {
		return ErrorEmptyPhoneNumber
	}

	// check if the phone number is too short
	if len(fullPhoneNumber) < 10 {
		return ErrorTooShortPhoneNumber
	}

	// check if the phone number is too long
	if len(fullPhoneNumber) > 15 {
		return ErrorTooLongPhoneNumber
	}

	return nil
}

// NormalicePhoneNumber removes any non-digit characters from the phone number
// while preserving the leading '+' for the country code.
func NormalicePhoneNumber(fullPhoneNumber string) string {
	var normalized strings.Builder
	for i, r := range fullPhoneNumber {
		if r >= '0' && r <= '9' {
			normalized.WriteRune(r)
		} else if i == 0 && r == '+' {
			normalized.WriteRune(r)
		}
	}
	return normalized.String()
}
