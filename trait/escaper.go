package trait



type EscaperT[T any] interface {
	Escape(raw T) T
	Unescape(cooked T) T
}

type Escaper interface {
	Escape(string) string
	Unescape(string) string
}


