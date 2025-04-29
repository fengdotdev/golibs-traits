package trait_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
	"github.com/fengdotdev/golibs-traits/exampletypes/resultex"
	"github.com/fengdotdev/golibs-traits/trait"
)

func TestResult_ContentString(t *testing.T) {

	type NetworkResult = trait.Result[string]
	var _ NetworkResult = &resultex.NetworkResult{}

	t.Run("Successful Result", func(t *testing.T) {

		var successful NetworkResult = resultex.NewValidNetworkResult()

		assert.EqualWithMessage(t, successful.String(), "NetworkResult: Success", "Expected the content to be 'Success'")

		t.Run("ValueOrPanic", func(t *testing.T) {
			assert.AssertNotPanicWithMessage(t, func() {
				assert.EqualWithMessage(t, successful.ValueOrPanic(), "Success", "Expected the content to be 'Success'")
			}, "Expected no panic when accessing the value of a successful result")
		})

		t.Run("Value", func(t *testing.T) {
			assert.EqualWithMessage(t, successful.Value(), "Success", "Expected the content to be 'Success'")
		})

		t.Run("ValueOr", func(t *testing.T) {
			assert.EqualWithMessage(t, successful.ValueOr("Default"), "Success", "Expected the content to be 'Success'")
		})

		t.Run("ValueOrErr", func(t *testing.T) {
			successfulValue, err := successful.ValueOrErr()
			assert.EqualWithMessage(t, successfulValue, "Success", "Expected the content to be 'Success'")
			assert.NilWithMessage(t, err, "Expected no error")
			assert.EqualWithMessage(t, successful.Error(), nil, "Expected no error")
		})

		t.Run("Error", func(t *testing.T) {
			assert.NilWithMessage(t, successful.Error(), "Expected no error")
		})

		t.Run("IsValid", func(t *testing.T) {
			assert.TrueWithMessage(t, successful.IsValid(), "Expected the result to be successful")
		})

		t.Run("String", func(t *testing.T) {
			assert.EqualWithMessage(t, successful.String(), "NetworkResult: Success", "Expected the content to be 'Success'")
		})

	})

	t.Run("Failed Result", func(t *testing.T) {

		// Create a new NetworkResult with an error
		var failed NetworkResult = resultex.NewInvalidNetworkResult()

		assert.EqualWithMessage(t, failed.String(), "NetworkResultError: network error", "Expected the content to be empty")

		t.Run("ValueOrPanic", func(t *testing.T) {
			assert.AssertPanicWithMessage(t, func() {
				failed.ValueOrPanic()
			}, "Expected panic when accessing the value of a failed result")
		})

		t.Run("Value", func(t *testing.T) {
			assert.EqualWithMessage(t, failed.Value(), "", "Expected the content to be empty")
		})

		t.Run("ValueOr", func(t *testing.T) {
			assert.EqualWithMessage(t, failed.ValueOr("Default"), "Default", "Expected the content to be 'Default'")
		})

		t.Run("ValueOrErr", func(t *testing.T) {
			failedValue, err := failed.ValueOrErr()
			assert.EqualWithMessage(t, failedValue, "", "Expected the content to be empty")
			assert.NotNilWithMessage(t, err, "Expected an error")
			assert.ErrorWithMessage(t, err, "Expected an error")
		})

		t.Run("Error", func(t *testing.T) {
			assert.NotNilWithMessage(t, failed.Error(), "Expected an error")
			assert.ErrorWithMessage(t, failed.Error(), "Expected an error")
		})

		t.Run("IsValid", func(t *testing.T) {
			assert.FalseWithMessage(t, failed.IsValid(), "Expected the result to be invalid")
		})

		t.Run("String", func(t *testing.T) {
			assert.EqualWithMessage(t, failed.String(), "NetworkResultError: network error", "Expected the content to be empty")
		})

	})

}
