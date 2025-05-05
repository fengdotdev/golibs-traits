package validableex_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/validableex"
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
		{"u@corto.c", "El dominio de nivel superior es demasiado corto (debe tener al menos dos letras)."},
		{"usuario con espacio@dominio.com", "La parte del usuario contiene un espacio en blanco."},
		{"usuario[parentesis]@dominio.com", "La parte del usuario contiene caracteres no permitidos (corchetes)."},
	}
)

func TestEmail(t *testing.T) {

	t.Run("Valid Emails", func(t *testing.T) {
		for _, email := range validEmails {
			t.Run(email, func(t *testing.T) {
				e := validableex.NewEmail(email)
				t.Log(email)
				assert.FalseWithMessage(t, e.IsValid(), "Email %s should not be valid", email)
				assert.TrueWithMessage(t, e.IsInvalid(), "Email %s should be invalid", email)
				assert.ErrorWithMessage(t, e.Reason(), "Email %s should have a reason", email)
				assert.EqualWithMessage(t, e.String(), email, "Email %s should be equal to the raw email", email)

				e.Validate()
				assert.TrueWithMessage(t, e.IsValid(), "Email %s should be valid", email)
				assert.FalseWithMessage(t, e.IsInvalid(), "Email %s should not be invalid", email)
				t.Log(e.Reason())
				assert.NotErrorWithMessage(t, e.Reason(), "Email %s should not have a reason", email)
			})
		}
	})

	t.Run("Invalid Emails", func(t *testing.T) {

		for _, email := range invalidEmails {
			t.Run(email.Email, func(t *testing.T) {
				e := validableex.NewEmail(email.Email)
				assert.FalseWithMessage(t, e.IsValid(), "Email %s should not be valid", email.Email)
				assert.TrueWithMessage(t, e.IsInvalid(), "Email %s should be invalid", email.Email)
				assert.ErrorWithMessage(t, e.Reason(), "Email %s should have a reason", email.Email)
				assert.EqualWithMessage(t, e.String(), email.Email, "Email %s should be equal to the raw email", email.Email)

				e.Validate()
				assert.FalseWithMessage(t, e.IsValid(), "Email %s should not be valid", email.Email)
				assert.TrueWithMessage(t, e.IsInvalid(), "Email %s should be invalid", email.Email)
				assert.ErrorWithMessage(t, e.Reason(), "Email %s should have a reason", email.Email)
			})
		}

	})
}
