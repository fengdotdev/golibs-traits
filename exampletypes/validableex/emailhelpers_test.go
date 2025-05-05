package validableex

import (
	"strings"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestEmail_Validator(t *testing.T) {	
	

}

func TestEmail_IsEmailEmpty(t *testing.T) {

	t.Run("empty email", func(t *testing.T) {

		email := ""

		empty := IsEmailEmpty(email)
		assert.TrueWithMessage(t, empty, "Email %s should not be empty", email)

	})

	t.Run("not empty email", func(t *testing.T) {

		email := "foo@foo.foo"
		notempty := IsEmailEmpty(email)
		assert.FalseWithMessage(t, notempty, "Email %s should not be empty", email)
	})
}

func TestEmail_EmailContainArrobas(t *testing.T) {

	// only one arroba allowed

	t.Run("empty email", func(t *testing.T) {
		email := ""
		zeroArrobas := EmailContainArrobas(email)
		assert.FalseWithMessage(t, zeroArrobas, "Email %s should not contain arrobas", email)
	})

	t.Run("no arrobas", func(t *testing.T) {
		email := "foofoo.foo"
		zeroArrobas := EmailContainArrobas(email)
		assert.FalseWithMessage(t, zeroArrobas, "Email %s should not contain arrobas", email)
	})
	t.Run("one arroba", func(t *testing.T) {
		email := "foo@foo.com"
		oneArroba := EmailContainArrobas(email)
		assert.TrueWithMessage(t, oneArroba, "Email %s should contain arrobas", email)
	})

	t.Run("two arrobas", func(t *testing.T) {
		email := "foo@@foo.com"
		twoArrobas := EmailContainArrobas(email)
		assert.FalseWithMessage(t, twoArrobas, "Email %s should not contain arrobas", email)
	})
	t.Run("three arrobas", func(t *testing.T) {
		email := "foo@@@foo.com"
		threeArrobas := EmailContainArrobas(email)
		assert.FalseWithMessage(t, threeArrobas, "Email %s should not contain arrobas", email)
	})

	t.Run("300 arrobas", func(t *testing.T) {
		email := strings.Repeat("foo@", 300) + "foo.com"
		overflow := EmailContainArrobas(email)
		assert.FalseWithMessage(t, overflow, "Email %s should not contain arrobas", email)

	})
}

func TestEmail_EmailContainDomain(t *testing.T) {

	t.Run("not domain", func(t *testing.T) {
		email := "foo@"
		domain := EmailContainDomain(email)
		assert.FalseWithMessage(t, domain, "Email %s should not contain domain", email)
	})

	t.Run("domain", func(t *testing.T) {
		email := "foo@x.com"

		domain := EmailContainDomain(email)
		assert.TrueWithMessage(t, domain, "Email %s should contain domain", email)
	})

	t.Run("domain with subdomain", func(t *testing.T) {
		email := "foo@foo.com.com"
		domain := EmailContainDomain(email)
		assert.TrueWithMessage(t, domain, "Email %s should contain domain", email)
	})
	// more edge cases

}

func TestEmail_EmailHaveForbiddenChars(t *testing.T) {
	t.Run("empty email", func(t *testing.T) {
		email := ""
		forbidden := EmailHaveForbiddenChars(email)
		assert.FalseWithMessage(t, forbidden, "Email %s should not have forbidden chars", email)
	})

	t.Run("no forbidden chars", func(t *testing.T) {
		email := "fo.o_fo+fo.o@foo.com"

		forbidden := EmailHaveForbiddenChars(email)
		assert.FalseWithMessage(t, forbidden, "Email %s should not have forbidden chars", email)
	})

	// test all the forbidden chars
	for _, char := range forbiddenEmailChars {
		t.Run("forbidden char "+char, func(t *testing.T) {
			email := "foo" + char + "foo.com"
			forbidden := EmailHaveForbiddenChars(email)
			assert.TrueWithMessage(t, forbidden, "Email %s should have forbidden chars", email)
		})
	}

}

var forbiddenEmailChars = []string{
	" ",  // Space
	"\t", // Tab
	"\n", // Newline
	"\r", // Carriage return
	",",  // Comma
	";",  // Semicolon
	":",  // Colon
	"\\", // Backslash
	"/",  // Slash
	"(",  // Opening parenthesis
	")",  // Closing parenthesis
	"[",  // Opening bracket
	"]",  // Closing bracket
	"{",  // Opening brace
	"}",  // Closing brace
	"<",  // Less than
	">",  // Greater than
	"|",  // Vertical bar
	"`",  // Backtick
	"^",  // Caret
	"*",  // Asterisk
	"!",  // Exclamation mark
	"#",  // Hash
	"$",  // Dollar
	"%",  // Percent
	"&",  // Ampersand
	"=",  // Equal
	"?",  // Question mark
	"~",  // Tilde
}
