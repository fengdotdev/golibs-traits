package golibstraits

type Validable[SELF any] interface {
	IsValid() bool
	IsInvalid() bool
}

type ValidableDeep[SELF any] interface {
	Validable[SELF]

	Validate() // asume inner checking
	Validator() func(SELF) (bool, error) // validator function external or nil
	ValidatorDefault() func(SELF) (bool, error) // default validator function

	IsDefaultValidator() bool                 // checks if the validator is the default one
	IsExternalValidator() bool                // checks if the validator is the external one
	SetValidator(fn func(SELF) (bool, error)) // sets the validator function
	SetValidatorNil()                         // sets the validator function to nil
}
