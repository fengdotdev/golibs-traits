package trait

// NOT READY FOR PRODUCTION
type Binarer interface {
	ToBinary() ([]byte, error)
	FromBinary([]byte) error
}

