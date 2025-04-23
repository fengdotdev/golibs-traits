package trait

type Cloneable[SELF any] interface {
	Clone() SELF // returns a new instance of the same type if is some singleton must panic
}
