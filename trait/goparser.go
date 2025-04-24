package trait


type GoParser[Origin any,Go any] interface {
	ParseGo(origin Origin) (Go, error)
}