package trait

// NOT READY FOR PRODUCTION
type GoParser[Origin any,Go any] interface {
	ParseGo(origin Origin) (Go, error)
}