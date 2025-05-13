package crudex

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.CRUD[string, any] = &BigMap[string, any]{}

type BigMap[K Indexable, V any] struct {
	container map[K]Map[K, V]
}

// All implements trait.CRUD.
func (b *BigMap[K, V]) All() map[K]V {
	panic("unimplemented")
}

// Count implements trait.CRUD.
func (b *BigMap[K, V]) Count(term string) (int, error) {
	panic("unimplemented")
}

// Create implements trait.CRUD.
func (b *BigMap[K, V]) Create(id K, item V) error {
	panic("unimplemented")
}

// Delete implements trait.CRUD.
func (b *BigMap[K, V]) Delete(id K) error {
	panic("unimplemented")
}

// Exists implements trait.CRUD.
func (b *BigMap[K, V]) Exists(id K) (bool, error) {
	panic("unimplemented")
}

// Keys implements trait.CRUD.
func (b *BigMap[K, V]) Keys() []K {
	panic("unimplemented")
}

// Len implements trait.CRUD.
func (b *BigMap[K, V]) Len() int {
	panic("unimplemented")
}

// Read implements trait.CRUD.
func (b *BigMap[K, V]) Read(id K) (V, error) {
	panic("unimplemented")
}

// Search implements trait.CRUD.
func (b *BigMap[K, V]) Search(term string, where string) ([]V, error) {
	panic("unimplemented")
}

// SearchAll implements trait.CRUD.
func (b *BigMap[K, V]) SearchAll(term string) ([]V, error) {
	panic("unimplemented")
}

// Update implements trait.CRUD.
func (b *BigMap[K, V]) Update(id K, item V) error {
	panic("unimplemented")
}

// Values implements trait.CRUD.
func (b *BigMap[K, V]) Values() []V {
	panic("unimplemented")
}

// constructor
func NewBigMap[K Indexable, V any]() *BigMap[K, V] {
	return &BigMap[K, V]{
		container: make(map[K]Map[K, V]),
	}
}
