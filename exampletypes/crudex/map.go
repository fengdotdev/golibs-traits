package crudex

import (
	"errors"

	"github.com/fengdotdev/golibs-traits/exampletypes/handyfuncs"
	"github.com/fengdotdev/golibs-traits/trait"
)

var (
	ErrorIdTypeAlreadyInited = errors.New("ID type already initialized")
	ErrIdTypeMismatch        = errors.New("ID type mismatch")
	ErrAlreadyExists         = errors.New("item already exists")
	ErrNotFound              = errors.New("item not found")
	ErrInvalidID             = errors.New("invalid ID")
	ErrInvalidItem           = errors.New("invalid item")
	ErrInvalidTerm           = errors.New("invalid search term")
	ErrInvalidWhere          = errors.New("invalid search where")
	ErrInvalidContext        = errors.New("invalid context")
)

var _ trait.CRUD[string, any] = &Map[string, any]{}


type Indexable interface {
    ~int | ~string | ~float64 
}

type Map[K Indexable, V any] struct {
	container map[K]V
}

//constructor




func NewMap[K Indexable, V any]() *Map[K, V] {

	return &Map[K, V]{
		container: make(map[K]V),
	}
}

// All implements trait.CRUD.
func (m *Map[K, V]) All() map[K]V {
	return m.container
}

// Count implements trait.CRUD.
func (m *Map[K, V]) Count(term string) (int, error) {
	if term == "" {
		return 0, ErrInvalidTerm
	}

	count := 0
	for _, item := range m.container {

		switch v := any(item).(type) {
		case string:
			if handyfuncs.LookupStringIn(term, v) {
				count++
			}
		default:
		}
	}
	return count, nil
}

// Create implements trait.CRUD.
func (m *Map[K, V]) Create(id K, item V) error {

	if _, exists := m.container[id]; exists {
		return ErrAlreadyExists
	}
	m.container[id] = item
	return nil

}

// Delete implements trait.CRUD.
func (m *Map[K, V]) Delete(id K) error {
	if _, exists := m.container[id]; !exists {
		return ErrNotFound
	}
	delete(m.container, id)
	return nil
}

// Exists implements trait.CRUD.
func (m *Map[K, V]) Exists(id K) (bool, error) {

	_, exists := m.container[id]
	return exists, nil
}

// Keys implements trait.CRUD.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.container))
	for k := range m.container {
		keys = append(keys, k)
	}
	return keys
}

// Len implements trait.CRUD.
func (m *Map[K, V]) Len() int {
	return len(m.container)
}

// Read implements trait.CRUD.
func (m *Map[K, V]) Read(id K) (V, error) {

	var zero V
	if _, exists := m.container[id]; !exists {
		return zero, ErrNotFound
	}
	return m.container[id], nil
}

// Search implements trait.CRUD.
// this function is not supported for this map implementation
// use SearchAll instead
func (m *Map[K, V]) Search(term string, where string) ([]V, error) {
	return nil, errors.ErrUnsupported
}

// SearchAll implements trait.CRUD.
// SearchAll searches for all items that match the term in the specified field
func (m *Map[K, V]) SearchAll(term string) ([]V, error) {
	if term == "" {
		return nil, ErrInvalidTerm
	}

	results := make([]V, 0)
	for _, item := range m.container {

		switch v := any(item).(type) {
		case string:
			if handyfuncs.LookupStringIn(term, v) {
				results = append(results, item)
			}
		default:
		}
	}
	return results, nil
}

// Update implements trait.CRUD.
func (m *Map[K, V]) Update(id K, item V) error {
	if _, exists := m.container[id]; !exists {
		return ErrNotFound
	}
	m.container[id] = item
	return nil
}

// Values implements trait.CRUD.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, len(m.container))
	for _, v := range m.container {
		values = append(values, v)
	}
	return values
}
