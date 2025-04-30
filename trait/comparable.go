package trait


// NOT READY FOR PRODUCTION
type Comparable interface {

	// AreEqual checks if the current instance is equal to another instance of the same type.
	AreEqual(other any) bool

	// AreNotEqual checks if the current instance is not equal to another instance of the same type.
	AreNotEqual(other any) bool

	// AreSameType checks if the current instance is of the same type as another instance.
	AreSameType(other any) bool

	// AreNotSameType checks if the current instance is not of the same type as another instance.
	AreNotSameType(other any) bool
}


