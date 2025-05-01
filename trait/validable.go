package trait

// Validable must be a container of a type that can be validated
// asume that u need to conserve the data in the container valid or not
// by default, the data is not validated
// also the type must be protected against nil values or zero values ex: Mytype{} will be invalid

type Validable interface {
	IsValid() bool
	IsInvalid() bool
	Validate()   // asume inner checking
	Reason() error // reason of the validation
}

// NOT READY FOR PRODUCTION
type ValidableDeep[SELF any] interface {
	Validable
               
	Validator() func(SELF) (bool, error)        // validator function external or nil
	ValidatorDefault() func(SELF) (bool, error) // default validator function

	IsDefaultValidator() bool                 // checks if the validator is the default one
	IsExternalValidator() bool                // checks if the validator is the external one
	SetValidator(fn func(SELF) (bool, error)) // sets the validator function
	SetValidatorNil()                         // sets the validator function to nil
}
