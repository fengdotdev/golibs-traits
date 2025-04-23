package golibstraits

type MapTypeStringerSerializer interface {
	MapTypeStringerDeserializable
	MapTypeStringerSerializable
}

type MapTypeStringerSerializable interface {
	ToMapTypeStringer() (map[string]TypeStringer, error)
}

type MapTypeStringerDeserializable interface {
	FromMapTypeStringer(m map[string]TypeStringer) error
}



