package trait


// NOT READY FOR PRODUCTION
type Typeable[SELF any] interface {
	GetType() SELF // must be a empty object o zero object
	GetTypeName() string
	AreSameType(other any) bool
}
