package trait
// NOT READY FOR PRODUCTION
type MapTypeStringerSerializer interface {
	MapTypeStringerDeserializable
	MapTypeStringerSerializable
}

// NOT READY FOR PRODUCTION
type MapTypeStringerSerializable interface {
	ToMapTypeStringer() (map[string]TypeStringer, error)
}

// NOT READY FOR PRODUCTION
type MapTypeStringerDeserializable interface {
	FromMapTypeStringer(m map[string]TypeStringer) error
}



