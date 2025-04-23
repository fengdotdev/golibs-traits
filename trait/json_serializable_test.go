package trait_test

import (
	"encoding/json"
	"testing"

	"github.com/fengdotdev/golibs-traits/trait"
)

type Foo struct {
	Name string `json:"Name"`
}

func (f *Foo) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(f)
	return string(jsonBytes), err
}

func (f *Foo) FromJSON(s string) error {
	return json.Unmarshal([]byte(s), f)
}

func Test_JSONTrait(t *testing.T) {
	var f trait.JSONTrait = &Foo{Name: "John"}

	jsonStr, err := f.ToJSON()

	if err != nil {
		t.Errorf("ToJSON() failed: %v", err)
	}
	if jsonStr != `{"Name":"John"}` {
		t.Errorf("ToJSON() returned wrong value: %v", jsonStr)
	}

	var f2 Foo
	err = f2.FromJSON(jsonStr)

	if err != nil {
		t.Errorf("FromJSON() failed: %v", err)
	}
	if f2.Name != "John" {
		t.Errorf("FromJSON() returned wrong value: %v", f2.Name)
	}

	t.Logf("JSONTrait test passed")
}
