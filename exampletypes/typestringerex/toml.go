package typestringerex

import "github.com/fengdotdev/golibs-traits/trait"

var _ = trait.TypeStringer(&TOML{})

type TOML struct {
	content string
}

func NewToML(content string) *TOML {
	return &TOML{
		content: content,
	}
}


// TypeName implements trait.TypeStringer.
func (t *TOML) TypeName() string {
	return "toml"
}

// ValueString implements trait.TypeStringer.
func (t *TOML) ValueString() string {
	return t.content
}
