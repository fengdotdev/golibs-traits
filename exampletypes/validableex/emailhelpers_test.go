package validableex

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

var (
	validEmails = []string{
		"usuario@dominio.com",
		"otro.usuario123@subdominio.ejemplo.cl",
		"mi_correo+etiqueta@proveedor.net",
		"primera.letra@dominio-con-guion.info",
		"a@b.c",
	}

	invalidEmails = []struct {
		Email   string
		Comment string
	}{
		{"sinarroba.com", "Falta el símbolo '@' que separa el usuario del dominio."},
		{"usuario@", "Falta la parte del dominio después del '@'."},
		{"@dominio.com", "Falta la parte del usuario antes del '@'."},
		{"usuario@.com", "El dominio no puede empezar con un punto."},
		{"usuario@dominio..com", "El dominio contiene dos puntos consecutivos."},
		{"usuario@dominio com", "El dominio contiene un espacio en blanco."},
		{"usuario@dominio-.com", "El dominio no puede terminar con un guion."},
		{"usuario@-dominio.com", "El dominio no puede empezar con un guion."},
		{"usuario@muyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy.com", "La longitud del dominio excede el límite práctico."},
		{"usuario con espacio@dominio.com", "La parte del usuario contiene un espacio en blanco."},
		{"usuario[parentesis]@dominio.com", "La parte del usuario contiene caracteres no permitidos (corchetes)."},
	}
)

func TestEmail_Validator_Valid(t *testing.T) {

	for _, email := range validEmails {
		t.Run(email, func(t *testing.T) {
			valid, err := EmailValidator(email)
			assert.TrueWithMessage(t, valid, "Email %s should be valid", email)
			assert.NoErrorWithMessage(t, err, fmt.Sprintf("Email %s should not have a reason", email))
		})
	}

}


func TestEmail_Validator_Invalid(t *testing.T) {

	for _, email := range invalidEmails {
		t.Run(email.Email, func(t *testing.T) {
			valid, err := EmailValidator(email.Email)
			assert.FalseWithMessage(t, valid, "Email %s should not be valid", email.Email)
			assert.ErrorWithMessage(t, err, fmt.Sprintf("Email %s should have a reason", email.Email))
		})
	}
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
