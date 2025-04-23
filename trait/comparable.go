package trait

type Comparable[SELF any] interface {
	AreEqual(other SELF) bool
	AreNotEqual(other SELF) bool
	AreSameType(other any) bool
	AreNotSameType(other any) bool
}
