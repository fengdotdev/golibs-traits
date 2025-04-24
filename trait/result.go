package trait

type Result[CONTENT any] interface {
	// ValueOrPanic returns the underlying value.
	// It panics if the value is not considered valid.
	// Use this method when you are certain the Result holds a valid value.
	ValueOrPanic() CONTENT


	// Value returns the underlying value.
	// this returns the value regardless of its validity.
	// if the value is not valid, it will return the zero value of CONTENT or nil.
	// else, it will return the valid value.
	// this not panic if the value is not valid.

	Value() CONTENT


	// ValueOr returns the underlying value if it's valid; otherwise, it returns the provided 'or' value.
	// This method offers a safe way to access the value with a fallback.
	ValueOr(or CONTENT) CONTENT

	// ValueOrErr returns the underlying value and an error.
	// If the value is not valid, the returned value will be the zero value of CONTENT,
	// and the error will provide the reason for the invalid state.
	ValueOrErr() (CONTENT, error)

	// Error returns the error associated with the Result, if the value is not valid.
	// If the value is valid, this method returns nil.
	Error() error

	// IsValid returns true if the Result holds a valid value; otherwise, it returns false.
	// This method allows you to explicitly check the validity of the Result before accessing its value.
	IsValid() bool

	// String returns a string representation of the underlying value.
	// The format of the string will depend on the type of CONTENT.
	String() string
}