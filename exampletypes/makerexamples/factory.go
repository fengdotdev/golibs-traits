package makerexamples

import (
	"errors"
	"fmt"

	"github.com/fengdotdev/golibs-traits/trait"
)

var customTypesMap = make(map[string]any)

func RegisterCustomType(name string, obj any) {
	if _, ok := customTypesMap[name]; ok {
		panic("custom type already exists")
	}
	customTypesMap[name] = obj
}


var ErrCustomTypeNotFound = errors.New("custom type not found")

func GetCustomType[T any](name string) (T, error) {
    obj, ok := customTypesMap[name]
    if !ok {
        var zero T // Get the zero value of the type
        return zero, fmt.Errorf("custom type %s not found: %w", name, ErrCustomTypeNotFound)
    }
    // Type assertion is still needed
    if val, ok := obj.(T); ok {
        return val, nil
    }
    var zero T
    return zero, fmt.Errorf("custom type %s has incorrect underlying type", name)
}


type Factory struct {
	blueprints map[string]any
}

// blueprint objects must implement Maker[T] interface { Make(args ..any) T }
func NewFactory(blueprints ...any) *Factory {
	f := &Factory{
		blueprints: make(map[string]any),
	}

	for _, b := range blueprints {
		nameof := fmt.Sprintf("%T", b)

		if _, ok := f.blueprints[nameof]; ok {
			panic("blueprint already exists")
		}
		f.blueprints[nameof] = b
	}

	return f
}

func (f *Factory) Makeit(obj string) (any, error) {
	blueprint, ok := f.blueprints[obj]
	if !ok {
		return nil, fmt.Errorf("blueprint %s not found", obj)
	}

	T := fmt.Sprintf("%T", blueprint)
	if T == "any" {
		return nil, fmt.Errorf("blueprint %s is not a valid type", obj)
	}

	maker, ok := blueprint.(trait.Maker[any])
	if !ok {
		//return nil, fmt.Errorf("blueprint %s does not implement Maker", obj)
	}

	return maker.Make(), nil
}
