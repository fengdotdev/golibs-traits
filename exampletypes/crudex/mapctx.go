package crudex

import "sync"

type MapWithContext[K Indexable, V any] struct {
	container map[K]V
	mu        sync.RWMutex
}
