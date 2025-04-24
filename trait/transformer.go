package trait

type Transformer[In any, Out any] interface {
	Transform(in In) (Out, error)
	ReverseTransform(out Out) (In, error)
}
