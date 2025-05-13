package crudex

import (
	"errors"
	"sync"

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

type Map[K Indexable, V any] struct {
	container map[K]V
	mu        sync.RWMutex
}

//constructor

func NewMap[K Indexable, V any]() *Map[K, V] {

	return &Map[K, V]{
		container: make(map[K]V),
		mu:        sync.RWMutex{},
	}
}

// All implements trait.CRUD.
func (m *Map[K, V]) All() map[K]V {

	// todo check this may is a better implementation

	m.mu.RLock()
	defer m.mu.RUnlock()

	copyMap := make(map[K]V, len(m.container))
	for key, value := range m.container {
		copyMap[key] = value
	}
	return copyMap
}

// Create implements trait.CRUD.
func (m *Map[K, V]) Create(id K, item V) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, exists := m.container[id]; exists {
		return ErrAlreadyExists
	}
	m.container[id] = item
	return nil

}

// Delete implements trait.CRUD.
func (m *Map[K, V]) Delete(id K) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, exists := m.container[id]; !exists {
		return ErrNotFound
	}
	delete(m.container, id)
	return nil
}

// Exists implements trait.CRUD.
func (m *Map[K, V]) Exists(id K) (bool, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, exists := m.container[id]; !exists {
		return false, ErrNotFound
	}
	return true, nil
}

// Keys implements trait.CRUD.
func (m *Map[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	keys := make([]K, 0, len(m.container))
	for k := range m.container {
		keys = append(keys, k)
	}
	return keys
}

// Len implements trait.CRUD.
func (m *Map[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.container)
}

// Read implements trait.CRUD.
func (m *Map[K, V]) Read(id K) (V, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var zero V
	if _, exists := m.container[id]; !exists {
		return zero, ErrNotFound
	}
	return m.container[id], nil
}

// Update implements trait.CRUD.
func (m *Map[K, V]) Update(id K, item V) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, exists := m.container[id]; !exists {
		return ErrNotFound
	}
	m.container[id] = item
	return nil
}

// Values implements trait.CRUD.
func (m *Map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	values := make([]V, 0, len(m.container))
	for _, v := range m.container {
		values = append(values, v)
	}
	return values
}

// Iterate implements trait.CRUD.
// Iterate iterates over the map and applies the provided function to each key-value pair.
// If the function returns an error, the iteration stops.
func (m *Map[K, V]) Iterate(fn func(K, V) error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for k, v := range m.container {
		if err := fn(k, v); err != nil {
			return
		}
	}
}

// Clean implements trait.CRUD.
// Clean clears the map.
func (m *Map[K, V]) Clean() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.container = make(map[K]V)
}

// Populate populates the map with the provided items.
// It locks the map for writing, so it should be used with caution in concurrent environments.
// this overrides the current map content and may replace some items
func (m *Map[K, V]) Populate(items map[K]V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for k, v := range items {
		// ignore map.copy dont use it
		m.container[k] = v
	}
}
