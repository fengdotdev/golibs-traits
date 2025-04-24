package trait

type Equality interface {
	Comparable

	Equal(other any, treshhold float64) bool
	NotEqual(other any, treshhold float64) bool

	// ex: dog and person may have same name

	Similar(other any, treshhold float64) map[string]bool
	Similarity(other any) map[string]float64

	// ex: dog and person may have name
	// but but dog have owner
	// and person have pet
	CommonFields(other any) map[string]bool
	DifferentFields(other any) map[string]bool
}
