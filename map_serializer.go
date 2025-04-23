package golibstraits

type MapSerializer interface {
	MapDeserializable
	MapSerializable
}

type MapSerializable interface {
	ToMap() (map[string]any, error)
}

type MapDeserializable interface {
	FromMap(m map[string]any) error
}

