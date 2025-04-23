package trait

// ex:  1 =>  typename: "int", value: "1"
// ex fancymap => typename: "json", value: "{\"fancy\": \"map\"}"
// ex: binarydata => typename: "encode64", value: "SGVsbG8gV29ybGQh"
type TypeStringer interface {
	TypeName() string
	ValueString() string
}
