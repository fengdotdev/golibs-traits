package golibstraits

type JSONSerializable interface {
	ToJSON() (string, error)
}

type JSONDeserializable interface {
	FromJSON(s string) error
}

type JSONTrait interface {
	JSONDeserializable
	JSONSerializable
}
