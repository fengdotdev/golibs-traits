package golibstraits

type Result[CONTENT any] interface {
	ValueOrPanic() CONTENT      // returns the value or panics if the value is not valid
	ValueOr(or CONTENT) CONTENT // returns the value or the given default value
	ValueOrErr() (CONTENT, error)
	Error() error   // returns the error if the value is not valid
	IsValid() bool  // checks if the value is valid
	String() string // returns the string representation of the value
}


