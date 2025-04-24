package trait

type Binarer interface {
	ToBinary() ([]byte, error)
	FromBinary([]byte) error
}

