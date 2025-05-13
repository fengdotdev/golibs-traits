package trait

import "context"

type Indexable interface {
	~int | ~string | ~float64
}
type CRUD1[K Indexable, V any] interface {
	Create(id K, item V) error
	Read(id K) (V, error)
	Update(id K, item V) error
	Delete(id K) error
	Exists(id K) (bool, error)
	Len() int
	Keys() []K
	Values() []V
	All() map[K]V
}


type CRUD[K Indexable, V any] interface {
	Create(id K, item V) error
	Read(id K) (V, error)
	Update(id K, item V) error
	Delete(id K) error
	Exists(id K) (bool, error)
	Len() int
	Keys() []K
	Values() []V
	All() map[K]V
	Count(term string) (int, error)
	Search(from string, term string, where func(item V) bool) ([]V, error)
	SearchAll(from string, term string) ([]V, error)
}
type CRUDWithCTX[K Indexable, V any] interface {
	Create(ctx context.Context, id K, item V) error
	Read(ctx context.Context, id K) (V, error)
	Update(ctx context.Context, id K, item V) error
	Delete(ctx context.Context, id K) error
	Exists(ctx context.Context, id K) (bool, error)
	Len(ctx context.Context) int
	Keys(ctx context.Context) []K
	Values(ctx context.Context) []V
	All(ctx context.Context) map[K]V
	Count(ctx context.Context, term string) (int, error)
	Search(ctx context.Context, term string, where func(item V) bool, from ...string) ([]V, error)
	SearchAll(ctx context.Context, term string) ([]V, error)
}
