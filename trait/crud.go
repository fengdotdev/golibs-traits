package trait

import "context"

type Indexable interface {
	~int | ~string | ~float64
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
	Iterate(fn func(K, V) (stop bool, err error)) error
	Clean()
	Populate(map[K]V)
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
	Iterate(ctx context.Context, fn func(K, V) (stop bool, err error)) error
	Clean(ctx context.Context)
	Populate(ctx context.Context, items map[K]V)
}
