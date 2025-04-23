package golibstraits

type MapSerialization interface {
	MapDeserializable
	MapSerializable
}

type MapSerializable interface {
	ToMap() (map[string]any, error)
}

type MapDeserializable interface {
	FromMap(m map[string]any) error
}

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

// ex:  typename: "int", value: 1
type TypeStringer interface {
	TypeName() string
	ValueString() string
}
