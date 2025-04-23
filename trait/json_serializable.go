package trait

type JSONSerializer interface {
	JSONDeserializable
	JSONSerializable
}

type JSONSerializable interface {
	ToJSON() (string, error)
}

type JSONDeserializable interface {
	FromJSON(s string) error
}

// must be deprecated if favor or JSONSerialization
type JSONTrait interface {
	JSONDeserializable
	JSONSerializable
}
