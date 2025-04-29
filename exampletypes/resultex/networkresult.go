package resultex

import (
	"fmt"

	"github.com/fengdotdev/golibs-traits/trait"
)

// These lines are compile-time checks to ensure that NetworkResult correctly implements the trait.Result interface for string.
// They verify that any pointer to NetworkResult and any instance of NetworkResult can be treated as a trait.Result[string].
var _ trait.Result[string] = (*NetworkResult)(nil)
var _ trait.Result[string] = &NetworkResult{}

// NetworkResult is a concrete implementation of the trait.Result interface, specifically designed to hold the outcome
// of network operations that might succeed with a string value or fail with an error.
type NetworkResult struct {
	// err stores any error that occurred during the network operation.
	// It will be nil if the operation was successful.
	err error
	// val holds the string value obtained from a successful network operation.
	// It will be the zero value of string ("") if the operation failed.
	val string
}

// Constructors

// NewNetworkResult creates a new NetworkResult with the given value and error.
// Use this constructor when the outcome of an operation might be a successful value or an error.
func NewNetworkResult(val string, err error) *NetworkResult {
	return &NetworkResult{
		err: err,
		val: val,
	}
}

// NewNetworkResultWithError creates a new NetworkResult representing a failed operation with the given error.
// The value of the result will be the zero value for string ("").
func NewNetworkResultWithError(err error) *NetworkResult {
	return &NetworkResult{
		err: err,
		val: "",
	}
}

// Constructors for Tests

// NewValidNetworkResult creates a new NetworkResult representing a successful operation for testing purposes.
// It has a nil error and a default successful string value ("Success").
func NewValidNetworkResult() *NetworkResult {
	return &NetworkResult{
		err: nil,
		val: "Success",
	}
}

// NewInvalidNetworkResult creates a new NetworkResult representing a failed operation for testing purposes.
// It has a predefined error ("network error") and an empty string value.
func NewInvalidNetworkResult() *NetworkResult {
	return &NetworkResult{
		err: fmt.Errorf("network error"),
		val: "",
	}
}

// methods

// ValueOrPanic implements the ValueOrPanic method of the trait.Result interface.
// It returns the underlying string value if the NetworkResult is valid.
// If the result is not valid (i.e., an error occurred), it panics with the associated error.
// This method should be used in scenarios where an invalid result is considered an unrecoverable program error.
func (n *NetworkResult) ValueOrPanic() string {
	if n.IsValid() {
		return n.val
	}
	panic(n.err)
}

// Value implements the Value method of the trait.Result interface.
// It returns the underlying string value regardless of the validity of the NetworkResult.
// If the result is not valid, it will return the zero value of string ("").
func (n *NetworkResult) Value() string {
	return n.val
}

// ValueOr implements the ValueOr method of the trait.Result interface.
// It returns the underlying string value if the NetworkResult is valid.
// If the result is not valid (i.e., an error occurred), it returns the provided 'or' string value as a default.
func (n *NetworkResult) ValueOr(or string) string {
	if n.IsValid() {
		return n.val
	}
	return or
}

// ValueOrErr implements the ValueOrErr method of the trait.Result interface.
// It returns the underlying string value and a nil error if the NetworkResult is valid.
// If the result is not valid, it returns the zero value of string ("") and the associated error.
// This method forces the caller to explicitly handle the potential error.
func (n *NetworkResult) ValueOrErr() (string, error) {
	if n.IsValid() {
		return n.val, nil
	}
	return "", n.err
}

// Error implements the Error method of the trait.Result interface.
// It returns the error associated with the NetworkResult. If the network operation was successful, it returns nil.
func (n *NetworkResult) Error() error {
	return n.err
}

// IsValid implements the IsValid method of the trait.Result interface.
// It returns true if the NetworkResult represents a successful operation (i.e., the 'err' field is nil),
// and false otherwise.
func (n *NetworkResult) IsValid() bool {
	return n.err == nil
}

// String implements the String method of the trait.Result interface.
// It provides a human-readable string representation of the NetworkResult.
// If the result is valid, NetworkResult: <value> is returned, where <value> is the underlying string value.
// If the result is an error, NetworkResultError: <error> is returned, where <error> is the error message.
func (n *NetworkResult) String() string {
	if n.IsValid() {
		return fmt.Sprintf("NetworkResult: %s", n.val)
	}
	return fmt.Sprintf("NetworkResultError: %s", n.err)
}
