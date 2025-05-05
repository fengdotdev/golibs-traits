package trait

// Validable must be a container of a type that can be validated
// asume that u need to conserve the data in the container valid or not
// by default, the data is not validated
// also the type must be protected against nil values or zero values ex: Mytype{} will be invalid

type Validable[T any] interface {
	IsValid() bool
	IsInvalid() bool
	Validate()     // asume inner checking
	Reason() error // reason of the validation
	Value() T
}

// NOT READY FOR PRODUCTION
type ValidableDeep[T any] interface {
	Validable[T]

	Validator() func(T) (bool, error)        // validator function external or nil
	ValidatorDefault() func(T) (bool, error) // default validator function

	IsDefaultValidator() bool              // checks if the validator is the default one
	IsExternalValidator() bool             // checks if the validator is the external one
	SetValidator(fn func(T) (bool, error)) // sets the validator function
	SetValidatorNil()                      // sets the validator function to nil
}
