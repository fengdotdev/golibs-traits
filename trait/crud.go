package trait

import "context"

type CRUD[T any] interface {
	Create(id any,item T) (T, error)
	Read(id any) (T, error)
	Update(id any, item T) error
	Delete(id any) error
	Exists(id any) (bool, error)
	Count(term string) (int, error)
	Search(term string, where string) ([]T, error)
}
type CRUDWithCTX[T any] interface {
	Create(ctx context.Context, id any, item T) error
	Read(ctx context.Context, id any) (T, error)
	Update(ctx context.Context, id any, item T) error
	Delete(ctx context.Context, id any) error
	Exists(ctx context.Context, id any) (bool, error)
	Count(ctx context.Context, term string) (int, error)
	Search(ctx context.Context, term string, where string) ([]T, error)
}
