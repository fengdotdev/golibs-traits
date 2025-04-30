package trait


// NOT READY FOR PRODUCTION
type MapSerializer interface {
	MapDeserializable
	MapSerializable
}
// NOT READY FOR PRODUCTION
type MapSerializable interface {
	ToMap() (map[string]any, error)
}
// NOT READY FOR PRODUCTION
type MapDeserializable interface {
	FromMap(m map[string]any) error
}

