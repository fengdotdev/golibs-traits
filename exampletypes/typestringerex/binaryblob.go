package typestringerex

import (
	"encoding/base64"

	"github.com/fengdotdev/golibs-traits/trait"
)

var _ = trait.TypeStringer(&BinaryBlob{})

type BinaryBlob struct {
	content []byte
}


// NewBinaryBlob creates a new BinaryBlob from a byte slice
func NewBinaryBlob(content []byte) *BinaryBlob {
	return &BinaryBlob{
		content: content,
	}
}


// NewBinaryBlobFromString creates a new BinaryBlob from a base64 encoded string
func NewBinaryBlobFromString(content string) (*BinaryBlob, error) {
	// decode from base64
	decodedContent, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}
	// create a new BinaryBlob with the decoded content
	// and return it

	return &BinaryBlob{
		content: decodedContent,
	}, nil

}

// Content returns the content of the BinaryBlob as a byte slice
func (b *BinaryBlob) Content() []byte {
	return b.content
}

// SetContent sets the content of the BinaryBlob
// to the given byte slice
func (b *BinaryBlob) SetContent(content []byte) {
	b.content = content
}

// TypeName implements trait.TypeStringer.
// return the name of the type
// in this case, it is "binaryblob"
func (b *BinaryBlob) TypeName() string {
	return "binaryblob"
}

// ValueString implements trait.TypeStringer.
// encode the content to base64
// and return it as a string
func (b *BinaryBlob) ValueString() string {

	// encode to base64
	return base64.StdEncoding.EncodeToString(b.content)
}
