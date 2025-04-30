package trait


// NOT READY FOR PRODUCTION
type EscaperT[T any] interface {
	Escape(raw T) T
	Unescape(cooked T) T
}

// NOT READY FOR PRODUCTION
type Escaper interface {
	Escape(string) string
	Unescape(string) string
}


